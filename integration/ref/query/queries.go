package query

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/models"
	"github.com/jalphad/gosqlgen/integration/ref/query/ast"
	"github.com/jalphad/gosqlgen/integration/ref/query/builder"
	"github.com/jalphad/gosqlgen/integration/ref/query/comments"
	"github.com/jalphad/gosqlgen/integration/ref/query/post_tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/posts"
	"github.com/jalphad/gosqlgen/integration/ref/query/tags"
	"github.com/jalphad/gosqlgen/integration/ref/query/users"
)

type updateBatchSource int

const (
	updateBatchSourceUnnest updateBatchSource = iota
	updateBatchSourceValues
)

type UpdateOption func(*updateOptions) error

type updateOptions struct {
	columns []ast.NamedExpression
	source  updateBatchSource
}

type SelectOption[T any] func(*selectOptions[T]) error

type selectOptions[T any] struct {
	columns []ast.Projection[T]
}

func SelectColumns[T any](columns ...ast.Projection[T]) SelectOption[T] {
	return func(options *selectOptions[T]) error {
		if len(columns) == 0 {
			return fmt.Errorf("query.SelectColumns requires at least one column")
		}
		for _, column := range columns {
			if column == nil {
				return fmt.Errorf("query.SelectColumns received a nil column")
			}
		}
		options.columns = columns
		return nil
	}
}

func newSelectOptions[T any](opts ...SelectOption[T]) (selectOptions[T], error) {
	var options selectOptions[T]
	for _, opt := range opts {
		if opt != nil {
			if err := opt(&options); err != nil {
				return selectOptions[T]{}, err
			}
		}
	}
	return options, nil
}

func UpdateColumns(columns ...ast.NamedExpression) UpdateOption {
	return func(options *updateOptions) error {
		if len(columns) == 0 {
			return fmt.Errorf("query.UpdateColumns requires at least one column")
		}
		options.columns = columns
		return nil
	}
}

func UpdateWithValues() UpdateOption {
	return func(options *updateOptions) error {
		options.source = updateBatchSourceValues
		return nil
	}
}

func UpdateWithUnnest() UpdateOption {
	return func(options *updateOptions) error {
		options.source = updateBatchSourceUnnest
		return nil
	}
}

func newUpdateOptions(opts ...UpdateOption) (updateOptions, error) {
	options := updateOptions{
		source: updateBatchSourceValues,
	}
	for _, opt := range opts {
		if opt != nil {
			if err := opt(&options); err != nil {
				return updateOptions{}, err
			}
		}
	}
	return options, nil
}

// CommentsDtoSelectOne selects a single comments DTO by primary key
func CommentsDtoSelectOne(pool *pgxpool.Pool, dto *models.CommentsDto, opts ...SelectOption[models.CommentsDto]) (builder.SelectFinalizeQuery[models.CommentsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = comments.Into.AllColumns()
	}
	return comments.NewQuery(pool).
		Select(columns...).
		Where(comments.Id().Eq(Val(*dto.Id))), nil
}

// CommentsDtoSelectMany selects comments DTOs
func CommentsDtoSelectMany(pool *pgxpool.Pool, opts ...SelectOption[models.CommentsDto]) (builder.SelectJoinQuery[models.CommentsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = comments.Into.AllColumns()
	}
	return comments.NewQuery(pool).
		Select(columns...), nil
}

// CommentsDtoDeleteOne deletes a single comments DTO by primary key
func CommentsDtoDeleteOne(pool *pgxpool.Pool, dto *models.CommentsDto) builder.DeleteFinalizeQuery[models.CommentsDto] {
	return comments.NewQuery(pool).
		Delete().
		Where(comments.Id().Eq(Val(*dto.Id)))
}

// CommentsDtoDeleteMany deletes comments DTOs by primary key
func CommentsDtoDeleteMany(pool *pgxpool.Pool, dtos models.CommentsDtos) builder.DeleteFinalizeQuery[models.CommentsDto] {
	v := comments.As("v", primaryKeyCommentsDtoDeleteValueColumns...)
	return comments.NewQuery(pool).
		Delete().
		Using(Values(CommentsDtoDeleteValuesProvider{dtos: dtos}).As(v.Alias)).
		Where(comments.Id().Eq(v.Id()))
}

var primaryKeyCommentsDtoDeleteValueColumns = []ast.NamedExpression{
	comments.Id(),
}

type CommentsDtoDeleteValuesProvider struct {
	dtos models.CommentsDtos
}

func (p CommentsDtoDeleteValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p CommentsDtoDeleteValuesProvider) ColumnCount() int {
	return 1
}

func (p CommentsDtoDeleteValuesProvider) Value(row int, column int) (any, error) {
	switch column {
	case 0:
		return *p.dtos[row].Id, nil
	default:
		return nil, fmt.Errorf("VALUES column %d has no value", column)
	}
}

func (p CommentsDtoDeleteValuesProvider) ColumnSQLType(column int) (string, error) {
	switch column {
	case 0:
		return "integer", nil
	default:
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
}

// CommentsDtoInsertOne inserts a single comments DTO
func CommentsDtoInsertOne(pool *pgxpool.Pool, dto *models.CommentsDto) builder.InsertFinalizeQuery[models.CommentsDto] {
	return comments.NewQuery(pool).
		Insert(
			comments.PostId(),
			comments.UserId(),
			comments.Content(),
		).
		Returning(comments.Id()).
		Values(dto)
}

// CommentsDtoInsertMany inserts multiple comments DTOs
func CommentsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.CommentsDto) builder.InsertFinalizeQuery[models.CommentsDto] {
	return comments.NewQuery(pool).
		Insert(
			comments.PostId(),
			comments.UserId(),
			comments.Content(),
		).
		Returning(comments.Id()).
		Values(dtos...)
}

type CommentsDtoUpdateValueExtractor func(*models.CommentsDto) any

var defaultCommentsDtoUpdateValueColumns = []ast.NamedExpression{
	comments.Content(),
	comments.CreatedAt(),
	comments.Id(),
	comments.IsApproved(),
	comments.PostId(),
	comments.TestDate(),
	comments.UserId(),
}

var defaultCommentsDtoUpdateValueExtractors = []CommentsDtoUpdateValueExtractor{
	func(dto *models.CommentsDto) any { return dto.Content },
	func(dto *models.CommentsDto) any { return dto.CreatedAt },
	func(dto *models.CommentsDto) any { return *dto.Id },
	func(dto *models.CommentsDto) any { return dto.IsApproved },
	func(dto *models.CommentsDto) any { return dto.PostId },
	func(dto *models.CommentsDto) any { return dto.TestDate },
	func(dto *models.CommentsDto) any { return dto.UserId },
}

var defaultCommentsDtoUpdateValueCastTypes = []string{
	"text",
	"timestamp",
	"integer",
	"boolean",
	"integer",
	"date",
	"uuid",
}

