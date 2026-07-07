package wide_records

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/jalphad/gosqlgen/integration/updatebench/gen/models"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/builder"
)

// NewQuery returns a query builder for wide_records
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.WideRecordsDto] {
	return builder.NewKnownTableBuilder[models.WideRecordsDto](pool, Table())
}

var Into = intoWideRecordsDto{}

type intoWideRecordsDto struct{}

type forWideRecordsDto[E models.ExportsWideRecordsDto[T], T any] struct {
	alias *Alias
	err   error
}

// TODO(go1.27): if type parameters on methods are available, consider
// supporting alias.For[*resultRow]() as the primary alias API.
func For[E models.ExportsWideRecordsDto[T], T any](alias ...*Alias) forWideRecordsDto[E, T] {
	switch len(alias) {
	case 0:
		return forWideRecordsDto[E, T]{}
	case 1:
		if alias[0] == nil || alias[0].Alias == nil {
			return forWideRecordsDto[E, T]{err: fmt.Errorf("wide_records.For requires a non-nil alias")}
		}
		return forWideRecordsDto[E, T]{alias: alias[0]}
	default:
		return forWideRecordsDto[E, T]{err: fmt.Errorf("wide_records.For accepts at most one alias")}
	}
}

func Table() *ast.TableSource {
	return ast.NewTableSource("wide_records")
}

func Id() *ast.UUIDColumnProjection[models.WideRecordsDto, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		"wide_records",
		"id",
		func(w *models.WideRecordsDto) **uuid.UUID {
			return &w.Id
		},
	)
}

func Col001() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_001",
		func(w *models.WideRecordsDto) *string {
			return &w.Col001
		},
	)
}

func Col002() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_002",
		func(w *models.WideRecordsDto) *string {
			return &w.Col002
		},
	)
}

func Col003() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_003",
		func(w *models.WideRecordsDto) *string {
			return &w.Col003
		},
	)
}

func Col004() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_004",
		func(w *models.WideRecordsDto) *string {
			return &w.Col004
		},
	)
}

func Col005() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_005",
		func(w *models.WideRecordsDto) *string {
			return &w.Col005
		},
	)
}

func Col006() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_006",
		func(w *models.WideRecordsDto) *string {
			return &w.Col006
		},
	)
}

func Col007() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_007",
		func(w *models.WideRecordsDto) *string {
			return &w.Col007
		},
	)
}

func Col008() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_008",
		func(w *models.WideRecordsDto) *string {
			return &w.Col008
		},
	)
}

func Col009() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_009",
		func(w *models.WideRecordsDto) *string {
			return &w.Col009
		},
	)
}

func Col010() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_010",
		func(w *models.WideRecordsDto) *string {
			return &w.Col010
		},
	)
}

func Col011() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_011",
		func(w *models.WideRecordsDto) *string {
			return &w.Col011
		},
	)
}

func Col012() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_012",
		func(w *models.WideRecordsDto) *string {
			return &w.Col012
		},
	)
}

func Col013() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_013",
		func(w *models.WideRecordsDto) *string {
			return &w.Col013
		},
	)
}

func Col014() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_014",
		func(w *models.WideRecordsDto) *string {
			return &w.Col014
		},
	)
}

func Col015() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_015",
		func(w *models.WideRecordsDto) *string {
			return &w.Col015
		},
	)
}

func Col016() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_016",
		func(w *models.WideRecordsDto) *string {
			return &w.Col016
		},
	)
}

func Col017() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_017",
		func(w *models.WideRecordsDto) *string {
			return &w.Col017
		},
	)
}

func Col018() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_018",
		func(w *models.WideRecordsDto) *string {
			return &w.Col018
		},
	)
}

func Col019() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_019",
		func(w *models.WideRecordsDto) *string {
			return &w.Col019
		},
	)
}

func Col020() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_020",
		func(w *models.WideRecordsDto) *string {
			return &w.Col020
		},
	)
}

func Col021() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_021",
		func(w *models.WideRecordsDto) *string {
			return &w.Col021
		},
	)
}

func Col022() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_022",
		func(w *models.WideRecordsDto) *string {
			return &w.Col022
		},
	)
}

func Col023() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_023",
		func(w *models.WideRecordsDto) *string {
			return &w.Col023
		},
	)
}

func Col024() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_024",
		func(w *models.WideRecordsDto) *string {
			return &w.Col024
		},
	)
}

func Col025() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_025",
		func(w *models.WideRecordsDto) *string {
			return &w.Col025
		},
	)
}

func Col026() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_026",
		func(w *models.WideRecordsDto) *string {
			return &w.Col026
		},
	)
}

func Col027() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_027",
		func(w *models.WideRecordsDto) *string {
			return &w.Col027
		},
	)
}

func Col028() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_028",
		func(w *models.WideRecordsDto) *string {
			return &w.Col028
		},
	)
}

func Col029() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_029",
		func(w *models.WideRecordsDto) *string {
			return &w.Col029
		},
	)
}

func Col030() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_030",
		func(w *models.WideRecordsDto) *string {
			return &w.Col030
		},
	)
}

func Col031() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_031",
		func(w *models.WideRecordsDto) *string {
			return &w.Col031
		},
	)
}

func Col032() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_032",
		func(w *models.WideRecordsDto) *string {
			return &w.Col032
		},
	)
}

func Col033() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_033",
		func(w *models.WideRecordsDto) *string {
			return &w.Col033
		},
	)
}

func Col034() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_034",
		func(w *models.WideRecordsDto) *string {
			return &w.Col034
		},
	)
}

func Col035() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_035",
		func(w *models.WideRecordsDto) *string {
			return &w.Col035
		},
	)
}

func Col036() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_036",
		func(w *models.WideRecordsDto) *string {
			return &w.Col036
		},
	)
}

func Col037() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_037",
		func(w *models.WideRecordsDto) *string {
			return &w.Col037
		},
	)
}

