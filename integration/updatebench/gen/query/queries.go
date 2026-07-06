package query

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/models"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/ast"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/builder"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/users"
	"github.com/jalphad/gosqlgen/integration/updatebench/gen/query/wide_records"
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

// UsersDtoInsertOne inserts a single users DTO
func UsersDtoInsertOne(pool *pgxpool.Pool, dto *models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Insert(
			users.Id(),
		).
		Returning(users.Id()).
		Values(dto)
}

// UsersDtoInsertMany inserts multiple users DTOs
func UsersDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.UsersDto) builder.InsertFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Insert(
			users.Id(),
		).
		Returning(users.Id()).
		Values(dtos...)
}

// WideRecordsDtoInsertOne inserts a single wide_records DTO
func WideRecordsDtoInsertOne(pool *pgxpool.Pool, dto *models.WideRecordsDto) builder.InsertFinalizeQuery[models.WideRecordsDto] {
	return wide_records.NewQuery(pool).
		Insert(
			wide_records.Id(),
			wide_records.Col001(),
			wide_records.Col002(),
			wide_records.Col003(),
			wide_records.Col004(),
			wide_records.Col005(),
			wide_records.Col006(),
			wide_records.Col007(),
			wide_records.Col008(),
			wide_records.Col009(),
			wide_records.Col010(),
			wide_records.Col011(),
			wide_records.Col012(),
			wide_records.Col013(),
			wide_records.Col014(),
			wide_records.Col015(),
			wide_records.Col016(),
			wide_records.Col017(),
			wide_records.Col018(),
			wide_records.Col019(),
			wide_records.Col020(),
			wide_records.Col021(),
			wide_records.Col022(),
			wide_records.Col023(),
			wide_records.Col024(),
			wide_records.Col025(),
			wide_records.Col026(),
			wide_records.Col027(),
			wide_records.Col028(),
			wide_records.Col029(),
			wide_records.Col030(),
			wide_records.Col031(),
			wide_records.Col032(),
			wide_records.Col033(),
			wide_records.Col034(),
			wide_records.Col035(),
			wide_records.Col036(),
			wide_records.Col037(),
			wide_records.Col038(),
			wide_records.Col039(),
			wide_records.Col040(),
			wide_records.Col041(),
			wide_records.Col042(),
			wide_records.Col043(),
			wide_records.Col044(),
			wide_records.Col045(),
			wide_records.Col046(),
			wide_records.Col047(),
			wide_records.Col048(),
			wide_records.Col049(),
			wide_records.Col050(),
			wide_records.Col051(),
			wide_records.Col052(),
			wide_records.Col053(),
			wide_records.Col054(),
			wide_records.Col055(),
			wide_records.Col056(),
			wide_records.Col057(),
			wide_records.Col058(),
			wide_records.Col059(),
			wide_records.Col060(),
			wide_records.Col061(),
			wide_records.Col062(),
			wide_records.Col063(),
			wide_records.Col064(),
			wide_records.Col065(),
			wide_records.Col066(),
			wide_records.Col067(),
			wide_records.Col068(),
			wide_records.Col069(),
			wide_records.Col070(),
			wide_records.Col071(),
			wide_records.Col072(),
			wide_records.Col073(),
			wide_records.Col074(),
			wide_records.Col075(),
			wide_records.Col076(),
			wide_records.Col077(),
			wide_records.Col078(),
			wide_records.Col079(),
			wide_records.Col080(),
			wide_records.Col081(),
			wide_records.Col082(),
			wide_records.Col083(),
			wide_records.Col084(),
			wide_records.Col085(),
			wide_records.Col086(),
			wide_records.Col087(),
			wide_records.Col088(),
			wide_records.Col089(),
			wide_records.Col090(),
			wide_records.Col091(),
			wide_records.Col092(),
			wide_records.Col093(),
			wide_records.Col094(),
			wide_records.Col095(),
			wide_records.Col096(),
			wide_records.Col097(),
			wide_records.Col098(),
			wide_records.Col099(),
			wide_records.Col100(),
		).
		Returning(wide_records.Id()).
		Values(dto)
}

// WideRecordsDtoInsertMany inserts multiple wide_records DTOs
func WideRecordsDtoInsertMany(pool *pgxpool.Pool, dtos ...*models.WideRecordsDto) builder.InsertFinalizeQuery[models.WideRecordsDto] {
	return wide_records.NewQuery(pool).
		Insert(
			wide_records.Id(),
			wide_records.Col001(),
			wide_records.Col002(),
			wide_records.Col003(),
			wide_records.Col004(),
			wide_records.Col005(),
			wide_records.Col006(),
			wide_records.Col007(),
			wide_records.Col008(),
			wide_records.Col009(),
			wide_records.Col010(),
			wide_records.Col011(),
			wide_records.Col012(),
			wide_records.Col013(),
			wide_records.Col014(),
			wide_records.Col015(),
			wide_records.Col016(),
			wide_records.Col017(),
			wide_records.Col018(),
			wide_records.Col019(),
			wide_records.Col020(),
			wide_records.Col021(),
			wide_records.Col022(),
			wide_records.Col023(),
			wide_records.Col024(),
			wide_records.Col025(),
			wide_records.Col026(),
			wide_records.Col027(),
			wide_records.Col028(),
			wide_records.Col029(),
			wide_records.Col030(),
			wide_records.Col031(),
			wide_records.Col032(),
			wide_records.Col033(),
			wide_records.Col034(),
			wide_records.Col035(),
			wide_records.Col036(),
			wide_records.Col037(),
			wide_records.Col038(),
			wide_records.Col039(),
			wide_records.Col040(),
			wide_records.Col041(),
			wide_records.Col042(),
			wide_records.Col043(),
			wide_records.Col044(),
			wide_records.Col045(),
			wide_records.Col046(),
			wide_records.Col047(),
			wide_records.Col048(),
			wide_records.Col049(),
			wide_records.Col050(),
			wide_records.Col051(),
			wide_records.Col052(),
			wide_records.Col053(),
			wide_records.Col054(),
			wide_records.Col055(),
			wide_records.Col056(),
			wide_records.Col057(),
			wide_records.Col058(),
			wide_records.Col059(),
			wide_records.Col060(),
			wide_records.Col061(),
			wide_records.Col062(),
			wide_records.Col063(),
			wide_records.Col064(),
			wide_records.Col065(),
			wide_records.Col066(),
			wide_records.Col067(),
			wide_records.Col068(),
			wide_records.Col069(),
			wide_records.Col070(),
			wide_records.Col071(),
			wide_records.Col072(),
			wide_records.Col073(),
			wide_records.Col074(),
			wide_records.Col075(),
			wide_records.Col076(),
			wide_records.Col077(),
			wide_records.Col078(),
			wide_records.Col079(),
			wide_records.Col080(),
			wide_records.Col081(),
			wide_records.Col082(),
			wide_records.Col083(),
			wide_records.Col084(),
			wide_records.Col085(),
			wide_records.Col086(),
			wide_records.Col087(),
			wide_records.Col088(),
			wide_records.Col089(),
			wide_records.Col090(),
			wide_records.Col091(),
			wide_records.Col092(),
			wide_records.Col093(),
			wide_records.Col094(),
			wide_records.Col095(),
			wide_records.Col096(),
			wide_records.Col097(),
			wide_records.Col098(),
			wide_records.Col099(),
			wide_records.Col100(),
		).
		Returning(wide_records.Id()).
		Values(dtos...)
}

type WideRecordsDtoUpdateValueExtractor func(*models.WideRecordsDto) any

