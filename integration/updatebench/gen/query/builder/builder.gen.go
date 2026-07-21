package builder

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
)

// KnownTableBuilder is the builder for known tables
type KnownTableBuilder[T any] struct {
	pool  *pgxpool.Pool
	tx    pgx.Tx
	table ast.NamedTableExpression
	with  []*ast.CTE
}

func NewKnownTableBuilder[T any](pool *pgxpool.Pool, table ast.NamedTableExpression) *KnownTableBuilder[T] {
	return &KnownTableBuilder[T]{
		pool:  pool,
		table: table,
	}
}

type StatementBuilder struct {
	table ast.NamedTableExpression
	with  []*ast.CTE
}

func NewStatementBuilder(table ast.NamedTableExpression) *StatementBuilder {
	return &StatementBuilder{table: table}
}

func (b *StatementBuilder) With(ctes ...*ast.CTE) StatementStartQuery {
	b.with = append(b.with, ctes...)
	return b
}

func (b *StatementBuilder) Select(columns ...ast.NamedExpression) StatementSelectJoinQuery {
	return &StatementSelectBuilder{
		stmt: &ast.SelectStatement{
			With:       b.with,
			SelectList: columns,
			From:       b.table,
		},
	}
}

func (b *StatementBuilder) Update(toSet ...ast.UpdateSetExpr) StatementUpdateFromQuery {
	return &StatementUpdateBuilder{
		stmt: &ast.UpdateStatement{
			With:    b.with,
			Table:   b.table.GetName(),
			SetList: toSet,
		},
	}
}

func (b *StatementBuilder) Delete() StatementDeleteUsingQuery {
	return &StatementDeleteBuilder{
		stmt: &ast.DeleteStatement{
			With:  b.with,
			Table: b.table.GetName(),
		},
	}
}

func (b *KnownTableBuilder[T]) WithTx(tx pgx.Tx) KnownTableStartQuery[T] {
	b.tx = tx
	return b
}

func (b *KnownTableBuilder[T]) With(ctes ...*ast.CTE) KnownTableStartQuery[T] {
	b.with = append(b.with, ctes...)
	return b
}

func (b *KnownTableBuilder[T]) Insert(columns ...ast.NamedExpression) InsertSourceQuery[T] {
	return &InsertBuilder[T]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.InsertStatement{
			With:  b.with,
			Table: b.table.GetName(),
			Into:  columns,
		},
	}
}

func (b *KnownTableBuilder[T]) Select(projections ...ast.Projection[T]) SelectJoinQuery[T] {
	selectList := make([]ast.NamedExpression, 0, len(projections))
	for _, projection := range projections {
		selectList = append(selectList, projection)
	}
	return &SelectBuilder[T]{
		pool:        b.pool,
		tx:          b.tx,
		projections: projections,
		stmt: &ast.SelectStatement{
			With:       b.with,
			SelectList: selectList,
			From:       b.table,
		},
	}
}

func (b *KnownTableBuilder[T]) Update(toSet ...ast.UpdateSetExpr) UpdateFromQuery[T] {
	return &UpdateBuilder[T]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.UpdateStatement{
			With:    b.with,
			Table:   b.table.GetName(),
			SetList: toSet,
		},
	}
}

func (b *KnownTableBuilder[T]) Delete() DeleteUsingQuery[T] {
	return &DeleteBuilder[T]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.DeleteStatement{
			With:  b.with,
			Table: b.table.GetName(),
		},
	}
}

type SelectBuilder[T any] struct {
	pool        *pgxpool.Pool
	tx          pgx.Tx
	stmt        *ast.SelectStatement
	projections []ast.Projection[T]
}

func (b *SelectBuilder[T]) WithTx(tx pgx.Tx) *SelectBuilder[T] {
	b.tx = tx
	return b
}

func (b *SelectBuilder[T]) With(ctes ...*ast.CTE) SelectJoinQuery[T] {
	selectWith(b.stmt, ctes...)
	return b
}

func (b *SelectBuilder[T]) From(table ast.TableExpression) SelectWhereQuery[T] {
	selectFrom(b.stmt, table)
	return b
}