func Col038() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_038",
		func(w *models.WideRecordsDto) *string {
			return &w.Col038
		},
	)
}

func Col039() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_039",
		func(w *models.WideRecordsDto) *string {
			return &w.Col039
		},
	)
}

func Col040() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_040",
		func(w *models.WideRecordsDto) *string {
			return &w.Col040
		},
	)
}

func Col041() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_041",
		func(w *models.WideRecordsDto) *string {
			return &w.Col041
		},
	)
}

func Col042() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_042",
		func(w *models.WideRecordsDto) *string {
			return &w.Col042
		},
	)
}

func Col043() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_043",
		func(w *models.WideRecordsDto) *string {
			return &w.Col043
		},
	)
}

func Col044() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_044",
		func(w *models.WideRecordsDto) *string {
			return &w.Col044
		},
	)
}

func Col045() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_045",
		func(w *models.WideRecordsDto) *string {
			return &w.Col045
		},
	)
}

func Col046() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_046",
		func(w *models.WideRecordsDto) *string {
			return &w.Col046
		},
	)
}

func Col047() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_047",
		func(w *models.WideRecordsDto) *string {
			return &w.Col047
		},
	)
}

func Col048() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_048",
		func(w *models.WideRecordsDto) *string {
			return &w.Col048
		},
	)
}

func Col049() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_049",
		func(w *models.WideRecordsDto) *string {
			return &w.Col049
		},
	)
}

func Col050() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_050",
		func(w *models.WideRecordsDto) *string {
			return &w.Col050
		},
	)
}

func Col051() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_051",
		func(w *models.WideRecordsDto) *string {
			return &w.Col051
		},
	)
}

func Col052() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_052",
		func(w *models.WideRecordsDto) *string {
			return &w.Col052
		},
	)
}

func Col053() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_053",
		func(w *models.WideRecordsDto) *string {
			return &w.Col053
		},
	)
}

func Col054() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_054",
		func(w *models.WideRecordsDto) *string {
			return &w.Col054
		},
	)
}

func Col055() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_055",
		func(w *models.WideRecordsDto) *string {
			return &w.Col055
		},
	)
}

func Col056() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_056",
		func(w *models.WideRecordsDto) *string {
			return &w.Col056
		},
	)
}

func Col057() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_057",
		func(w *models.WideRecordsDto) *string {
			return &w.Col057
		},
	)
}

func Col058() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_058",
		func(w *models.WideRecordsDto) *string {
			return &w.Col058
		},
	)
}

func Col059() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_059",
		func(w *models.WideRecordsDto) *string {
			return &w.Col059
		},
	)
}

func Col060() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_060",
		func(w *models.WideRecordsDto) *string {
			return &w.Col060
		},
	)
}

func Col061() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_061",
		func(w *models.WideRecordsDto) *string {
			return &w.Col061
		},
	)
}

func Col062() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_062",
		func(w *models.WideRecordsDto) *string {
			return &w.Col062
		},
	)
}

func Col063() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_063",
		func(w *models.WideRecordsDto) *string {
			return &w.Col063
		},
	)
}

func Col064() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_064",
		func(w *models.WideRecordsDto) *string {
			return &w.Col064
		},
	)
}

func Col065() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_065",
		func(w *models.WideRecordsDto) *string {
			return &w.Col065
		},
	)
}

func Col066() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_066",
		func(w *models.WideRecordsDto) *string {
			return &w.Col066
		},
	)
}

func Col067() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_067",
		func(w *models.WideRecordsDto) *string {
			return &w.Col067
		},
	)
}

func Col068() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_068",
		func(w *models.WideRecordsDto) *string {
			return &w.Col068
		},
	)
}

func Col069() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_069",
		func(w *models.WideRecordsDto) *string {
			return &w.Col069
		},
	)
}

func Col070() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_070",
		func(w *models.WideRecordsDto) *string {
			return &w.Col070
		},
	)
}

func Col071() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_071",
		func(w *models.WideRecordsDto) *string {
			return &w.Col071
		},
	)
}

func Col072() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_072",
		func(w *models.WideRecordsDto) *string {
			return &w.Col072
		},
	)
}

func Col073() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_073",
		func(w *models.WideRecordsDto) *string {
			return &w.Col073
		},
	)
}

func Col074() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_074",
		func(w *models.WideRecordsDto) *string {
			return &w.Col074
		},
	)
}

func Col075() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_075",
		func(w *models.WideRecordsDto) *string {
			return &w.Col075
		},
	)
}

func Col076() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_076",
		func(w *models.WideRecordsDto) *string {
			return &w.Col076
		},
	)
}

func Col077() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_077",
		func(w *models.WideRecordsDto) *string {
			return &w.Col077
		},
	)
}

func Col078() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_078",
		func(w *models.WideRecordsDto) *string {
			return &w.Col078
		},
	)
}

func Col079() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_079",
		func(w *models.WideRecordsDto) *string {
			return &w.Col079
		},
	)
}

func Col080() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_080",
		func(w *models.WideRecordsDto) *string {
			return &w.Col080
		},
	)
}

func Col081() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_081",
		func(w *models.WideRecordsDto) *string {
			return &w.Col081
		},
	)
}

func Col082() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_082",
		func(w *models.WideRecordsDto) *string {
			return &w.Col082
		},
	)
}

func Col083() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_083",
		func(w *models.WideRecordsDto) *string {
			return &w.Col083
		},
	)
}

func Col084() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_084",
		func(w *models.WideRecordsDto) *string {
			return &w.Col084
		},
	)
}

func Col085() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_085",
		func(w *models.WideRecordsDto) *string {
			return &w.Col085
		},
	)
}

func Col086() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_086",
		func(w *models.WideRecordsDto) *string {
			return &w.Col086
		},
	)
}

func Col087() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_087",
		func(w *models.WideRecordsDto) *string {
			return &w.Col087
		},
	)
}

func Col088() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_088",
		func(w *models.WideRecordsDto) *string {
			return &w.Col088
		},
	)
}