var defaultWideRecordsDtoUpdateValueColumns = []ast.NamedExpression{
	wide_records.Col001(),
	wide_records.Col002(),
	wide_records.Col003(),
	wide_records.Col004(),
	wide_records.Col005(),
	wide_records.Col006(),
	wide_records.Col007(),
	wide_records.Col008(),
	wide_records.Col009(),
	wide_records.Col010(),
	wide_records.Col011(),
	wide_records.Col012(),
	wide_records.Col013(),
	wide_records.Col014(),
	wide_records.Col015(),
	wide_records.Col016(),
	wide_records.Col017(),
	wide_records.Col018(),
	wide_records.Col019(),
	wide_records.Col020(),
	wide_records.Col021(),
	wide_records.Col022(),
	wide_records.Col023(),
	wide_records.Col024(),
	wide_records.Col025(),
	wide_records.Col026(),
	wide_records.Col027(),
	wide_records.Col028(),
	wide_records.Col029(),
	wide_records.Col030(),
	wide_records.Col031(),
	wide_records.Col032(),
	wide_records.Col033(),
	wide_records.Col034(),
	wide_records.Col035(),
	wide_records.Col036(),
	wide_records.Col037(),
	wide_records.Col038(),
	wide_records.Col039(),
	wide_records.Col040(),
	wide_records.Col041(),
	wide_records.Col042(),
	wide_records.Col043(),
	wide_records.Col044(),
	wide_records.Col045(),
	wide_records.Col046(),
	wide_records.Col047(),
	wide_records.Col048(),
	wide_records.Col049(),
	wide_records.Col050(),
	wide_records.Col051(),
	wide_records.Col052(),
	wide_records.Col053(),
	wide_records.Col054(),
	wide_records.Col055(),
	wide_records.Col056(),
	wide_records.Col057(),
	wide_records.Col058(),
	wide_records.Col059(),
	wide_records.Col060(),
	wide_records.Col061(),
	wide_records.Col062(),
	wide_records.Col063(),
	wide_records.Col064(),
	wide_records.Col065(),
	wide_records.Col066(),
	wide_records.Col067(),
	wide_records.Col068(),
	wide_records.Col069(),
	wide_records.Col070(),
	wide_records.Col071(),
	wide_records.Col072(),
	wide_records.Col073(),
	wide_records.Col074(),
	wide_records.Col075(),
	wide_records.Col076(),
	wide_records.Col077(),
	wide_records.Col078(),
	wide_records.Col079(),
	wide_records.Col080(),
	wide_records.Col081(),
	wide_records.Col082(),
	wide_records.Col083(),
	wide_records.Col084(),
	wide_records.Col085(),
	wide_records.Col086(),
	wide_records.Col087(),
	wide_records.Col088(),
	wide_records.Col089(),
	wide_records.Col090(),
	wide_records.Col091(),
	wide_records.Col092(),
	wide_records.Col093(),
	wide_records.Col094(),
	wide_records.Col095(),
	wide_records.Col096(),
	wide_records.Col097(),
	wide_records.Col098(),
	wide_records.Col099(),
	wide_records.Col100(),
	wide_records.Id(),
}

var defaultWideRecordsDtoUpdateValueExtractors = []WideRecordsDtoUpdateValueExtractor{
	func(dto *models.WideRecordsDto) any { return dto.Col001 },
	func(dto *models.WideRecordsDto) any { return dto.Col002 },
	func(dto *models.WideRecordsDto) any { return dto.Col003 },
	func(dto *models.WideRecordsDto) any { return dto.Col004 },
	func(dto *models.WideRecordsDto) any { return dto.Col005 },
	func(dto *models.WideRecordsDto) any { return dto.Col006 },
	func(dto *models.WideRecordsDto) any { return dto.Col007 },
	func(dto *models.WideRecordsDto) any { return dto.Col008 },
	func(dto *models.WideRecordsDto) any { return dto.Col009 },
	func(dto *models.WideRecordsDto) any { return dto.Col010 },
	func(dto *models.WideRecordsDto) any { return dto.Col011 },
	func(dto *models.WideRecordsDto) any { return dto.Col012 },
	func(dto *models.WideRecordsDto) any { return dto.Col013 },
	func(dto *models.WideRecordsDto) any { return dto.Col014 },
	func(dto *models.WideRecordsDto) any { return dto.Col015 },
	func(dto *models.WideRecordsDto) any { return dto.Col016 },
	func(dto *models.WideRecordsDto) any { return dto.Col017 },
	func(dto *models.WideRecordsDto) any { return dto.Col018 },
	func(dto *models.WideRecordsDto) any { return dto.Col019 },
	func(dto *models.WideRecordsDto) any { return dto.Col020 },
	func(dto *models.WideRecordsDto) any { return dto.Col021 },
	func(dto *models.WideRecordsDto) any { return dto.Col022 },
	func(dto *models.WideRecordsDto) any { return dto.Col023 },
	func(dto *models.WideRecordsDto) any { return dto.Col024 },
	func(dto *models.WideRecordsDto) any { return dto.Col025 },
	func(dto *models.WideRecordsDto) any { return dto.Col026 },
	func(dto *models.WideRecordsDto) any { return dto.Col027 },
	func(dto *models.WideRecordsDto) any { return dto.Col028 },
	func(dto *models.WideRecordsDto) any { return dto.Col029 },
	func(dto *models.WideRecordsDto) any { return dto.Col030 },
	func(dto *models.WideRecordsDto) any { return dto.Col031 },
	func(dto *models.WideRecordsDto) any { return dto.Col032 },
	func(dto *models.WideRecordsDto) any { return dto.Col033 },
	func(dto *models.WideRecordsDto) any { return dto.Col034 },
	func(dto *models.WideRecordsDto) any { return dto.Col035 },
	func(dto *models.WideRecordsDto) any { return dto.Col036 },
	func(dto *models.WideRecordsDto) any { return dto.Col037 },
	func(dto *models.WideRecordsDto) any { return dto.Col038 },
	func(dto *models.WideRecordsDto) any { return dto.Col039 },
	func(dto *models.WideRecordsDto) any { return dto.Col040 },
	func(dto *models.WideRecordsDto) any { return dto.Col041 },
	func(dto *models.WideRecordsDto) any { return dto.Col042 },
	func(dto *models.WideRecordsDto) any { return dto.Col043 },
	func(dto *models.WideRecordsDto) any { return dto.Col044 },
	func(dto *models.WideRecordsDto) any { return dto.Col045 },
	func(dto *models.WideRecordsDto) any { return dto.Col046 },
	func(dto *models.WideRecordsDto) any { return dto.Col047 },
	func(dto *models.WideRecordsDto) any { return dto.Col048 },
	func(dto *models.WideRecordsDto) any { return dto.Col049 },
	func(dto *models.WideRecordsDto) any { return dto.Col050 },
	func(dto *models.WideRecordsDto) any { return dto.Col051 },
	func(dto *models.WideRecordsDto) any { return dto.Col052 },
	func(dto *models.WideRecordsDto) any { return dto.Col053 },
	func(dto *models.WideRecordsDto) any { return dto.Col054 },
	func(dto *models.WideRecordsDto) any { return dto.Col055 },
	func(dto *models.WideRecordsDto) any { return dto.Col056 },
	func(dto *models.WideRecordsDto) any { return dto.Col057 },
	func(dto *models.WideRecordsDto) any { return dto.Col058 },
	func(dto *models.WideRecordsDto) any { return dto.Col059 },
	func(dto *models.WideRecordsDto) any { return dto.Col060 },
	func(dto *models.WideRecordsDto) any { return dto.Col061 },
	func(dto *models.WideRecordsDto) any { return dto.Col062 },
	func(dto *models.WideRecordsDto) any { return dto.Col063 },
	func(dto *models.WideRecordsDto) any { return dto.Col064 },
	func(dto *models.WideRecordsDto) any { return dto.Col065 },
	func(dto *models.WideRecordsDto) any { return dto.Col066 },
	func(dto *models.WideRecordsDto) any { return dto.Col067 },
	func(dto *models.WideRecordsDto) any { return dto.Col068 },
	func(dto *models.WideRecordsDto) any { return dto.Col069 },
	func(dto *models.WideRecordsDto) any { return dto.Col070 },
	func(dto *models.WideRecordsDto) any { return dto.Col071 },
	func(dto *models.WideRecordsDto) any { return dto.Col072 },
	func(dto *models.WideRecordsDto) any { return dto.Col073 },
	func(dto *models.WideRecordsDto) any { return dto.Col074 },
	func(dto *models.WideRecordsDto) any { return dto.Col075 },
	func(dto *models.WideRecordsDto) any { return dto.Col076 },
	func(dto *models.WideRecordsDto) any { return dto.Col077 },
	func(dto *models.WideRecordsDto) any { return dto.Col078 },
	func(dto *models.WideRecordsDto) any { return dto.Col079 },
	func(dto *models.WideRecordsDto) any { return dto.Col080 },
	func(dto *models.WideRecordsDto) any { return dto.Col081 },
	func(dto *models.WideRecordsDto) any { return dto.Col082 },
	func(dto *models.WideRecordsDto) any { return dto.Col083 },
	func(dto *models.WideRecordsDto) any { return dto.Col084 },
	func(dto *models.WideRecordsDto) any { return dto.Col085 },
	func(dto *models.WideRecordsDto) any { return dto.Col086 },
	func(dto *models.WideRecordsDto) any { return dto.Col087 },
	func(dto *models.WideRecordsDto) any { return dto.Col088 },
	func(dto *models.WideRecordsDto) any { return dto.Col089 },
	func(dto *models.WideRecordsDto) any { return dto.Col090 },
	func(dto *models.WideRecordsDto) any { return dto.Col091 },
	func(dto *models.WideRecordsDto) any { return dto.Col092 },
	func(dto *models.WideRecordsDto) any { return dto.Col093 },
	func(dto *models.WideRecordsDto) any { return dto.Col094 },
	func(dto *models.WideRecordsDto) any { return dto.Col095 },
	func(dto *models.WideRecordsDto) any { return dto.Col096 },
	func(dto *models.WideRecordsDto) any { return dto.Col097 },
	func(dto *models.WideRecordsDto) any { return dto.Col098 },
	func(dto *models.WideRecordsDto) any { return dto.Col099 },
	func(dto *models.WideRecordsDto) any { return dto.Col100 },
	func(dto *models.WideRecordsDto) any { return *dto.Id },
}