var defaultCommentsDtoUpdateColumnIndexes = []int{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

var primaryKeyCommentsDtoUpdateValueColumns = []ast.NamedExpression{
	comments.Id(),
}

var primaryKeyCommentsDtoUpdateValueExtractors = []CommentsDtoUpdateValueExtractor{
	func(dto *models.CommentsDto) any { return *dto.Id },
}

var primaryKeyCommentsDtoUpdateValueCastTypes = []string{
	"integer",
}

var primaryKeyCommentsDtoUpdateColumnIndexes = []int{
	2,
}

var defaultCommentsDtoUpdateSets = []ast.UpdateSetExpr{
	Set(comments.Content()).To(ast.SetType[string](ast.NewColumnNode("v", "content"))),
	Set(comments.CreatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "created_at"))),
	Set(comments.IsApproved()).To(ast.SetType[bool](ast.NewColumnNode("v", "is_approved"))),
	Set(comments.PostId()).To(ast.SetType[int64](ast.NewColumnNode("v", "post_id"))),
	Set(comments.TestDate()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "test_date"))),
	Set(comments.UserId()).To(ast.SetType[uuid.UUID](ast.NewColumnNode("v", "user_id"))),
}

var defaultCommentsDtoUpdateSetByValueColumn = []ast.UpdateSetExpr{
	Set(comments.Content()).To(ast.SetType[string](ast.NewColumnNode("v", "content"))),
	Set(comments.CreatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "created_at"))),
	nil,
	Set(comments.IsApproved()).To(ast.SetType[bool](ast.NewColumnNode("v", "is_approved"))),
	Set(comments.PostId()).To(ast.SetType[int64](ast.NewColumnNode("v", "post_id"))),
	Set(comments.TestDate()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "test_date"))),
	Set(comments.UserId()).To(ast.SetType[uuid.UUID](ast.NewColumnNode("v", "user_id"))),
}

func CommentsDtoUpdateOne(pool *pgxpool.Pool, dto *models.CommentsDto, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.CommentsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	updateColumns, _, _, _, _, err := selectedCommentsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	return comments.NewQuery(pool).
		Update(SetTo(dto, updateCommentsDtoProjections(updateColumns)...)).
		Where(comments.Id().Eq(Val(*dto.Id))), nil
}

func CommentsDtoUpdateMany(pool *pgxpool.Pool, dtos models.CommentsDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.CommentsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, err := selectedCommentsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	v := comments.As("v", valueColumns...)
	if len(sets) == 0 {
		sets = append(sets, ast.UpdateSetList{})
	}
	var from ast.NamedTableExpression = Values(CommentsDtoUpdateValuesProvider{
		dtos:       dtos,
		extractors: valueExtractors,
		castTypes:  valueCastTypes,
	}).As(v.Alias)
	if options.source == updateBatchSourceUnnest {
		from = Unnest(CommentsDtoUpdateUnnestExpressions(dtos, updateColumnIndexes)...).As(v.Alias)
	}
	return comments.NewQuery(pool).
		Update(sets...).
		From(from).
		Where(comments.Id().Eq(v.Id())), nil
}

func selectedCommentsDtoUpdateExpressions(options updateOptions) ([]ast.NamedExpression, []CommentsDtoUpdateValueExtractor, []string, []int, []ast.UpdateSetExpr, error) {
	if options.columns == nil {
		return defaultCommentsDtoUpdateValueColumns, defaultCommentsDtoUpdateValueExtractors, defaultCommentsDtoUpdateValueCastTypes, defaultCommentsDtoUpdateColumnIndexes, defaultCommentsDtoUpdateSets, nil
	}

	requested := make([]ast.NamedExpression, len(options.columns))
	copy(requested, options.columns)
	for _, column := range requested {
		if column == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("query.UpdateColumns received a nil column")
		}
	}
	slices.SortFunc(requested, func(a, b ast.NamedExpression) int {
		if a.Name() < b.Name() {
			return -1
		}
		if a.Name() > b.Name() {
			return 1
		}
		return 0
	})
	requested = slices.CompactFunc(requested, func(a, b ast.NamedExpression) bool {
		return a.Name() == b.Name()
	})

	valueColumns := make([]ast.NamedExpression, 0, len(requested)+1)
	valueColumns = append(valueColumns, primaryKeyCommentsDtoUpdateValueColumns...)
	valueExtractors := make([]CommentsDtoUpdateValueExtractor, 0, len(requested)+1)
	valueExtractors = append(valueExtractors, primaryKeyCommentsDtoUpdateValueExtractors...)
	valueCastTypes := make([]string, 0, len(requested)+1)
	valueCastTypes = append(valueCastTypes, primaryKeyCommentsDtoUpdateValueCastTypes...)
	updateColumnIndexes := make([]int, 0, len(requested)+1)
	updateColumnIndexes = append(updateColumnIndexes, primaryKeyCommentsDtoUpdateColumnIndexes...)
	sets := make([]ast.UpdateSetExpr, 0, len(requested))

	knownIndex := 0
	for _, column := range requested {
		name := column.Name()
		for knownIndex < len(defaultCommentsDtoUpdateValueColumns) && defaultCommentsDtoUpdateValueColumns[knownIndex].Name() < name {
			knownIndex++
		}
		if knownIndex == len(defaultCommentsDtoUpdateValueColumns) || defaultCommentsDtoUpdateValueColumns[knownIndex].Name() != name {
			return nil, nil, nil, nil, nil, fmt.Errorf("unknown update column %q", name)
		}
		set := defaultCommentsDtoUpdateSetByValueColumn[knownIndex]
		if set == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		valueColumns = append(valueColumns, defaultCommentsDtoUpdateValueColumns[knownIndex])
		valueExtractors = append(valueExtractors, defaultCommentsDtoUpdateValueExtractors[knownIndex])
		valueCastTypes = append(valueCastTypes, defaultCommentsDtoUpdateValueCastTypes[knownIndex])
		updateColumnIndexes = append(updateColumnIndexes, knownIndex)
		sets = append(sets, set)
	}
	return valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, nil
}

func updateCommentsDtoProjections(columns []ast.NamedExpression) []ast.Projection[models.CommentsDto] {
	projections := make([]ast.Projection[models.CommentsDto], 0, len(columns))
	for _, column := range columns {
		switch column.Name() {
		case "post_id":
			projections = append(projections, comments.PostId())
		case "user_id":
			projections = append(projections, comments.UserId())
		case "content":
			projections = append(projections, comments.Content())
		case "is_approved":
			projections = append(projections, comments.IsApproved())
		case "created_at":
			projections = append(projections, comments.CreatedAt())
		case "test_date":
			projections = append(projections, comments.TestDate())
		}
	}
	return projections
}

func defaultCommentsDtoUpdateColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		comments.Content(),
		comments.CreatedAt(),
		comments.IsApproved(),
		comments.PostId(),
		comments.TestDate(),
		comments.UserId(),
	}
}