func Col089() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_089",
		func(w *models.WideRecordsDto) *string {
			return &w.Col089
		},
	)
}

func Col090() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_090",
		func(w *models.WideRecordsDto) *string {
			return &w.Col090
		},
	)
}

func Col091() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_091",
		func(w *models.WideRecordsDto) *string {
			return &w.Col091
		},
	)
}

func Col092() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_092",
		func(w *models.WideRecordsDto) *string {
			return &w.Col092
		},
	)
}

func Col093() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_093",
		func(w *models.WideRecordsDto) *string {
			return &w.Col093
		},
	)
}

func Col094() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_094",
		func(w *models.WideRecordsDto) *string {
			return &w.Col094
		},
	)
}

func Col095() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_095",
		func(w *models.WideRecordsDto) *string {
			return &w.Col095
		},
	)
}

func Col096() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_096",
		func(w *models.WideRecordsDto) *string {
			return &w.Col096
		},
	)
}

func Col097() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_097",
		func(w *models.WideRecordsDto) *string {
			return &w.Col097
		},
	)
}

func Col098() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_098",
		func(w *models.WideRecordsDto) *string {
			return &w.Col098
		},
	)
}

func Col099() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_099",
		func(w *models.WideRecordsDto) *string {
			return &w.Col099
		},
	)
}

func Col100() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return ast.NewStringColumnProjection(
		"wide_records",
		"col_100",
		func(w *models.WideRecordsDto) *string {
			return &w.Col100
		},
	)
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Col001(),
		Col002(),
		Col003(),
		Col004(),
		Col005(),
		Col006(),
		Col007(),
		Col008(),
		Col009(),
		Col010(),
		Col011(),
		Col012(),
		Col013(),
		Col014(),
		Col015(),
		Col016(),
		Col017(),
		Col018(),
		Col019(),
		Col020(),
		Col021(),
		Col022(),
		Col023(),
		Col024(),
		Col025(),
		Col026(),
		Col027(),
		Col028(),
		Col029(),
		Col030(),
		Col031(),
		Col032(),
		Col033(),
		Col034(),
		Col035(),
		Col036(),
		Col037(),
		Col038(),
		Col039(),
		Col040(),
		Col041(),
		Col042(),
		Col043(),
		Col044(),
		Col045(),
		Col046(),
		Col047(),
		Col048(),
		Col049(),
		Col050(),
		Col051(),
		Col052(),
		Col053(),
		Col054(),
		Col055(),
		Col056(),
		Col057(),
		Col058(),
		Col059(),
		Col060(),
		Col061(),
		Col062(),
		Col063(),
		Col064(),
		Col065(),
		Col066(),
		Col067(),
		Col068(),
		Col069(),
		Col070(),
		Col071(),
		Col072(),
		Col073(),
		Col074(),
		Col075(),
		Col076(),
		Col077(),
		Col078(),
		Col079(),
		Col080(),
		Col081(),
		Col082(),
		Col083(),
		Col084(),
		Col085(),
		Col086(),
		Col087(),
		Col088(),
		Col089(),
		Col090(),
		Col091(),
		Col092(),
		Col093(),
		Col094(),
		Col095(),
		Col096(),
		Col097(),
		Col098(),
		Col099(),
		Col100(),
	}
}

func (intoWideRecordsDto) Id() ast.Projection[models.WideRecordsDto] {
	return Id()
}

func (intoWideRecordsDto) Col001() ast.Projection[models.WideRecordsDto] {
	return Col001()
}

func (intoWideRecordsDto) Col002() ast.Projection[models.WideRecordsDto] {
	return Col002()
}

func (intoWideRecordsDto) Col003() ast.Projection[models.WideRecordsDto] {
	return Col003()
}

func (intoWideRecordsDto) Col004() ast.Projection[models.WideRecordsDto] {
	return Col004()
}

func (intoWideRecordsDto) Col005() ast.Projection[models.WideRecordsDto] {
	return Col005()
}

func (intoWideRecordsDto) Col006() ast.Projection[models.WideRecordsDto] {
	return Col006()
}

func (intoWideRecordsDto) Col007() ast.Projection[models.WideRecordsDto] {
	return Col007()
}

func (intoWideRecordsDto) Col008() ast.Projection[models.WideRecordsDto] {
	return Col008()
}

func (intoWideRecordsDto) Col009() ast.Projection[models.WideRecordsDto] {
	return Col009()
}

func (intoWideRecordsDto) Col010() ast.Projection[models.WideRecordsDto] {
	return Col010()
}

func (intoWideRecordsDto) Col011() ast.Projection[models.WideRecordsDto] {
	return Col011()
}

func (intoWideRecordsDto) Col012() ast.Projection[models.WideRecordsDto] {
	return Col012()
}

func (intoWideRecordsDto) Col013() ast.Projection[models.WideRecordsDto] {
	return Col013()
}

func (intoWideRecordsDto) Col014() ast.Projection[models.WideRecordsDto] {
	return Col014()
}

func (intoWideRecordsDto) Col015() ast.Projection[models.WideRecordsDto] {
	return Col015()
}

func (intoWideRecordsDto) Col016() ast.Projection[models.WideRecordsDto] {
	return Col016()
}

func (intoWideRecordsDto) Col017() ast.Projection[models.WideRecordsDto] {
	return Col017()
}

func (intoWideRecordsDto) Col018() ast.Projection[models.WideRecordsDto] {
	return Col018()
}

func (intoWideRecordsDto) Col019() ast.Projection[models.WideRecordsDto] {
	return Col019()
}

func (intoWideRecordsDto) Col020() ast.Projection[models.WideRecordsDto] {
	return Col020()
}

func (intoWideRecordsDto) Col021() ast.Projection[models.WideRecordsDto] {
	return Col021()
}

func (intoWideRecordsDto) Col022() ast.Projection[models.WideRecordsDto] {
	return Col022()
}