var defaultWideRecordsDtoUpdateValueCastTypes = []string{
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"text",
	"uuid",
}

var defaultWideRecordsDtoUpdateColumnIndexes = []int{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	15,
	16,
	17,
	18,
	19,
	20,
	21,
	22,
	23,
	24,
	25,
	26,
	27,
	28,
	29,
	30,
	31,
	32,
	33,
	34,
	35,
	36,
	37,
	38,
	39,
	40,
	41,
	42,
	43,
	44,
	45,
	46,
	47,
	48,
	49,
	50,
	51,
	52,
	53,
	54,
	55,
	56,
	57,
	58,
	59,
	60,
	61,
	62,
	63,
	64,
	65,
	66,
	67,
	68,
	69,
	70,
	71,
	72,
	73,
	74,
	75,
	76,
	77,
	78,
	79,
	80,
	81,
	82,
	83,
	84,
	85,
	86,
	87,
	88,
	89,
	90,
	91,
	92,
	93,
	94,
	95,
	96,
	97,
	98,
	99,
	100,
}

var primaryKeyWideRecordsDtoUpdateValueColumns = []ast.NamedExpression{
	wide_records.Id(),
}

var primaryKeyWideRecordsDtoUpdateValueExtractors = []WideRecordsDtoUpdateValueExtractor{
	func(dto *models.WideRecordsDto) any { return *dto.Id },
}

var primaryKeyWideRecordsDtoUpdateValueCastTypes = []string{
	"uuid",
}

var primaryKeyWideRecordsDtoUpdateColumnIndexes = []int{
	100,
}

