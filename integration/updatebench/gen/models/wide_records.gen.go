package models

import (
	"strings"

	"github.com/google/uuid"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
)

type WideRecordsDtos []*WideRecordsDto

func (w WideRecordsDtos) Id() ast.OfType[[]uuid.UUID] {
	out := make([]*uuid.UUID, len(w))
	for idx, dto := range w {
		out[idx] = dto.Id
	}
	return ast.SetType[[]uuid.UUID](ast.NewLiteralExpression(out))
}

func (w WideRecordsDtos) Col001() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col001
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col002() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col002
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col003() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col003
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col004() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col004
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col005() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col005
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col006() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col006
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col007() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col007
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col008() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col008
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col009() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col009
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col010() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col010
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col011() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col011
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col012() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col012
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col013() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col013
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col014() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col014
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col015() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col015
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col016() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col016
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col017() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col017
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col018() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col018
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col019() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col019
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col020() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col020
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col021() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col021
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col022() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col022
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col023() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col023
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col024() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col024
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col025() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col025
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col026() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col026
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col027() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col027
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col028() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col028
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col029() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col029
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col030() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col030
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col031() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col031
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col032() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col032
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col033() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col033
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col034() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col034
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col035() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col035
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col036() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col036
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col037() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col037
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col038() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col038
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col039() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col039
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col040() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col040
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col041() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col041
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col042() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col042
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col043() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col043
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col044() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col044
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col045() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col045
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col046() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col046
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col047() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col047
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col048() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col048
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col049() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col049
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col050() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col050
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col051() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col051
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col052() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col052
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col053() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col053
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col054() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col054
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col055() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col055
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col056() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col056
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col057() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col057
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col058() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col058
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col059() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col059
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col060() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col060
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col061() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col061
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col062() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col062
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col063() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col063
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col064() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col064
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col065() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col065
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col066() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col066
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col067() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col067
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col068() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col068
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col069() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col069
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col070() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col070
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col071() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col071
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col072() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col072
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col073() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col073
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col074() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col074
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col075() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col075
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col076() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col076
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col077() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col077
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col078() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col078
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col079() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col079
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col080() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col080
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col081() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col081
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col082() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col082
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col083() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col083
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col084() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col084
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col085() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col085
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col086() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col086
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col087() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col087
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col088() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col088
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col089() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col089
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col090() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col090
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col091() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col091
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col092() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col092
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col093() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col093
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col094() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col094
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col095() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col095
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col096() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col096
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col097() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col097
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col098() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col098
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col099() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col099
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) Col100() ast.OfType[[]string] {
	out := make([]string, len(w))
	for idx, dto := range w {
		out[idx] = dto.Col100
	}
	return ast.NewSQLType(out)
}

