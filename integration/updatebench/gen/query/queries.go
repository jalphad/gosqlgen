package query

import (
	"fmt"

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

type UpdateOption func(*updateOptions)

type updateOptions struct {
	columns    []ast.NamedExpression
	columnsSet bool
	source     updateBatchSource
}

func UpdateColumns(columns ...ast.NamedExpression) UpdateOption {
	return func(options *updateOptions) {
		options.columns = columns
		options.columnsSet = true
	}
}

func UpdateWithValues() UpdateOption {
	return func(options *updateOptions) {
		options.source = updateBatchSourceValues
	}
}

func newUpdateOptions(opts ...UpdateOption) updateOptions {
	options := updateOptions{
		source: updateBatchSourceUnnest,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	return options
}

func resolveUpdateColumns(requested []ast.NamedExpression, columnsSet bool, allColumns []ast.NamedExpression, updateColumns []ast.NamedExpression) ([]ast.NamedExpression, error) {
	if !columnsSet {
		return updateColumns, nil
	}
	if len(requested) == 0 {
		return []ast.NamedExpression{}, nil
	}

	allColumnNames := columnNameSet(allColumns)
	updateColumnNames := columnNameSet(updateColumns)
	requestedColumnNames := make(map[string]struct{}, len(requested))
	for _, column := range requested {
		if column == nil {
			return nil, fmt.Errorf("query.UpdateColumns received a nil column")
		}
		name := column.Name()
		if _, ok := allColumnNames[name]; !ok {
			return nil, fmt.Errorf("unknown update column %q", name)
		}
		if _, ok := updateColumnNames[name]; !ok {
			return nil, fmt.Errorf("cannot update primary key column %q", name)
		}
		requestedColumnNames[name] = struct{}{}
	}

	selected := make([]ast.NamedExpression, 0, len(requestedColumnNames))
	for _, column := range updateColumns {
		if _, ok := requestedColumnNames[column.Name()]; ok {
			selected = append(selected, column)
		}
	}
	return selected, nil
}

func columnNameSet(columns []ast.NamedExpression) map[string]struct{} {
	names := make(map[string]struct{}, len(columns))
	for _, column := range columns {
		if column != nil {
			names[column.Name()] = struct{}{}
		}
	}
	return names
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

func WideRecordsDtoUpdateOne(pool *pgxpool.Pool, dto *models.WideRecordsDto, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.WideRecordsDto], error) {
	options := newUpdateOptions(opts...)
	updateColumns, err := selectedWideRecordsDtoUpdateColumns(options)
	if err != nil {
		return nil, err
	}
	return wide_records.NewQuery(pool).
		Update(SetTo(dto, updateWideRecordsDtoProjections(updateColumns)...)).
		Where(wide_records.Id().Eq(Val(*dto.Id))), nil
}

func WideRecordsDtoUpdateMany(pool *pgxpool.Pool, dtos models.WideRecordsDtos, opts ...UpdateOption) (builder.UpdateFinalizeQuery[models.WideRecordsDto], error) {
	return wideRecordsDtoUpdateManyFast(pool, dtos, opts...)
}

func selectedWideRecordsDtoUpdateColumns(options updateOptions) ([]ast.NamedExpression, error) {
	return resolveUpdateColumns(options.columns, options.columnsSet, wide_records.AllColumns(), defaultWideRecordsDtoUpdateColumns())
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

func DeleteUser(pool *pgxpool.Pool, userId uuid.UUID) builder.DeleteFinalizeQuery[models.UsersDto] {
	return users.NewQuery(pool).
		Delete().
		Where(users.Id().Eq(Val(userId)))
}