var defaultWideRecordsDtoUpdateSets = []ast.UpdateSetExpr{
	Set(wide_records.Col001()).To(ast.SetType[string](ast.NewColumnNode("v", "col_001"))),
	Set(wide_records.Col002()).To(ast.SetType[string](ast.NewColumnNode("v", "col_002"))),
	Set(wide_records.Col003()).To(ast.SetType[string](ast.NewColumnNode("v", "col_003"))),
	Set(wide_records.Col004()).To(ast.SetType[string](ast.NewColumnNode("v", "col_004"))),
	Set(wide_records.Col005()).To(ast.SetType[string](ast.NewColumnNode("v", "col_005"))),
	Set(wide_records.Col006()).To(ast.SetType[string](ast.NewColumnNode("v", "col_006"))),
	Set(wide_records.Col007()).To(ast.SetType[string](ast.NewColumnNode("v", "col_007"))),
	Set(wide_records.Col008()).To(ast.SetType[string](ast.NewColumnNode("v", "col_008"))),
	Set(wide_records.Col009()).To(ast.SetType[string](ast.NewColumnNode("v", "col_009"))),
	Set(wide_records.Col010()).To(ast.SetType[string](ast.NewColumnNode("v", "col_010"))),
	Set(wide_records.Col011()).To(ast.SetType[string](ast.NewColumnNode("v", "col_011"))),
	Set(wide_records.Col012()).To(ast.SetType[string](ast.NewColumnNode("v", "col_012"))),
	Set(wide_records.Col013()).To(ast.SetType[string](ast.NewColumnNode("v", "col_013"))),
	Set(wide_records.Col014()).To(ast.SetType[string](ast.NewColumnNode("v", "col_014"))),
	Set(wide_records.Col015()).To(ast.SetType[string](ast.NewColumnNode("v", "col_015"))),
	Set(wide_records.Col016()).To(ast.SetType[string](ast.NewColumnNode("v", "col_016"))),
	Set(wide_records.Col017()).To(ast.SetType[string](ast.NewColumnNode("v", "col_017"))),
	Set(wide_records.Col018()).To(ast.SetType[string](ast.NewColumnNode("v", "col_018"))),
	Set(wide_records.Col019()).To(ast.SetType[string](ast.NewColumnNode("v", "col_019"))),
	Set(wide_records.Col020()).To(ast.SetType[string](ast.NewColumnNode("v", "col_020"))),
	Set(wide_records.Col021()).To(ast.SetType[string](ast.NewColumnNode("v", "col_021"))),
	Set(wide_records.Col022()).To(ast.SetType[string](ast.NewColumnNode("v", "col_022"))),
	Set(wide_records.Col023()).To(ast.SetType[string](ast.NewColumnNode("v", "col_023"))),
	Set(wide_records.Col024()).To(ast.SetType[string](ast.NewColumnNode("v", "col_024"))),
	Set(wide_records.Col025()).To(ast.SetType[string](ast.NewColumnNode("v", "col_025"))),
	Set(wide_records.Col026()).To(ast.SetType[string](ast.NewColumnNode("v", "col_026"))),
	Set(wide_records.Col027()).To(ast.SetType[string](ast.NewColumnNode("v", "col_027"))),
	Set(wide_records.Col028()).To(ast.SetType[string](ast.NewColumnNode("v", "col_028"))),
	Set(wide_records.Col029()).To(ast.SetType[string](ast.NewColumnNode("v", "col_029"))),
	Set(wide_records.Col030()).To(ast.SetType[string](ast.NewColumnNode("v", "col_030"))),
	Set(wide_records.Col031()).To(ast.SetType[string](ast.NewColumnNode("v", "col_031"))),
	Set(wide_records.Col032()).To(ast.SetType[string](ast.NewColumnNode("v", "col_032"))),
	Set(wide_records.Col033()).To(ast.SetType[string](ast.NewColumnNode("v", "col_033"))),
	Set(wide_records.Col034()).To(ast.SetType[string](ast.NewColumnNode("v", "col_034"))),
	Set(wide_records.Col035()).To(ast.SetType[string](ast.NewColumnNode("v", "col_035"))),
	Set(wide_records.Col036()).To(ast.SetType[string](ast.NewColumnNode("v", "col_036"))),
	Set(wide_records.Col037()).To(ast.SetType[string](ast.NewColumnNode("v", "col_037"))),
	Set(wide_records.Col038()).To(ast.SetType[string](ast.NewColumnNode("v", "col_038"))),
	Set(wide_records.Col039()).To(ast.SetType[string](ast.NewColumnNode("v", "col_039"))),
	Set(wide_records.Col040()).To(ast.SetType[string](ast.NewColumnNode("v", "col_040"))),
	Set(wide_records.Col041()).To(ast.SetType[string](ast.NewColumnNode("v", "col_041"))),
	Set(wide_records.Col042()).To(ast.SetType[string](ast.NewColumnNode("v", "col_042"))),
	Set(wide_records.Col043()).To(ast.SetType[string](ast.NewColumnNode("v", "col_043"))),
	Set(wide_records.Col044()).To(ast.SetType[string](ast.NewColumnNode("v", "col_044"))),
	Set(wide_records.Col045()).To(ast.SetType[string](ast.NewColumnNode("v", "col_045"))),
	Set(wide_records.Col046()).To(ast.SetType[string](ast.NewColumnNode("v", "col_046"))),
	Set(wide_records.Col047()).To(ast.SetType[string](ast.NewColumnNode("v", "col_047"))),
	Set(wide_records.Col048()).To(ast.SetType[string](ast.NewColumnNode("v", "col_048"))),
	Set(wide_records.Col049()).To(ast.SetType[string](ast.NewColumnNode("v", "col_049"))),
	Set(wide_records.Col050()).To(ast.SetType[string](ast.NewColumnNode("v", "col_050"))),
	Set(wide_records.Col051()).To(ast.SetType[string](ast.NewColumnNode("v", "col_051"))),
	Set(wide_records.Col052()).To(ast.SetType[string](ast.NewColumnNode("v", "col_052"))),
	Set(wide_records.Col053()).To(ast.SetType[string](ast.NewColumnNode("v", "col_053"))),
	Set(wide_records.Col054()).To(ast.SetType[string](ast.NewColumnNode("v", "col_054"))),
	Set(wide_records.Col055()).To(ast.SetType[string](ast.NewColumnNode("v", "col_055"))),
	Set(wide_records.Col056()).To(ast.SetType[string](ast.NewColumnNode("v", "col_056"))),
	Set(wide_records.Col057()).To(ast.SetType[string](ast.NewColumnNode("v", "col_057"))),
	Set(wide_records.Col058()).To(ast.SetType[string](ast.NewColumnNode("v", "col_058"))),
	Set(wide_records.Col059()).To(ast.SetType[string](ast.NewColumnNode("v", "col_059"))),
	Set(wide_records.Col060()).To(ast.SetType[string](ast.NewColumnNode("v", "col_060"))),
	Set(wide_records.Col061()).To(ast.SetType[string](ast.NewColumnNode("v", "col_061"))),
	Set(wide_records.Col062()).To(ast.SetType[string](ast.NewColumnNode("v", "col_062"))),
	Set(wide_records.Col063()).To(ast.SetType[string](ast.NewColumnNode("v", "col_063"))),
	Set(wide_records.Col064()).To(ast.SetType[string](ast.NewColumnNode("v", "col_064"))),
	Set(wide_records.Col065()).To(ast.SetType[string](ast.NewColumnNode("v", "col_065"))),
	Set(wide_records.Col066()).To(ast.SetType[string](ast.NewColumnNode("v", "col_066"))),
	Set(wide_records.Col067()).To(ast.SetType[string](ast.NewColumnNode("v", "col_067"))),
	Set(wide_records.Col068()).To(ast.SetType[string](ast.NewColumnNode("v", "col_068"))),
	Set(wide_records.Col069()).To(ast.SetType[string](ast.NewColumnNode("v", "col_069"))),
	Set(wide_records.Col070()).To(ast.SetType[string](ast.NewColumnNode("v", "col_070"))),
	Set(wide_records.Col071()).To(ast.SetType[string](ast.NewColumnNode("v", "col_071"))),
	Set(wide_records.Col072()).To(ast.SetType[string](ast.NewColumnNode("v", "col_072"))),
	Set(wide_records.Col073()).To(ast.SetType[string](ast.NewColumnNode("v", "col_073"))),
	Set(wide_records.Col074()).To(ast.SetType[string](ast.NewColumnNode("v", "col_074"))),
	Set(wide_records.Col075()).To(ast.SetType[string](ast.NewColumnNode("v", "col_075"))),
	Set(wide_records.Col076()).To(ast.SetType[string](ast.NewColumnNode("v", "col_076"))),
	Set(wide_records.Col077()).To(ast.SetType[string](ast.NewColumnNode("v", "col_077"))),
	Set(wide_records.Col078()).To(ast.SetType[string](ast.NewColumnNode("v", "col_078"))),
	Set(wide_records.Col079()).To(ast.SetType[string](ast.NewColumnNode("v", "col_079"))),
	Set(wide_records.Col080()).To(ast.SetType[string](ast.NewColumnNode("v", "col_080"))),
	Set(wide_records.Col081()).To(ast.SetType[string](ast.NewColumnNode("v", "col_081"))),
	Set(wide_records.Col082()).To(ast.SetType[string](ast.NewColumnNode("v", "col_082"))),
	Set(wide_records.Col083()).To(ast.SetType[string](ast.NewColumnNode("v", "col_083"))),
	Set(wide_records.Col084()).To(ast.SetType[string](ast.NewColumnNode("v", "col_084"))),
	Set(wide_records.Col085()).To(ast.SetType[string](ast.NewColumnNode("v", "col_085"))),
	Set(wide_records.Col086()).To(ast.SetType[string](ast.NewColumnNode("v", "col_086"))),
	Set(wide_records.Col087()).To(ast.SetType[string](ast.NewColumnNode("v", "col_087"))),
	Set(wide_records.Col088()).To(ast.SetType[string](ast.NewColumnNode("v", "col_088"))),
	Set(wide_records.Col089()).To(ast.SetType[string](ast.NewColumnNode("v", "col_089"))),
	Set(wide_records.Col090()).To(ast.SetType[string](ast.NewColumnNode("v", "col_090"))),
	Set(wide_records.Col091()).To(ast.SetType[string](ast.NewColumnNode("v", "col_091"))),
	Set(wide_records.Col092()).To(ast.SetType[string](ast.NewColumnNode("v", "col_092"))),
	Set(wide_records.Col093()).To(ast.SetType[string](ast.NewColumnNode("v", "col_093"))),
	Set(wide_records.Col094()).To(ast.SetType[string](ast.NewColumnNode("v", "col_094"))),
	Set(wide_records.Col095()).To(ast.SetType[string](ast.NewColumnNode("v", "col_095"))),
	Set(wide_records.Col096()).To(ast.SetType[string](ast.NewColumnNode("v", "col_096"))),
	Set(wide_records.Col097()).To(ast.SetType[string](ast.NewColumnNode("v", "col_097"))),
	Set(wide_records.Col098()).To(ast.SetType[string](ast.NewColumnNode("v", "col_098"))),
	Set(wide_records.Col099()).To(ast.SetType[string](ast.NewColumnNode("v", "col_099"))),
	Set(wide_records.Col100()).To(ast.SetType[string](ast.NewColumnNode("v", "col_100"))),
}

