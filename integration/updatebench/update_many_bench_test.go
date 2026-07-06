package updatebench

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/models"
	genquery "github.com/jalphad/gosqlgen/integration/updatebench/gen/query"
)

const benchmarkRowCount = 200

func BenchmarkWideRecordsUpdateManyConstructUnnest200Rows100Columns(b *testing.B) {
	dtos := makeWideRecordDtos(benchmarkRowCount)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query, err := genquery.WideRecordsDtoUpdateMany(nil, dtos)
		if err != nil {
			b.Fatal(err)
		}
		if query == nil {
			b.Fatal("nil query")
		}
	}
}

func BenchmarkWideRecordsUpdateManyConstructValues200Rows100Columns(b *testing.B) {
	dtos := makeWideRecordDtos(benchmarkRowCount)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query, err := genquery.WideRecordsDtoUpdateMany(nil, dtos, genquery.UpdateWithValues())
		if err != nil {
			b.Fatal(err)
		}
		if query == nil {
			b.Fatal("nil query")
		}
	}
}

func BenchmarkWideRecordsUpdateManyUnnest200Rows100Columns(b *testing.B) {
	dtos := makeWideRecordDtos(benchmarkRowCount)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query, err := genquery.WideRecordsDtoUpdateMany(nil, dtos)
		if err != nil {
			b.Fatal(err)
		}
		sql, args, err := query.ToSql()
		if err != nil {
			b.Fatal(err)
		}
		if len(sql) == 0 || len(args) != 101 {
			b.Fatalf("unexpected query output: sql length %d, args %d", len(sql), len(args))
		}
	}
}

func BenchmarkWideRecordsUpdateManyValues200Rows100Columns(b *testing.B) {
	dtos := makeWideRecordDtos(benchmarkRowCount)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		query, err := genquery.WideRecordsDtoUpdateMany(nil, dtos, genquery.UpdateWithValues())
		if err != nil {
			b.Fatal(err)
		}
		sql, args, err := query.ToSql()
		if err != nil {
			b.Fatal(err)
		}
		if len(sql) == 0 || len(args) != 20200 {
			b.Fatalf("unexpected query output: sql length %d, args %d", len(sql), len(args))
		}
	}
}

func makeWideRecordDtos(count int) models.WideRecordsDtos {
	dtos := make(models.WideRecordsDtos, count)
	for i := range dtos {
		id := uuid.New()
		dtos[i] = &models.WideRecordsDto{
			Id: &id,
		}
	}
	return dtos
}