func (intoWideRecordsDto) Col023() ast.Projection[models.WideRecordsDto] {
	return Col023()
}

func (intoWideRecordsDto) Col024() ast.Projection[models.WideRecordsDto] {
	return Col024()
}

func (intoWideRecordsDto) Col025() ast.Projection[models.WideRecordsDto] {
	return Col025()
}

func (intoWideRecordsDto) Col026() ast.Projection[models.WideRecordsDto] {
	return Col026()
}

func (intoWideRecordsDto) Col027() ast.Projection[models.WideRecordsDto] {
	return Col027()
}

func (intoWideRecordsDto) Col028() ast.Projection[models.WideRecordsDto] {
	return Col028()
}

func (intoWideRecordsDto) Col029() ast.Projection[models.WideRecordsDto] {
	return Col029()
}

func (intoWideRecordsDto) Col030() ast.Projection[models.WideRecordsDto] {
	return Col030()
}

func (intoWideRecordsDto) Col031() ast.Projection[models.WideRecordsDto] {
	return Col031()
}

func (intoWideRecordsDto) Col032() ast.Projection[models.WideRecordsDto] {
	return Col032()
}

func (intoWideRecordsDto) Col033() ast.Projection[models.WideRecordsDto] {
	return Col033()
}

func (intoWideRecordsDto) Col034() ast.Projection[models.WideRecordsDto] {
	return Col034()
}

func (intoWideRecordsDto) Col035() ast.Projection[models.WideRecordsDto] {
	return Col035()
}

func (intoWideRecordsDto) Col036() ast.Projection[models.WideRecordsDto] {
	return Col036()
}

func (intoWideRecordsDto) Col037() ast.Projection[models.WideRecordsDto] {
	return Col037()
}

func (intoWideRecordsDto) Col038() ast.Projection[models.WideRecordsDto] {
	return Col038()
}

func (intoWideRecordsDto) Col039() ast.Projection[models.WideRecordsDto] {
	return Col039()
}

func (intoWideRecordsDto) Col040() ast.Projection[models.WideRecordsDto] {
	return Col040()
}

func (intoWideRecordsDto) Col041() ast.Projection[models.WideRecordsDto] {
	return Col041()
}

func (intoWideRecordsDto) Col042() ast.Projection[models.WideRecordsDto] {
	return Col042()
}

func (intoWideRecordsDto) Col043() ast.Projection[models.WideRecordsDto] {
	return Col043()
}

func (intoWideRecordsDto) Col044() ast.Projection[models.WideRecordsDto] {
	return Col044()
}

func (intoWideRecordsDto) Col045() ast.Projection[models.WideRecordsDto] {
	return Col045()
}

func (intoWideRecordsDto) Col046() ast.Projection[models.WideRecordsDto] {
	return Col046()
}

func (intoWideRecordsDto) Col047() ast.Projection[models.WideRecordsDto] {
	return Col047()
}

func (intoWideRecordsDto) Col048() ast.Projection[models.WideRecordsDto] {
	return Col048()
}

func (intoWideRecordsDto) Col049() ast.Projection[models.WideRecordsDto] {
	return Col049()
}

func (intoWideRecordsDto) Col050() ast.Projection[models.WideRecordsDto] {
	return Col050()
}

func (intoWideRecordsDto) Col051() ast.Projection[models.WideRecordsDto] {
	return Col051()
}

func (intoWideRecordsDto) Col052() ast.Projection[models.WideRecordsDto] {
	return Col052()
}

func (intoWideRecordsDto) Col053() ast.Projection[models.WideRecordsDto] {
	return Col053()
}

func (intoWideRecordsDto) Col054() ast.Projection[models.WideRecordsDto] {
	return Col054()
}

func (intoWideRecordsDto) Col055() ast.Projection[models.WideRecordsDto] {
	return Col055()
}

func (intoWideRecordsDto) Col056() ast.Projection[models.WideRecordsDto] {
	return Col056()
}

func (intoWideRecordsDto) Col057() ast.Projection[models.WideRecordsDto] {
	return Col057()
}

func (intoWideRecordsDto) Col058() ast.Projection[models.WideRecordsDto] {
	return Col058()
}

func (intoWideRecordsDto) Col059() ast.Projection[models.WideRecordsDto] {
	return Col059()
}

func (intoWideRecordsDto) Col060() ast.Projection[models.WideRecordsDto] {
	return Col060()
}

func (intoWideRecordsDto) Col061() ast.Projection[models.WideRecordsDto] {
	return Col061()
}

func (intoWideRecordsDto) Col062() ast.Projection[models.WideRecordsDto] {
	return Col062()
}

func (intoWideRecordsDto) Col063() ast.Projection[models.WideRecordsDto] {
	return Col063()
}

func (intoWideRecordsDto) Col064() ast.Projection[models.WideRecordsDto] {
	return Col064()
}

func (intoWideRecordsDto) Col065() ast.Projection[models.WideRecordsDto] {
	return Col065()
}

func (intoWideRecordsDto) Col066() ast.Projection[models.WideRecordsDto] {
	return Col066()
}

func (intoWideRecordsDto) Col067() ast.Projection[models.WideRecordsDto] {
	return Col067()
}

func (intoWideRecordsDto) Col068() ast.Projection[models.WideRecordsDto] {
	return Col068()
}

func (intoWideRecordsDto) Col069() ast.Projection[models.WideRecordsDto] {
	return Col069()
}

func (intoWideRecordsDto) Col070() ast.Projection[models.WideRecordsDto] {
	return Col070()
}

func (intoWideRecordsDto) Col071() ast.Projection[models.WideRecordsDto] {
	return Col071()
}

func (intoWideRecordsDto) Col072() ast.Projection[models.WideRecordsDto] {
	return Col072()
}

func (intoWideRecordsDto) Col073() ast.Projection[models.WideRecordsDto] {
	return Col073()
}

func (intoWideRecordsDto) Col074() ast.Projection[models.WideRecordsDto] {
	return Col074()
}