var defaultWideRecordsDtoUpdateSetByValueColumn = []ast.UpdateSetExpr{
	Set(wide_records.Col001()).To(ast.SetType[string](ast.NewColumnNode("v", "col_001"))),
	Set(wide_records.Col002()).To(ast.SetType[string](ast.NewColumnNode("v", "col_002"))),
	Set(wide_records.Col003()).To(ast.SetType[string](ast.NewColumnNode("v", "col_003"))),
	Set(wide_records.Col004()).To(ast.SetType[string](ast.NewColumnNode("v", "col_004"))),
	Set(wide_records.Col005()).To(ast.SetType[string](ast.NewColumnNode("v", "col_005"))),
	Set(wide_records.Col006()).To(ast.SetType[string](ast.NewColumnNode("v", "col_006"))),
	Set(wide_records.Col007()).To(ast.SetType[string](ast.NewColumnNode("v", "col_007"))),
	Set(wide_records.Col008()).To(ast.SetType[string](ast.NewColumnNode("v", "col_008"))),
	Set(wide_records.Col009()).To(ast.SetType[string](ast.NewColumnNode("v", "col_009"))),
	Set(wide_records.Col010()).To(ast.SetType[string](ast.NewColumnNode("v", "col_010"))),
	Set(wide_records.Col011()).To(ast.SetType[string](ast.NewColumnNode("v", "col_011"))),
	Set(wide_records.Col012()).To(ast.SetType[string](ast.NewColumnNode("v", "col_012"))),
	Set(wide_records.Col013()).To(ast.SetType[string](ast.NewColumnNode("v", "col_013"))),
	Set(wide_records.Col014()).To(ast.SetType[string](ast.NewColumnNode("v", "col_014"))),
	Set(wide_records.Col015()).To(ast.SetType[string](ast.NewColumnNode("v", "col_015"))),
	Set(wide_records.Col016()).To(ast.SetType[string](ast.NewColumnNode("v", "col_016"))),
	Set(wide_records.Col017()).To(ast.SetType[string](ast.NewColumnNode("v", "col_017"))),
	Set(wide_records.Col018()).To(ast.SetType[string](ast.NewColumnNode("v", "col_018"))),
	Set(wide_records.Col019()).To(ast.SetType[string](ast.NewColumnNode("v", "col_019"))),
	Set(wide_records.Col020()).To(ast.SetType[string](ast.NewColumnNode("v", "col_020"))),
	Set(wide_records.Col021()).To(ast.SetType[string](ast.NewColumnNode("v", "col_021"))),
	Set(wide_records.Col022()).To(ast.SetType[string](ast.NewColumnNode("v", "col_022"))),
	Set(wide_records.Col023()).To(ast.SetType[string](ast.NewColumnNode("v", "col_023"))),
	Set(wide_records.Col024()).To(ast.SetType[string](ast.NewColumnNode("v", "col_024"))),
	Set(wide_records.Col025()).To(ast.SetType[string](ast.NewColumnNode("v", "col_025"))),
	Set(wide_records.Col026()).To(ast.SetType[string](ast.NewColumnNode("v", "col_026"))),
	Set(wide_records.Col027()).To(ast.SetType[string](ast.NewColumnNode("v", "col_027"))),
	Set(wide_records.Col028()).To(ast.SetType[string](ast.NewColumnNode("v", "col_028"))),
	Set(wide_records.Col029()).To(ast.SetType[string](ast.NewColumnNode("v", "col_029"))),
	Set(wide_records.Col030()).To(ast.SetType[string](ast.NewColumnNode("v", "col_030"))),
	Set(wide_records.Col031()).To(ast.SetType[string](ast.NewColumnNode("v", "col_031"))),
	Set(wide_records.Col032()).To(ast.SetType[string](ast.NewColumnNode("v", "col_032"))),
	Set(wide_records.Col033()).To(ast.SetType[string](ast.NewColumnNode("v", "col_033"))),
	Set(wide_records.Col034()).To(ast.SetType[string](ast.NewColumnNode("v", "col_034"))),
	Set(wide_records.Col035()).To(ast.SetType[string](ast.NewColumnNode("v", "col_035"))),
	Set(wide_records.Col036()).To(ast.SetType[string](ast.NewColumnNode("v", "col_036"))),
	Set(wide_records.Col037()).To(ast.SetType[string](ast.NewColumnNode("v", "col_037"))),
	Set(wide_records.Col038()).To(ast.SetType[string](ast.NewColumnNode("v", "col_038"))),
	Set(wide_records.Col039()).To(ast.SetType[string](ast.NewColumnNode("v", "col_039"))),
	Set(wide_records.Col040()).To(ast.SetType[string](ast.NewColumnNode("v", "col_040"))),
	Set(wide_records.Col041()).To(ast.SetType[string](ast.NewColumnNode("v", "col_041"))),
	Set(wide_records.Col042()).To(ast.SetType[string](ast.NewColumnNode("v", "col_042"))),
	Set(wide_records.Col043()).To(ast.SetType[string](ast.NewColumnNode("v", "col_043"))),
	Set(wide_records.Col044()).To(ast.SetType[string](ast.NewColumnNode("v", "col_044"))),
	Set(wide_records.Col045()).To(ast.SetType[string](ast.NewColumnNode("v", "col_045"))),
	Set(wide_records.Col046()).To(ast.SetType[string](ast.NewColumnNode("v", "col_046"))),
	Set(wide_records.Col047()).To(ast.SetType[string](ast.NewColumnNode("v", "col_047"))),
	Set(wide_records.Col048()).To(ast.SetType[string](ast.NewColumnNode("v", "col_048"))),
	Set(wide_records.Col049()).To(ast.SetType[string](ast.NewColumnNode("v", "col_049"))),
	Set(wide_records.Col050()).To(ast.SetType[string](ast.NewColumnNode("v", "col_050"))),
	Set(wide_records.Col051()).To(ast.SetType[string](ast.NewColumnNode("v", "col_051"))),
	Set(wide_records.Col052()).To(ast.SetType[string](ast.NewColumnNode("v", "col_052"))),
	Set(wide_records.Col053()).To(ast.SetType[string](ast.NewColumnNode("v", "col_053"))),
	Set(wide_records.Col054()).To(ast.SetType[string](ast.NewColumnNode("v", "col_054"))),
	Set(wide_records.Col055()).To(ast.SetType[string](ast.NewColumnNode("v", "col_055"))),
	Set(wide_records.Col056()).To(ast.SetType[string](ast.NewColumnNode("v", "col_056"))),
	Set(wide_records.Col057()).To(ast.SetType[string](ast.NewColumnNode("v", "col_057"))),
	Set(wide_records.Col058()).To(ast.SetType[string](ast.NewColumnNode("v", "col_058"))),
	Set(wide_records.Col059()).To(ast.SetType[string](ast.NewColumnNode("v", "col_059"))),
	Set(wide_records.Col060()).To(ast.SetType[string](ast.NewColumnNode("v", "col_060"))),
	Set(wide_records.Col061()).To(ast.SetType[string](ast.NewColumnNode("v", "col_061"))),
	Set(wide_records.Col062()).To(ast.SetType[string](ast.NewColumnNode("v", "col_062"))),
	Set(wide_records.Col063()).To(ast.SetType[string](ast.NewColumnNode("v", "col_063"))),
	Set(wide_records.Col064()).To(ast.SetType[string](ast.NewColumnNode("v", "col_064"))),
	Set(wide_records.Col065()).To(ast.SetType[string](ast.NewColumnNode("v", "col_065"))),
	Set(wide_records.Col066()).To(ast.SetType[string](ast.NewColumnNode("v", "col_066"))),
	Set(wide_records.Col067()).To(ast.SetType[string](ast.NewColumnNode("v", "col_067"))),
	Set(wide_records.Col068()).To(ast.SetType[string](ast.NewColumnNode("v", "col_068"))),
	Set(wide_records.Col069()).To(ast.SetType[string](ast.NewColumnNode("v", "col_069"))),
	Set(wide_records.Col070()).To(ast.SetType[string](ast.NewColumnNode("v", "col_070"))),
	Set(wide_records.Col071()).To(ast.SetType[string](ast.NewColumnNode("v", "col_071"))),
	Set(wide_records.Col072()).To(ast.SetType[string](ast.NewColumnNode("v", "col_072"))),
	Set(wide_records.Col073()).To(ast.SetType[string](ast.NewColumnNode("v", "col_073"))),
	Set(wide_records.Col074()).To(ast.SetType[string](ast.NewColumnNode("v", "col_074"))),
	Set(wide_records.Col075()).To(ast.SetType[string](ast.NewColumnNode("v", "col_075"))),
	Set(wide_records.Col076()).To(ast.SetType[string](ast.NewColumnNode("v", "col_076"))),
	Set(wide_records.Col077()).To(ast.SetType[string](ast.NewColumnNode("v", "col_077"))),
	Set(wide_records.Col078()).To(ast.SetType[string](ast.NewColumnNode("v", "col_078"))),
	Set(wide_records.Col079()).To(ast.SetType[string](ast.NewColumnNode("v", "col_079"))),
	Set(wide_records.Col080()).To(ast.SetType[string](ast.NewColumnNode("v", "col_080"))),
	Set(wide_records.Col081()).To(ast.SetType[string](ast.NewColumnNode("v", "col_081"))),
	Set(wide_records.Col082()).To(ast.SetType[string](ast.NewColumnNode("v", "col_082"))),
	Set(wide_records.Col083()).To(ast.SetType[string](ast.NewColumnNode("v", "col_083"))),
	Set(wide_records.Col084()).To(ast.SetType[string](ast.NewColumnNode("v", "col_084"))),
	Set(wide_records.Col085()).To(ast.SetType[string](ast.NewColumnNode("v", "col_085"))),
	Set(wide_records.Col086()).To(ast.SetType[string](ast.NewColumnNode("v", "col_086"))),
	Set(wide_records.Col087()).To(ast.SetType[string](ast.NewColumnNode("v", "col_087"))),
	Set(wide_records.Col088()).To(ast.SetType[string](ast.NewColumnNode("v", "col_088"))),
	Set(wide_records.Col089()).To(ast.SetType[string](ast.NewColumnNode("v", "col_089"))),
	Set(wide_records.Col090()).To(ast.SetType[string](ast.NewColumnNode("v", "col_090"))),
	Set(wide_records.Col091()).To(ast.SetType[string](ast.NewColumnNode("v", "col_091"))),
	Set(wide_records.Col092()).To(ast.SetType[string](ast.NewColumnNode("v", "col_092"))),
	Set(wide_records.Col093()).To(ast.SetType[string](ast.NewColumnNode("v", "col_093"))),
	Set(wide_records.Col094()).To(ast.SetType[string](ast.NewColumnNode("v", "col_094"))),
	Set(wide_records.Col095()).To(ast.SetType[string](ast.NewColumnNode("v", "col_095"))),
	Set(wide_records.Col096()).To(ast.SetType[string](ast.NewColumnNode("v", "col_096"))),
	Set(wide_records.Col097()).To(ast.SetType[string](ast.NewColumnNode("v", "col_097"))),
	Set(wide_records.Col098()).To(ast.SetType[string](ast.NewColumnNode("v", "col_098"))),
	Set(wide_records.Col099()).To(ast.SetType[string](ast.NewColumnNode("v", "col_099"))),
	Set(wide_records.Col100()).To(ast.SetType[string](ast.NewColumnNode("v", "col_100"))),
	nil,
}

