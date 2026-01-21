package builder

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type DTO[T any] interface {
	*T
	ScanInto(row pgx.Row, stmt ast.SqlStatement) error
	GetArg(column ast.NamedExpression) (ast.Expression, error)
	TableName() string
}

// KnownTableBuilder is the builder for known tables
type KnownTableBuilder[T any, O DTO[T]] struct {
	pool      *pgxpool.Pool
	tx        pgx.Tx
	tableName string
}

func NewKnownTableBuilder[T any, O DTO[T]](dto O, pool *pgxpool.Pool) *KnownTableBuilder[T, O] {
	return &KnownTableBuilder[T, O]{
		pool:      pool,
		tableName: dto.TableName(),
	}
}

func (b *KnownTableBuilder[T, O]) WithTx(tx pgx.Tx) KnownTableStartQuery[T, O] {
	b.tx = tx
	return b
}

func (b *KnownTableBuilder[T, O]) Insert(columns ...ast.NamedExpression) InsertOnConflictQuery[T, O] {
	return &InsertBuilder[T, O]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.InsertStatement{
			Table: b.tableName,
			Into:  columns,
		},
	}
}

func (b *KnownTableBuilder[T, O]) Select(columns ...ast.NamedExpression) SelectJoinQuery[T] {
	return &SelectBuilder[T, O]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.SelectStatement{
			SelectList: columns,
			From:       &ast.TableSource{Table: b.tableName},
		},
	}
}

func (b *KnownTableBuilder[T, O]) Update(columns ...ast.NamedExpression) UpdateFromQuery[T, O] {
	setList := make([]*ast.SetKV, len(columns))
	for i, column := range columns {
		setList[i] = &ast.SetKV{Key: column}
	}

	return &UpdateBuilder[T, O]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.UpdateStatement{
			Table:   b.tableName,
			SetList: setList,
		},
	}
}

func (b *KnownTableBuilder[T, O]) Delete() DeleteUsingQuery[T, O] {
	return &DeleteBuilder[T, O]{
		pool: b.pool,
		tx:   b.tx,
		stmt: &ast.DeleteStatement{
			Table: b.tableName,
		},
	}
}

type SelectBuilder[T any, O DTO[T]] struct {
	pool   *pgxpool.Pool
	tx     pgx.Tx
	stmt   *ast.SelectStatement
	params []any
}

func (b *SelectBuilder[T, O]) WithTx(tx pgx.Tx) *SelectBuilder[T, O] {
	b.tx = tx
	return b
}

func (b *SelectBuilder[T, O]) Select(columns ...ast.NamedExpression) SelectFromQuery[T] {
	b.stmt.SelectList = append(b.stmt.SelectList, columns...)
	return b
}

func (b *SelectBuilder[T, O]) From(table *ast.TableSource) SelectWhereQuery[T] {
	b.stmt.From = table
	return b
}

func (b *SelectBuilder[T, O]) Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) SelectJoinQuery[T] {
	b.stmt.From.Join(joinType, table, expr)
	return b
}

func (b *SelectBuilder[T, O]) Where(expr ast.OfType[bool]) SelectGroupByQuery[T] {
	b.stmt.Where = expr
	return b
}

func (b *SelectBuilder[T, O]) GroupBy(columns ...ast.Expression) SelectHavingQuery[T] {
	b.stmt.GroupBy = columns
	return b
}

func (b *SelectBuilder[T, O]) Having(expr ast.OfType[bool]) SelectOrderByQuery[T] {
	b.stmt.Having = expr
	return b
}

func (b *SelectBuilder[T, O]) OrderBy(orderBy ...*ast.OrderByItem) SelectPagingQuery[T] {
	b.stmt.OrderBy = orderBy
	return b
}

func (b *SelectBuilder[T, O]) Limit(limit int) SelectPagingQuery[T] {
	if b.stmt.Limit == nil {
		b.stmt.Limit = &ast.LimitClause{}
	}
	b.stmt.Limit.Limit = limit
	return b
}

func (b *SelectBuilder[T, O]) Offset(offset int) SelectPagingQuery[T] {
	if b.stmt.Limit == nil {
		b.stmt.Limit = &ast.LimitClause{}
	}
	b.stmt.Limit.Offset = offset
	return b
}

func (b *SelectBuilder[T, O]) ToSql() string {
	return ast.Render(b.stmt, &b.params)
}