func (w WideRecordsDtos) GetValues(columns ...ast.NamedExpression) []ast.Expression {
	im := make([]ast.Expression, 0, len(w)*len(columns))
	for _, column := range columns {
		columnLower := strings.ToLower(column.Name())

		if columnLower == "id" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Id))
			}
		}

		if columnLower == "col_001" || columnLower == "col001" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col001))
			}
		}

		if columnLower == "col_002" || columnLower == "col002" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col002))
			}
		}

		if columnLower == "col_003" || columnLower == "col003" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col003))
			}
		}

		if columnLower == "col_004" || columnLower == "col004" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col004))
			}
		}

		if columnLower == "col_005" || columnLower == "col005" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col005))
			}
		}

		if columnLower == "col_006" || columnLower == "col006" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col006))
			}
		}

		if columnLower == "col_007" || columnLower == "col007" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col007))
			}
		}

		if columnLower == "col_008" || columnLower == "col008" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col008))
			}
		}

		if columnLower == "col_009" || columnLower == "col009" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col009))
			}
		}

		if columnLower == "col_010" || columnLower == "col010" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col010))
			}
		}

		if columnLower == "col_011" || columnLower == "col011" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col011))
			}
		}

		if columnLower == "col_012" || columnLower == "col012" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col012))
			}
		}

		if columnLower == "col_013" || columnLower == "col013" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col013))
			}
		}

		if columnLower == "col_014" || columnLower == "col014" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col014))
			}
		}

		if columnLower == "col_015" || columnLower == "col015" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col015))
			}
		}

		if columnLower == "col_016" || columnLower == "col016" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col016))
			}
		}

		if columnLower == "col_017" || columnLower == "col017" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col017))
			}
		}

		if columnLower == "col_018" || columnLower == "col018" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col018))
			}
		}

		if columnLower == "col_019" || columnLower == "col019" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col019))
			}
		}

		if columnLower == "col_020" || columnLower == "col020" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col020))
			}
		}

		if columnLower == "col_021" || columnLower == "col021" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col021))
			}
		}

		if columnLower == "col_022" || columnLower == "col022" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col022))
			}
		}

		if columnLower == "col_023" || columnLower == "col023" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col023))
			}
		}

		if columnLower == "col_024" || columnLower == "col024" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col024))
			}
		}

		if columnLower == "col_025" || columnLower == "col025" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col025))
			}
		}

		if columnLower == "col_026" || columnLower == "col026" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col026))
			}
		}

		if columnLower == "col_027" || columnLower == "col027" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col027))
			}
		}

		if columnLower == "col_028" || columnLower == "col028" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col028))
			}
		}

		if columnLower == "col_029" || columnLower == "col029" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col029))
			}
		}

		if columnLower == "col_030" || columnLower == "col030" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col030))
			}
		}

		if columnLower == "col_031" || columnLower == "col031" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col031))
			}
		}

		if columnLower == "col_032" || columnLower == "col032" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col032))
			}
		}

		if columnLower == "col_033" || columnLower == "col033" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col033))
			}
		}

		if columnLower == "col_034" || columnLower == "col034" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col034))
			}
		}

		if columnLower == "col_035" || columnLower == "col035" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col035))
			}
		}

		if columnLower == "col_036" || columnLower == "col036" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col036))
			}
		}

		if columnLower == "col_037" || columnLower == "col037" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col037))
			}
		}

		if columnLower == "col_038" || columnLower == "col038" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col038))
			}
		}

		if columnLower == "col_039" || columnLower == "col039" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col039))
			}
		}

		if columnLower == "col_040" || columnLower == "col040" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col040))
			}
		}

		if columnLower == "col_041" || columnLower == "col041" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col041))
			}
		}

		if columnLower == "col_042" || columnLower == "col042" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col042))
			}
		}

		if columnLower == "col_043" || columnLower == "col043" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col043))
			}
		}

		if columnLower == "col_044" || columnLower == "col044" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col044))
			}
		}

		if columnLower == "col_045" || columnLower == "col045" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col045))
			}
		}

		if columnLower == "col_046" || columnLower == "col046" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col046))
			}
		}

		if columnLower == "col_047" || columnLower == "col047" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col047))
			}
		}

		if columnLower == "col_048" || columnLower == "col048" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col048))
			}
		}

		if columnLower == "col_049" || columnLower == "col049" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col049))
			}
		}

		if columnLower == "col_050" || columnLower == "col050" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col050))
			}
		}

		if columnLower == "col_051" || columnLower == "col051" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col051))
			}
		}

		if columnLower == "col_052" || columnLower == "col052" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col052))
			}
		}

		if columnLower == "col_053" || columnLower == "col053" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col053))
			}
		}

		if columnLower == "col_054" || columnLower == "col054" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col054))
			}
		}

		if columnLower == "col_055" || columnLower == "col055" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col055))
			}
		}

		if columnLower == "col_056" || columnLower == "col056" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col056))
			}
		}

		if columnLower == "col_057" || columnLower == "col057" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col057))
			}
		}

		if columnLower == "col_058" || columnLower == "col058" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col058))
			}
		}

		if columnLower == "col_059" || columnLower == "col059" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col059))
			}
		}

		if columnLower == "col_060" || columnLower == "col060" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col060))
			}
		}

		if columnLower == "col_061" || columnLower == "col061" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col061))
			}
		}

		if columnLower == "col_062" || columnLower == "col062" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col062))
			}
		}

		if columnLower == "col_063" || columnLower == "col063" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col063))
			}
		}

		if columnLower == "col_064" || columnLower == "col064" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col064))
			}
		}

		if columnLower == "col_065" || columnLower == "col065" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col065))
			}
		}

		if columnLower == "col_066" || columnLower == "col066" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col066))
			}
		}

		if columnLower == "col_067" || columnLower == "col067" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col067))
			}
		}

		if columnLower == "col_068" || columnLower == "col068" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col068))
			}
		}

		if columnLower == "col_069" || columnLower == "col069" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col069))
			}
		}

		if columnLower == "col_070" || columnLower == "col070" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col070))
			}
		}

		if columnLower == "col_071" || columnLower == "col071" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col071))
			}
		}

		if columnLower == "col_072" || columnLower == "col072" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col072))
			}
		}

		if columnLower == "col_073" || columnLower == "col073" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col073))
			}
		}

		if columnLower == "col_074" || columnLower == "col074" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col074))
			}
		}

		if columnLower == "col_075" || columnLower == "col075" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col075))
			}
		}

		if columnLower == "col_076" || columnLower == "col076" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col076))
			}
		}

		if columnLower == "col_077" || columnLower == "col077" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col077))
			}
		}

		if columnLower == "col_078" || columnLower == "col078" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col078))
			}
		}

		if columnLower == "col_079" || columnLower == "col079" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col079))
			}
		}

		if columnLower == "col_080" || columnLower == "col080" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col080))
			}
		}

		if columnLower == "col_081" || columnLower == "col081" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col081))
			}
		}

		if columnLower == "col_082" || columnLower == "col082" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col082))
			}
		}

		if columnLower == "col_083" || columnLower == "col083" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col083))
			}
		}

		if columnLower == "col_084" || columnLower == "col084" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col084))
			}
		}

		if columnLower == "col_085" || columnLower == "col085" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col085))
			}
		}

		if columnLower == "col_086" || columnLower == "col086" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col086))
			}
		}

		if columnLower == "col_087" || columnLower == "col087" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col087))
			}
		}

		if columnLower == "col_088" || columnLower == "col088" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col088))
			}
		}

		if columnLower == "col_089" || columnLower == "col089" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col089))
			}
		}

		if columnLower == "col_090" || columnLower == "col090" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col090))
			}
		}

		if columnLower == "col_091" || columnLower == "col091" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col091))
			}
		}

		if columnLower == "col_092" || columnLower == "col092" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col092))
			}
		}

		if columnLower == "col_093" || columnLower == "col093" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col093))
			}
		}

		if columnLower == "col_094" || columnLower == "col094" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col094))
			}
		}

		if columnLower == "col_095" || columnLower == "col095" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col095))
			}
		}

		if columnLower == "col_096" || columnLower == "col096" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col096))
			}
		}

		if columnLower == "col_097" || columnLower == "col097" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col097))
			}
		}

		if columnLower == "col_098" || columnLower == "col098" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col098))
			}
		}

		if columnLower == "col_099" || columnLower == "col099" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col099))
			}
		}

		if columnLower == "col_100" || columnLower == "col100" {
			for _, dto := range w {
				im = append(im, ast.NewLiteralExpression(dto.Col100))
			}
		}
	}

	values := make([]ast.Expression, 0, len(w))
	for i := 0; i < len(w); i++ {
		grouped := make([]ast.Expression, len(columns))
		for j := range columns {
			grouped[j] = im[j*len(w)+i]
		}

		values = append(values, ast.NewGroupedExpression(grouped...))
	}

	return values
}