func (b *SelectBuilder[T]) Join(join *ast.JoinExpr) SelectJoinQuery[T] {
	selectJoin(b.stmt, join)
	return b
}

func (b *SelectBuilder[T]) Where(expr ast.OfType[bool]) SelectGroupByQuery[T] {
	selectWhere(b.stmt, expr)
	return b
}

func (b *SelectBuilder[T]) GroupBy(columns ...ast.Expression) SelectHavingQuery[T] {
	selectGroupBy(b.stmt, columns...)
	return b
}

func (b *SelectBuilder[T]) Having(expr ast.OfType[bool]) SelectOrderByQuery[T] {
	selectHaving(b.stmt, expr)
	return b
}

func (b *SelectBuilder[T]) OrderBy(orderBy ...*ast.OrderByItem) SelectPagingQuery[T] {
	selectOrderBy(b.stmt, orderBy...)
	return b
}

func (b *SelectBuilder[T]) Limit(limit int) SelectPagingQuery[T] {
	selectLimit(b.stmt, limit)
	return b
}

func (b *SelectBuilder[T]) Offset(offset int) SelectPagingQuery[T] {
	selectOffset(b.stmt, offset)
	return b
}

func (b *SelectBuilder[T]) ToSql() (string, []any, error) {
	params := make([]any, 0)
	sql, err := ast.RenderWithContext(b.stmt, &params, &ast.QueryContext{PrimaryTable: ast.TableExpressionName(b.stmt.From)})
	return sql, params, err
}

func (b *SelectBuilder[T]) Statement() ast.SqlStatement {
	return b.stmt
}

func (b *SelectBuilder[T]) SelectStatement() *ast.SelectStatement {
	return b.stmt
}

