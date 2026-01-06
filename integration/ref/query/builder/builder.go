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

	sb                    selectBuilder[T, O]
	SelectGroupByQuery[T] //To make sure KnownTableBuilder implements the full interface
}

func NewKnownTableBuilder[T any, O DTO[T]](dto O, pool *pgxpool.Pool) *KnownTableBuilder[T, O] {
	return &KnownTableBuilder[T, O]{
		pool:      pool,
		tableName: dto.TableName(),
	}
}

func (qb *KnownTableBuilder[T, O]) WithTx(tx pgx.Tx) KnownTableStartQuery[T, O] {
	qb.tx = tx
	return qb
}

func (qb *KnownTableBuilder[T, O]) Insert(columns ...ast.NamedExpression) InsertOnConflictQuery[T, O] {
	return &InsertBuilder[T, O]{
		pool: qb.pool,
		tx:   qb.tx,
		stmt: &ast.InsertStatement{
			Table: qb.tableName,
			Into:  columns,
		},
	}
}

func (qb *KnownTableBuilder[T, O]) Update(columns ...ast.NamedExpression) UpdateFromQuery[T, O] {
	setList := make([]*ast.SetKV, len(columns))
	for i, column := range columns {
		setList[i] = &ast.SetKV{Key: column}
	}

	return &UpdateBuilder[T, O]{
		pool: qb.pool,
		tx:   qb.tx,
		stmt: &ast.UpdateStatement{
			Table:   qb.tableName,
			SetList: setList,
		},
		params: nil,
	}
}

func (qb *KnownTableBuilder[T, O]) Select(columns ...ast.NamedExpression) SelectJoinQuery[T] {
	if qb.sb.stmt == nil {
		qb.sb.stmt = &ast.SelectStatement{
			SelectList: columns,
			From:       &ast.TableSource{Name: qb.tableName},
		}
	}
	qb.sb.stmt.SelectList = append(qb.sb.stmt.SelectList, columns...)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) SelectJoinQuery[T] {
	_ = qb.sb.stmt.From.Join(joinType, table, expr)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Where(expr ast.OfType[bool]) SelectGroupByQuery[T] {
	qb.sb.stmt.Where = expr
	qb.sb.pool = qb.pool
	qb.sb.tx = qb.tx
	return &qb.sb
}

type SelectBuilder[T any, O DTO[T]] struct {
	pool   *pgxpool.Pool
	tx     pgx.Tx
	stmt   *ast.SelectStatement
	params []any
}

func (qb *SelectBuilder[T, O]) WithTx(tx pgx.Tx) *SelectBuilder[T, O] {
	qb.tx = tx
	return qb
}

func (qb *SelectBuilder[T, O]) Select(columns ...ast.NamedExpression) SelectFromQuery[T] {
	qb.stmt.SelectList = append(qb.stmt.SelectList, columns...)
	return qb
}

func (qb *SelectBuilder[T, O]) From(table *ast.TableSource) SelectWhereQuery[T] {
	qb.stmt.From = table
	return qb
}

func (qb *SelectBuilder[T, O]) Where(expr ast.OfType[bool]) SelectGroupByQuery[T] {
	qb.stmt.Where = expr
	return qb
}

func (qb *SelectBuilder[T, O]) GroupBy(columns ...ast.Expression) SelectHavingQuery[T] {
	qb.stmt.GroupBy = columns
	return qb
}

func (qb *SelectBuilder[T, O]) Having(expr ast.OfType[bool]) SelectOrderByQuery[T] {
	qb.stmt.Having = expr
	return qb
}

func (qb *SelectBuilder[T, O]) OrderBy(orderBy ...*ast.OrderByItem) SelectPagingQuery[T] {
	qb.stmt.OrderBy = orderBy
	return qb
}

func (qb *SelectBuilder[T, O]) Limit(limit int) SelectPagingQuery[T] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Limit = limit
	return qb
}

func (qb *SelectBuilder[T, O]) Offset(offset int) SelectPagingQuery[T] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Offset = offset
	return qb
}

func (qb *SelectBuilder[T, O]) ToSql() string {
	return ast.Render(qb.stmt, &qb.params)
}

func (qb *SelectBuilder[T, O]) Find(ctx context.Context) ([]T, error) {
	query := qb.ToSql()
	args := qb.params

	var rows pgx.Rows
	var err error
	if qb.tx != nil {
		rows, err = qb.tx.Query(ctx, query, args...)
	} else {
		rows, err = qb.pool.Query(ctx, query, args...)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		var t T
		result := O(&t)
		if err = result.ScanInto(rows, qb.stmt); err != nil {
			return nil, err
		}
		results = append(results, t)
	}

	return results, rows.Err()
}