// WideRecordsDto represents the wide_records table
type WideRecordsDto struct {
	Id     *uuid.UUID `db:"id" json:"id"`
	Col001 string     `db:"col_001" json:"col_001"`
	Col002 string     `db:"col_002" json:"col_002"`
	Col003 string     `db:"col_003" json:"col_003"`
	Col004 string     `db:"col_004" json:"col_004"`
	Col005 string     `db:"col_005" json:"col_005"`
	Col006 string     `db:"col_006" json:"col_006"`
	Col007 string     `db:"col_007" json:"col_007"`
	Col008 string     `db:"col_008" json:"col_008"`
	Col009 string     `db:"col_009" json:"col_009"`
	Col010 string     `db:"col_010" json:"col_010"`
	Col011 string     `db:"col_011" json:"col_011"`
	Col012 string     `db:"col_012" json:"col_012"`
	Col013 string     `db:"col_013" json:"col_013"`
	Col014 string     `db:"col_014" json:"col_014"`
	Col015 string     `db:"col_015" json:"col_015"`
	Col016 string     `db:"col_016" json:"col_016"`
	Col017 string     `db:"col_017" json:"col_017"`
	Col018 string     `db:"col_018" json:"col_018"`
	Col019 string     `db:"col_019" json:"col_019"`
	Col020 string     `db:"col_020" json:"col_020"`
	Col021 string     `db:"col_021" json:"col_021"`
	Col022 string     `db:"col_022" json:"col_022"`
	Col023 string     `db:"col_023" json:"col_023"`
	Col024 string     `db:"col_024" json:"col_024"`
	Col025 string     `db:"col_025" json:"col_025"`
	Col026 string     `db:"col_026" json:"col_026"`
	Col027 string     `db:"col_027" json:"col_027"`
	Col028 string     `db:"col_028" json:"col_028"`
	Col029 string     `db:"col_029" json:"col_029"`
	Col030 string     `db:"col_030" json:"col_030"`
	Col031 string     `db:"col_031" json:"col_031"`
	Col032 string     `db:"col_032" json:"col_032"`
	Col033 string     `db:"col_033" json:"col_033"`
	Col034 string     `db:"col_034" json:"col_034"`
	Col035 string     `db:"col_035" json:"col_035"`
	Col036 string     `db:"col_036" json:"col_036"`
	Col037 string     `db:"col_037" json:"col_037"`
	Col038 string     `db:"col_038" json:"col_038"`
	Col039 string     `db:"col_039" json:"col_039"`
	Col040 string     `db:"col_040" json:"col_040"`
	Col041 string     `db:"col_041" json:"col_041"`
	Col042 string     `db:"col_042" json:"col_042"`
	Col043 string     `db:"col_043" json:"col_043"`
	Col044 string     `db:"col_044" json:"col_044"`
	Col045 string     `db:"col_045" json:"col_045"`
	Col046 string     `db:"col_046" json:"col_046"`
	Col047 string     `db:"col_047" json:"col_047"`
	Col048 string     `db:"col_048" json:"col_048"`
	Col049 string     `db:"col_049" json:"col_049"`
	Col050 string     `db:"col_050" json:"col_050"`
	Col051 string     `db:"col_051" json:"col_051"`
	Col052 string     `db:"col_052" json:"col_052"`
	Col053 string     `db:"col_053" json:"col_053"`
	Col054 string     `db:"col_054" json:"col_054"`
	Col055 string     `db:"col_055" json:"col_055"`
	Col056 string     `db:"col_056" json:"col_056"`
	Col057 string     `db:"col_057" json:"col_057"`
	Col058 string     `db:"col_058" json:"col_058"`
	Col059 string     `db:"col_059" json:"col_059"`
	Col060 string     `db:"col_060" json:"col_060"`
	Col061 string     `db:"col_061" json:"col_061"`
	Col062 string     `db:"col_062" json:"col_062"`
	Col063 string     `db:"col_063" json:"col_063"`
	Col064 string     `db:"col_064" json:"col_064"`
	Col065 string     `db:"col_065" json:"col_065"`
	Col066 string     `db:"col_066" json:"col_066"`
	Col067 string     `db:"col_067" json:"col_067"`
	Col068 string     `db:"col_068" json:"col_068"`
	Col069 string     `db:"col_069" json:"col_069"`
	Col070 string     `db:"col_070" json:"col_070"`
	Col071 string     `db:"col_071" json:"col_071"`
	Col072 string     `db:"col_072" json:"col_072"`
	Col073 string     `db:"col_073" json:"col_073"`
	Col074 string     `db:"col_074" json:"col_074"`
	Col075 string     `db:"col_075" json:"col_075"`
	Col076 string     `db:"col_076" json:"col_076"`
	Col077 string     `db:"col_077" json:"col_077"`
	Col078 string     `db:"col_078" json:"col_078"`
	Col079 string     `db:"col_079" json:"col_079"`
	Col080 string     `db:"col_080" json:"col_080"`
	Col081 string     `db:"col_081" json:"col_081"`
	Col082 string     `db:"col_082" json:"col_082"`
	Col083 string     `db:"col_083" json:"col_083"`
	Col084 string     `db:"col_084" json:"col_084"`
	Col085 string     `db:"col_085" json:"col_085"`
	Col086 string     `db:"col_086" json:"col_086"`
	Col087 string     `db:"col_087" json:"col_087"`
	Col088 string     `db:"col_088" json:"col_088"`
	Col089 string     `db:"col_089" json:"col_089"`
	Col090 string     `db:"col_090" json:"col_090"`
	Col091 string     `db:"col_091" json:"col_091"`
	Col092 string     `db:"col_092" json:"col_092"`
	Col093 string     `db:"col_093" json:"col_093"`
	Col094 string     `db:"col_094" json:"col_094"`
	Col095 string     `db:"col_095" json:"col_095"`
	Col096 string     `db:"col_096" json:"col_096"`
	Col097 string     `db:"col_097" json:"col_097"`
	Col098 string     `db:"col_098" json:"col_098"`
	Col099 string     `db:"col_099" json:"col_099"`
	Col100 string     `db:"col_100" json:"col_100"`
}

type ExportsWideRecordsDto[T any] interface {
	*T
	GetWideRecordsDto() *WideRecordsDto
}