func (b *SelectBuilder[T]) Find(ctx context.Context) ([]T, error) {
	queryCtx := &ast.QueryContext{
		PrimaryTable: ast.TableExpressionName(b.stmt.From),
	}

	args := make([]any, 0)
	query, err := ast.RenderWithContext(b.stmt, &args, queryCtx)
	if err != nil {
		return nil, err
	}

	var rows pgx.Rows
	if b.tx != nil {
		rows, err = b.tx.Query(ctx, query, args...)
	} else {
		rows, err = b.pool.Query(ctx, query, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var t T
		bindings := bindProjections(&t, b.projections)
		if err = rows.Scan(scanDestinations(bindings)...); err != nil {
			return nil, err
		}
		if err = assignProjections(bindings); err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, rows.Err()
}

func (b *SelectBuilder[T]) FindOne(ctx context.Context) (T, error) {
	b.Limit(1)
	results, err := b.Find(ctx)
	if err != nil {
		var t T
		return t, err
	}
	if len(results) == 0 {
		var t T
		return t, pgx.ErrNoRows
	}
	return results[0], nil
}

type StatementSelectBuilder struct {
	stmt *ast.SelectStatement
}

func (b *StatementSelectBuilder) With(ctes ...*ast.CTE) StatementSelectJoinQuery {
	selectWith(b.stmt, ctes...)
	return b
}

func (b *StatementSelectBuilder) From(table ast.TableExpression) StatementSelectWhereQuery {
	selectFrom(b.stmt, table)
	return b
}

func (b *StatementSelectBuilder) Join(join *ast.JoinExpr) StatementSelectJoinQuery {
	selectJoin(b.stmt, join)
	return b
}

func (b *StatementSelectBuilder) Where(expr ast.OfType[bool]) StatementSelectGroupByQuery {
	selectWhere(b.stmt, expr)
	return b
}

func (b *StatementSelectBuilder) GroupBy(columns ...ast.Expression) StatementSelectHavingQuery {
	selectGroupBy(b.stmt, columns...)
	return b
}

func (b *StatementSelectBuilder) Having(expr ast.OfType[bool]) StatementSelectOrderByQuery {
	selectHaving(b.stmt, expr)
	return b
}

func (b *StatementSelectBuilder) OrderBy(orderBy ...*ast.OrderByItem) StatementSelectPagingQuery {
	selectOrderBy(b.stmt, orderBy...)
	return b
}

func (b *StatementSelectBuilder) Limit(limit int) StatementSelectPagingQuery {
	selectLimit(b.stmt, limit)
	return b
}

func (b *StatementSelectBuilder) Offset(offset int) StatementSelectPagingQuery {
	selectOffset(b.stmt, offset)
	return b
}

func (b *StatementSelectBuilder) Statement() ast.SqlStatement {
	return b.stmt
}

func (b *StatementSelectBuilder) SelectStatement() *ast.SelectStatement {
	return b.stmt
}

func selectWith(stmt *ast.SelectStatement, ctes ...*ast.CTE) {
	stmt.With = append(stmt.With, ctes...)
}

func selectFrom(stmt *ast.SelectStatement, table ast.TableExpression) {
	stmt.From = table
}

func selectJoin(stmt *ast.SelectStatement, join *ast.JoinExpr) {
	stmt.From = stmt.From.Join(join)
}

func selectWhere(stmt *ast.SelectStatement, expr ast.OfType[bool]) {
	stmt.Where = expr
}

func selectGroupBy(stmt *ast.SelectStatement, columns ...ast.Expression) {
	stmt.GroupBy = columns
}

func selectHaving(stmt *ast.SelectStatement, expr ast.OfType[bool]) {
	stmt.Having = expr
}

func selectOrderBy(stmt *ast.SelectStatement, orderBy ...*ast.OrderByItem) {
	stmt.OrderBy = orderBy
}

func selectLimit(stmt *ast.SelectStatement, limit int) {
	if stmt.Limit == nil {
		stmt.Limit = &ast.LimitClause{}
	}
	stmt.Limit.Limit = limit
}

func selectOffset(stmt *ast.SelectStatement, offset int) {
	if stmt.Limit == nil {
		stmt.Limit = &ast.LimitClause{}
	}
	stmt.Limit.Offset = offset
}

func projectionsToNamedExpressions[T any](projections []ast.Projection[T]) []ast.NamedExpression {
	named := make([]ast.NamedExpression, 0, len(projections))
	for _, projection := range projections {
		named = append(named, projection)
	}
	return named
}

func scanProjections[T any](target *T, projections []ast.Projection[T], row pgx.Row) error {
	bindings := bindProjections(target, projections)
	if err := row.Scan(scanDestinations(bindings)...); err != nil {
		return err
	}
	return assignProjections(bindings)
}

func bindProjections[T any](target *T, projections []ast.Projection[T]) []ast.ScanBinding {
	bindings := make([]ast.ScanBinding, 0, len(projections))
	for _, projection := range projections {
		bindings = append(bindings, projection.BindScan(target))
	}
	return bindings
}

func scanDestinations(bindings []ast.ScanBinding) []any {
	scanDest := make([]any, 0, len(bindings))
	for _, binding := range bindings {
		scanDest = append(scanDest, binding.Destination())
	}
	return scanDest
}

func assignProjections(bindings []ast.ScanBinding) error {
	for _, binding := range bindings {
		if err := binding.Assign(); err != nil {
			return err
		}
	}
	return nil
}

type InsertBuilder[T any] struct {
	pool      *pgxpool.Pool
	tx        pgx.Tx
	stmt      *ast.InsertStatement
	returning []ast.Projection[T]
}

func (b *InsertBuilder[T]) Insert(columns ...ast.NamedExpression) InsertSourceQuery[T] {
	b.stmt.Into = append(b.stmt.Into, columns...)
	return b
}

func (b *InsertBuilder[T]) With(ctes ...*ast.CTE) InsertQuery[T] {
	b.stmt.With = append(b.stmt.With, ctes...)
	return b
}

func (b *InsertBuilder[T]) OnConflict(columns ...ast.NamedExpression) InsertOnConflictDoQuery[T] {
	b.stmt.OnConflict = &ast.Conflict{
		Columns: columns,
	}
	return b
}

func (b *InsertBuilder[T]) Do(expr ast.OnConflictDoExpression) InsertReturningQuery[T] {
	b.stmt.OnConflict.Action = expr
	return b
}

func (b *InsertBuilder[T]) Returning(projections ...ast.Projection[T]) InsertFinalizeQuery[T] {
	b.returning = projections
	b.stmt.Returning = projectionsToNamedExpressions(projections)
	return b
}

func (b *InsertBuilder[T]) Values(values ...*T) InsertOnConflictQuery[T] {
	columns := b.stmt.Into
	provider := &insertValuesProvider[T]{
		rows:        values,
		projections: make([]ast.Projection[T], len(columns)),
	}
	for i, column := range columns {
		if projection, ok := column.(ast.Projection[T]); ok {
			provider.projections[i] = projection
			continue
		}
		if provider.err == nil {
			if column == nil {
				provider.err = fmt.Errorf("INSERT column %d is nil", i)
			} else {
				provider.err = fmt.Errorf("INSERT column %q is not a projection for the inserted type", column.GetName())
			}
		}
	}
	b.stmt.Values = ast.NewValuesTable(provider)

	return b
}

type insertValuesProvider[T any] struct {
	rows        []*T
	projections []ast.Projection[T]
	err         error
}

func (p *insertValuesProvider[T]) RowCount() int {
	return len(p.rows)
}

func (p *insertValuesProvider[T]) ColumnCount() int {
	return len(p.projections)
}

func (p *insertValuesProvider[T]) Value(row int, column int) (any, error) {
	if p.err != nil {
		return nil, p.err
	}
	if row < 0 || row >= len(p.rows) {
		return nil, fmt.Errorf("INSERT VALUES row %d is out of range", row)
	}
	if p.rows[row] == nil {
		return nil, fmt.Errorf("INSERT VALUES row %d is nil", row)
	}
	if column < 0 || column >= len(p.projections) {
		return nil, fmt.Errorf("INSERT VALUES column %d is out of range", column)
	}
	return p.projections[column].Value(p.rows[row])
}

func (b *InsertBuilder[T]) Select(query StatementSelectFinalizeQuery) InsertOnConflictQuery[T] {
	if query != nil {
		b.stmt.Select = query.SelectStatement()
	}
	return b
}

// Exec executes the insert and returns its affected-row count and any RETURNING rows.
func (b *InsertBuilder[T]) Exec(ctx context.Context) (int64, []T, error) {
	args := make([]any, 0)
	queryCtx := &ast.QueryContext{
		PrimaryTable: b.stmt.Table,
	}
	qry, err := ast.RenderWithContext(b.stmt, &args, queryCtx)
	if err != nil {
		return 0, nil, err
	}

	if len(b.stmt.Returning) == 0 {
		var tag pgconn.CommandTag
		if b.tx != nil {
			tag, err = b.tx.Exec(ctx, qry, args...)
		} else {
			tag, err = b.pool.Exec(ctx, qry, args...)
		}
		if err != nil {
			return 0, nil, err
		}
		return tag.RowsAffected(), nil, nil
	}

	var rows pgx.Rows
	if b.tx != nil {
		rows, err = b.tx.Query(ctx, qry, args...)
	} else {
		rows, err = b.pool.Query(ctx, qry, args...)
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var result T
		if err = scanProjections(&result, b.returning, rows); err != nil {
			return rows.CommandTag().RowsAffected(), results, err
		}
		results = append(results, result)
	}

	return rows.CommandTag().RowsAffected(), results, rows.Err()
}

func (b *InsertBuilder[T]) ToSql() (string, []any, error) {
	args := make([]any, 0)
	sql, err := ast.RenderWithContext(b.stmt, &args, &ast.QueryContext{PrimaryTable: b.stmt.Table})
	return sql, args, err
}

func (b *InsertBuilder[T]) Statement() ast.SqlStatement {
	return b.stmt
}

type UpdateBuilder[T any] struct {
	pool      *pgxpool.Pool
	tx        pgx.Tx
	stmt      *ast.UpdateStatement
	returning []ast.Projection[T]
}

func NewUpdateBuilder[T any](pool *pgxpool.Pool, table string) *UpdateBuilder[T] {
	return &UpdateBuilder[T]{
		pool: pool,
		stmt: &ast.UpdateStatement{
			Table: table,
		},
	}
}

func (b *UpdateBuilder[T]) Update(toSet ...ast.UpdateSetExpr) UpdateFromQuery[T] {
	updateSet(b.stmt, toSet...)
	return b
}

func (b *UpdateBuilder[T]) With(ctes ...*ast.CTE) UpdateQuery[T] {
	updateWith(b.stmt, ctes...)
	return b
}

func (b *UpdateBuilder[T]) From(expr ast.NamedTableExpression) UpdateWhereQuery[T] {
	updateFrom(b.stmt, expr)
	return b
}

func (b *UpdateBuilder[T]) Where(expr ast.OfType[bool]) UpdateReturningQuery[T] {
	updateWhere(b.stmt, expr)
	return b
}

func (b *UpdateBuilder[T]) Returning(projections ...ast.Projection[T]) UpdateFinalizeQuery[T] {
	b.returning = projections
	updateReturning(b.stmt, projectionsToNamedExpressions(projections)...)
	return b
}

// Exec executes the update statement and returns the result.
//
// Since UPDATE statements potentially return data (RETURNING...), the signature of this method
// allows for this. []T will be nil if no data is returned from the query.
//
// Note that it's possible to affect multiple rows in a single query. In this case,
// values provided through 'update' are applied to all rows.
func (b *UpdateBuilder[T]) Exec(ctx context.Context) (int64, []T, error) {
	args := make([]any, 0, len(b.stmt.SetList)+1) // pre-allocate provided values to set + 1 where clause
	queryCtx := &ast.QueryContext{
		PrimaryTable: b.stmt.Table,
	}
	query, err := ast.RenderWithContext(b.stmt, &args, queryCtx)
	if err != nil {
		return 0, nil, err
	}
	if len(b.stmt.Returning) == 0 {
		var tag pgconn.CommandTag
		if b.tx != nil {
			tag, err = b.tx.Exec(ctx, query, args...)
		} else {
			tag, err = b.pool.Exec(ctx, query, args...)
		}

		return tag.RowsAffected(), nil, err
	}

	var rows pgx.Rows
	if b.tx != nil {
		rows, err = b.tx.Query(ctx, query, args...)
	} else {
		rows, err = b.pool.Query(ctx, query, args...)
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var t T
		if err = scanProjections(&t, b.returning, rows); err != nil {
			return rows.CommandTag().RowsAffected(), results, err
		}
		results = append(results, t)
	}

	return rows.CommandTag().RowsAffected(), results, rows.Err()
}

func (b *UpdateBuilder[T]) ToSql() (string, []any, error) {
	params := make([]any, 0, len(b.stmt.SetList)+1)
	sql, err := ast.RenderWithContext(b.stmt, &params, &ast.QueryContext{PrimaryTable: b.stmt.Table})
	return sql, params, err
}

func (b *UpdateBuilder[T]) Statement() ast.SqlStatement {
	return b.stmt
}

type StatementUpdateBuilder struct {
	stmt *ast.UpdateStatement
}

func (b *StatementUpdateBuilder) With(ctes ...*ast.CTE) StatementUpdateFromQuery {
	updateWith(b.stmt, ctes...)
	return b
}

func (b *StatementUpdateBuilder) From(expr ast.NamedTableExpression) StatementUpdateWhereQuery {
	updateFrom(b.stmt, expr)
	return b
}

func (b *StatementUpdateBuilder) Where(expr ast.OfType[bool]) StatementUpdateReturningQuery {
	updateWhere(b.stmt, expr)
	return b
}

func (b *StatementUpdateBuilder) Returning(columns ...ast.NamedExpression) StatementFinalizeQuery {
	updateReturning(b.stmt, columns...)
	return b
}

func (b *StatementUpdateBuilder) Statement() ast.SqlStatement {
	return b.stmt
}

func updateSet(stmt *ast.UpdateStatement, toSet ...ast.UpdateSetExpr) {
	stmt.SetList = toSet
}

func updateWith(stmt *ast.UpdateStatement, ctes ...*ast.CTE) {
	stmt.With = append(stmt.With, ctes...)
}

func updateFrom(stmt *ast.UpdateStatement, expr ast.NamedTableExpression) {
	stmt.From = expr
}

func updateWhere(stmt *ast.UpdateStatement, expr ast.OfType[bool]) {
	stmt.Where = expr
}

func updateReturning(stmt *ast.UpdateStatement, columns ...ast.NamedExpression) {
	stmt.Returning = columns
}

type DeleteBuilder[T any] struct {
	pool      *pgxpool.Pool
	tx        pgx.Tx
	stmt      *ast.DeleteStatement
	returning []ast.Projection[T]
}

func NewDeleteBuilder[T any](pool *pgxpool.Pool, table string) *DeleteBuilder[T] {
	return &DeleteBuilder[T]{
		pool: pool,
		stmt: &ast.DeleteStatement{
			Table: table,
		},
	}
}

func (b *DeleteBuilder[T]) Delete(from string) DeleteUsingQuery[T] {
	deleteFrom(b.stmt, from)
	return b
}

func (b *DeleteBuilder[T]) Using(tables ...ast.NamedExpression) DeleteWhereQuery[T] {
	deleteUsing(b.stmt, tables...)
	return b
}

func (b *DeleteBuilder[T]) Where(expr ast.OfType[bool]) DeleteReturningQuery[T] {
	deleteWhere(b.stmt, expr)
	return b
}

func (b *DeleteBuilder[T]) Returning(projections ...ast.Projection[T]) DeleteFinalizeQuery[T] {
	b.returning = projections
	deleteReturning(b.stmt, projectionsToNamedExpressions(projections)...)
	return b
}

func (b *DeleteBuilder[T]) Exec(ctx context.Context) (int64, []T, error) {
	args := make([]any, 0)
	queryCtx := &ast.QueryContext{}
	query, err := ast.RenderWithContext(b.stmt, &args, queryCtx)
	if err != nil {
		return 0, nil, err
	}
	if len(b.stmt.Returning) == 0 {
		var tag pgconn.CommandTag
		if b.tx != nil {
			tag, err = b.tx.Exec(ctx, query, args...)
		} else {
			tag, err = b.pool.Exec(ctx, query, args...)
		}

		return tag.RowsAffected(), nil, err
	}

	var rows pgx.Rows
	if b.tx != nil {
		rows, err = b.tx.Query(ctx, query, args...)
	} else {
		rows, err = b.pool.Query(ctx, query, args...)
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var t T
		if err = scanProjections(&t, b.returning, rows); err != nil {
			return rows.CommandTag().RowsAffected(), results, err
		}
		results = append(results, t)
	}

	return rows.CommandTag().RowsAffected(), results, rows.Err()
}

func (b *DeleteBuilder[T]) ToSql() (string, []any, error) {
	params := make([]any, 0)
	sql, err := ast.RenderWithContext(b.stmt, &params, &ast.QueryContext{PrimaryTable: b.stmt.Table})
	return sql, params, err
}

func (b *DeleteBuilder[T]) Statement() ast.SqlStatement {
	return b.stmt
}

type StatementDeleteBuilder struct {
	stmt *ast.DeleteStatement
}

func (b *StatementDeleteBuilder) Using(tables ...ast.NamedExpression) StatementDeleteWhereQuery {
	deleteUsing(b.stmt, tables...)
	return b
}

func (b *StatementDeleteBuilder) Where(expr ast.OfType[bool]) StatementDeleteReturningQuery {
	deleteWhere(b.stmt, expr)
	return b
}

func (b *StatementDeleteBuilder) Returning(columns ...ast.NamedExpression) StatementFinalizeQuery {
	deleteReturning(b.stmt, columns...)
	return b
}

func (b *StatementDeleteBuilder) Statement() ast.SqlStatement {
	return b.stmt
}

func deleteFrom(stmt *ast.DeleteStatement, from string) {
	stmt.Table = from
}

func deleteUsing(stmt *ast.DeleteStatement, tables ...ast.NamedExpression) {
	stmt.Using = tables
}

func deleteWhere(stmt *ast.DeleteStatement, expr ast.OfType[bool]) {
	stmt.Where = expr
}

func deleteReturning(stmt *ast.DeleteStatement, columns ...ast.NamedExpression) {
	stmt.Returning = columns
}