func (b *SelectBuilder[T, O]) Find(ctx context.Context) ([]T, error) {
	query := b.ToSql()
	args := b.params

	var rows pgx.Rows
	var err error
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
		result := O(&t)
		if err = result.ScanInto(rows, b.stmt); err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, rows.Err()
}

func (b *SelectBuilder[T, O]) FindOne(ctx context.Context) (T, error) {
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

type InsertBuilder[T any, O DTO[T]] struct {
	pool   *pgxpool.Pool
	tx     pgx.Tx
	stmt   *ast.InsertStatement
	params []any
}

func (b *InsertBuilder[T, O]) Insert(columns ...ast.NamedExpression) InsertOnConflictQuery[T, O] {
	b.stmt.Into = append(b.stmt.Into, columns...)
	return b
}

func (b *InsertBuilder[T, O]) OnConflict(columns ...ast.NamedExpression) InsertOnConflictDoQuery[T, O] {
	b.stmt.OnConflict = &ast.Conflict{
		Columns: columns,
	}
	return b
}

func (b *InsertBuilder[T, O]) Do(expr ast.OnConflictDoExpression) InsertReturningQuery[T, O] {
	b.stmt.OnConflict.Action = expr
	return b
}

func (b *InsertBuilder[T, O]) Returning(columns ...ast.NamedExpression) InsertFinalizeQuery[T, O] {
	b.stmt.Returning = columns
	return b
}

// Exec inserts a new record
func (b *InsertBuilder[T, O]) Exec(ctx context.Context, record O) error {
	for _, column := range b.stmt.Into {
		value, err := record.GetArg(column)
		if err != nil {
			return fmt.Errorf("error executing insert query: %w", err)
		}
		b.stmt.Values = append(b.stmt.Values, value)
	}
	query := b.ToSql()
	args := b.params

	if len(args) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	var row pgx.Row
	if b.tx != nil {
		row = b.tx.QueryRow(ctx, query, args...)
	} else {
		row = b.pool.QueryRow(ctx, query, args...)
	}
	var err error
	if len(b.stmt.Returning) > 0 {
		if err = record.ScanInto(row, b.stmt); err != nil {
			return err
		}
	}

	return nil
}

func (b *InsertBuilder[T, O]) ToSql() string {
	return ast.Render(b.stmt, &b.params)
}

// ExecBatch inserts multiple records efficiently using pgx batch
func (b *InsertBuilder[T, O]) ExecBatch(ctx context.Context, records []T) error {
	if len(records) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	for _, record := range records {
		dto := O(&record)
		for _, column := range b.stmt.Into {
			value, err := dto.GetArg(column)
			if err != nil {
				return fmt.Errorf("error executing insert query: %w", err)
			}
			b.stmt.Values = append(b.stmt.Values, value)
		}
		b.params = []any{}
		query := b.ToSql()
		args := b.params
		batch.Queue(query, args...)
	}

	var br pgx.BatchResults
	if b.tx != nil {
		br = b.tx.SendBatch(ctx, batch)
	} else {
		br = b.pool.SendBatch(ctx, batch)
	}
	defer br.Close()

	// Scan returned PKs (if any) back into records
	if len(b.stmt.Returning) > 0 {
		for _, record := range records {
			dto := O(&record)
			if err := dto.ScanInto(br.QueryRow().(pgx.Rows), b.stmt); err != nil {
				return fmt.Errorf("failed to scan batch result %+v: %w", record, err)
			}
		}
	}

	return br.Close()
}

type UpdateBuilder[T any, O DTO[T]] struct {
	pool *pgxpool.Pool
	tx   pgx.Tx
	stmt *ast.UpdateStatement
}

func NewUpdateBuilder[T any, O DTO[T]](pool *pgxpool.Pool, table string) *UpdateBuilder[T, O] {
	return &UpdateBuilder[T, O]{
		pool: pool,
		stmt: &ast.UpdateStatement{
			Table: table,
		},
	}
}

func (b *UpdateBuilder[T, O]) Set(columns ...ast.NamedExpression) UpdateFromQuery[T, O] {
	setKv := make([]*ast.SetKV, 0, len(columns))
	for _, column := range columns {
		setKv = append(setKv, &ast.SetKV{Key: column})
	}

	b.stmt.SetList = setKv
	return b
}

func (b *UpdateBuilder[T, O]) From(table *ast.TableSource) UpdateWhereQuery[T, O] {
	b.stmt.From = table
	return b
}

func (b *UpdateBuilder[T, O]) Where(expr ast.OfType[bool]) UpdateReturningQuery[T, O] {
	b.stmt.Where = expr
	return b
}

func (b *UpdateBuilder[T, O]) Returning(columns ...ast.NamedExpression) UpdateFinalizeQuery[T, O] {
	b.stmt.Returning = columns
	return b
}

// Exec executes the update statement and returns the result.
//
// Since UPDATE statements potentially return data (RETURNING...), the signature of this method
// allows for this. []T will be nil if no data is returned from the query.
//
// Note that it's possible to affect multiple rows in a single query. In this case,
// values provided through 'update' are applied to all rows.
func (b *UpdateBuilder[T, O]) Exec(ctx context.Context, record O) (int64, []T, error) {
	var err error
	for _, set := range b.stmt.SetList {
		if set.Value == nil {
			set.Value, err = record.GetArg(set.Key)
			if err != nil {
				return 0, nil, err
			}
		}
	}

	args := make([]any, 0, len(b.stmt.SetList)+1) // pre-allocate provided values to set + 1 where clause
	query := ast.Render(b.stmt, &args)
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
		result := O(&t)
		if err = result.ScanInto(rows, b.stmt); err != nil {
			return rows.CommandTag().RowsAffected(), results, err
		}
		results = append(results, t)
	}

	return rows.CommandTag().RowsAffected(), results, rows.Err()
}