func (qb *SelectBuilder[T, O]) FindOne(ctx context.Context) (T, error) {
	qb.Limit(1)
	results, err := qb.Find(ctx)
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

func (qb *InsertBuilder[T, O]) Insert(columns ...ast.NamedExpression) InsertOnConflictQuery[T, O] {
	qb.stmt.Into = append(qb.stmt.Into, columns...)
	return qb
}

func (qb *InsertBuilder[T, O]) OnConflict(columns ...ast.NamedExpression) InsertOnConflictDoQuery[T, O] {
	qb.stmt.OnConflict = &ast.Conflict{
		Columns: columns,
	}
	return qb
}

func (qb *InsertBuilder[T, O]) Do(expr ast.OnConflictDoExpression) InsertReturningQuery[T, O] {
	qb.stmt.OnConflict.Action = expr
	return qb
}

func (qb *InsertBuilder[T, O]) Returning(columns ...ast.NamedExpression) InsertFinalizeQuery[T, O] {
	qb.stmt.Returning = columns
	return qb
}

// Exec inserts a new record
func (qb *InsertBuilder[T, O]) Exec(ctx context.Context, record O) error {
	for _, column := range qb.stmt.Into {
		value, err := record.GetArg(column)
		if err != nil {
			return fmt.Errorf("error executing insert query: %w", err)
		}
		qb.stmt.Values = append(qb.stmt.Values, value)
	}
	query := qb.ToSql()
	args := qb.params

	if len(args) == 0 {
		return fmt.Errorf("no values provided for insert")
	}

	var row pgx.Row
	if qb.tx != nil {
		row = qb.tx.QueryRow(ctx, query, args...)
	} else {
		row = qb.pool.QueryRow(ctx, query, args...)
	}
	var err error
	if len(qb.stmt.Returning) > 0 {
		if err = record.ScanInto(row, qb.stmt); err != nil {
			return err
		}
	}

	return nil
}

func (qb *InsertBuilder[T, O]) ToSql() string {
	return ast.Render(qb.stmt, &qb.params)
}

// ExecBatch inserts multiple records efficiently using pgx batch
func (qb *InsertBuilder[T, O]) ExecBatch(ctx context.Context, records []T) error {
	if len(records) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	for _, record := range records {
		dto := O(&record)
		for _, column := range qb.stmt.Into {
			value, err := dto.GetArg(column)
			if err != nil {
				return fmt.Errorf("error executing insert query: %w", err)
			}
			qb.stmt.Values = append(qb.stmt.Values, value)
		}
		qb.params = []any{}
		query := qb.ToSql()
		args := qb.params
		batch.Queue(query, args...)
	}

	var br pgx.BatchResults
	if qb.tx != nil {
		br = qb.tx.SendBatch(ctx, batch)
	} else {
		br = qb.pool.SendBatch(ctx, batch)
	}
	defer br.Close()

	// Scan returned PKs (if any) back into records
	if len(qb.stmt.Returning) > 0 {
		for _, record := range records {
			dto := O(&record)
			if err := dto.ScanInto(br.QueryRow().(pgx.Rows), qb.stmt); err != nil {
				return fmt.Errorf("failed to scan batch result %+v: %w", record, err)
			}
		}
	}

	return br.Close()
}

type UpdateBuilder[T any, O DTO[T]] struct {
	pool   *pgxpool.Pool
	tx     pgx.Tx
	stmt   *ast.UpdateStatement
	params []any
}

func NewUpdateBuilder[T any, O DTO[T]](pool *pgxpool.Pool) *UpdateBuilder[T, O] {
	var o O
	return &UpdateBuilder[T, O]{
		pool: pool,
		stmt: &ast.UpdateStatement{
			Table: o.TableName(),
		},
		params: make([]any, 0),
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
func (b *UpdateBuilder[T, O]) Exec(ctx context.Context, update O) (int64, []T, error) {
	var err error
	for _, set := range b.stmt.SetList {
		if set.Value == nil {
			set.Value, err = update.GetArg(set.Key)
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
	return ast.Render(b.stmt, &b.params)
}

//// Delete deletes matching records
//func (qb *SelectBuilder) Delete(ctx context.Context) (int64, error) {
//	var query strings.Builder
//	var args []interface{}
//	argIndex := 1
//
//	query.WriteString("DELETE FROM users")
//
//	// Add WHERE conditions
//	if len(qb.conditions) > 0 {
//		query.WriteString(" WHERE ")
//		whereStr, whereArgs := qb.buildConditions(qb.conditions, &argIndex)
//		query.WriteString(whereStr)
//		args = append(args, whereArgs...)
//	}
//
//	var tag pgconn.CommandTag
//	var err error
//
//	if qb.tx != nil {
//		tag, err = qb.tx.Exec(ctx, query.String(), args...)
//	} else {
//		tag, err = qb.pool.Exec(ctx, query.String(), args...)
//	}
//
//	if err != nil {
//		return 0, err
//	}
//
//	return tag.RowsAffected(), nil
//}

type selectBuilder[T any, O DTO[T]] = SelectBuilder[T, O]