func CommentsDtoUpdateUnnestExpressions(dtos models.CommentsDtos, columnIndexes []int) []ast.Expression {
	unnestExpressions := make([]ast.Expression, len(columnIndexes))
	for i, columnIndex := range columnIndexes {
		switch columnIndex {
		case 0:
			unnestExpressions[i] = Cast(dtos.Content()).AsTextArray()
		case 1:
			unnestExpressions[i] = Cast(dtos.CreatedAt()).AsTimestampArray()
		case 2:
			unnestExpressions[i] = Cast(dtos.Id()).AsIntegerArray()
		case 3:
			unnestExpressions[i] = Cast(dtos.IsApproved()).AsBooleanArray()
		case 4:
			unnestExpressions[i] = Cast(dtos.PostId()).AsIntegerArray()
		case 5:
			unnestExpressions[i] = Cast(dtos.TestDate()).AsDateArray()
		case 6:
			unnestExpressions[i] = Cast(dtos.UserId()).AsUUIDArray()
		}
	}
	return unnestExpressions
}

type CommentsDtoUpdateValuesProvider struct {
	dtos       models.CommentsDtos
	extractors []CommentsDtoUpdateValueExtractor
	castTypes  []string
}

func (p CommentsDtoUpdateValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p CommentsDtoUpdateValuesProvider) ColumnCount() int {
	return len(p.extractors)
}

func (p CommentsDtoUpdateValuesProvider) Value(row int, column int) (any, error) {
	return p.extractors[column](p.dtos[row]), nil
}

func (p CommentsDtoUpdateValuesProvider) ColumnSQLType(column int) (string, error) {
	if column < 0 || column >= len(p.castTypes) {
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
	return p.castTypes[column], nil
}

// PostTagsDtoSelectOne selects a single post_tags DTO by primary key
func PostTagsDtoSelectOne(pool *pgxpool.Pool, dto *models.PostTagsDto, opts ...SelectOption[models.PostTagsDto]) (builder.SelectFinalizeQuery[models.PostTagsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = post_tags.Into.AllColumns()
	}
	return post_tags.NewQuery(pool).
		Select(columns...).
		Where(post_tags.PostId().Eq(Val(dto.PostId)).And(post_tags.TagId().Eq(Val(dto.TagId)))), nil
}

// PostTagsDtoSelectMany selects post_tags DTOs
func PostTagsDtoSelectMany(pool *pgxpool.Pool, opts ...SelectOption[models.PostTagsDto]) (builder.SelectJoinQuery[models.PostTagsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = post_tags.Into.AllColumns()
	}
	return post_tags.NewQuery(pool).
		Select(columns...), nil
}

// PostTagsDtoDeleteOne deletes a single post_tags DTO by primary key
func PostTagsDtoDeleteOne(pool *pgxpool.Pool, dto *models.PostTagsDto) builder.DeleteFinalizeQuery[models.PostTagsDto] {
	return post_tags.NewQuery(pool).
		Delete().
		Where(post_tags.PostId().Eq(Val(dto.PostId)).And(post_tags.TagId().Eq(Val(dto.TagId))))
}

// PostTagsDtoDeleteMany deletes post_tags DTOs by primary key
func PostTagsDtoDeleteMany(pool *pgxpool.Pool, dtos models.PostTagsDtos) builder.DeleteFinalizeQuery[models.PostTagsDto] {
	v := post_tags.As("v", primaryKeyPostTagsDtoDeleteValueColumns...)
	return post_tags.NewQuery(pool).
		Delete().
		Using(Values(PostTagsDtoDeleteValuesProvider{dtos: dtos}).As(v.Alias)).
		Where(post_tags.PostId().Eq(v.PostId()).And(post_tags.TagId().Eq(v.TagId())))
}

var primaryKeyPostTagsDtoDeleteValueColumns = []ast.NamedExpression{
	post_tags.PostId(),
	post_tags.TagId(),
}

type PostTagsDtoDeleteValuesProvider struct {
	dtos models.PostTagsDtos
}

func (p PostTagsDtoDeleteValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p PostTagsDtoDeleteValuesProvider) ColumnCount() int {
	return 2
}

func (p PostTagsDtoDeleteValuesProvider) Value(row int, column int) (any, error) {
	switch column {
	case 0:
		return p.dtos[row].PostId, nil
	case 1:
		return p.dtos[row].TagId, nil
	default:
		return nil, fmt.Errorf("VALUES column %d has no value", column)
	}
}

func (p PostTagsDtoDeleteValuesProvider) ColumnSQLType(column int) (string, error) {
	switch column {
	case 0:
		return "integer", nil
	case 1:
		return "integer", nil
	default:
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
}

// PostTagsDtoInsertOne inserts a single post_tags DTO
func PostTagsDtoInsertOne(pool *pgxpool.Pool, dto *models.PostTagsDto) builder.InsertFinalizeQuery[models.PostTagsDto] {
	return post_tags.NewQuery(pool).
		Insert(
			post_tags.PostId(),
			post_tags.TagId(),
		).
		Returning(post_tags.PostId(), post_tags.TagId()).
		Values(dto)
}

// PostTagsDtoInsertMany inserts multiple post_tags DTOs
func PostTagsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.PostTagsDto) builder.InsertFinalizeQuery[models.PostTagsDto] {
	return post_tags.NewQuery(pool).
		Insert(
			post_tags.PostId(),
			post_tags.TagId(),
		).
		Returning(post_tags.PostId(), post_tags.TagId()).
		Values(dtos...)
}

// PostsDtoSelectOne selects a single posts DTO by primary key
func PostsDtoSelectOne(pool *pgxpool.Pool, dto *models.PostsDto, opts ...SelectOption[models.PostsDto]) (builder.SelectFinalizeQuery[models.PostsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = posts.Into.AllColumns()
	}
	return posts.NewQuery(pool).
		Select(columns...).
		Where(posts.Id().Eq(Val(*dto.Id))), nil
}

// PostsDtoSelectMany selects posts DTOs
func PostsDtoSelectMany(pool *pgxpool.Pool, opts ...SelectOption[models.PostsDto]) (builder.SelectJoinQuery[models.PostsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = posts.Into.AllColumns()
	}
	return posts.NewQuery(pool).
		Select(columns...), nil
}

// PostsDtoDeleteOne deletes a single posts DTO by primary key
func PostsDtoDeleteOne(pool *pgxpool.Pool, dto *models.PostsDto) builder.DeleteFinalizeQuery[models.PostsDto] {
	return posts.NewQuery(pool).
		Delete().
		Where(posts.Id().Eq(Val(*dto.Id)))
}

// PostsDtoDeleteMany deletes posts DTOs by primary key
func PostsDtoDeleteMany(pool *pgxpool.Pool, dtos models.PostsDtos) builder.DeleteFinalizeQuery[models.PostsDto] {
	v := posts.As("v", primaryKeyPostsDtoDeleteValueColumns...)
	return posts.NewQuery(pool).
		Delete().
		Using(Values(PostsDtoDeleteValuesProvider{dtos: dtos}).As(v.Alias)).
		Where(posts.Id().Eq(v.Id()))
}

var primaryKeyPostsDtoDeleteValueColumns = []ast.NamedExpression{
	posts.Id(),
}

type PostsDtoDeleteValuesProvider struct {
	dtos models.PostsDtos
}

func (p PostsDtoDeleteValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p PostsDtoDeleteValuesProvider) ColumnCount() int {
	return 1
}

func (p PostsDtoDeleteValuesProvider) Value(row int, column int) (any, error) {
	switch column {
	case 0:
		return *p.dtos[row].Id, nil
	default:
		return nil, fmt.Errorf("VALUES column %d has no value", column)
	}
}

func (p PostsDtoDeleteValuesProvider) ColumnSQLType(column int) (string, error) {
	switch column {
	case 0:
		return "integer", nil
	default:
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
}

// PostsDtoInsertOne inserts a single posts DTO
func PostsDtoInsertOne(pool *pgxpool.Pool, dto *models.PostsDto) builder.InsertFinalizeQuery[models.PostsDto] {
	return posts.NewQuery(pool).
		Insert(
			posts.UserId(),
			posts.Title(),
			posts.Content(),
			posts.PublishedAt(),
		).
		Returning(posts.Id()).
		Values(dto)
}

// PostsDtoInsertMany inserts multiple posts DTOs
func PostsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.PostsDto) builder.InsertFinalizeQuery[models.PostsDto] {
	return posts.NewQuery(pool).
		Insert(
			posts.UserId(),
			posts.Title(),
			posts.Content(),
			posts.PublishedAt(),
		).
		Returning(posts.Id()).
		Values(dtos...)
}

type PostsDtoUpdateValueExtractor func(*models.PostsDto) any

var defaultPostsDtoUpdateValueColumns = []ast.NamedExpression{
	posts.Content(),
	posts.CreatedAt(),
	posts.Id(),
	posts.PublishedAt(),
	posts.Status(),
	posts.Title(),
	posts.UpdatedAt(),
	posts.UserId(),
	posts.ViewCount(),
}

var defaultPostsDtoUpdateValueExtractors = []PostsDtoUpdateValueExtractor{
	func(dto *models.PostsDto) any { return dto.Content },
	func(dto *models.PostsDto) any { return dto.CreatedAt },
	func(dto *models.PostsDto) any { return *dto.Id },
	func(dto *models.PostsDto) any { return dto.PublishedAt },
	func(dto *models.PostsDto) any { return dto.Status },
	func(dto *models.PostsDto) any { return dto.Title },
	func(dto *models.PostsDto) any { return dto.UpdatedAt },
	func(dto *models.PostsDto) any { return dto.UserId },
	func(dto *models.PostsDto) any { return dto.ViewCount },
}

var defaultPostsDtoUpdateValueCastTypes = []string{
	"text",
	"timestamp",
	"integer",
	"timestamp",
	"text",
	"text",
	"timestamp",
	"uuid",
	"integer",
}

var defaultPostsDtoUpdateColumnIndexes = []int{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
}

var primaryKeyPostsDtoUpdateValueColumns = []ast.NamedExpression{
	posts.Id(),
}

var primaryKeyPostsDtoUpdateValueExtractors = []PostsDtoUpdateValueExtractor{
	func(dto *models.PostsDto) any { return *dto.Id },
}

var primaryKeyPostsDtoUpdateValueCastTypes = []string{
	"integer",
}

var primaryKeyPostsDtoUpdateColumnIndexes = []int{
	2,
}

var defaultPostsDtoUpdateSets = []ast.UpdateSetExpr{
	Set(posts.Content()).To(ast.SetType[string](ast.NewColumnNode("v", "content"))),
	Set(posts.CreatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "created_at"))),
	Set(posts.PublishedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "published_at"))),
	Set(posts.Status()).To(ast.SetType[string](ast.NewColumnNode("v", "status"))),
	Set(posts.Title()).To(ast.SetType[string](ast.NewColumnNode("v", "title"))),
	Set(posts.UpdatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "updated_at"))),
	Set(posts.UserId()).To(ast.SetType[uuid.UUID](ast.NewColumnNode("v", "user_id"))),
	Set(posts.ViewCount()).To(ast.SetType[int64](ast.NewColumnNode("v", "view_count"))),
}