func (b *UpdateBuilder[T, O]) ToSql() string {
	params := make([]any, 0, len(b.stmt.SetList)+1)
	return ast.Render(b.stmt, &params)
}

func (b *UpdateBuilder[T, O]) ExecBatch(ctx context.Context, records []T) error {
	if len(records) == 0 {
		return nil
	}

	var err error
	batch := &pgx.Batch{}
	for _, record := range records {
		dto := O(&record)
		for _, set := range b.stmt.SetList {
			if set.Value == nil {
				set.Value, err = dto.GetArg(set.Key)
				if err != nil {
					return err
				}
			}
		}
		args := make([]any, 0, len(b.stmt.SetList)+1) // pre-allocate provided values to set + 1 where clause
		query := ast.Render(b.stmt, &args)
		batch.Queue(query, args...)
	}

	var br pgx.BatchResults
	if b.tx != nil {
		br = b.tx.SendBatch(ctx, batch)
	} else {
		br = b.pool.SendBatch(ctx, batch)
	}
	defer br.Close()

	// Scan returned PKs (if any) back into records
	if len(b.stmt.Returning) > 0 {
		for _, record := range records {
			dto := O(&record)
			if err := dto.ScanInto(br.QueryRow().(pgx.Rows), b.stmt); err != nil {
				return fmt.Errorf("failed to scan batch result %+v: %w", record, err)
			}
		}
	}

	return br.Close()
}

type DeleteBuilder[T any, O DTO[T]] struct {
	pool *pgxpool.Pool
	tx   pgx.Tx
	stmt *ast.DeleteStatement
}

func NewDeleteBuilder[T any, O DTO[T]](pool *pgxpool.Pool, table string) *DeleteBuilder[T, O] {
	return &DeleteBuilder[T, O]{
		pool: pool,
		stmt: &ast.DeleteStatement{
			Table: table,
		},
	}
}

func (b *DeleteBuilder[T, O]) Delete(from string) DeleteUsingQuery[T, O] {
	b.stmt.Table = from
	return b
}

func (b *DeleteBuilder[T, O]) Using(tables ...ast.NamedExpression) DeleteWhereQuery[T, O] {
	b.stmt.Using = tables
	return b
}

func (b *DeleteBuilder[T, O]) Where(expr ast.OfType[bool]) DeleteReturningQuery[T, O] {
	b.stmt.Where = expr
	return b
}

func (b *DeleteBuilder[T, O]) Returning(columns ...ast.NamedExpression) DeleteFinalizeQuery[T, O] {
	b.stmt.Returning = columns
	return b
}

func (b *DeleteBuilder[T, O]) Exec(ctx context.Context) (int64, []T, error) {
	var err error
	args := make([]any, 0)
	query := ast.Render(b.stmt, &args)
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
		result := O(&t)
		if err = result.ScanInto(rows, b.stmt); err != nil {
			return rows.CommandTag().RowsAffected(), results, err
		}
		results = append(results, t)
	}

	return rows.CommandTag().RowsAffected(), results, rows.Err()
}

func (b *DeleteBuilder[T, O]) ToSql() string {
	params := make([]any, 0)
	return ast.Render(b.stmt, &params)
}