func (intoWideRecordsDto) Col075() ast.Projection[models.WideRecordsDto] {
	return Col075()
}

func (intoWideRecordsDto) Col076() ast.Projection[models.WideRecordsDto] {
	return Col076()
}

func (intoWideRecordsDto) Col077() ast.Projection[models.WideRecordsDto] {
	return Col077()
}

func (intoWideRecordsDto) Col078() ast.Projection[models.WideRecordsDto] {
	return Col078()
}

func (intoWideRecordsDto) Col079() ast.Projection[models.WideRecordsDto] {
	return Col079()
}

func (intoWideRecordsDto) Col080() ast.Projection[models.WideRecordsDto] {
	return Col080()
}

func (intoWideRecordsDto) Col081() ast.Projection[models.WideRecordsDto] {
	return Col081()
}

func (intoWideRecordsDto) Col082() ast.Projection[models.WideRecordsDto] {
	return Col082()
}

func (intoWideRecordsDto) Col083() ast.Projection[models.WideRecordsDto] {
	return Col083()
}

func (intoWideRecordsDto) Col084() ast.Projection[models.WideRecordsDto] {
	return Col084()
}

func (intoWideRecordsDto) Col085() ast.Projection[models.WideRecordsDto] {
	return Col085()
}

func (intoWideRecordsDto) Col086() ast.Projection[models.WideRecordsDto] {
	return Col086()
}

func (intoWideRecordsDto) Col087() ast.Projection[models.WideRecordsDto] {
	return Col087()
}

func (intoWideRecordsDto) Col088() ast.Projection[models.WideRecordsDto] {
	return Col088()
}

func (intoWideRecordsDto) Col089() ast.Projection[models.WideRecordsDto] {
	return Col089()
}

func (intoWideRecordsDto) Col090() ast.Projection[models.WideRecordsDto] {
	return Col090()
}

func (intoWideRecordsDto) Col091() ast.Projection[models.WideRecordsDto] {
	return Col091()
}

func (intoWideRecordsDto) Col092() ast.Projection[models.WideRecordsDto] {
	return Col092()
}

func (intoWideRecordsDto) Col093() ast.Projection[models.WideRecordsDto] {
	return Col093()
}

func (intoWideRecordsDto) Col094() ast.Projection[models.WideRecordsDto] {
	return Col094()
}

func (intoWideRecordsDto) Col095() ast.Projection[models.WideRecordsDto] {
	return Col095()
}

func (intoWideRecordsDto) Col096() ast.Projection[models.WideRecordsDto] {
	return Col096()
}

func (intoWideRecordsDto) Col097() ast.Projection[models.WideRecordsDto] {
	return Col097()
}

func (intoWideRecordsDto) Col098() ast.Projection[models.WideRecordsDto] {
	return Col098()
}

func (intoWideRecordsDto) Col099() ast.Projection[models.WideRecordsDto] {
	return Col099()
}

func (intoWideRecordsDto) Col100() ast.Projection[models.WideRecordsDto] {
	return Col100()
}

func (intoWideRecordsDto) AllColumns() []ast.Projection[models.WideRecordsDto] {
	return []ast.Projection[models.WideRecordsDto]{
		Into.Id(),
		Into.Col001(),
		Into.Col002(),
		Into.Col003(),
		Into.Col004(),
		Into.Col005(),
		Into.Col006(),
		Into.Col007(),
		Into.Col008(),
		Into.Col009(),
		Into.Col010(),
		Into.Col011(),
		Into.Col012(),
		Into.Col013(),
		Into.Col014(),
		Into.Col015(),
		Into.Col016(),
		Into.Col017(),
		Into.Col018(),
		Into.Col019(),
		Into.Col020(),
		Into.Col021(),
		Into.Col022(),
		Into.Col023(),
		Into.Col024(),
		Into.Col025(),
		Into.Col026(),
		Into.Col027(),
		Into.Col028(),
		Into.Col029(),
		Into.Col030(),
		Into.Col031(),
		Into.Col032(),
		Into.Col033(),
		Into.Col034(),
		Into.Col035(),
		Into.Col036(),
		Into.Col037(),
		Into.Col038(),
		Into.Col039(),
		Into.Col040(),
		Into.Col041(),
		Into.Col042(),
		Into.Col043(),
		Into.Col044(),
		Into.Col045(),
		Into.Col046(),
		Into.Col047(),
		Into.Col048(),
		Into.Col049(),
		Into.Col050(),
		Into.Col051(),
		Into.Col052(),
		Into.Col053(),
		Into.Col054(),
		Into.Col055(),
		Into.Col056(),
		Into.Col057(),
		Into.Col058(),
		Into.Col059(),
		Into.Col060(),
		Into.Col061(),
		Into.Col062(),
		Into.Col063(),
		Into.Col064(),
		Into.Col065(),
		Into.Col066(),
		Into.Col067(),
		Into.Col068(),
		Into.Col069(),
		Into.Col070(),
		Into.Col071(),
		Into.Col072(),
		Into.Col073(),
		Into.Col074(),
		Into.Col075(),
		Into.Col076(),
		Into.Col077(),
		Into.Col078(),
		Into.Col079(),
		Into.Col080(),
		Into.Col081(),
		Into.Col082(),
		Into.Col083(),
		Into.Col084(),
		Into.Col085(),
		Into.Col086(),
		Into.Col087(),
		Into.Col088(),
		Into.Col089(),
		Into.Col090(),
		Into.Col091(),
		Into.Col092(),
		Into.Col093(),
		Into.Col094(),
		Into.Col095(),
		Into.Col096(),
		Into.Col097(),
		Into.Col098(),
		Into.Col099(),
		Into.Col100(),
	}
}