var defaultPostsDtoUpdateSetByValueColumn = []ast.UpdateSetExpr{
	Set(posts.Content()).To(ast.SetType[string](ast.NewColumnNode("v", "content"))),
	Set(posts.CreatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "created_at"))),
	nil,
	Set(posts.PublishedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "published_at"))),
	Set(posts.Status()).To(ast.SetType[string](ast.NewColumnNode("v", "status"))),
	Set(posts.Title()).To(ast.SetType[string](ast.NewColumnNode("v", "title"))),
	Set(posts.UpdatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "updated_at"))),
	Set(posts.UserId()).To(ast.SetType[uuid.UUID](ast.NewColumnNode("v", "user_id"))),
	Set(posts.ViewCount()).To(ast.SetType[int64](ast.NewColumnNode("v", "view_count"))),
}

func PostsDtoUpdateOne(pool *pgxpool.Pool, dto *models.PostsDto, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.PostsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	updateColumns, _, _, _, _, err := selectedPostsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	return posts.NewQuery(pool).
		Update(SetTo(dto, updatePostsDtoProjections(updateColumns)...)).
		Where(posts.Id().Eq(Val(*dto.Id))), nil
}

func PostsDtoUpdateMany(pool *pgxpool.Pool, dtos models.PostsDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.PostsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, err := selectedPostsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	v := posts.As("v", valueColumns...)
	if len(sets) == 0 {
		sets = append(sets, ast.UpdateSetList{})
	}
	var from ast.NamedTableExpression = Values(PostsDtoUpdateValuesProvider{
		dtos:       dtos,
		extractors: valueExtractors,
		castTypes:  valueCastTypes,
	}).As(v.Alias)
	if options.source == updateBatchSourceUnnest {
		from = Unnest(PostsDtoUpdateUnnestExpressions(dtos, updateColumnIndexes)...).As(v.Alias)
	}
	return posts.NewQuery(pool).
		Update(sets...).
		From(from).
		Where(posts.Id().Eq(v.Id())), nil
}

