package query

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/models"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/builder"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/wide_records"
)

type wideRecordsUpdateColumnExpr struct {
	value  ast.NamedExpression
	unnest ast.Expression
	set    ast.UpdateSetExpr
}

func wideRecordsDtoUpdateManyFast(pool *pgxpool.Pool, dtos models.WideRecordsDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.WideRecordsDto], error) {
	options := newUpdateOptions(opts...)
	valueColumns, unnestExpressions, sets, err := wideRecordsDtoUpdateManyExpressions(dtos, options)
	if err != nil {
		return nil, err
	}
	if len(sets) == 0 {
		sets = append(sets, ast.UpdateSetList{})
	}
	v := wide_records.As("v", valueColumns...)
	var from ast.NamedTableExpression = Unnest(unnestExpressions...).As(v.Alias)
	if options.source == updateBatchSourceValues {
		from = Values(dtos.GetValues(valueColumns...)...).As(v.Alias)
	}
	return wide_records.NewQuery(pool).
		Update(sets...).
		From(from).
		Where(ast.NewUUIDColumnExpression("wide_records", "id").Eq(ast.NewUUIDColumnExpression("v", "id"))), nil
}

func wideRecordsDtoUpdateManyExpressions(dtos models.WideRecordsDtos, options updateOptions) ([]ast.NamedExpression, []ast.Expression, []ast.UpdateSetExpr, error) {
	if !options.columnsSet {
		return wideRecordsDtoDefaultUpdateManyExpressions(dtos)
	}
	if len(options.columns) == 0 {
		return []ast.NamedExpression{ast.NewUUIDColumnExpression("wide_records", "id")}, []ast.Expression{Cast(dtos.Id()).AsUUIDArray()}, nil, nil
	}
	expressions := wideRecordsDtoUpdateManyExpressionMap(dtos)
	valueColumns := make([]ast.NamedExpression, 1, len(options.columns)+1)
	unnestExpressions := make([]ast.Expression, 1, len(options.columns)+1)
	sets := make([]ast.UpdateSetExpr, 0, len(options.columns))
	valueColumns[0] = ast.NewUUIDColumnExpression("wide_records", "id")
	unnestExpressions[0] = Cast(dtos.Id()).AsUUIDArray()
	seen := make(map[string]struct{}, len(options.columns))
	for _, column := range options.columns {
		if column == nil {
			return nil, nil, nil, fmt.Errorf("query.UpdateColumns received a nil column")
		}
		name := column.Name()
		if name == "id" {
			return nil, nil, nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		if _, ok := seen[name]; ok {
			continue
		}
		expr, ok := expressions[name]
		if !ok {
			return nil, nil, nil, fmt.Errorf("unknown update column %q", name)
		}
		seen[name] = struct{}{}
		valueColumns = append(valueColumns, expr.value)
		unnestExpressions = append(unnestExpressions, expr.unnest)
		sets = append(sets, expr.set)
	}
	return valueColumns, unnestExpressions, sets, nil
}

func wideRecordsDtoDefaultUpdateManyExpressions(dtos models.WideRecordsDtos) ([]ast.NamedExpression, []ast.Expression, []ast.UpdateSetExpr, error) {
	valueColumns := []ast.NamedExpression{
		ast.NewUUIDColumnExpression("wide_records", "id"),
		ast.NewStringColumnExpression("wide_records", "col_001"),
		ast.NewStringColumnExpression("wide_records", "col_002"),
		ast.NewStringColumnExpression("wide_records", "col_003"),
		ast.NewStringColumnExpression("wide_records", "col_004"),
		ast.NewStringColumnExpression("wide_records", "col_005"),
		ast.NewStringColumnExpression("wide_records", "col_006"),
		ast.NewStringColumnExpression("wide_records", "col_007"),
		ast.NewStringColumnExpression("wide_records", "col_008"),
		ast.NewStringColumnExpression("wide_records", "col_009"),
		ast.NewStringColumnExpression("wide_records", "col_010"),
		ast.NewStringColumnExpression("wide_records", "col_011"),
		ast.NewStringColumnExpression("wide_records", "col_012"),
		ast.NewStringColumnExpression("wide_records", "col_013"),
		ast.NewStringColumnExpression("wide_records", "col_014"),
		ast.NewStringColumnExpression("wide_records", "col_015"),
		ast.NewStringColumnExpression("wide_records", "col_016"),
		ast.NewStringColumnExpression("wide_records", "col_017"),
		ast.NewStringColumnExpression("wide_records", "col_018"),
		ast.NewStringColumnExpression("wide_records", "col_019"),
		ast.NewStringColumnExpression("wide_records", "col_020"),
		ast.NewStringColumnExpression("wide_records", "col_021"),
		ast.NewStringColumnExpression("wide_records", "col_022"),
		ast.NewStringColumnExpression("wide_records", "col_023"),
		ast.NewStringColumnExpression("wide_records", "col_024"),
		ast.NewStringColumnExpression("wide_records", "col_025"),
		ast.NewStringColumnExpression("wide_records", "col_026"),
		ast.NewStringColumnExpression("wide_records", "col_027"),
		ast.NewStringColumnExpression("wide_records", "col_028"),
		ast.NewStringColumnExpression("wide_records", "col_029"),
		ast.NewStringColumnExpression("wide_records", "col_030"),
		ast.NewStringColumnExpression("wide_records", "col_031"),
		ast.NewStringColumnExpression("wide_records", "col_032"),
		ast.NewStringColumnExpression("wide_records", "col_033"),
		ast.NewStringColumnExpression("wide_records", "col_034"),
		ast.NewStringColumnExpression("wide_records", "col_035"),
		ast.NewStringColumnExpression("wide_records", "col_036"),
		ast.NewStringColumnExpression("wide_records", "col_037"),
		ast.NewStringColumnExpression("wide_records", "col_038"),
		ast.NewStringColumnExpression("wide_records", "col_039"),
		ast.NewStringColumnExpression("wide_records", "col_040"),
		ast.NewStringColumnExpression("wide_records", "col_041"),
		ast.NewStringColumnExpression("wide_records", "col_042"),
		ast.NewStringColumnExpression("wide_records", "col_043"),
		ast.NewStringColumnExpression("wide_records", "col_044"),
		ast.NewStringColumnExpression("wide_records", "col_045"),
		ast.NewStringColumnExpression("wide_records", "col_046"),
		ast.NewStringColumnExpression("wide_records", "col_047"),
		ast.NewStringColumnExpression("wide_records", "col_048"),
		ast.NewStringColumnExpression("wide_records", "col_049"),
		ast.NewStringColumnExpression("wide_records", "col_050"),
		ast.NewStringColumnExpression("wide_records", "col_051"),
		ast.NewStringColumnExpression("wide_records", "col_052"),
		ast.NewStringColumnExpression("wide_records", "col_053"),
		ast.NewStringColumnExpression("wide_records", "col_054"),
		ast.NewStringColumnExpression("wide_records", "col_055"),
		ast.NewStringColumnExpression("wide_records", "col_056"),
		ast.NewStringColumnExpression("wide_records", "col_057"),
		ast.NewStringColumnExpression("wide_records", "col_058"),
		ast.NewStringColumnExpression("wide_records", "col_059"),
		ast.NewStringColumnExpression("wide_records", "col_060"),
		ast.NewStringColumnExpression("wide_records", "col_061"),
		ast.NewStringColumnExpression("wide_records", "col_062"),
		ast.NewStringColumnExpression("wide_records", "col_063"),
		ast.NewStringColumnExpression("wide_records", "col_064"),
		ast.NewStringColumnExpression("wide_records", "col_065"),
		ast.NewStringColumnExpression("wide_records", "col_066"),
		ast.NewStringColumnExpression("wide_records", "col_067"),
		ast.NewStringColumnExpression("wide_records", "col_068"),
		ast.NewStringColumnExpression("wide_records", "col_069"),
		ast.NewStringColumnExpression("wide_records", "col_070"),
		ast.NewStringColumnExpression("wide_records", "col_071"),
		ast.NewStringColumnExpression("wide_records", "col_072"),
		ast.NewStringColumnExpression("wide_records", "col_073"),
		ast.NewStringColumnExpression("wide_records", "col_074"),
		ast.NewStringColumnExpression("wide_records", "col_075"),
		ast.NewStringColumnExpression("wide_records", "col_076"),
		ast.NewStringColumnExpression("wide_records", "col_077"),
		ast.NewStringColumnExpression("wide_records", "col_078"),
		ast.NewStringColumnExpression("wide_records", "col_079"),
		ast.NewStringColumnExpression("wide_records", "col_080"),
		ast.NewStringColumnExpression("wide_records", "col_081"),
		ast.NewStringColumnExpression("wide_records", "col_082"),
		ast.NewStringColumnExpression("wide_records", "col_083"),
		ast.NewStringColumnExpression("wide_records", "col_084"),
		ast.NewStringColumnExpression("wide_records", "col_085"),
		ast.NewStringColumnExpression("wide_records", "col_086"),
		ast.NewStringColumnExpression("wide_records", "col_087"),
		ast.NewStringColumnExpression("wide_records", "col_088"),
		ast.NewStringColumnExpression("wide_records", "col_089"),
		ast.NewStringColumnExpression("wide_records", "col_090"),
		ast.NewStringColumnExpression("wide_records", "col_091"),
		ast.NewStringColumnExpression("wide_records", "col_092"),
		ast.NewStringColumnExpression("wide_records", "col_093"),
		ast.NewStringColumnExpression("wide_records", "col_094"),
		ast.NewStringColumnExpression("wide_records", "col_095"),
		ast.NewStringColumnExpression("wide_records", "col_096"),
		ast.NewStringColumnExpression("wide_records", "col_097"),
		ast.NewStringColumnExpression("wide_records", "col_098"),
		ast.NewStringColumnExpression("wide_records", "col_099"),
		ast.NewStringColumnExpression("wide_records", "col_100"),
	}
	unnestExpressions := []ast.Expression{
		Cast(dtos.Id()).AsUUIDArray(),
		Cast(dtos.Col001()).AsTextArray(),
		Cast(dtos.Col002()).AsTextArray(),
		Cast(dtos.Col003()).AsTextArray(),
		Cast(dtos.Col004()).AsTextArray(),
		Cast(dtos.Col005()).AsTextArray(),
		Cast(dtos.Col006()).AsTextArray(),
		Cast(dtos.Col007()).AsTextArray(),
		Cast(dtos.Col008()).AsTextArray(),
		Cast(dtos.Col009()).AsTextArray(),
		Cast(dtos.Col010()).AsTextArray(),
		Cast(dtos.Col011()).AsTextArray(),
		Cast(dtos.Col012()).AsTextArray(),
		Cast(dtos.Col013()).AsTextArray(),
		Cast(dtos.Col014()).AsTextArray(),
		Cast(dtos.Col015()).AsTextArray(),
		Cast(dtos.Col016()).AsTextArray(),
		Cast(dtos.Col017()).AsTextArray(),
		Cast(dtos.Col018()).AsTextArray(),
		Cast(dtos.Col019()).AsTextArray(),
		Cast(dtos.Col020()).AsTextArray(),
		Cast(dtos.Col021()).AsTextArray(),
		Cast(dtos.Col022()).AsTextArray(),
		Cast(dtos.Col023()).AsTextArray(),
		Cast(dtos.Col024()).AsTextArray(),
		Cast(dtos.Col025()).AsTextArray(),
		Cast(dtos.Col026()).AsTextArray(),
		Cast(dtos.Col027()).AsTextArray(),
		Cast(dtos.Col028()).AsTextArray(),
		Cast(dtos.Col029()).AsTextArray(),
		Cast(dtos.Col030()).AsTextArray(),
		Cast(dtos.Col031()).AsTextArray(),
		Cast(dtos.Col032()).AsTextArray(),
		Cast(dtos.Col033()).AsTextArray(),
		Cast(dtos.Col034()).AsTextArray(),
		Cast(dtos.Col035()).AsTextArray(),
		Cast(dtos.Col036()).AsTextArray(),
		Cast(dtos.Col037()).AsTextArray(),
		Cast(dtos.Col038()).AsTextArray(),
		Cast(dtos.Col039()).AsTextArray(),
		Cast(dtos.Col040()).AsTextArray(),
		Cast(dtos.Col041()).AsTextArray(),
		Cast(dtos.Col042()).AsTextArray(),
		Cast(dtos.Col043()).AsTextArray(),
		Cast(dtos.Col044()).AsTextArray(),
		Cast(dtos.Col045()).AsTextArray(),
		Cast(dtos.Col046()).AsTextArray(),
		Cast(dtos.Col047()).AsTextArray(),
		Cast(dtos.Col048()).AsTextArray(),
		Cast(dtos.Col049()).AsTextArray(),
		Cast(dtos.Col050()).AsTextArray(),
		Cast(dtos.Col051()).AsTextArray(),
		Cast(dtos.Col052()).AsTextArray(),
		Cast(dtos.Col053()).AsTextArray(),
		Cast(dtos.Col054()).AsTextArray(),
		Cast(dtos.Col055()).AsTextArray(),
		Cast(dtos.Col056()).AsTextArray(),
		Cast(dtos.Col057()).AsTextArray(),
		Cast(dtos.Col058()).AsTextArray(),
		Cast(dtos.Col059()).AsTextArray(),
		Cast(dtos.Col060()).AsTextArray(),
		Cast(dtos.Col061()).AsTextArray(),
		Cast(dtos.Col062()).AsTextArray(),
		Cast(dtos.Col063()).AsTextArray(),
		Cast(dtos.Col064()).AsTextArray(),
		Cast(dtos.Col065()).AsTextArray(),
		Cast(dtos.Col066()).AsTextArray(),
		Cast(dtos.Col067()).AsTextArray(),
		Cast(dtos.Col068()).AsTextArray(),
		Cast(dtos.Col069()).AsTextArray(),
		Cast(dtos.Col070()).AsTextArray(),
		Cast(dtos.Col071()).AsTextArray(),
		Cast(dtos.Col072()).AsTextArray(),
		Cast(dtos.Col073()).AsTextArray(),
		Cast(dtos.Col074()).AsTextArray(),
		Cast(dtos.Col075()).AsTextArray(),
		Cast(dtos.Col076()).AsTextArray(),
		Cast(dtos.Col077()).AsTextArray(),
		Cast(dtos.Col078()).AsTextArray(),
		Cast(dtos.Col079()).AsTextArray(),
		Cast(dtos.Col080()).AsTextArray(),
		Cast(dtos.Col081()).AsTextArray(),
		Cast(dtos.Col082()).AsTextArray(),
		Cast(dtos.Col083()).AsTextArray(),
		Cast(dtos.Col084()).AsTextArray(),
		Cast(dtos.Col085()).AsTextArray(),
		Cast(dtos.Col086()).AsTextArray(),
		Cast(dtos.Col087()).AsTextArray(),
		Cast(dtos.Col088()).AsTextArray(),
		Cast(dtos.Col089()).AsTextArray(),
		Cast(dtos.Col090()).AsTextArray(),
		Cast(dtos.Col091()).AsTextArray(),
		Cast(dtos.Col092()).AsTextArray(),
		Cast(dtos.Col093()).AsTextArray(),
		Cast(dtos.Col094()).AsTextArray(),
		Cast(dtos.Col095()).AsTextArray(),
		Cast(dtos.Col096()).AsTextArray(),
		Cast(dtos.Col097()).AsTextArray(),
		Cast(dtos.Col098()).AsTextArray(),
		Cast(dtos.Col099()).AsTextArray(),
		Cast(dtos.Col100()).AsTextArray(),
	}
	sets := []ast.UpdateSetExpr{
		Set(ast.NewStringColumnExpression("wide_records", "col_001")).To(ast.NewStringColumnExpression("v", "col_001")),
		Set(ast.NewStringColumnExpression("wide_records", "col_002")).To(ast.NewStringColumnExpression("v", "col_002")),
		Set(ast.NewStringColumnExpression("wide_records", "col_003")).To(ast.NewStringColumnExpression("v", "col_003")),
		Set(ast.NewStringColumnExpression("wide_records", "col_004")).To(ast.NewStringColumnExpression("v", "col_004")),
		Set(ast.NewStringColumnExpression("wide_records", "col_005")).To(ast.NewStringColumnExpression("v", "col_005")),
		Set(ast.NewStringColumnExpression("wide_records", "col_006")).To(ast.NewStringColumnExpression("v", "col_006")),
		Set(ast.NewStringColumnExpression("wide_records", "col_007")).To(ast.NewStringColumnExpression("v", "col_007")),
		Set(ast.NewStringColumnExpression("wide_records", "col_008")).To(ast.NewStringColumnExpression("v", "col_008")),
		Set(ast.NewStringColumnExpression("wide_records", "col_009")).To(ast.NewStringColumnExpression("v", "col_009")),
		Set(ast.NewStringColumnExpression("wide_records", "col_010")).To(ast.NewStringColumnExpression("v", "col_010")),
		Set(ast.NewStringColumnExpression("wide_records", "col_011")).To(ast.NewStringColumnExpression("v", "col_011")),
		Set(ast.NewStringColumnExpression("wide_records", "col_012")).To(ast.NewStringColumnExpression("v", "col_012")),
		Set(ast.NewStringColumnExpression("wide_records", "col_013")).To(ast.NewStringColumnExpression("v", "col_013")),
		Set(ast.NewStringColumnExpression("wide_records", "col_014")).To(ast.NewStringColumnExpression("v", "col_014")),
		Set(ast.NewStringColumnExpression("wide_records", "col_015")).To(ast.NewStringColumnExpression("v", "col_015")),
		Set(ast.NewStringColumnExpression("wide_records", "col_016")).To(ast.NewStringColumnExpression("v", "col_016")),
		Set(ast.NewStringColumnExpression("wide_records", "col_017")).To(ast.NewStringColumnExpression("v", "col_017")),
		Set(ast.NewStringColumnExpression("wide_records", "col_018")).To(ast.NewStringColumnExpression("v", "col_018")),
		Set(ast.NewStringColumnExpression("wide_records", "col_019")).To(ast.NewStringColumnExpression("v", "col_019")),
		Set(ast.NewStringColumnExpression("wide_records", "col_020")).To(ast.NewStringColumnExpression("v", "col_020")),
		Set(ast.NewStringColumnExpression("wide_records", "col_021")).To(ast.NewStringColumnExpression("v", "col_021")),
		Set(ast.NewStringColumnExpression("wide_records", "col_022")).To(ast.NewStringColumnExpression("v", "col_022")),
		Set(ast.NewStringColumnExpression("wide_records", "col_023")).To(ast.NewStringColumnExpression("v", "col_023")),
		Set(ast.NewStringColumnExpression("wide_records", "col_024")).To(ast.NewStringColumnExpression("v", "col_024")),
		Set(ast.NewStringColumnExpression("wide_records", "col_025")).To(ast.NewStringColumnExpression("v", "col_025")),
		Set(ast.NewStringColumnExpression("wide_records", "col_026")).To(ast.NewStringColumnExpression("v", "col_026")),
		Set(ast.NewStringColumnExpression("wide_records", "col_027")).To(ast.NewStringColumnExpression("v", "col_027")),
		Set(ast.NewStringColumnExpression("wide_records", "col_028")).To(ast.NewStringColumnExpression("v", "col_028")),
		Set(ast.NewStringColumnExpression("wide_records", "col_029")).To(ast.NewStringColumnExpression("v", "col_029")),
		Set(ast.NewStringColumnExpression("wide_records", "col_030")).To(ast.NewStringColumnExpression("v", "col_030")),
		Set(ast.NewStringColumnExpression("wide_records", "col_031")).To(ast.NewStringColumnExpression("v", "col_031")),
		Set(ast.NewStringColumnExpression("wide_records", "col_032")).To(ast.NewStringColumnExpression("v", "col_032")),
		Set(ast.NewStringColumnExpression("wide_records", "col_033")).To(ast.NewStringColumnExpression("v", "col_033")),
		Set(ast.NewStringColumnExpression("wide_records", "col_034")).To(ast.NewStringColumnExpression("v", "col_034")),
		Set(ast.NewStringColumnExpression("wide_records", "col_035")).To(ast.NewStringColumnExpression("v", "col_035")),
		Set(ast.NewStringColumnExpression("wide_records", "col_036")).To(ast.NewStringColumnExpression("v", "col_036")),
		Set(ast.NewStringColumnExpression("wide_records", "col_037")).To(ast.NewStringColumnExpression("v", "col_037")),
		Set(ast.NewStringColumnExpression("wide_records", "col_038")).To(ast.NewStringColumnExpression("v", "col_038")),
		Set(ast.NewStringColumnExpression("wide_records", "col_039")).To(ast.NewStringColumnExpression("v", "col_039")),
		Set(ast.NewStringColumnExpression("wide_records", "col_040")).To(ast.NewStringColumnExpression("v", "col_040")),
		Set(ast.NewStringColumnExpression("wide_records", "col_041")).To(ast.NewStringColumnExpression("v", "col_041")),
		Set(ast.NewStringColumnExpression("wide_records", "col_042")).To(ast.NewStringColumnExpression("v", "col_042")),
		Set(ast.NewStringColumnExpression("wide_records", "col_043")).To(ast.NewStringColumnExpression("v", "col_043")),
		Set(ast.NewStringColumnExpression("wide_records", "col_044")).To(ast.NewStringColumnExpression("v", "col_044")),
		Set(ast.NewStringColumnExpression("wide_records", "col_045")).To(ast.NewStringColumnExpression("v", "col_045")),
		Set(ast.NewStringColumnExpression("wide_records", "col_046")).To(ast.NewStringColumnExpression("v", "col_046")),
		Set(ast.NewStringColumnExpression("wide_records", "col_047")).To(ast.NewStringColumnExpression("v", "col_047")),
		Set(ast.NewStringColumnExpression("wide_records", "col_048")).To(ast.NewStringColumnExpression("v", "col_048")),
		Set(ast.NewStringColumnExpression("wide_records", "col_049")).To(ast.NewStringColumnExpression("v", "col_049")),
		Set(ast.NewStringColumnExpression("wide_records", "col_050")).To(ast.NewStringColumnExpression("v", "col_050")),
		Set(ast.NewStringColumnExpression("wide_records", "col_051")).To(ast.NewStringColumnExpression("v", "col_051")),
		Set(ast.NewStringColumnExpression("wide_records", "col_052")).To(ast.NewStringColumnExpression("v", "col_052")),
		Set(ast.NewStringColumnExpression("wide_records", "col_053")).To(ast.NewStringColumnExpression("v", "col_053")),
		Set(ast.NewStringColumnExpression("wide_records", "col_054")).To(ast.NewStringColumnExpression("v", "col_054")),
		Set(ast.NewStringColumnExpression("wide_records", "col_055")).To(ast.NewStringColumnExpression("v", "col_055")),
		Set(ast.NewStringColumnExpression("wide_records", "col_056")).To(ast.NewStringColumnExpression("v", "col_056")),
		Set(ast.NewStringColumnExpression("wide_records", "col_057")).To(ast.NewStringColumnExpression("v", "col_057")),
		Set(ast.NewStringColumnExpression("wide_records", "col_058")).To(ast.NewStringColumnExpression("v", "col_058")),
		Set(ast.NewStringColumnExpression("wide_records", "col_059")).To(ast.NewStringColumnExpression("v", "col_059")),
		Set(ast.NewStringColumnExpression("wide_records", "col_060")).To(ast.NewStringColumnExpression("v", "col_060")),
		Set(ast.NewStringColumnExpression("wide_records", "col_061")).To(ast.NewStringColumnExpression("v", "col_061")),
		Set(ast.NewStringColumnExpression("wide_records", "col_062")).To(ast.NewStringColumnExpression("v", "col_062")),
		Set(ast.NewStringColumnExpression("wide_records", "col_063")).To(ast.NewStringColumnExpression("v", "col_063")),
		Set(ast.NewStringColumnExpression("wide_records", "col_064")).To(ast.NewStringColumnExpression("v", "col_064")),
		Set(ast.NewStringColumnExpression("wide_records", "col_065")).To(ast.NewStringColumnExpression("v", "col_065")),
		Set(ast.NewStringColumnExpression("wide_records", "col_066")).To(ast.NewStringColumnExpression("v", "col_066")),
		Set(ast.NewStringColumnExpression("wide_records", "col_067")).To(ast.NewStringColumnExpression("v", "col_067")),
		Set(ast.NewStringColumnExpression("wide_records", "col_068")).To(ast.NewStringColumnExpression("v", "col_068")),
		Set(ast.NewStringColumnExpression("wide_records", "col_069")).To(ast.NewStringColumnExpression("v", "col_069")),
		Set(ast.NewStringColumnExpression("wide_records", "col_070")).To(ast.NewStringColumnExpression("v", "col_070")),
		Set(ast.NewStringColumnExpression("wide_records", "col_071")).To(ast.NewStringColumnExpression("v", "col_071")),
		Set(ast.NewStringColumnExpression("wide_records", "col_072")).To(ast.NewStringColumnExpression("v", "col_072")),
		Set(ast.NewStringColumnExpression("wide_records", "col_073")).To(ast.NewStringColumnExpression("v", "col_073")),
		Set(ast.NewStringColumnExpression("wide_records", "col_074")).To(ast.NewStringColumnExpression("v", "col_074")),
		Set(ast.NewStringColumnExpression("wide_records", "col_075")).To(ast.NewStringColumnExpression("v", "col_075")),
		Set(ast.NewStringColumnExpression("wide_records", "col_076")).To(ast.NewStringColumnExpression("v", "col_076")),
		Set(ast.NewStringColumnExpression("wide_records", "col_077")).To(ast.NewStringColumnExpression("v", "col_077")),
		Set(ast.NewStringColumnExpression("wide_records", "col_078")).To(ast.NewStringColumnExpression("v", "col_078")),
		Set(ast.NewStringColumnExpression("wide_records", "col_079")).To(ast.NewStringColumnExpression("v", "col_079")),
		Set(ast.NewStringColumnExpression("wide_records", "col_080")).To(ast.NewStringColumnExpression("v", "col_080")),
		Set(ast.NewStringColumnExpression("wide_records", "col_081")).To(ast.NewStringColumnExpression("v", "col_081")),
		Set(ast.NewStringColumnExpression("wide_records", "col_082")).To(ast.NewStringColumnExpression("v", "col_082")),
		Set(ast.NewStringColumnExpression("wide_records", "col_083")).To(ast.NewStringColumnExpression("v", "col_083")),
		Set(ast.NewStringColumnExpression("wide_records", "col_084")).To(ast.NewStringColumnExpression("v", "col_084")),
		Set(ast.NewStringColumnExpression("wide_records", "col_085")).To(ast.NewStringColumnExpression("v", "col_085")),
		Set(ast.NewStringColumnExpression("wide_records", "col_086")).To(ast.NewStringColumnExpression("v", "col_086")),
		Set(ast.NewStringColumnExpression("wide_records", "col_087")).To(ast.NewStringColumnExpression("v", "col_087")),
		Set(ast.NewStringColumnExpression("wide_records", "col_088")).To(ast.NewStringColumnExpression("v", "col_088")),
		Set(ast.NewStringColumnExpression("wide_records", "col_089")).To(ast.NewStringColumnExpression("v", "col_089")),
		Set(ast.NewStringColumnExpression("wide_records", "col_090")).To(ast.NewStringColumnExpression("v", "col_090")),
		Set(ast.NewStringColumnExpression("wide_records", "col_091")).To(ast.NewStringColumnExpression("v", "col_091")),
		Set(ast.NewStringColumnExpression("wide_records", "col_092")).To(ast.NewStringColumnExpression("v", "col_092")),
		Set(ast.NewStringColumnExpression("wide_records", "col_093")).To(ast.NewStringColumnExpression("v", "col_093")),
		Set(ast.NewStringColumnExpression("wide_records", "col_094")).To(ast.NewStringColumnExpression("v", "col_094")),
		Set(ast.NewStringColumnExpression("wide_records", "col_095")).To(ast.NewStringColumnExpression("v", "col_095")),
		Set(ast.NewStringColumnExpression("wide_records", "col_096")).To(ast.NewStringColumnExpression("v", "col_096")),
		Set(ast.NewStringColumnExpression("wide_records", "col_097")).To(ast.NewStringColumnExpression("v", "col_097")),
		Set(ast.NewStringColumnExpression("wide_records", "col_098")).To(ast.NewStringColumnExpression("v", "col_098")),
		Set(ast.NewStringColumnExpression("wide_records", "col_099")).To(ast.NewStringColumnExpression("v", "col_099")),
		Set(ast.NewStringColumnExpression("wide_records", "col_100")).To(ast.NewStringColumnExpression("v", "col_100")),
	}
	return valueColumns, unnestExpressions, sets, nil
}

func wideRecordsDtoUpdateManyExpressionMap(dtos models.WideRecordsDtos) map[string]wideRecordsUpdateColumnExpr {
	return map[string]wideRecordsUpdateColumnExpr{
		"col_001": {
			value:  ast.NewStringColumnExpression("wide_records", "col_001"),
			unnest: Cast(dtos.Col001()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_001")).To(ast.NewStringColumnExpression("v", "col_001")),
		},
		"col_002": {
			value:  ast.NewStringColumnExpression("wide_records", "col_002"),
			unnest: Cast(dtos.Col002()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_002")).To(ast.NewStringColumnExpression("v", "col_002")),
		},
		"col_003": {
			value:  ast.NewStringColumnExpression("wide_records", "col_003"),
			unnest: Cast(dtos.Col003()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_003")).To(ast.NewStringColumnExpression("v", "col_003")),
		},
		"col_004": {
			value:  ast.NewStringColumnExpression("wide_records", "col_004"),
			unnest: Cast(dtos.Col004()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_004")).To(ast.NewStringColumnExpression("v", "col_004")),
		},
		"col_005": {
			value:  ast.NewStringColumnExpression("wide_records", "col_005"),
			unnest: Cast(dtos.Col005()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_005")).To(ast.NewStringColumnExpression("v", "col_005")),
		},
		"col_006": {
			value:  ast.NewStringColumnExpression("wide_records", "col_006"),
			unnest: Cast(dtos.Col006()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_006")).To(ast.NewStringColumnExpression("v", "col_006")),
		},
		"col_007": {
			value:  ast.NewStringColumnExpression("wide_records", "col_007"),
			unnest: Cast(dtos.Col007()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_007")).To(ast.NewStringColumnExpression("v", "col_007")),
		},
		"col_008": {
			value:  ast.NewStringColumnExpression("wide_records", "col_008"),
			unnest: Cast(dtos.Col008()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_008")).To(ast.NewStringColumnExpression("v", "col_008")),
		},
		"col_009": {
			value:  ast.NewStringColumnExpression("wide_records", "col_009"),
			unnest: Cast(dtos.Col009()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_009")).To(ast.NewStringColumnExpression("v", "col_009")),
		},
		"col_010": {
			value:  ast.NewStringColumnExpression("wide_records", "col_010"),
			unnest: Cast(dtos.Col010()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_010")).To(ast.NewStringColumnExpression("v", "col_010")),
		},
		"col_011": {
			value:  ast.NewStringColumnExpression("wide_records", "col_011"),
			unnest: Cast(dtos.Col011()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_011")).To(ast.NewStringColumnExpression("v", "col_011")),
		},
		"col_012": {
			value:  ast.NewStringColumnExpression("wide_records", "col_012"),
			unnest: Cast(dtos.Col012()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_012")).To(ast.NewStringColumnExpression("v", "col_012")),
		},
		"col_013": {
			value:  ast.NewStringColumnExpression("wide_records", "col_013"),
			unnest: Cast(dtos.Col013()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_013")).To(ast.NewStringColumnExpression("v", "col_013")),
		},
		"col_014": {
			value:  ast.NewStringColumnExpression("wide_records", "col_014"),
			unnest: Cast(dtos.Col014()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_014")).To(ast.NewStringColumnExpression("v", "col_014")),
		},
		"col_015": {
			value:  ast.NewStringColumnExpression("wide_records", "col_015"),
			unnest: Cast(dtos.Col015()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_015")).To(ast.NewStringColumnExpression("v", "col_015")),
		},
		"col_016": {
			value:  ast.NewStringColumnExpression("wide_records", "col_016"),
			unnest: Cast(dtos.Col016()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_016")).To(ast.NewStringColumnExpression("v", "col_016")),
		},
		"col_017": {
			value:  ast.NewStringColumnExpression("wide_records", "col_017"),
			unnest: Cast(dtos.Col017()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_017")).To(ast.NewStringColumnExpression("v", "col_017")),
		},
		"col_018": {
			value:  ast.NewStringColumnExpression("wide_records", "col_018"),
			unnest: Cast(dtos.Col018()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_018")).To(ast.NewStringColumnExpression("v", "col_018")),
		},
		"col_019": {
			value:  ast.NewStringColumnExpression("wide_records", "col_019"),
			unnest: Cast(dtos.Col019()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_019")).To(ast.NewStringColumnExpression("v", "col_019")),
		},
		"col_020": {
			value:  ast.NewStringColumnExpression("wide_records", "col_020"),
			unnest: Cast(dtos.Col020()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_020")).To(ast.NewStringColumnExpression("v", "col_020")),
		},
		"col_021": {
			value:  ast.NewStringColumnExpression("wide_records", "col_021"),
			unnest: Cast(dtos.Col021()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_021")).To(ast.NewStringColumnExpression("v", "col_021")),
		},
		"col_022": {
			value:  ast.NewStringColumnExpression("wide_records", "col_022"),
			unnest: Cast(dtos.Col022()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_022")).To(ast.NewStringColumnExpression("v", "col_022")),
		},
		"col_023": {
			value:  ast.NewStringColumnExpression("wide_records", "col_023"),
			unnest: Cast(dtos.Col023()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_023")).To(ast.NewStringColumnExpression("v", "col_023")),
		},
		"col_024": {
			value:  ast.NewStringColumnExpression("wide_records", "col_024"),
			unnest: Cast(dtos.Col024()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_024")).To(ast.NewStringColumnExpression("v", "col_024")),
		},
		"col_025": {
			value:  ast.NewStringColumnExpression("wide_records", "col_025"),
			unnest: Cast(dtos.Col025()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_025")).To(ast.NewStringColumnExpression("v", "col_025")),
		},
		"col_026": {
			value:  ast.NewStringColumnExpression("wide_records", "col_026"),
			unnest: Cast(dtos.Col026()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_026")).To(ast.NewStringColumnExpression("v", "col_026")),
		},
		"col_027": {
			value:  ast.NewStringColumnExpression("wide_records", "col_027"),
			unnest: Cast(dtos.Col027()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_027")).To(ast.NewStringColumnExpression("v", "col_027")),
		},
		"col_028": {
			value:  ast.NewStringColumnExpression("wide_records", "col_028"),
			unnest: Cast(dtos.Col028()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_028")).To(ast.NewStringColumnExpression("v", "col_028")),
		},
		"col_029": {
			value:  ast.NewStringColumnExpression("wide_records", "col_029"),
			unnest: Cast(dtos.Col029()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_029")).To(ast.NewStringColumnExpression("v", "col_029")),
		},
		"col_030": {
			value:  ast.NewStringColumnExpression("wide_records", "col_030"),
			unnest: Cast(dtos.Col030()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_030")).To(ast.NewStringColumnExpression("v", "col_030")),
		},
		"col_031": {
			value:  ast.NewStringColumnExpression("wide_records", "col_031"),
			unnest: Cast(dtos.Col031()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_031")).To(ast.NewStringColumnExpression("v", "col_031")),
		},
		"col_032": {
			value:  ast.NewStringColumnExpression("wide_records", "col_032"),
			unnest: Cast(dtos.Col032()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_032")).To(ast.NewStringColumnExpression("v", "col_032")),
		},
		"col_033": {
			value:  ast.NewStringColumnExpression("wide_records", "col_033"),
			unnest: Cast(dtos.Col033()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_033")).To(ast.NewStringColumnExpression("v", "col_033")),
		},
		"col_034": {
			value:  ast.NewStringColumnExpression("wide_records", "col_034"),
			unnest: Cast(dtos.Col034()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_034")).To(ast.NewStringColumnExpression("v", "col_034")),
		},
		"col_035": {
			value:  ast.NewStringColumnExpression("wide_records", "col_035"),
			unnest: Cast(dtos.Col035()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_035")).To(ast.NewStringColumnExpression("v", "col_035")),
		},
		"col_036": {
			value:  ast.NewStringColumnExpression("wide_records", "col_036"),
			unnest: Cast(dtos.Col036()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_036")).To(ast.NewStringColumnExpression("v", "col_036")),
		},
		"col_037": {
			value:  ast.NewStringColumnExpression("wide_records", "col_037"),
			unnest: Cast(dtos.Col037()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_037")).To(ast.NewStringColumnExpression("v", "col_037")),
		},
		"col_038": {
			value:  ast.NewStringColumnExpression("wide_records", "col_038"),
			unnest: Cast(dtos.Col038()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_038")).To(ast.NewStringColumnExpression("v", "col_038")),
		},
		"col_039": {
			value:  ast.NewStringColumnExpression("wide_records", "col_039"),
			unnest: Cast(dtos.Col039()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_039")).To(ast.NewStringColumnExpression("v", "col_039")),
		},
		"col_040": {
			value:  ast.NewStringColumnExpression("wide_records", "col_040"),
			unnest: Cast(dtos.Col040()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_040")).To(ast.NewStringColumnExpression("v", "col_040")),
		},
		"col_041": {
			value:  ast.NewStringColumnExpression("wide_records", "col_041"),
			unnest: Cast(dtos.Col041()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_041")).To(ast.NewStringColumnExpression("v", "col_041")),
		},
		"col_042": {
			value:  ast.NewStringColumnExpression("wide_records", "col_042"),
			unnest: Cast(dtos.Col042()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_042")).To(ast.NewStringColumnExpression("v", "col_042")),
		},
		"col_043": {
			value:  ast.NewStringColumnExpression("wide_records", "col_043"),
			unnest: Cast(dtos.Col043()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_043")).To(ast.NewStringColumnExpression("v", "col_043")),
		},
		"col_044": {
			value:  ast.NewStringColumnExpression("wide_records", "col_044"),
			unnest: Cast(dtos.Col044()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_044")).To(ast.NewStringColumnExpression("v", "col_044")),
		},
		"col_045": {
			value:  ast.NewStringColumnExpression("wide_records", "col_045"),
			unnest: Cast(dtos.Col045()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_045")).To(ast.NewStringColumnExpression("v", "col_045")),
		},
		"col_046": {
			value:  ast.NewStringColumnExpression("wide_records", "col_046"),
			unnest: Cast(dtos.Col046()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_046")).To(ast.NewStringColumnExpression("v", "col_046")),
		},
		"col_047": {
			value:  ast.NewStringColumnExpression("wide_records", "col_047"),
			unnest: Cast(dtos.Col047()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_047")).To(ast.NewStringColumnExpression("v", "col_047")),
		},
		"col_048": {
			value:  ast.NewStringColumnExpression("wide_records", "col_048"),
			unnest: Cast(dtos.Col048()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_048")).To(ast.NewStringColumnExpression("v", "col_048")),
		},
		"col_049": {
			value:  ast.NewStringColumnExpression("wide_records", "col_049"),
			unnest: Cast(dtos.Col049()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_049")).To(ast.NewStringColumnExpression("v", "col_049")),
		},
		"col_050": {
			value:  ast.NewStringColumnExpression("wide_records", "col_050"),
			unnest: Cast(dtos.Col050()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_050")).To(ast.NewStringColumnExpression("v", "col_050")),
		},
		"col_051": {
			value:  ast.NewStringColumnExpression("wide_records", "col_051"),
			unnest: Cast(dtos.Col051()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_051")).To(ast.NewStringColumnExpression("v", "col_051")),
		},
		"col_052": {
			value:  ast.NewStringColumnExpression("wide_records", "col_052"),
			unnest: Cast(dtos.Col052()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_052")).To(ast.NewStringColumnExpression("v", "col_052")),
		},
		"col_053": {
			value:  ast.NewStringColumnExpression("wide_records", "col_053"),
			unnest: Cast(dtos.Col053()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_053")).To(ast.NewStringColumnExpression("v", "col_053")),
		},
		"col_054": {
			value:  ast.NewStringColumnExpression("wide_records", "col_054"),
			unnest: Cast(dtos.Col054()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_054")).To(ast.NewStringColumnExpression("v", "col_054")),
		},
		"col_055": {
			value:  ast.NewStringColumnExpression("wide_records", "col_055"),
			unnest: Cast(dtos.Col055()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_055")).To(ast.NewStringColumnExpression("v", "col_055")),
		},
		"col_056": {
			value:  ast.NewStringColumnExpression("wide_records", "col_056"),
			unnest: Cast(dtos.Col056()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_056")).To(ast.NewStringColumnExpression("v", "col_056")),
		},
		"col_057": {
			value:  ast.NewStringColumnExpression("wide_records", "col_057"),
			unnest: Cast(dtos.Col057()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_057")).To(ast.NewStringColumnExpression("v", "col_057")),
		},
		"col_058": {
			value:  ast.NewStringColumnExpression("wide_records", "col_058"),
			unnest: Cast(dtos.Col058()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_058")).To(ast.NewStringColumnExpression("v", "col_058")),
		},
		"col_059": {
			value:  ast.NewStringColumnExpression("wide_records", "col_059"),
			unnest: Cast(dtos.Col059()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_059")).To(ast.NewStringColumnExpression("v", "col_059")),
		},
		"col_060": {
			value:  ast.NewStringColumnExpression("wide_records", "col_060"),
			unnest: Cast(dtos.Col060()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_060")).To(ast.NewStringColumnExpression("v", "col_060")),
		},
		"col_061": {
			value:  ast.NewStringColumnExpression("wide_records", "col_061"),
			unnest: Cast(dtos.Col061()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_061")).To(ast.NewStringColumnExpression("v", "col_061")),
		},
		"col_062": {
			value:  ast.NewStringColumnExpression("wide_records", "col_062"),
			unnest: Cast(dtos.Col062()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_062")).To(ast.NewStringColumnExpression("v", "col_062")),
		},
		"col_063": {
			value:  ast.NewStringColumnExpression("wide_records", "col_063"),
			unnest: Cast(dtos.Col063()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_063")).To(ast.NewStringColumnExpression("v", "col_063")),
		},
		"col_064": {
			value:  ast.NewStringColumnExpression("wide_records", "col_064"),
			unnest: Cast(dtos.Col064()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_064")).To(ast.NewStringColumnExpression("v", "col_064")),
		},
		"col_065": {
			value:  ast.NewStringColumnExpression("wide_records", "col_065"),
			unnest: Cast(dtos.Col065()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_065")).To(ast.NewStringColumnExpression("v", "col_065")),
		},
		"col_066": {
			value:  ast.NewStringColumnExpression("wide_records", "col_066"),
			unnest: Cast(dtos.Col066()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_066")).To(ast.NewStringColumnExpression("v", "col_066")),
		},
		"col_067": {
			value:  ast.NewStringColumnExpression("wide_records", "col_067"),
			unnest: Cast(dtos.Col067()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_067")).To(ast.NewStringColumnExpression("v", "col_067")),
		},
		"col_068": {
			value:  ast.NewStringColumnExpression("wide_records", "col_068"),
			unnest: Cast(dtos.Col068()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_068")).To(ast.NewStringColumnExpression("v", "col_068")),
		},
		"col_069": {
			value:  ast.NewStringColumnExpression("wide_records", "col_069"),
			unnest: Cast(dtos.Col069()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_069")).To(ast.NewStringColumnExpression("v", "col_069")),
		},
		"col_070": {
			value:  ast.NewStringColumnExpression("wide_records", "col_070"),
			unnest: Cast(dtos.Col070()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_070")).To(ast.NewStringColumnExpression("v", "col_070")),
		},
		"col_071": {
			value:  ast.NewStringColumnExpression("wide_records", "col_071"),
			unnest: Cast(dtos.Col071()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_071")).To(ast.NewStringColumnExpression("v", "col_071")),
		},
		"col_072": {
			value:  ast.NewStringColumnExpression("wide_records", "col_072"),
			unnest: Cast(dtos.Col072()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_072")).To(ast.NewStringColumnExpression("v", "col_072")),
		},
		"col_073": {
			value:  ast.NewStringColumnExpression("wide_records", "col_073"),
			unnest: Cast(dtos.Col073()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_073")).To(ast.NewStringColumnExpression("v", "col_073")),
		},
		"col_074": {
			value:  ast.NewStringColumnExpression("wide_records", "col_074"),
			unnest: Cast(dtos.Col074()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_074")).To(ast.NewStringColumnExpression("v", "col_074")),
		},
		"col_075": {
			value:  ast.NewStringColumnExpression("wide_records", "col_075"),
			unnest: Cast(dtos.Col075()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_075")).To(ast.NewStringColumnExpression("v", "col_075")),
		},
		"col_076": {
			value:  ast.NewStringColumnExpression("wide_records", "col_076"),
			unnest: Cast(dtos.Col076()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_076")).To(ast.NewStringColumnExpression("v", "col_076")),
		},
		"col_077": {
			value:  ast.NewStringColumnExpression("wide_records", "col_077"),
			unnest: Cast(dtos.Col077()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_077")).To(ast.NewStringColumnExpression("v", "col_077")),
		},
		"col_078": {
			value:  ast.NewStringColumnExpression("wide_records", "col_078"),
			unnest: Cast(dtos.Col078()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_078")).To(ast.NewStringColumnExpression("v", "col_078")),
		},
		"col_079": {
			value:  ast.NewStringColumnExpression("wide_records", "col_079"),
			unnest: Cast(dtos.Col079()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_079")).To(ast.NewStringColumnExpression("v", "col_079")),
		},
		"col_080": {
			value:  ast.NewStringColumnExpression("wide_records", "col_080"),
			unnest: Cast(dtos.Col080()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_080")).To(ast.NewStringColumnExpression("v", "col_080")),
		},
		"col_081": {
			value:  ast.NewStringColumnExpression("wide_records", "col_081"),
			unnest: Cast(dtos.Col081()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_081")).To(ast.NewStringColumnExpression("v", "col_081")),
		},
		"col_082": {
			value:  ast.NewStringColumnExpression("wide_records", "col_082"),
			unnest: Cast(dtos.Col082()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_082")).To(ast.NewStringColumnExpression("v", "col_082")),
		},
		"col_083": {
			value:  ast.NewStringColumnExpression("wide_records", "col_083"),
			unnest: Cast(dtos.Col083()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_083")).To(ast.NewStringColumnExpression("v", "col_083")),
		},
		"col_084": {
			value:  ast.NewStringColumnExpression("wide_records", "col_084"),
			unnest: Cast(dtos.Col084()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_084")).To(ast.NewStringColumnExpression("v", "col_084")),
		},
		"col_085": {
			value:  ast.NewStringColumnExpression("wide_records", "col_085"),
			unnest: Cast(dtos.Col085()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_085")).To(ast.NewStringColumnExpression("v", "col_085")),
		},
		"col_086": {
			value:  ast.NewStringColumnExpression("wide_records", "col_086"),
			unnest: Cast(dtos.Col086()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_086")).To(ast.NewStringColumnExpression("v", "col_086")),
		},
		"col_087": {
			value:  ast.NewStringColumnExpression("wide_records", "col_087"),
			unnest: Cast(dtos.Col087()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_087")).To(ast.NewStringColumnExpression("v", "col_087")),
		},
		"col_088": {
			value:  ast.NewStringColumnExpression("wide_records", "col_088"),
			unnest: Cast(dtos.Col088()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_088")).To(ast.NewStringColumnExpression("v", "col_088")),
		},
		"col_089": {
			value:  ast.NewStringColumnExpression("wide_records", "col_089"),
			unnest: Cast(dtos.Col089()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_089")).To(ast.NewStringColumnExpression("v", "col_089")),
		},
		"col_090": {
			value:  ast.NewStringColumnExpression("wide_records", "col_090"),
			unnest: Cast(dtos.Col090()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_090")).To(ast.NewStringColumnExpression("v", "col_090")),
		},
		"col_091": {
			value:  ast.NewStringColumnExpression("wide_records", "col_091"),
			unnest: Cast(dtos.Col091()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_091")).To(ast.NewStringColumnExpression("v", "col_091")),
		},
		"col_092": {
			value:  ast.NewStringColumnExpression("wide_records", "col_092"),
			unnest: Cast(dtos.Col092()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_092")).To(ast.NewStringColumnExpression("v", "col_092")),
		},
		"col_093": {
			value:  ast.NewStringColumnExpression("wide_records", "col_093"),
			unnest: Cast(dtos.Col093()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_093")).To(ast.NewStringColumnExpression("v", "col_093")),
		},
		"col_094": {
			value:  ast.NewStringColumnExpression("wide_records", "col_094"),
			unnest: Cast(dtos.Col094()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_094")).To(ast.NewStringColumnExpression("v", "col_094")),
		},
		"col_095": {
			value:  ast.NewStringColumnExpression("wide_records", "col_095"),
			unnest: Cast(dtos.Col095()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_095")).To(ast.NewStringColumnExpression("v", "col_095")),
		},
		"col_096": {
			value:  ast.NewStringColumnExpression("wide_records", "col_096"),
			unnest: Cast(dtos.Col096()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_096")).To(ast.NewStringColumnExpression("v", "col_096")),
		},
		"col_097": {
			value:  ast.NewStringColumnExpression("wide_records", "col_097"),
			unnest: Cast(dtos.Col097()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_097")).To(ast.NewStringColumnExpression("v", "col_097")),
		},
		"col_098": {
			value:  ast.NewStringColumnExpression("wide_records", "col_098"),
			unnest: Cast(dtos.Col098()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_098")).To(ast.NewStringColumnExpression("v", "col_098")),
		},
		"col_099": {
			value:  ast.NewStringColumnExpression("wide_records", "col_099"),
			unnest: Cast(dtos.Col099()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_099")).To(ast.NewStringColumnExpression("v", "col_099")),
		},
		"col_100": {
			value:  ast.NewStringColumnExpression("wide_records", "col_100"),
			unnest: Cast(dtos.Col100()).AsTextArray(),
			set:    Set(ast.NewStringColumnExpression("wide_records", "col_100")).To(ast.NewStringColumnExpression("v", "col_100")),
		},
	}
}