func WideRecordsDtoUpdateOne(pool *pgxpool.Pool, dto *models.WideRecordsDto, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.WideRecordsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	updateColumns, _, _, _, _, err := selectedWideRecordsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	return wide_records.NewQuery(pool).
		Update(SetTo(dto, updateWideRecordsDtoProjections(updateColumns)...)).
		Where(wide_records.Id().Eq(Val(*dto.Id))), nil
}

func WideRecordsDtoUpdateMany(pool *pgxpool.Pool, dtos models.WideRecordsDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.WideRecordsDto], error) {
	options, err := newUpdateOptions(opts...)
	if err != nil {
		return nil, err
	}
	valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, err := selectedWideRecordsDtoUpdateExpressions(options)
	if err != nil {
		return nil, err
	}
	v := wide_records.As("v", valueColumns...)
	if len(sets) == 0 {
		sets = append(sets, ast.UpdateSetList{})
	}
	var from ast.NamedTableExpression = Values(WideRecordsDtoUpdateValuesProvider{
		dtos:       dtos,
		extractors: valueExtractors,
		castTypes:  valueCastTypes,
	}).As(v.Alias)
	if options.source == updateBatchSourceUnnest {
		from = Unnest(WideRecordsDtoUpdateUnnestExpressions(dtos, updateColumnIndexes)...).As(v.Alias)
	}
	return wide_records.NewQuery(pool).
		Update(sets...).
		From(from).
		Where(wide_records.Id().Eq(v.Id())), nil
}

func selectedWideRecordsDtoUpdateExpressions(options updateOptions) ([]ast.NamedExpression, []WideRecordsDtoUpdateValueExtractor, []string, []int, []ast.UpdateSetExpr, error) {
	if options.columns == nil {
		return defaultWideRecordsDtoUpdateValueColumns, defaultWideRecordsDtoUpdateValueExtractors, defaultWideRecordsDtoUpdateValueCastTypes, defaultWideRecordsDtoUpdateColumnIndexes, defaultWideRecordsDtoUpdateSets, nil
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
	valueColumns = append(valueColumns, primaryKeyWideRecordsDtoUpdateValueColumns...)
	valueExtractors := make([]WideRecordsDtoUpdateValueExtractor, 0, len(requested)+1)
	valueExtractors = append(valueExtractors, primaryKeyWideRecordsDtoUpdateValueExtractors...)
	valueCastTypes := make([]string, 0, len(requested)+1)
	valueCastTypes = append(valueCastTypes, primaryKeyWideRecordsDtoUpdateValueCastTypes...)
	updateColumnIndexes := make([]int, 0, len(requested)+1)
	updateColumnIndexes = append(updateColumnIndexes, primaryKeyWideRecordsDtoUpdateColumnIndexes...)
	sets := make([]ast.UpdateSetExpr, 0, len(requested))

	knownIndex := 0
	for _, column := range requested {
		name := column.Name()
		for knownIndex < len(defaultWideRecordsDtoUpdateValueColumns) && defaultWideRecordsDtoUpdateValueColumns[knownIndex].Name() < name {
			knownIndex++
		}
		if knownIndex == len(defaultWideRecordsDtoUpdateValueColumns) || defaultWideRecordsDtoUpdateValueColumns[knownIndex].Name() != name {
			return nil, nil, nil, nil, nil, fmt.Errorf("unknown update column %q", name)
		}
		set := defaultWideRecordsDtoUpdateSetByValueColumn[knownIndex]
		if set == nil {
			return nil, nil, nil, nil, nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		valueColumns = append(valueColumns, defaultWideRecordsDtoUpdateValueColumns[knownIndex])
		valueExtractors = append(valueExtractors, defaultWideRecordsDtoUpdateValueExtractors[knownIndex])
		valueCastTypes = append(valueCastTypes, defaultWideRecordsDtoUpdateValueCastTypes[knownIndex])
		updateColumnIndexes = append(updateColumnIndexes, knownIndex)
		sets = append(sets, set)
	}
	return valueColumns, valueExtractors, valueCastTypes, updateColumnIndexes, sets, nil
}

func updateWideRecordsDtoProjections(columns []ast.NamedExpression) []ast.Projection[models.WideRecordsDto] {
	projections := make([]ast.Projection[models.WideRecordsDto], 0, len(columns))
	for _, column := range columns {
		switch column.Name() {
		case "col_001":
			projections = append(projections, wide_records.Col001())
		case "col_002":
			projections = append(projections, wide_records.Col002())
		case "col_003":
			projections = append(projections, wide_records.Col003())
		case "col_004":
			projections = append(projections, wide_records.Col004())
		case "col_005":
			projections = append(projections, wide_records.Col005())
		case "col_006":
			projections = append(projections, wide_records.Col006())
		case "col_007":
			projections = append(projections, wide_records.Col007())
		case "col_008":
			projections = append(projections, wide_records.Col008())
		case "col_009":
			projections = append(projections, wide_records.Col009())
		case "col_010":
			projections = append(projections, wide_records.Col010())
		case "col_011":
			projections = append(projections, wide_records.Col011())
		case "col_012":
			projections = append(projections, wide_records.Col012())
		case "col_013":
			projections = append(projections, wide_records.Col013())
		case "col_014":
			projections = append(projections, wide_records.Col014())
		case "col_015":
			projections = append(projections, wide_records.Col015())
		case "col_016":
			projections = append(projections, wide_records.Col016())
		case "col_017":
			projections = append(projections, wide_records.Col017())
		case "col_018":
			projections = append(projections, wide_records.Col018())
		case "col_019":
			projections = append(projections, wide_records.Col019())
		case "col_020":
			projections = append(projections, wide_records.Col020())
		case "col_021":
			projections = append(projections, wide_records.Col021())
		case "col_022":
			projections = append(projections, wide_records.Col022())
		case "col_023":
			projections = append(projections, wide_records.Col023())
		case "col_024":
			projections = append(projections, wide_records.Col024())
		case "col_025":
			projections = append(projections, wide_records.Col025())
		case "col_026":
			projections = append(projections, wide_records.Col026())
		case "col_027":
			projections = append(projections, wide_records.Col027())
		case "col_028":
			projections = append(projections, wide_records.Col028())
		case "col_029":
			projections = append(projections, wide_records.Col029())
		case "col_030":
			projections = append(projections, wide_records.Col030())
		case "col_031":
			projections = append(projections, wide_records.Col031())
		case "col_032":
			projections = append(projections, wide_records.Col032())
		case "col_033":
			projections = append(projections, wide_records.Col033())
		case "col_034":
			projections = append(projections, wide_records.Col034())
		case "col_035":
			projections = append(projections, wide_records.Col035())
		case "col_036":
			projections = append(projections, wide_records.Col036())
		case "col_037":
			projections = append(projections, wide_records.Col037())
		case "col_038":
			projections = append(projections, wide_records.Col038())
		case "col_039":
			projections = append(projections, wide_records.Col039())
		case "col_040":
			projections = append(projections, wide_records.Col040())
		case "col_041":
			projections = append(projections, wide_records.Col041())
		case "col_042":
			projections = append(projections, wide_records.Col042())
		case "col_043":
			projections = append(projections, wide_records.Col043())
		case "col_044":
			projections = append(projections, wide_records.Col044())
		case "col_045":
			projections = append(projections, wide_records.Col045())
		case "col_046":
			projections = append(projections, wide_records.Col046())
		case "col_047":
			projections = append(projections, wide_records.Col047())
		case "col_048":
			projections = append(projections, wide_records.Col048())
		case "col_049":
			projections = append(projections, wide_records.Col049())
		case "col_050":
			projections = append(projections, wide_records.Col050())
		case "col_051":
			projections = append(projections, wide_records.Col051())
		case "col_052":
			projections = append(projections, wide_records.Col052())
		case "col_053":
			projections = append(projections, wide_records.Col053())
		case "col_054":
			projections = append(projections, wide_records.Col054())
		case "col_055":
			projections = append(projections, wide_records.Col055())
		case "col_056":
			projections = append(projections, wide_records.Col056())
		case "col_057":
			projections = append(projections, wide_records.Col057())
		case "col_058":
			projections = append(projections, wide_records.Col058())
		case "col_059":
			projections = append(projections, wide_records.Col059())
		case "col_060":
			projections = append(projections, wide_records.Col060())
		case "col_061":
			projections = append(projections, wide_records.Col061())
		case "col_062":
			projections = append(projections, wide_records.Col062())
		case "col_063":
			projections = append(projections, wide_records.Col063())
		case "col_064":
			projections = append(projections, wide_records.Col064())
		case "col_065":
			projections = append(projections, wide_records.Col065())
		case "col_066":
			projections = append(projections, wide_records.Col066())
		case "col_067":
			projections = append(projections, wide_records.Col067())
		case "col_068":
			projections = append(projections, wide_records.Col068())
		case "col_069":
			projections = append(projections, wide_records.Col069())
		case "col_070":
			projections = append(projections, wide_records.Col070())
		case "col_071":
			projections = append(projections, wide_records.Col071())
		case "col_072":
			projections = append(projections, wide_records.Col072())
		case "col_073":
			projections = append(projections, wide_records.Col073())
		case "col_074":
			projections = append(projections, wide_records.Col074())
		case "col_075":
			projections = append(projections, wide_records.Col075())
		case "col_076":
			projections = append(projections, wide_records.Col076())
		case "col_077":
			projections = append(projections, wide_records.Col077())
		case "col_078":
			projections = append(projections, wide_records.Col078())
		case "col_079":
			projections = append(projections, wide_records.Col079())
		case "col_080":
			projections = append(projections, wide_records.Col080())
		case "col_081":
			projections = append(projections, wide_records.Col081())
		case "col_082":
			projections = append(projections, wide_records.Col082())
		case "col_083":
			projections = append(projections, wide_records.Col083())
		case "col_084":
			projections = append(projections, wide_records.Col084())
		case "col_085":
			projections = append(projections, wide_records.Col085())
		case "col_086":
			projections = append(projections, wide_records.Col086())
		case "col_087":
			projections = append(projections, wide_records.Col087())
		case "col_088":
			projections = append(projections, wide_records.Col088())
		case "col_089":
			projections = append(projections, wide_records.Col089())
		case "col_090":
			projections = append(projections, wide_records.Col090())
		case "col_091":
			projections = append(projections, wide_records.Col091())
		case "col_092":
			projections = append(projections, wide_records.Col092())
		case "col_093":
			projections = append(projections, wide_records.Col093())
		case "col_094":
			projections = append(projections, wide_records.Col094())
		case "col_095":
			projections = append(projections, wide_records.Col095())
		case "col_096":
			projections = append(projections, wide_records.Col096())
		case "col_097":
			projections = append(projections, wide_records.Col097())
		case "col_098":
			projections = append(projections, wide_records.Col098())
		case "col_099":
			projections = append(projections, wide_records.Col099())
		case "col_100":
			projections = append(projections, wide_records.Col100())
		}
	}
	return projections
}

func defaultWideRecordsDtoUpdateColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		wide_records.Col001(),
		wide_records.Col002(),
		wide_records.Col003(),
		wide_records.Col004(),
		wide_records.Col005(),
		wide_records.Col006(),
		wide_records.Col007(),
		wide_records.Col008(),
		wide_records.Col009(),
		wide_records.Col010(),
		wide_records.Col011(),
		wide_records.Col012(),
		wide_records.Col013(),
		wide_records.Col014(),
		wide_records.Col015(),
		wide_records.Col016(),
		wide_records.Col017(),
		wide_records.Col018(),
		wide_records.Col019(),
		wide_records.Col020(),
		wide_records.Col021(),
		wide_records.Col022(),
		wide_records.Col023(),
		wide_records.Col024(),
		wide_records.Col025(),
		wide_records.Col026(),
		wide_records.Col027(),
		wide_records.Col028(),
		wide_records.Col029(),
		wide_records.Col030(),
		wide_records.Col031(),
		wide_records.Col032(),
		wide_records.Col033(),
		wide_records.Col034(),
		wide_records.Col035(),
		wide_records.Col036(),
		wide_records.Col037(),
		wide_records.Col038(),
		wide_records.Col039(),
		wide_records.Col040(),
		wide_records.Col041(),
		wide_records.Col042(),
		wide_records.Col043(),
		wide_records.Col044(),
		wide_records.Col045(),
		wide_records.Col046(),
		wide_records.Col047(),
		wide_records.Col048(),
		wide_records.Col049(),
		wide_records.Col050(),
		wide_records.Col051(),
		wide_records.Col052(),
		wide_records.Col053(),
		wide_records.Col054(),
		wide_records.Col055(),
		wide_records.Col056(),
		wide_records.Col057(),
		wide_records.Col058(),
		wide_records.Col059(),
		wide_records.Col060(),
		wide_records.Col061(),
		wide_records.Col062(),
		wide_records.Col063(),
		wide_records.Col064(),
		wide_records.Col065(),
		wide_records.Col066(),
		wide_records.Col067(),
		wide_records.Col068(),
		wide_records.Col069(),
		wide_records.Col070(),
		wide_records.Col071(),
		wide_records.Col072(),
		wide_records.Col073(),
		wide_records.Col074(),
		wide_records.Col075(),
		wide_records.Col076(),
		wide_records.Col077(),
		wide_records.Col078(),
		wide_records.Col079(),
		wide_records.Col080(),
		wide_records.Col081(),
		wide_records.Col082(),
		wide_records.Col083(),
		wide_records.Col084(),
		wide_records.Col085(),
		wide_records.Col086(),
		wide_records.Col087(),
		wide_records.Col088(),
		wide_records.Col089(),
		wide_records.Col090(),
		wide_records.Col091(),
		wide_records.Col092(),
		wide_records.Col093(),
		wide_records.Col094(),
		wide_records.Col095(),
		wide_records.Col096(),
		wide_records.Col097(),
		wide_records.Col098(),
		wide_records.Col099(),
		wide_records.Col100(),
	}
}