func selectedPostsDtoUpdateExpressions(options updateOptions) ([]ast.NamedExpression, []PostsDtoUpdateValueExtractor, []string, []int, []ast.UpdateSetExpr, error) {
	if options.columns == nil {
		return defaultPostsDtoUpdateValueColumns, defaultPostsDtoUpdateValueExtractors, defaultPostsDtoUpdateValueCastTypes, defaultPostsDtoUpdateColumnIndexes, defaultPostsDtoUpdateSets, nil
	}

	requested := make([]ast.NamedExpression, len(options.columns))
	copy(requested, options.columns)
	for _, column := range requested {
		if column == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("query.UpdateColumns received a nil column")
		}
	}
	slices.SortFunc(requested, func(a, b ast.NamedExpression) int {
		if a.Name() < b.Name() {
			return -1
		}
		if a.Name() > b.Name() {
			return 1
		}
		return 0
	})
	requested = slices.CompactFunc(requested, func(a, b ast.NamedExpression) bool {
		return a.Name() == b.Name()
	})

	valueColumns := make([]ast.NamedExpression, 0, len(requested)+1)
	valueColumns = append(valueColumns, primaryKeyPostsDtoUpdateValueColumns...)
	valueExtractors := make([]PostsDtoUpdateValueExtractor, 0, len(requested)+1)
	valueExtractors = append(valueExtractors, primaryKeyPostsDtoUpdateValueExtractors...)
	valueCastTypes := make([]string, 0, len(requested)+1)
	valueCastTypes = append(valueCastTypes, primaryKeyPostsDtoUpdateValueCastTypes...)
	updateColumnIndexes := make([]int, 0, len(requested)+1)
	updateColumnIndexes = append(updateColumnIndexes, primaryKeyPostsDtoUpdateColumnIndexes...)
	sets := make([]ast.UpdateSetExpr, 0, len(requested))

	knownIndex := 0
	for _, column := range requested {
		name := column.Name()
		for knownIndex < len(defaultPostsDtoUpdateValueColumns) && defaultPostsDtoUpdateValueColumns[knownIndex].Name() < name {
			knownIndex++
		}
		if knownIndex == len(defaultPostsDtoUpdateValueColumns) || defaultPostsDtoUpdateValueColumns[knownIndex].Name() != name {
			return nil, nil, nil, nil, nil, fmt.Errorf("unknown update column %q", name)
		}
		set := defaultPostsDtoUpdateSetByValueColumn[knownIndex]
		if set == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		valueColumns = append(valueColumns, defaultPostsDtoUpdateValueColumns[knownIndex])
		valueExtractors = append(valueExtractors, defaultPostsDtoUpdateValueExtractors[knownIndex])
		valueCastTypes = append(valueCastTypes, defaultPostsDtoUpdateValueCastTypes[knownIndex])
		updateColumnIndexes = append(updateColumnIndexes, knownIndex)
		sets = append(sets, set)
	}
	return valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, nil
}

func updatePostsDtoProjections(columns []ast.NamedExpression) []ast.Projection[models.PostsDto] {
	projections := make([]ast.Projection[models.PostsDto], 0, len(columns))
	for _, column := range columns {
		switch column.Name() {
		case "user_id":
			projections = append(projections, posts.UserId())
		case "title":
			projections = append(projections, posts.Title())
		case "content":
			projections = append(projections, posts.Content())
		case "status":
			projections = append(projections, posts.Status())
		case "published_at":
			projections = append(projections, posts.PublishedAt())
		case "view_count":
			projections = append(projections, posts.ViewCount())
		case "created_at":
			projections = append(projections, posts.CreatedAt())
		case "updated_at":
			projections = append(projections, posts.UpdatedAt())
		}
	}
	return projections
}

func defaultPostsDtoUpdateColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		posts.Content(),
		posts.CreatedAt(),
		posts.PublishedAt(),
		posts.Status(),
		posts.Title(),
		posts.UpdatedAt(),
		posts.UserId(),
		posts.ViewCount(),
	}
}

func PostsDtoUpdateUnnestExpressions(dtos models.PostsDtos, columnIndexes []int) []ast.Expression {
	unnestExpressions := make([]ast.Expression, len(columnIndexes))
	for i, columnIndex := range columnIndexes {
		switch columnIndex {
		case 0:
			unnestExpressions[i] = Cast(dtos.Content()).AsTextArray()
		case 1:
			unnestExpressions[i] = Cast(dtos.CreatedAt()).AsTimestampArray()
		case 2:
			unnestExpressions[i] = Cast(dtos.Id()).AsIntegerArray()
		case 3:
			unnestExpressions[i] = Cast(dtos.PublishedAt()).AsTimestampArray()
		case 4:
			unnestExpressions[i] = Cast(dtos.Status()).AsTextArray()
		case 5:
			unnestExpressions[i] = Cast(dtos.Title()).AsTextArray()
		case 6:
			unnestExpressions[i] = Cast(dtos.UpdatedAt()).AsTimestampArray()
		case 7:
			unnestExpressions[i] = Cast(dtos.UserId()).AsUUIDArray()
		case 8:
			unnestExpressions[i] = Cast(dtos.ViewCount()).AsIntegerArray()
		}
	}
	return unnestExpressions
}

type PostsDtoUpdateValuesProvider struct {
	dtos       models.PostsDtos
	extractors []PostsDtoUpdateValueExtractor
	castTypes  []string
}

func (p PostsDtoUpdateValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p PostsDtoUpdateValuesProvider) ColumnCount() int {
	return len(p.extractors)
}

func (p PostsDtoUpdateValuesProvider) Value(row int, column int) (any, error) {
	return p.extractors[column](p.dtos[row]), nil
}