func (f forWideRecordsDto[E, T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	column := Id()
	projection := ast.NewUUIDColumnProjection[T, **uuid.UUID](
		f.tableName(),
		column.Name(),
		func(t *T) **uuid.UUID {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Id
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", f.alias.Alias.Name()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col001() *ast.StringColumnProjection[T, *string] {
	column := Col001()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col001
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_001' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col002() *ast.StringColumnProjection[T, *string] {
	column := Col002()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col002
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_002' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col003() *ast.StringColumnProjection[T, *string] {
	column := Col003()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col003
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_003' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col004() *ast.StringColumnProjection[T, *string] {
	column := Col004()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col004
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_004' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col005() *ast.StringColumnProjection[T, *string] {
	column := Col005()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col005
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_005' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col006() *ast.StringColumnProjection[T, *string] {
	column := Col006()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col006
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_006' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col007() *ast.StringColumnProjection[T, *string] {
	column := Col007()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col007
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_007' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col008() *ast.StringColumnProjection[T, *string] {
	column := Col008()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col008
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_008' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col009() *ast.StringColumnProjection[T, *string] {
	column := Col009()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col009
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_009' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col010() *ast.StringColumnProjection[T, *string] {
	column := Col010()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col010
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_010' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col011() *ast.StringColumnProjection[T, *string] {
	column := Col011()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col011
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_011' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col012() *ast.StringColumnProjection[T, *string] {
	column := Col012()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col012
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_012' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col013() *ast.StringColumnProjection[T, *string] {
	column := Col013()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col013
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_013' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col014() *ast.StringColumnProjection[T, *string] {
	column := Col014()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col014
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_014' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col015() *ast.StringColumnProjection[T, *string] {
	column := Col015()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col015
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_015' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col016() *ast.StringColumnProjection[T, *string] {
	column := Col016()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col016
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_016' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col017() *ast.StringColumnProjection[T, *string] {
	column := Col017()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col017
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_017' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col018() *ast.StringColumnProjection[T, *string] {
	column := Col018()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col018
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_018' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col019() *ast.StringColumnProjection[T, *string] {
	column := Col019()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col019
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_019' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col020() *ast.StringColumnProjection[T, *string] {
	column := Col020()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col020
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_020' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col021() *ast.StringColumnProjection[T, *string] {
	column := Col021()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col021
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_021' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col022() *ast.StringColumnProjection[T, *string] {
	column := Col022()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col022
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_022' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col023() *ast.StringColumnProjection[T, *string] {
	column := Col023()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col023
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_023' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col024() *ast.StringColumnProjection[T, *string] {
	column := Col024()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col024
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_024' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col025() *ast.StringColumnProjection[T, *string] {
	column := Col025()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col025
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_025' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col026() *ast.StringColumnProjection[T, *string] {
	column := Col026()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col026
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_026' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col027() *ast.StringColumnProjection[T, *string] {
	column := Col027()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col027
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_027' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col028() *ast.StringColumnProjection[T, *string] {
	column := Col028()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col028
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_028' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col029() *ast.StringColumnProjection[T, *string] {
	column := Col029()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col029
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_029' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col030() *ast.StringColumnProjection[T, *string] {
	column := Col030()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col030
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_030' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col031() *ast.StringColumnProjection[T, *string] {
	column := Col031()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col031
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_031' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col032() *ast.StringColumnProjection[T, *string] {
	column := Col032()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col032
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_032' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col033() *ast.StringColumnProjection[T, *string] {
	column := Col033()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col033
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_033' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col034() *ast.StringColumnProjection[T, *string] {
	column := Col034()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col034
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_034' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col035() *ast.StringColumnProjection[T, *string] {
	column := Col035()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col035
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_035' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col036() *ast.StringColumnProjection[T, *string] {
	column := Col036()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col036
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_036' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col037() *ast.StringColumnProjection[T, *string] {
	column := Col037()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col037
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_037' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col038() *ast.StringColumnProjection[T, *string] {
	column := Col038()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col038
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_038' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col039() *ast.StringColumnProjection[T, *string] {
	column := Col039()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col039
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_039' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col040() *ast.StringColumnProjection[T, *string] {
	column := Col040()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col040
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_040' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col041() *ast.StringColumnProjection[T, *string] {
	column := Col041()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col041
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_041' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col042() *ast.StringColumnProjection[T, *string] {
	column := Col042()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col042
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_042' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col043() *ast.StringColumnProjection[T, *string] {
	column := Col043()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col043
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_043' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col044() *ast.StringColumnProjection[T, *string] {
	column := Col044()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col044
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_044' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col045() *ast.StringColumnProjection[T, *string] {
	column := Col045()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col045
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_045' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col046() *ast.StringColumnProjection[T, *string] {
	column := Col046()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col046
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_046' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col047() *ast.StringColumnProjection[T, *string] {
	column := Col047()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col047
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_047' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col048() *ast.StringColumnProjection[T, *string] {
	column := Col048()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col048
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_048' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col049() *ast.StringColumnProjection[T, *string] {
	column := Col049()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col049
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_049' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col050() *ast.StringColumnProjection[T, *string] {
	column := Col050()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col050
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_050' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col051() *ast.StringColumnProjection[T, *string] {
	column := Col051()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col051
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_051' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col052() *ast.StringColumnProjection[T, *string] {
	column := Col052()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col052
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_052' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col053() *ast.StringColumnProjection[T, *string] {
	column := Col053()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col053
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_053' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col054() *ast.StringColumnProjection[T, *string] {
	column := Col054()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col054
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_054' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col055() *ast.StringColumnProjection[T, *string] {
	column := Col055()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col055
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_055' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col056() *ast.StringColumnProjection[T, *string] {
	column := Col056()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col056
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_056' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col057() *ast.StringColumnProjection[T, *string] {
	column := Col057()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col057
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_057' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col058() *ast.StringColumnProjection[T, *string] {
	column := Col058()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col058
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_058' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col059() *ast.StringColumnProjection[T, *string] {
	column := Col059()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col059
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_059' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col060() *ast.StringColumnProjection[T, *string] {
	column := Col060()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col060
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_060' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col061() *ast.StringColumnProjection[T, *string] {
	column := Col061()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col061
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_061' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col062() *ast.StringColumnProjection[T, *string] {
	column := Col062()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col062
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_062' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col063() *ast.StringColumnProjection[T, *string] {
	column := Col063()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col063
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_063' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col064() *ast.StringColumnProjection[T, *string] {
	column := Col064()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col064
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_064' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col065() *ast.StringColumnProjection[T, *string] {
	column := Col065()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col065
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_065' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col066() *ast.StringColumnProjection[T, *string] {
	column := Col066()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col066
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_066' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col067() *ast.StringColumnProjection[T, *string] {
	column := Col067()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col067
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_067' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col068() *ast.StringColumnProjection[T, *string] {
	column := Col068()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col068
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_068' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col069() *ast.StringColumnProjection[T, *string] {
	column := Col069()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col069
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_069' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col070() *ast.StringColumnProjection[T, *string] {
	column := Col070()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col070
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_070' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col071() *ast.StringColumnProjection[T, *string] {
	column := Col071()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col071
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_071' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col072() *ast.StringColumnProjection[T, *string] {
	column := Col072()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col072
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_072' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col073() *ast.StringColumnProjection[T, *string] {
	column := Col073()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col073
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_073' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col074() *ast.StringColumnProjection[T, *string] {
	column := Col074()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col074
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_074' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col075() *ast.StringColumnProjection[T, *string] {
	column := Col075()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col075
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_075' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col076() *ast.StringColumnProjection[T, *string] {
	column := Col076()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col076
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_076' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col077() *ast.StringColumnProjection[T, *string] {
	column := Col077()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col077
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_077' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col078() *ast.StringColumnProjection[T, *string] {
	column := Col078()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col078
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_078' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col079() *ast.StringColumnProjection[T, *string] {
	column := Col079()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col079
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_079' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col080() *ast.StringColumnProjection[T, *string] {
	column := Col080()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col080
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_080' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col081() *ast.StringColumnProjection[T, *string] {
	column := Col081()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col081
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_081' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col082() *ast.StringColumnProjection[T, *string] {
	column := Col082()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col082
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_082' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col083() *ast.StringColumnProjection[T, *string] {
	column := Col083()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col083
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_083' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col084() *ast.StringColumnProjection[T, *string] {
	column := Col084()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col084
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_084' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col085() *ast.StringColumnProjection[T, *string] {
	column := Col085()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col085
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_085' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col086() *ast.StringColumnProjection[T, *string] {
	column := Col086()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col086
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_086' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col087() *ast.StringColumnProjection[T, *string] {
	column := Col087()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col087
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_087' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col088() *ast.StringColumnProjection[T, *string] {
	column := Col088()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col088
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_088' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col089() *ast.StringColumnProjection[T, *string] {
	column := Col089()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col089
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_089' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col090() *ast.StringColumnProjection[T, *string] {
	column := Col090()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col090
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_090' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col091() *ast.StringColumnProjection[T, *string] {
	column := Col091()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col091
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_091' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col092() *ast.StringColumnProjection[T, *string] {
	column := Col092()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col092
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_092' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col093() *ast.StringColumnProjection[T, *string] {
	column := Col093()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col093
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_093' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col094() *ast.StringColumnProjection[T, *string] {
	column := Col094()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col094
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_094' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col095() *ast.StringColumnProjection[T, *string] {
	column := Col095()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col095
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_095' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col096() *ast.StringColumnProjection[T, *string] {
	column := Col096()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col096
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_096' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col097() *ast.StringColumnProjection[T, *string] {
	column := Col097()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col097
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_097' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col098() *ast.StringColumnProjection[T, *string] {
	column := Col098()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col098
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_098' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col099() *ast.StringColumnProjection[T, *string] {
	column := Col099()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col099
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_099' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) Col100() *ast.StringColumnProjection[T, *string] {
	column := Col100()
	projection := ast.NewStringColumnProjection[T, *string](
		f.tableName(),
		column.Name(),
		func(t *T) *string {
			e := E(t)
			dto := e.GetWideRecordsDto()
			return &dto.Col100
		},
	)
	if f.err != nil {
		errExpr := ast.NewErrorExpression(f.err)
		projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
		return projection
	}
	if f.alias == nil || slices.ContainsFunc(f.alias.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_100' in alias %s", f.alias.Alias.Name()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return projection
}

func (f forWideRecordsDto[E, T]) tableName() string {
	if f.alias == nil || f.alias.Alias == nil {
		return "wide_records"
	}
	return f.alias.Alias.Name()
}

func (f forWideRecordsDto[E, T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		f.Id(),
		f.Col001(),
		f.Col002(),
		f.Col003(),
		f.Col004(),
		f.Col005(),
		f.Col006(),
		f.Col007(),
		f.Col008(),
		f.Col009(),
		f.Col010(),
		f.Col011(),
		f.Col012(),
		f.Col013(),
		f.Col014(),
		f.Col015(),
		f.Col016(),
		f.Col017(),
		f.Col018(),
		f.Col019(),
		f.Col020(),
		f.Col021(),
		f.Col022(),
		f.Col023(),
		f.Col024(),
		f.Col025(),
		f.Col026(),
		f.Col027(),
		f.Col028(),
		f.Col029(),
		f.Col030(),
		f.Col031(),
		f.Col032(),
		f.Col033(),
		f.Col034(),
		f.Col035(),
		f.Col036(),
		f.Col037(),
		f.Col038(),
		f.Col039(),
		f.Col040(),
		f.Col041(),
		f.Col042(),
		f.Col043(),
		f.Col044(),
		f.Col045(),
		f.Col046(),
		f.Col047(),
		f.Col048(),
		f.Col049(),
		f.Col050(),
		f.Col051(),
		f.Col052(),
		f.Col053(),
		f.Col054(),
		f.Col055(),
		f.Col056(),
		f.Col057(),
		f.Col058(),
		f.Col059(),
		f.Col060(),
		f.Col061(),
		f.Col062(),
		f.Col063(),
		f.Col064(),
		f.Col065(),
		f.Col066(),
		f.Col067(),
		f.Col068(),
		f.Col069(),
		f.Col070(),
		f.Col071(),
		f.Col072(),
		f.Col073(),
		f.Col074(),
		f.Col075(),
		f.Col076(),
		f.Col077(),
		f.Col078(),
		f.Col079(),
		f.Col080(),
		f.Col081(),
		f.Col082(),
		f.Col083(),
		f.Col084(),
		f.Col085(),
		f.Col086(),
		f.Col087(),
		f.Col088(),
		f.Col089(),
		f.Col090(),
		f.Col091(),
		f.Col092(),
		f.Col093(),
		f.Col094(),
		f.Col095(),
		f.Col096(),
		f.Col097(),
		f.Col098(),
		f.Col099(),
		f.Col100(),
	}
}

type Alias struct {
	*ast.Alias
}

func As(name string, columns ...ast.NamedExpression) *Alias {
	return &Alias{
		Alias: ast.NewAlias(name, columns...),
	}
}

func (a *Alias) Id() *ast.UUIDColumnProjection[models.WideRecordsDto, **uuid.UUID] {
	column := Id()
	alias := ast.NewUUIDColumnProjection[models.WideRecordsDto, **uuid.UUID](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) **uuid.UUID {
			return &w.Id
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.Alias.Name()))
	alias.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col001() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col001()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col001
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_001' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col002() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col002()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col002
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_002' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col003() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col003()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col003
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_003' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col004() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col004()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col004
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_004' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col005() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col005()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col005
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_005' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col006() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col006()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col006
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_006' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col007() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col007()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col007
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_007' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col008() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col008()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col008
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_008' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col009() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col009()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col009
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_009' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col010() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col010()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col010
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_010' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col011() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col011()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col011
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_011' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col012() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col012()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col012
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_012' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col013() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col013()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col013
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_013' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col014() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col014()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col014
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_014' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col015() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col015()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col015
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_015' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col016() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col016()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col016
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_016' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col017() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col017()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col017
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_017' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col018() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col018()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col018
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_018' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col019() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col019()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col019
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_019' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col020() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col020()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col020
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_020' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col021() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col021()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col021
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_021' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col022() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col022()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col022
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_022' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col023() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col023()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col023
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_023' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col024() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col024()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col024
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_024' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col025() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col025()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col025
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_025' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col026() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col026()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col026
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_026' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col027() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col027()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col027
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_027' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col028() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col028()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col028
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_028' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col029() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col029()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col029
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_029' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col030() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col030()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col030
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_030' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col031() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col031()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col031
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_031' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col032() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col032()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col032
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_032' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col033() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col033()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col033
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_033' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col034() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col034()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col034
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_034' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col035() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col035()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col035
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_035' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col036() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col036()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col036
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_036' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col037() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col037()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col037
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_037' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col038() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col038()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col038
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_038' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col039() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col039()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col039
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_039' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col040() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col040()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col040
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_040' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col041() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col041()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col041
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_041' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col042() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col042()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col042
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_042' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col043() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col043()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col043
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_043' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col044() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col044()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col044
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_044' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col045() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col045()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col045
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_045' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col046() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col046()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col046
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_046' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col047() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col047()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col047
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_047' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col048() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col048()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col048
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_048' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col049() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col049()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col049
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_049' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col050() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col050()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col050
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_050' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col051() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col051()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col051
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_051' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col052() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col052()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col052
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_052' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col053() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col053()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col053
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_053' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col054() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col054()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col054
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_054' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col055() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col055()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col055
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_055' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col056() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col056()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col056
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_056' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col057() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col057()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col057
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_057' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col058() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col058()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col058
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_058' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col059() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col059()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col059
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_059' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col060() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col060()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col060
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_060' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col061() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col061()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col061
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_061' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col062() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col062()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col062
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_062' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col063() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col063()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col063
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_063' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col064() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col064()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col064
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_064' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col065() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col065()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col065
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_065' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col066() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col066()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col066
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_066' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col067() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col067()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col067
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_067' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col068() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col068()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col068
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_068' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col069() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col069()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col069
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_069' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col070() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col070()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col070
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_070' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col071() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col071()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col071
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_071' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col072() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col072()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col072
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_072' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col073() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col073()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col073
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_073' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col074() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col074()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col074
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_074' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col075() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col075()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col075
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_075' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col076() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col076()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col076
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_076' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col077() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col077()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col077
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_077' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col078() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col078()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col078
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_078' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col079() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col079()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col079
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_079' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col080() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col080()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col080
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_080' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col081() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col081()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col081
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_081' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col082() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col082()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col082
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_082' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col083() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col083()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col083
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_083' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col084() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col084()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col084
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_084' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col085() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col085()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col085
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_085' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col086() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col086()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col086
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_086' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col087() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col087()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col087
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_087' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col088() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col088()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col088
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_088' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col089() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col089()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col089
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_089' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col090() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col090()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col090
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_090' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col091() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col091()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col091
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_091' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col092() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col092()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col092
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_092' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col093() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col093()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col093
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_093' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col094() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col094()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col094
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_094' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col095() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col095()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col095
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_095' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col096() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col096()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col096
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_096' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col097() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col097()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col097
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_097' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col098() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col098()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col098
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_098' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col099() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col099()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col099
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_099' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}

func (a *Alias) Col100() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	column := Col100()
	alias := ast.NewStringColumnProjection[models.WideRecordsDto, *string](
		a.Alias.Name(),
		column.Name(),
		func(w *models.WideRecordsDto) *string {
			return &w.Col100
		},
	)
	if slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.Name() == column.Name()
	}) {
		return alias
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_100' in alias %s", a.Alias.Name()))
	alias.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(column.Name(), errExpr)
	return alias
}