func WideRecordsDtoUpdateUnnestExpressions(dtos models.WideRecordsDtos, columnIndexes []int) []ast.Expression {
	unnestExpressions := make([]ast.Expression, len(columnIndexes))
	for i, columnIndex := range columnIndexes {
		switch columnIndex {
		case 0:
			unnestExpressions[i] = Cast(dtos.Col001()).AsTextArray()
		case 1:
			unnestExpressions[i] = Cast(dtos.Col002()).AsTextArray()
		case 2:
			unnestExpressions[i] = Cast(dtos.Col003()).AsTextArray()
		case 3:
			unnestExpressions[i] = Cast(dtos.Col004()).AsTextArray()
		case 4:
			unnestExpressions[i] = Cast(dtos.Col005()).AsTextArray()
		case 5:
			unnestExpressions[i] = Cast(dtos.Col006()).AsTextArray()
		case 6:
			unnestExpressions[i] = Cast(dtos.Col007()).AsTextArray()
		case 7:
			unnestExpressions[i] = Cast(dtos.Col008()).AsTextArray()
		case 8:
			unnestExpressions[i] = Cast(dtos.Col009()).AsTextArray()
		case 9:
			unnestExpressions[i] = Cast(dtos.Col010()).AsTextArray()
		case 10:
			unnestExpressions[i] = Cast(dtos.Col011()).AsTextArray()
		case 11:
			unnestExpressions[i] = Cast(dtos.Col012()).AsTextArray()
		case 12:
			unnestExpressions[i] = Cast(dtos.Col013()).AsTextArray()
		case 13:
			unnestExpressions[i] = Cast(dtos.Col014()).AsTextArray()
		case 14:
			unnestExpressions[i] = Cast(dtos.Col015()).AsTextArray()
		case 15:
			unnestExpressions[i] = Cast(dtos.Col016()).AsTextArray()
		case 16:
			unnestExpressions[i] = Cast(dtos.Col017()).AsTextArray()
		case 17:
			unnestExpressions[i] = Cast(dtos.Col018()).AsTextArray()
		case 18:
			unnestExpressions[i] = Cast(dtos.Col019()).AsTextArray()
		case 19:
			unnestExpressions[i] = Cast(dtos.Col020()).AsTextArray()
		case 20:
			unnestExpressions[i] = Cast(dtos.Col021()).AsTextArray()
		case 21:
			unnestExpressions[i] = Cast(dtos.Col022()).AsTextArray()
		case 22:
			unnestExpressions[i] = Cast(dtos.Col023()).AsTextArray()
		case 23:
			unnestExpressions[i] = Cast(dtos.Col024()).AsTextArray()
		case 24:
			unnestExpressions[i] = Cast(dtos.Col025()).AsTextArray()
		case 25:
			unnestExpressions[i] = Cast(dtos.Col026()).AsTextArray()
		case 26:
			unnestExpressions[i] = Cast(dtos.Col027()).AsTextArray()
		case 27:
			unnestExpressions[i] = Cast(dtos.Col028()).AsTextArray()
		case 28:
			unnestExpressions[i] = Cast(dtos.Col029()).AsTextArray()
		case 29:
			unnestExpressions[i] = Cast(dtos.Col030()).AsTextArray()
		case 30:
			unnestExpressions[i] = Cast(dtos.Col031()).AsTextArray()
		case 31:
			unnestExpressions[i] = Cast(dtos.Col032()).AsTextArray()
		case 32:
			unnestExpressions[i] = Cast(dtos.Col033()).AsTextArray()
		case 33:
			unnestExpressions[i] = Cast(dtos.Col034()).AsTextArray()
		case 34:
			unnestExpressions[i] = Cast(dtos.Col035()).AsTextArray()
		case 35:
			unnestExpressions[i] = Cast(dtos.Col036()).AsTextArray()
		case 36:
			unnestExpressions[i] = Cast(dtos.Col037()).AsTextArray()
		case 37:
			unnestExpressions[i] = Cast(dtos.Col038()).AsTextArray()
		case 38:
			unnestExpressions[i] = Cast(dtos.Col039()).AsTextArray()
		case 39:
			unnestExpressions[i] = Cast(dtos.Col040()).AsTextArray()
		case 40:
			unnestExpressions[i] = Cast(dtos.Col041()).AsTextArray()
		case 41:
			unnestExpressions[i] = Cast(dtos.Col042()).AsTextArray()
		case 42:
			unnestExpressions[i] = Cast(dtos.Col043()).AsTextArray()
		case 43:
			unnestExpressions[i] = Cast(dtos.Col044()).AsTextArray()
		case 44:
			unnestExpressions[i] = Cast(dtos.Col045()).AsTextArray()
		case 45:
			unnestExpressions[i] = Cast(dtos.Col046()).AsTextArray()
		case 46:
			unnestExpressions[i] = Cast(dtos.Col047()).AsTextArray()
		case 47:
			unnestExpressions[i] = Cast(dtos.Col048()).AsTextArray()
		case 48:
			unnestExpressions[i] = Cast(dtos.Col049()).AsTextArray()
		case 49:
			unnestExpressions[i] = Cast(dtos.Col050()).AsTextArray()
		case 50:
			unnestExpressions[i] = Cast(dtos.Col051()).AsTextArray()
		case 51:
			unnestExpressions[i] = Cast(dtos.Col052()).AsTextArray()
		case 52:
			unnestExpressions[i] = Cast(dtos.Col053()).AsTextArray()
		case 53:
			unnestExpressions[i] = Cast(dtos.Col054()).AsTextArray()
		case 54:
			unnestExpressions[i] = Cast(dtos.Col055()).AsTextArray()
		case 55:
			unnestExpressions[i] = Cast(dtos.Col056()).AsTextArray()
		case 56:
			unnestExpressions[i] = Cast(dtos.Col057()).AsTextArray()
		case 57:
			unnestExpressions[i] = Cast(dtos.Col058()).AsTextArray()
		case 58:
			unnestExpressions[i] = Cast(dtos.Col059()).AsTextArray()
		case 59:
			unnestExpressions[i] = Cast(dtos.Col060()).AsTextArray()
		case 60:
			unnestExpressions[i] = Cast(dtos.Col061()).AsTextArray()
		case 61:
			unnestExpressions[i] = Cast(dtos.Col062()).AsTextArray()
		case 62:
			unnestExpressions[i] = Cast(dtos.Col063()).AsTextArray()
		case 63:
			unnestExpressions[i] = Cast(dtos.Col064()).AsTextArray()
		case 64:
			unnestExpressions[i] = Cast(dtos.Col065()).AsTextArray()
		case 65:
			unnestExpressions[i] = Cast(dtos.Col066()).AsTextArray()
		case 66:
			unnestExpressions[i] = Cast(dtos.Col067()).AsTextArray()
		case 67:
			unnestExpressions[i] = Cast(dtos.Col068()).AsTextArray()
		case 68:
			unnestExpressions[i] = Cast(dtos.Col069()).AsTextArray()
		case 69:
			unnestExpressions[i] = Cast(dtos.Col070()).AsTextArray()
		case 70:
			unnestExpressions[i] = Cast(dtos.Col071()).AsTextArray()
		case 71:
			unnestExpressions[i] = Cast(dtos.Col072()).AsTextArray()
		case 72:
			unnestExpressions[i] = Cast(dtos.Col073()).AsTextArray()
		case 73:
			unnestExpressions[i] = Cast(dtos.Col074()).AsTextArray()
		case 74:
			unnestExpressions[i] = Cast(dtos.Col075()).AsTextArray()
		case 75:
			unnestExpressions[i] = Cast(dtos.Col076()).AsTextArray()
		case 76:
			unnestExpressions[i] = Cast(dtos.Col077()).AsTextArray()
		case 77:
			unnestExpressions[i] = Cast(dtos.Col078()).AsTextArray()
		case 78:
			unnestExpressions[i] = Cast(dtos.Col079()).AsTextArray()
		case 79:
			unnestExpressions[i] = Cast(dtos.Col080()).AsTextArray()
		case 80:
			unnestExpressions[i] = Cast(dtos.Col081()).AsTextArray()
		case 81:
			unnestExpressions[i] = Cast(dtos.Col082()).AsTextArray()
		case 82:
			unnestExpressions[i] = Cast(dtos.Col083()).AsTextArray()
		case 83:
			unnestExpressions[i] = Cast(dtos.Col084()).AsTextArray()
		case 84:
			unnestExpressions[i] = Cast(dtos.Col085()).AsTextArray()
		case 85:
			unnestExpressions[i] = Cast(dtos.Col086()).AsTextArray()
		case 86:
			unnestExpressions[i] = Cast(dtos.Col087()).AsTextArray()
		case 87:
			unnestExpressions[i] = Cast(dtos.Col088()).AsTextArray()
		case 88:
			unnestExpressions[i] = Cast(dtos.Col089()).AsTextArray()
		case 89:
			unnestExpressions[i] = Cast(dtos.Col090()).AsTextArray()
		case 90:
			unnestExpressions[i] = Cast(dtos.Col091()).AsTextArray()
		case 91:
			unnestExpressions[i] = Cast(dtos.Col092()).AsTextArray()
		case 92:
			unnestExpressions[i] = Cast(dtos.Col093()).AsTextArray()
		case 93:
			unnestExpressions[i] = Cast(dtos.Col094()).AsTextArray()
		case 94:
			unnestExpressions[i] = Cast(dtos.Col095()).AsTextArray()
		case 95:
			unnestExpressions[i] = Cast(dtos.Col096()).AsTextArray()
		case 96:
			unnestExpressions[i] = Cast(dtos.Col097()).AsTextArray()
		case 97:
			unnestExpressions[i] = Cast(dtos.Col098()).AsTextArray()
		case 98:
			unnestExpressions[i] = Cast(dtos.Col099()).AsTextArray()
		case 99:
			unnestExpressions[i] = Cast(dtos.Col100()).AsTextArray()
		case 100:
			unnestExpressions[i] = Cast(dtos.Id()).AsUUIDArray()
		}
	}
	return unnestExpressions
}

type WideRecordsDtoUpdateValuesProvider struct {
	dtos       models.WideRecordsDtos
	extractors []WideRecordsDtoUpdateValueExtractor
	castTypes  []string
}

func (p WideRecordsDtoUpdateValuesProvider) RowCount() int {
	return len(p.dtos)
}

func (p WideRecordsDtoUpdateValuesProvider) ColumnCount() int {
	return len(p.extractors)
}

func (p WideRecordsDtoUpdateValuesProvider) Value(row int, column int) (any, error) {
	return p.extractors[column](p.dtos[row]), nil
}

func (p WideRecordsDtoUpdateValuesProvider) ColumnSQLType(column int) (string, error) {
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
