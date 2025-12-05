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
	ScanInto(rows pgx.Rows, stmt *ast.SelectStatement) error
	PrepareInsert() (string, []any, []any)
	GetArgs([]ast.NamedExpression) []any
	TableName() string
}

// KnownTableBuilder is the builder for known tables
type KnownTableBuilder[T any, O DTO[T]] struct {
	selectBuilder[T, O]
}

func NewKnownTableBuilder[T any, O DTO[T]](dto O, pool *pgxpool.Pool) *KnownTableBuilder[T, O] {
	return &KnownTableBuilder[T, O]{
		selectBuilder: SelectBuilder[T, O]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: dto.TableName()},
			},
		},
	}
}

func (qb *KnownTableBuilder[T, O]) WithTx(tx pgx.Tx) *KnownTableBuilder[T, O] {
	qb.selectBuilder.tx = tx
	return qb
}

func (qb *KnownTableBuilder[T, O]) Select(columns ...ast.NamedExpression) SelectJoinQuery[T] {
	qb.selectBuilder.stmt.SelectList = append(qb.selectBuilder.stmt.SelectList, columns...)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) SelectJoinQuery[T] {
	_ = qb.selectBuilder.stmt.From.Join(joinType, table, expr)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Where(expr ast.OfType[bool]) SelectGroupByQuery[T] {
	qb.selectBuilder.stmt.Where = expr
	return &qb.selectBuilder
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
		if err := result.ScanInto(rows, qb.stmt); err != nil {
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

// Insert inserts a new record
// Fields with nil values (defaults/sequences) are omitted, database handles them
func (qb *SelectBuilder[T, O]) Insert(ctx context.Context, record O) error {
	query, args, pk := record.PrepareInsert()
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
	if len(pk) > 0 {
		err = row.Scan(pk...)
	}

	return err
}

// InsertBatch inserts multiple records efficiently using pgx batch
func (qb *SelectBuilder[T, O]) InsertBatch(ctx context.Context, records []O) error {
	if len(records) == 0 {
		return nil
	}

	var pks [][]any
	batch := &pgx.Batch{}
	for _, record := range records {
		query, args, pk := record.PrepareInsert()

		if len(args) > 0 {
			pks = append(pks, pk)
			batch.Queue(query, args...)
		}
	}

	var br pgx.BatchResults
	if qb.tx != nil {
		br = qb.tx.SendBatch(ctx, batch)
	} else {
		br = qb.pool.SendBatch(ctx, batch)
	}
	defer br.Close()

	// Scan returned PKs (if any) back into records
	if len(pks[0]) > 0 {
		for i := range records {
			if err := br.QueryRow().Scan(pks[i]...); err != nil {
				return fmt.Errorf("failed to scan batch result %d: %w", i, err)
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

// Update updates a record using its primary key
func (b *UpdateBuilder[T, O]) Update(columns ...ast.NamedExpression) *UpdateBuilder[T, O] {
	setKv := make([]*ast.SetKV, 0, len(columns))
	for _, column := range columns {
		setKv = append(setKv, &ast.SetKV{Key: column})
	}

	b.stmt.SetList = setKv
	return b
}

func (b *UpdateBuilder[T, O]) Where(expr ast.OfType[bool]) *UpdateBuilder[T, O] {
	b.stmt.Where = expr
	return b
}

func (b *UpdateBuilder[T, O]) Exec(ctx context.Context, update O) (int64, error) {
	ce := make([]ast.NamedExpression, 0, len(b.stmt.SetList))
	for _, set := range b.stmt.SetList {
		ce = append(ce, set.Key)
	}

	args := update.GetArgs(ce)
	for i, set := range b.stmt.SetList {
		set.Value = args[i]
	}

	query := b.ToSql()
	args = b.params
	var err error
	var tag pgconn.CommandTag
	if b.tx != nil {
		tag, err = b.tx.Exec(ctx, query, args...)
	} else {
		tag, err = b.pool.Exec(ctx, query, args...)
	}

	return tag.RowsAffected(), err
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
