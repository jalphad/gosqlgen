package builder

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

type DTO[T any] interface {
	*T
	ScanInto(rows pgx.Rows, stmt *ast.SelectStatement) error
	PrepareInsert() (string, []any, []any)
	TableName() string
}

// KnownTableBuilder is the builder for known tables
type KnownTableBuilder[T any, O DTO[T]] struct {
	builder[T, O]
}

func NewKnownTableBuilder[T any, O DTO[T]](dto O, pool *pgxpool.Pool) *KnownTableBuilder[T, O] {
	return &KnownTableBuilder[T, O]{
		builder: Builder[T, O]{
			pool: pool,
			stmt: &ast.SelectStatement{
				From: &ast.TableSource{Name: dto.TableName()},
			},
		},
	}
}

func (qb *KnownTableBuilder[T, O]) WithTx(tx pgx.Tx) *KnownTableBuilder[T, O] {
	qb.builder.tx = tx
	return qb
}

func (qb *KnownTableBuilder[T, O]) Select(columns ...ast.NamedExpression) JoinQuery[T] {
	qb.builder.stmt.SelectList = append(qb.builder.stmt.SelectList, columns...)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Join(joinType ast.JoinType, table string, expr ast.OfType[bool]) JoinQuery[T] {
	_ = qb.builder.stmt.From.Join(joinType, table, expr)
	return qb
}

func (qb *KnownTableBuilder[T, O]) Where(expr ast.OfType[bool]) GroupByQuery[T] {
	qb.builder.stmt.Where = expr
	return &qb.builder
}

type Builder[T any, O DTO[T]] struct {
	pool   *pgxpool.Pool
	tx     pgx.Tx
	stmt   *ast.SelectStatement
	params []interface{}
}

func (qb *Builder[T, O]) WithTx(tx pgx.Tx) *Builder[T, O] {
	qb.tx = tx
	return qb
}

func (qb *Builder[T, O]) Select(columns ...ast.NamedExpression) FromQuery[T] {
	qb.stmt.SelectList = append(qb.stmt.SelectList, columns...)
	return qb
}

func (qb *Builder[T, O]) From(table *ast.TableSource) WhereQuery[T] {
	qb.stmt.From = table
	return qb
}

func (qb *Builder[T, O]) Where(expr ast.OfType[bool]) GroupByQuery[T] {
	qb.stmt.Where = expr
	return qb
}

func (qb *Builder[T, O]) GroupBy(columns ...ast.Expression) HavingQuery[T] {
	qb.stmt.GroupBy = columns
	return qb
}

func (qb *Builder[T, O]) Having(expr ast.OfType[bool]) OrderByQuery[T] {
	qb.stmt.Having = expr
	return qb
}

func (qb *Builder[T, O]) OrderBy(orderBy ...*ast.OrderByItem) PagingQuery[T] {
	qb.stmt.OrderBy = orderBy
	return qb
}

func (qb *Builder[T, O]) Limit(limit int) PagingQuery[T] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Limit = limit
	return qb
}

func (qb *Builder[T, O]) Offset(offset int) PagingQuery[T] {
	if qb.stmt.Limit == nil {
		qb.stmt.Limit = &ast.LimitClause{}
	}
	qb.stmt.Limit.Offset = offset
	return qb
}

func (qb *Builder[T, O]) ToSql() (string, []any) {
	sql := ast.Render(qb.stmt, &qb.params)
	return sql, qb.params
}

func (qb *Builder[T, O]) Find(ctx context.Context) ([]T, error) {
	query, args := qb.ToSql()

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

func (qb *Builder[T, O]) FindOne(ctx context.Context) (T, error) {
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
func (qb *Builder[T, O]) Insert(ctx context.Context, record O) error {
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
func (qb *Builder[T, O]) InsertBatch(ctx context.Context, records []O) error {
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

type builder[T any, O DTO[T]] = Builder[T, O]