func (p PostsDtoUpdateValuesProvider) ColumnSQLType(column int) (string, error) {
	if column < 0 || column >= len(p.castTypes) {
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
	return p.castTypes[column], nil
}

// TagsDtoSelectOne selects a single tags DTO by primary key
func TagsDtoSelectOne(pool *pgxpool.Pool, dto *models.TagsDto, opts ...SelectOption[models.TagsDto]) (builder.SelectFinalizeQuery[models.TagsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = tags.Into.AllColumns()
	}
	return tags.NewQuery(pool).
		Select(columns...).
		Where(tags.Id().Eq(Val(*dto.Id))), nil
}

// TagsDtoSelectMany selects tags DTOs
func TagsDtoSelectMany(pool *pgxpool.Pool, opts ...SelectOption[models.TagsDto]) (builder.SelectJoinQuery[models.TagsDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = tags.Into.AllColumns()
	}
	return tags.NewQuery(pool).
		Select(columns...), nil
}

// TagsDtoDeleteOne deletes a single tags DTO by primary key
func TagsDtoDeleteOne(pool *pgxpool.Pool, dto *models.TagsDto) builder.DeleteFinalizeQuery[models.TagsDto] {
	return tags.NewQuery(pool).
		Delete().
		Where(tags.Id().Eq(Val(*dto.Id)))
}

// TagsDtoDeleteMany deletes tags DTOs by primary key
func TagsDtoDeleteMany(pool *pgxpool.Pool, dtos models.TagsDtos) builder.DeleteFinalizeQuery[models.TagsDto] {
	v := tags.As("v", primaryKeyTagsDtoDeleteValueColumns...)
	return tags.NewQuery(pool).
		Delete().
		Using(Values(TagsDtoDeleteValuesProvider{dtos: dtos}).As(v.Alias)).
		Where(tags.Id().Eq(v.Id()))
}

var primaryKeyTagsDtoDeleteValueColumns = []ast.NamedExpression{
	tags.Id(),
}

type TagsDtoDeleteValuesProvider struct {
	dtos models.TagsDtos
}

func (p TagsDtoDeleteValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p TagsDtoDeleteValuesProvider) ColumnCount() int {
	return 1
}

func (p TagsDtoDeleteValuesProvider) Value(row int, column int) (any, error) {
	switch column {
	case 0:
		return *p.dtos[row].Id, nil
	default:
		return nil, fmt.Errorf("VALUES column %d has no value", column)
	}
}

func (p TagsDtoDeleteValuesProvider) ColumnSQLType(column int) (string, error) {
	switch column {
	case 0:
		return "integer", nil
	default:
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
}

// TagsDtoInsertOne inserts a single tags DTO
func TagsDtoInsertOne(pool *pgxpool.Pool, dto *models.TagsDto) builder.InsertFinalizeQuery[models.TagsDto] {
	return tags.NewQuery(pool).
		Insert(
			tags.Name(),
			tags.Slug(),
		).
		Returning(tags.Id()).
		Values(dto)
}

// TagsDtoInsertMany inserts multiple tags DTOs
func TagsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.TagsDto) builder.InsertFinalizeQuery[models.TagsDto] {
	return tags.NewQuery(pool).
		Insert(
			tags.Name(),
			tags.Slug(),
		).
		Returning(tags.Id()).
		Values(dtos...)
}

type TagsDtoUpdateValueExtractor func(*models.TagsDto) any

var defaultTagsDtoUpdateValueColumns = []ast.NamedExpression{
	tags.Id(),
	tags.Name(),
	tags.Slug(),
}

var defaultTagsDtoUpdateValueExtractors = []TagsDtoUpdateValueExtractor{
	func(dto *models.TagsDto) any { return *dto.Id },
	func(dto *models.TagsDto) any { return dto.Name },
	func(dto *models.TagsDto) any { return dto.Slug },
}

var defaultTagsDtoUpdateValueCastTypes = []string{
	"integer",
	"text",
	"text",
}

var defaultTagsDtoUpdateColumnIndexes = []int{
	0,
	1,
	2,
}

var primaryKeyTagsDtoUpdateValueColumns = []ast.NamedExpression{
	tags.Id(),
}

var primaryKeyTagsDtoUpdateValueExtractors = []TagsDtoUpdateValueExtractor{
	func(dto *models.TagsDto) any { return *dto.Id },
}

var primaryKeyTagsDtoUpdateValueCastTypes = []string{
	"integer",
}

var primaryKeyTagsDtoUpdateColumnIndexes = []int{
	0,
}

var defaultTagsDtoUpdateSets = []ast.UpdateSetExpr{
	Set(tags.Name()).To(ast.SetType[string](ast.NewColumnNode("v", "name"))),
	Set(tags.Slug()).To(ast.SetType[string](ast.NewColumnNode("v", "slug"))),
}

var defaultTagsDtoUpdateSetByValueColumn = []ast.UpdateSetExpr{
	nil,
	Set(tags.Name()).To(ast.SetType[string](ast.NewColumnNode("v", "name"))),
	Set(tags.Slug()).To(ast.SetType[string](ast.NewColumnNode("v", "slug"))),
}

func TagsDtoUpdateOne(pool *pgxpool.Pool, dto *models.TagsDto, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.TagsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	updateColumns, _, _, _, _, err := selectedTagsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	return tags.NewQuery(pool).
		Update(SetTo(dto, updateTagsDtoProjections(updateColumns)...)).
		Where(tags.Id().Eq(Val(*dto.Id))), nil
}

func TagsDtoUpdateMany(pool *pgxpool.Pool, dtos models.TagsDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.TagsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, err := selectedTagsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	v := tags.As("v", valueColumns...)
	if len(sets) == 0 {
		sets = append(sets, ast.UpdateSetList{})
	}
	var from ast.NamedTableExpression = Values(TagsDtoUpdateValuesProvider{
		dtos:       dtos,
		extractors: valueExtractors,
		castTypes:  valueCastTypes,
	}).As(v.Alias)
	if options.source == updateBatchSourceUnnest {
		from = Unnest(TagsDtoUpdateUnnestExpressions(dtos, updateColumnIndexes)...).As(v.Alias)
	}
	return tags.NewQuery(pool).
		Update(sets...).
		From(from).
		Where(tags.Id().Eq(v.Id())), nil
}

func selectedTagsDtoUpdateExpressions(options updateOptions) ([]ast.NamedExpression, []TagsDtoUpdateValueExtractor, []string, []int, []ast.UpdateSetExpr, error) {
	if options.columns == nil {
		return defaultTagsDtoUpdateValueColumns, defaultTagsDtoUpdateValueExtractors, defaultTagsDtoUpdateValueCastTypes, defaultTagsDtoUpdateColumnIndexes, defaultTagsDtoUpdateSets, nil
	}

	requested := make([]ast.NamedExpression, len(options.columns))
	copy(requested, options.columns)
	for _, column := range requested {
		if column == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("query.UpdateColumns received a nil column")
		}
	}
	slices.SortFunc(requested, func(a, b ast.NamedExpression) int {
		if a.Name() < b.Name() {
			return -1
		}
		if a.Name() > b.Name() {
			return 1
		}
		return 0
	})
	requested = slices.CompactFunc(requested, func(a, b ast.NamedExpression) bool {
		return a.Name() == b.Name()
	})

	valueColumns := make([]ast.NamedExpression, 0, len(requested)+1)
	valueColumns = append(valueColumns, primaryKeyTagsDtoUpdateValueColumns...)
	valueExtractors := make([]TagsDtoUpdateValueExtractor, 0, len(requested)+1)
	valueExtractors = append(valueExtractors, primaryKeyTagsDtoUpdateValueExtractors...)
	valueCastTypes := make([]string, 0, len(requested)+1)
	valueCastTypes = append(valueCastTypes, primaryKeyTagsDtoUpdateValueCastTypes...)
	updateColumnIndexes := make([]int, 0, len(requested)+1)
	updateColumnIndexes = append(updateColumnIndexes, primaryKeyTagsDtoUpdateColumnIndexes...)
	sets := make([]ast.UpdateSetExpr, 0, len(requested))

	knownIndex := 0
	for _, column := range requested {
		name := column.Name()
		for knownIndex < len(defaultTagsDtoUpdateValueColumns) && defaultTagsDtoUpdateValueColumns[knownIndex].Name() < name {
			knownIndex++
		}
		if knownIndex == len(defaultTagsDtoUpdateValueColumns) || defaultTagsDtoUpdateValueColumns[knownIndex].Name() != name {
			return nil, nil, nil, nil, nil, fmt.Errorf("unknown update column %q", name)
		}
		set := defaultTagsDtoUpdateSetByValueColumn[knownIndex]
		if set == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		valueColumns = append(valueColumns, defaultTagsDtoUpdateValueColumns[knownIndex])
		valueExtractors = append(valueExtractors, defaultTagsDtoUpdateValueExtractors[knownIndex])
		valueCastTypes = append(valueCastTypes, defaultTagsDtoUpdateValueCastTypes[knownIndex])
		updateColumnIndexes = append(updateColumnIndexes, knownIndex)
		sets = append(sets, set)
	}
	return valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, nil
}

func updateTagsDtoProjections(columns []ast.NamedExpression) []ast.Projection[models.TagsDto] {
	projections := make([]ast.Projection[models.TagsDto], 0, len(columns))
	for _, column := range columns {
		switch column.Name() {
		case "name":
			projections = append(projections, tags.Name())
		case "slug":
			projections = append(projections, tags.Slug())
		}
	}
	return projections
}

func defaultTagsDtoUpdateColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		tags.Name(),
		tags.Slug(),
	}
}

func TagsDtoUpdateUnnestExpressions(dtos models.TagsDtos, columnIndexes []int) []ast.Expression {
	unnestExpressions := make([]ast.Expression, len(columnIndexes))
	for i, columnIndex := range columnIndexes {
		switch columnIndex {
		case 0:
			unnestExpressions[i] = Cast(dtos.Id()).AsIntegerArray()
		case 1:
			unnestExpressions[i] = Cast(dtos.Name()).AsTextArray()
		case 2:
			unnestExpressions[i] = Cast(dtos.Slug()).AsTextArray()
		}
	}
	return unnestExpressions
}

type TagsDtoUpdateValuesProvider struct {
	dtos       models.TagsDtos
	extractors []TagsDtoUpdateValueExtractor
	castTypes  []string
}

func (p TagsDtoUpdateValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p TagsDtoUpdateValuesProvider) ColumnCount() int {
	return len(p.extractors)
}

func (p TagsDtoUpdateValuesProvider) Value(row int, column int) (any, error) {
	return p.extractors[column](p.dtos[row]), nil
}

func (p TagsDtoUpdateValuesProvider) ColumnSQLType(column int) (string, error) {
	if column < 0 || column >= len(p.castTypes) {
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
	return p.castTypes[column], nil
}

// UsersDtoSelectOne selects a single users DTO by primary key
func UsersDtoSelectOne(pool *pgxpool.Pool, dto *models.UsersDto, opts ...SelectOption[models.UsersDto]) (builder.SelectFinalizeQuery[models.UsersDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = users.Into.AllColumns()
	}
	return users.NewQuery(pool).
		Select(columns...).
		Where(users.Id().Eq(Val(*dto.Id))), nil
}

// UsersDtoSelectMany selects users DTOs
func UsersDtoSelectMany(pool *pgxpool.Pool, opts ...SelectOption[models.UsersDto]) (builder.SelectJoinQuery[models.UsersDto], error) {
	options, err := newSelectOptions(opts...)
	if err != nil {
		return nil, err
	}
	columns := options.columns
	if columns == nil {
		columns = users.Into.AllColumns()
	}
	return users.NewQuery(pool).
		Select(columns...), nil
}

// UsersDtoDeleteOne deletes a single users DTO by primary key
func UsersDtoDeleteOne(pool *pgxpool.Pool, dto *models.UsersDto) builder.DeleteFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Delete().
		Where(users.Id().Eq(Val(*dto.Id)))
}

// UsersDtoDeleteMany deletes users DTOs by primary key
func UsersDtoDeleteMany(pool *pgxpool.Pool, dtos models.UsersDtos) builder.DeleteFinalizeQuery[models.UsersDto] {
	v := users.As("v", primaryKeyUsersDtoDeleteValueColumns...)
	return users.NewQuery(pool).
		Delete().
		Using(Values(UsersDtoDeleteValuesProvider{dtos: dtos}).As(v.Alias)).
		Where(users.Id().Eq(v.Id()))
}

var primaryKeyUsersDtoDeleteValueColumns = []ast.NamedExpression{
	users.Id(),
}

type UsersDtoDeleteValuesProvider struct {
	dtos models.UsersDtos
}

func (p UsersDtoDeleteValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p UsersDtoDeleteValuesProvider) ColumnCount() int {
	return 1
}

func (p UsersDtoDeleteValuesProvider) Value(row int, column int) (any, error) {
	switch column {
	case 0:
		return *p.dtos[row].Id, nil
	default:
		return nil, fmt.Errorf("VALUES column %d has no value", column)
	}
}

func (p UsersDtoDeleteValuesProvider) ColumnSQLType(column int) (string, error) {
	switch column {
	case 0:
		return "uuid", nil
	default:
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
}

// UsersDtoInsertOne inserts a single users DTO
func UsersDtoInsertOne(pool *pgxpool.Pool, dto *models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
		).
		Returning(users.Id()).
		Values(dto)
}

// UsersDtoInsertMany inserts multiple users DTOs
func UsersDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Insert(
			users.Username(),
			users.Email(),
			users.FullName(),
		).
		Returning(users.Id()).
		Values(dtos...)
}

type UsersDtoUpdateValueExtractor func(*models.UsersDto) any

var defaultUsersDtoUpdateValueColumns = []ast.NamedExpression{
	users.CreatedAt(),
	users.Email(),
	users.FullName(),
	users.Id(),
	users.IsActive(),
	users.UpdatedAt(),
	users.Username(),
}

var defaultUsersDtoUpdateValueExtractors = []UsersDtoUpdateValueExtractor{
	func(dto *models.UsersDto) any { return dto.CreatedAt },
	func(dto *models.UsersDto) any { return dto.Email },
	func(dto *models.UsersDto) any { return dto.FullName },
	func(dto *models.UsersDto) any { return *dto.Id },
	func(dto *models.UsersDto) any { return dto.IsActive },
	func(dto *models.UsersDto) any { return dto.UpdatedAt },
	func(dto *models.UsersDto) any { return dto.Username },
}

var defaultUsersDtoUpdateValueCastTypes = []string{
	"timestamp",
	"text",
	"text",
	"uuid",
	"boolean",
	"timestamp",
	"text",
}

var defaultUsersDtoUpdateColumnIndexes = []int{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
}

var primaryKeyUsersDtoUpdateValueColumns = []ast.NamedExpression{
	users.Id(),
}

var primaryKeyUsersDtoUpdateValueExtractors = []UsersDtoUpdateValueExtractor{
	func(dto *models.UsersDto) any { return *dto.Id },
}

var primaryKeyUsersDtoUpdateValueCastTypes = []string{
	"uuid",
}

var primaryKeyUsersDtoUpdateColumnIndexes = []int{
	3,
}

var defaultUsersDtoUpdateSets = []ast.UpdateSetExpr{
	Set(users.CreatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "created_at"))),
	Set(users.Email()).To(ast.SetType[string](ast.NewColumnNode("v", "email"))),
	Set(users.FullName()).To(ast.SetType[string](ast.NewColumnNode("v", "full_name"))),
	Set(users.IsActive()).To(ast.SetType[bool](ast.NewColumnNode("v", "is_active"))),
	Set(users.UpdatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "updated_at"))),
	Set(users.Username()).To(ast.SetType[string](ast.NewColumnNode("v", "username"))),
}

var defaultUsersDtoUpdateSetByValueColumn = []ast.UpdateSetExpr{
	Set(users.CreatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "created_at"))),
	Set(users.Email()).To(ast.SetType[string](ast.NewColumnNode("v", "email"))),
	Set(users.FullName()).To(ast.SetType[string](ast.NewColumnNode("v", "full_name"))),
	nil,
	Set(users.IsActive()).To(ast.SetType[bool](ast.NewColumnNode("v", "is_active"))),
	Set(users.UpdatedAt()).To(ast.SetType[time.Time](ast.NewColumnNode("v", "updated_at"))),
	Set(users.Username()).To(ast.SetType[string](ast.NewColumnNode("v", "username"))),
}

func UsersDtoUpdateOne(pool *pgxpool.Pool, dto *models.UsersDto, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.UsersDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	updateColumns, _, _, _, _, err := selectedUsersDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	return users.NewQuery(pool).
		Update(SetTo(dto, updateUsersDtoProjections(updateColumns)...)).
		Where(users.Id().Eq(Val(*dto.Id))), nil
}

func UsersDtoUpdateMany(pool *pgxpool.Pool, dtos models.UsersDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.UsersDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, err := selectedUsersDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	v := users.As("v", valueColumns...)
	if len(sets) == 0 {
		sets = append(sets, ast.UpdateSetList{})
	}
	var from ast.NamedTableExpression = Values(UsersDtoUpdateValuesProvider{
		dtos:       dtos,
		extractors: valueExtractors,
		castTypes:  valueCastTypes,
	}).As(v.Alias)
	if options.source == updateBatchSourceUnnest {
		from = Unnest(UsersDtoUpdateUnnestExpressions(dtos, updateColumnIndexes)...).As(v.Alias)
	}
	return users.NewQuery(pool).
		Update(sets...).
		From(from).
		Where(users.Id().Eq(v.Id())), nil
}

func selectedUsersDtoUpdateExpressions(options updateOptions) ([]ast.NamedExpression, []UsersDtoUpdateValueExtractor, []string, []int, []ast.UpdateSetExpr, error) {
	if options.columns == nil {
		return defaultUsersDtoUpdateValueColumns, defaultUsersDtoUpdateValueExtractors, defaultUsersDtoUpdateValueCastTypes, defaultUsersDtoUpdateColumnIndexes, defaultUsersDtoUpdateSets, nil
	}

	requested := make([]ast.NamedExpression, len(options.columns))
	copy(requested, options.columns)
	for _, column := range requested {
		if column == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("query.UpdateColumns received a nil column")
		}
	}
	slices.SortFunc(requested, func(a, b ast.NamedExpression) int {
		if a.Name() < b.Name() {
			return -1
		}
		if a.Name() > b.Name() {
			return 1
		}
		return 0
	})
	requested = slices.CompactFunc(requested, func(a, b ast.NamedExpression) bool {
		return a.Name() == b.Name()
	})

	valueColumns := make([]ast.NamedExpression, 0, len(requested)+1)
	valueColumns = append(valueColumns, primaryKeyUsersDtoUpdateValueColumns...)
	valueExtractors := make([]UsersDtoUpdateValueExtractor, 0, len(requested)+1)
	valueExtractors = append(valueExtractors, primaryKeyUsersDtoUpdateValueExtractors...)
	valueCastTypes := make([]string, 0, len(requested)+1)
	valueCastTypes = append(valueCastTypes, primaryKeyUsersDtoUpdateValueCastTypes...)
	updateColumnIndexes := make([]int, 0, len(requested)+1)
	updateColumnIndexes = append(updateColumnIndexes, primaryKeyUsersDtoUpdateColumnIndexes...)
	sets := make([]ast.UpdateSetExpr, 0, len(requested))

	knownIndex := 0
	for _, column := range requested {
		name := column.Name()
		for knownIndex < len(defaultUsersDtoUpdateValueColumns) && defaultUsersDtoUpdateValueColumns[knownIndex].Name() < name {
			knownIndex++
		}
		if knownIndex == len(defaultUsersDtoUpdateValueColumns) || defaultUsersDtoUpdateValueColumns[knownIndex].Name() != name {
			return nil, nil, nil, nil, nil, fmt.Errorf("unknown update column %q", name)
		}
		set := defaultUsersDtoUpdateSetByValueColumn[knownIndex]
		if set == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		valueColumns = append(valueColumns, defaultUsersDtoUpdateValueColumns[knownIndex])
		valueExtractors = append(valueExtractors, defaultUsersDtoUpdateValueExtractors[knownIndex])
		valueCastTypes = append(valueCastTypes, defaultUsersDtoUpdateValueCastTypes[knownIndex])
		updateColumnIndexes = append(updateColumnIndexes, knownIndex)
		sets = append(sets, set)
	}
	return valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, nil
}

func updateUsersDtoProjections(columns []ast.NamedExpression) []ast.Projection[models.UsersDto] {
	projections := make([]ast.Projection[models.UsersDto], 0, len(columns))
	for _, column := range columns {
		switch column.Name() {
		case "username":
			projections = append(projections, users.Username())
		case "email":
			projections = append(projections, users.Email())
		case "full_name":
			projections = append(projections, users.FullName())
		case "created_at":
			projections = append(projections, users.CreatedAt())
		case "updated_at":
			projections = append(projections, users.UpdatedAt())
		case "is_active":
			projections = append(projections, users.IsActive())
		}
	}
	return projections
}

func defaultUsersDtoUpdateColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		users.CreatedAt(),
		users.Email(),
		users.FullName(),
		users.IsActive(),
		users.UpdatedAt(),
		users.Username(),
	}
}

func UsersDtoUpdateUnnestExpressions(dtos models.UsersDtos, columnIndexes []int) []ast.Expression {
	unnestExpressions := make([]ast.Expression, len(columnIndexes))
	for i, columnIndex := range columnIndexes {
		switch columnIndex {
		case 0:
			unnestExpressions[i] = Cast(dtos.CreatedAt()).AsTimestampArray()
		case 1:
			unnestExpressions[i] = Cast(dtos.Email()).AsTextArray()
		case 2:
			unnestExpressions[i] = Cast(dtos.FullName()).AsTextArray()
		case 3:
			unnestExpressions[i] = Cast(dtos.Id()).AsUUIDArray()
		case 4:
			unnestExpressions[i] = Cast(dtos.IsActive()).AsBooleanArray()
		case 5:
			unnestExpressions[i] = Cast(dtos.UpdatedAt()).AsTimestampArray()
		case 6:
			unnestExpressions[i] = Cast(dtos.Username()).AsTextArray()
		}
	}
	return unnestExpressions
}

type UsersDtoUpdateValuesProvider struct {
	dtos       models.UsersDtos
	extractors []UsersDtoUpdateValueExtractor
	castTypes  []string
}

func (p UsersDtoUpdateValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p UsersDtoUpdateValuesProvider) ColumnCount() int {
	return len(p.extractors)
}

func (p UsersDtoUpdateValuesProvider) Value(row int, column int) (any, error) {
	return p.extractors[column](p.dtos[row]), nil
}

func (p UsersDtoUpdateValuesProvider) ColumnSQLType(column int) (string, error) {
	if column < 0 || column >= len(p.castTypes) {
		return "", fmt.Errorf("VALUES column %d has no SQL type", column)
	}
	return p.castTypes[column], nil
}

func DeleteUser(pool *pgxpool.Pool, userId uuid.UUID) builder.DeleteFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Delete().
		Where(users.Id().Eq(Val(userId)))
}
