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

// NewQuery returns a query builder for "public"."wide_records"
func NewQuery(pool *pgxpool.Pool) *builder.KnownTableBuilder[models.WideRecordsDto] {
	return builder.NewKnownTableBuilder[models.WideRecordsDto](pool, Table())
}

var Into = intoWideRecordsDto{}

type intoWideRecordsDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource(`"public"."wide_records"`)
}

func Id() *ast.UUIDColumnProjection[models.WideRecordsDto, **uuid.UUID] {
	return newId(`"public"."wide_records"`, func(w *models.WideRecordsDto) **uuid.UUID {
		return &w.Id
	})
}

func Col001() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol001(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col001
	})
}

func Col002() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol002(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col002
	})
}

func Col003() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol003(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col003
	})
}

func Col004() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol004(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col004
	})
}

func Col005() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol005(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col005
	})
}

func Col006() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol006(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col006
	})
}

func Col007() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol007(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col007
	})
}

func Col008() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol008(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col008
	})
}

func Col009() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol009(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col009
	})
}

func Col010() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol010(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col010
	})
}

func Col011() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol011(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col011
	})
}

func Col012() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol012(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col012
	})
}

func Col013() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol013(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col013
	})
}

func Col014() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol014(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col014
	})
}

func Col015() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol015(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col015
	})
}

func Col016() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol016(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col016
	})
}

func Col017() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol017(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col017
	})
}

func Col018() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol018(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col018
	})
}

func Col019() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol019(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col019
	})
}

func Col020() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol020(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col020
	})
}

func Col021() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol021(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col021
	})
}

func Col022() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol022(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col022
	})
}

func Col023() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol023(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col023
	})
}

func Col024() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol024(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col024
	})
}

func Col025() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol025(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col025
	})
}

func Col026() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol026(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col026
	})
}

func Col027() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol027(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col027
	})
}

func Col028() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol028(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col028
	})
}

func Col029() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol029(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col029
	})
}

func Col030() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol030(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col030
	})
}

func Col031() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol031(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col031
	})
}

func Col032() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol032(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col032
	})
}

func Col033() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol033(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col033
	})
}

func Col034() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol034(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col034
	})
}

func Col035() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol035(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col035
	})
}

func Col036() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol036(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col036
	})
}

func Col037() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol037(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col037
	})
}

func Col038() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol038(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col038
	})
}

func Col039() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol039(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col039
	})
}

func Col040() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol040(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col040
	})
}

func Col041() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol041(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col041
	})
}

func Col042() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol042(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col042
	})
}

func Col043() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol043(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col043
	})
}

func Col044() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol044(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col044
	})
}

func Col045() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol045(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col045
	})
}

func Col046() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol046(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col046
	})
}

func Col047() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol047(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col047
	})
}

func Col048() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol048(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col048
	})
}

func Col049() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol049(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col049
	})
}

func Col050() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol050(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col050
	})
}

func Col051() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol051(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col051
	})
}

func Col052() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol052(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col052
	})
}

func Col053() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol053(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col053
	})
}

func Col054() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol054(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col054
	})
}

func Col055() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol055(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col055
	})
}

func Col056() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol056(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col056
	})
}

func Col057() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol057(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col057
	})
}

func Col058() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol058(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col058
	})
}

func Col059() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol059(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col059
	})
}

func Col060() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol060(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col060
	})
}

func Col061() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol061(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col061
	})
}

func Col062() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol062(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col062
	})
}

func Col063() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol063(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col063
	})
}

func Col064() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol064(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col064
	})
}

func Col065() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol065(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col065
	})
}

func Col066() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol066(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col066
	})
}

func Col067() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol067(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col067
	})
}

func Col068() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol068(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col068
	})
}

func Col069() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol069(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col069
	})
}

func Col070() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol070(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col070
	})
}

func Col071() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol071(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col071
	})
}

func Col072() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol072(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col072
	})
}

func Col073() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol073(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col073
	})
}

func Col074() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol074(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col074
	})
}

func Col075() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol075(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col075
	})
}

func Col076() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol076(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col076
	})
}

func Col077() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol077(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col077
	})
}

func Col078() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol078(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col078
	})
}

func Col079() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol079(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col079
	})
}

func Col080() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol080(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col080
	})
}

func Col081() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol081(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col081
	})
}

func Col082() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol082(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col082
	})
}

func Col083() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol083(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col083
	})
}

func Col084() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol084(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col084
	})
}

func Col085() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol085(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col085
	})
}

func Col086() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol086(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col086
	})
}

func Col087() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol087(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col087
	})
}

func Col088() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol088(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col088
	})
}

func Col089() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol089(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col089
	})
}

func Col090() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol090(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col090
	})
}

func Col091() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol091(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col091
	})
}

func Col092() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol092(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col092
	})
}

func Col093() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol093(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col093
	})
}

func Col094() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol094(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col094
	})
}

func Col095() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol095(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col095
	})
}

func Col096() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol096(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col096
	})
}

func Col097() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol097(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col097
	})
}

func Col098() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol098(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col098
	})
}

func Col099() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol099(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col099
	})
}

func Col100() *ast.StringColumnProjection[models.WideRecordsDto, *string] {
	return newCol100(`"public"."wide_records"`, func(w *models.WideRecordsDto) *string {
		return &w.Col100
	})
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

type forWideRecordsDto[T any] struct {
	dest func(*T) *models.WideRecordsDto
}

func For[T any](dest func(*T) *models.WideRecordsDto) forWideRecordsDto[T] {
	return forWideRecordsDto[T]{dest: dest}
}

func (f forWideRecordsDto[T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return newId(`"public"."wide_records"`, func(t *T) **uuid.UUID {
		dto := f.dest(t)
		return &dto.Id
	})
}

func (f forWideRecordsDto[T]) Col001() *ast.StringColumnProjection[T, *string] {
	return newCol001(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col001
	})
}

func (f forWideRecordsDto[T]) Col002() *ast.StringColumnProjection[T, *string] {
	return newCol002(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col002
	})
}

func (f forWideRecordsDto[T]) Col003() *ast.StringColumnProjection[T, *string] {
	return newCol003(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col003
	})
}

func (f forWideRecordsDto[T]) Col004() *ast.StringColumnProjection[T, *string] {
	return newCol004(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col004
	})
}

func (f forWideRecordsDto[T]) Col005() *ast.StringColumnProjection[T, *string] {
	return newCol005(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col005
	})
}

func (f forWideRecordsDto[T]) Col006() *ast.StringColumnProjection[T, *string] {
	return newCol006(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col006
	})
}

func (f forWideRecordsDto[T]) Col007() *ast.StringColumnProjection[T, *string] {
	return newCol007(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col007
	})
}

func (f forWideRecordsDto[T]) Col008() *ast.StringColumnProjection[T, *string] {
	return newCol008(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col008
	})
}

func (f forWideRecordsDto[T]) Col009() *ast.StringColumnProjection[T, *string] {
	return newCol009(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col009
	})
}

func (f forWideRecordsDto[T]) Col010() *ast.StringColumnProjection[T, *string] {
	return newCol010(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col010
	})
}

func (f forWideRecordsDto[T]) Col011() *ast.StringColumnProjection[T, *string] {
	return newCol011(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col011
	})
}

func (f forWideRecordsDto[T]) Col012() *ast.StringColumnProjection[T, *string] {
	return newCol012(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col012
	})
}

func (f forWideRecordsDto[T]) Col013() *ast.StringColumnProjection[T, *string] {
	return newCol013(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col013
	})
}

func (f forWideRecordsDto[T]) Col014() *ast.StringColumnProjection[T, *string] {
	return newCol014(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col014
	})
}

func (f forWideRecordsDto[T]) Col015() *ast.StringColumnProjection[T, *string] {
	return newCol015(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col015
	})
}

func (f forWideRecordsDto[T]) Col016() *ast.StringColumnProjection[T, *string] {
	return newCol016(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col016
	})
}

func (f forWideRecordsDto[T]) Col017() *ast.StringColumnProjection[T, *string] {
	return newCol017(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col017
	})
}

func (f forWideRecordsDto[T]) Col018() *ast.StringColumnProjection[T, *string] {
	return newCol018(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col018
	})
}

func (f forWideRecordsDto[T]) Col019() *ast.StringColumnProjection[T, *string] {
	return newCol019(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col019
	})
}

func (f forWideRecordsDto[T]) Col020() *ast.StringColumnProjection[T, *string] {
	return newCol020(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col020
	})
}

func (f forWideRecordsDto[T]) Col021() *ast.StringColumnProjection[T, *string] {
	return newCol021(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col021
	})
}

func (f forWideRecordsDto[T]) Col022() *ast.StringColumnProjection[T, *string] {
	return newCol022(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col022
	})
}

func (f forWideRecordsDto[T]) Col023() *ast.StringColumnProjection[T, *string] {
	return newCol023(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col023
	})
}

func (f forWideRecordsDto[T]) Col024() *ast.StringColumnProjection[T, *string] {
	return newCol024(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col024
	})
}

func (f forWideRecordsDto[T]) Col025() *ast.StringColumnProjection[T, *string] {
	return newCol025(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col025
	})
}

func (f forWideRecordsDto[T]) Col026() *ast.StringColumnProjection[T, *string] {
	return newCol026(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col026
	})
}

func (f forWideRecordsDto[T]) Col027() *ast.StringColumnProjection[T, *string] {
	return newCol027(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col027
	})
}

func (f forWideRecordsDto[T]) Col028() *ast.StringColumnProjection[T, *string] {
	return newCol028(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col028
	})
}

func (f forWideRecordsDto[T]) Col029() *ast.StringColumnProjection[T, *string] {
	return newCol029(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col029
	})
}

func (f forWideRecordsDto[T]) Col030() *ast.StringColumnProjection[T, *string] {
	return newCol030(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col030
	})
}

func (f forWideRecordsDto[T]) Col031() *ast.StringColumnProjection[T, *string] {
	return newCol031(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col031
	})
}

func (f forWideRecordsDto[T]) Col032() *ast.StringColumnProjection[T, *string] {
	return newCol032(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col032
	})
}

func (f forWideRecordsDto[T]) Col033() *ast.StringColumnProjection[T, *string] {
	return newCol033(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col033
	})
}

func (f forWideRecordsDto[T]) Col034() *ast.StringColumnProjection[T, *string] {
	return newCol034(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col034
	})
}

func (f forWideRecordsDto[T]) Col035() *ast.StringColumnProjection[T, *string] {
	return newCol035(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col035
	})
}

func (f forWideRecordsDto[T]) Col036() *ast.StringColumnProjection[T, *string] {
	return newCol036(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col036
	})
}

func (f forWideRecordsDto[T]) Col037() *ast.StringColumnProjection[T, *string] {
	return newCol037(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col037
	})
}

func (f forWideRecordsDto[T]) Col038() *ast.StringColumnProjection[T, *string] {
	return newCol038(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col038
	})
}

func (f forWideRecordsDto[T]) Col039() *ast.StringColumnProjection[T, *string] {
	return newCol039(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col039
	})
}

func (f forWideRecordsDto[T]) Col040() *ast.StringColumnProjection[T, *string] {
	return newCol040(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col040
	})
}

func (f forWideRecordsDto[T]) Col041() *ast.StringColumnProjection[T, *string] {
	return newCol041(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col041
	})
}

func (f forWideRecordsDto[T]) Col042() *ast.StringColumnProjection[T, *string] {
	return newCol042(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col042
	})
}

func (f forWideRecordsDto[T]) Col043() *ast.StringColumnProjection[T, *string] {
	return newCol043(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col043
	})
}

func (f forWideRecordsDto[T]) Col044() *ast.StringColumnProjection[T, *string] {
	return newCol044(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col044
	})
}

func (f forWideRecordsDto[T]) Col045() *ast.StringColumnProjection[T, *string] {
	return newCol045(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col045
	})
}

func (f forWideRecordsDto[T]) Col046() *ast.StringColumnProjection[T, *string] {
	return newCol046(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col046
	})
}

func (f forWideRecordsDto[T]) Col047() *ast.StringColumnProjection[T, *string] {
	return newCol047(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col047
	})
}

func (f forWideRecordsDto[T]) Col048() *ast.StringColumnProjection[T, *string] {
	return newCol048(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col048
	})
}

func (f forWideRecordsDto[T]) Col049() *ast.StringColumnProjection[T, *string] {
	return newCol049(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col049
	})
}

func (f forWideRecordsDto[T]) Col050() *ast.StringColumnProjection[T, *string] {
	return newCol050(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col050
	})
}

func (f forWideRecordsDto[T]) Col051() *ast.StringColumnProjection[T, *string] {
	return newCol051(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col051
	})
}

func (f forWideRecordsDto[T]) Col052() *ast.StringColumnProjection[T, *string] {
	return newCol052(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col052
	})
}

func (f forWideRecordsDto[T]) Col053() *ast.StringColumnProjection[T, *string] {
	return newCol053(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col053
	})
}

func (f forWideRecordsDto[T]) Col054() *ast.StringColumnProjection[T, *string] {
	return newCol054(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col054
	})
}

func (f forWideRecordsDto[T]) Col055() *ast.StringColumnProjection[T, *string] {
	return newCol055(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col055
	})
}

func (f forWideRecordsDto[T]) Col056() *ast.StringColumnProjection[T, *string] {
	return newCol056(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col056
	})
}

func (f forWideRecordsDto[T]) Col057() *ast.StringColumnProjection[T, *string] {
	return newCol057(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col057
	})
}

func (f forWideRecordsDto[T]) Col058() *ast.StringColumnProjection[T, *string] {
	return newCol058(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col058
	})
}

func (f forWideRecordsDto[T]) Col059() *ast.StringColumnProjection[T, *string] {
	return newCol059(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col059
	})
}

func (f forWideRecordsDto[T]) Col060() *ast.StringColumnProjection[T, *string] {
	return newCol060(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col060
	})
}

func (f forWideRecordsDto[T]) Col061() *ast.StringColumnProjection[T, *string] {
	return newCol061(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col061
	})
}

func (f forWideRecordsDto[T]) Col062() *ast.StringColumnProjection[T, *string] {
	return newCol062(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col062
	})
}

func (f forWideRecordsDto[T]) Col063() *ast.StringColumnProjection[T, *string] {
	return newCol063(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col063
	})
}

func (f forWideRecordsDto[T]) Col064() *ast.StringColumnProjection[T, *string] {
	return newCol064(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col064
	})
}

func (f forWideRecordsDto[T]) Col065() *ast.StringColumnProjection[T, *string] {
	return newCol065(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col065
	})
}

func (f forWideRecordsDto[T]) Col066() *ast.StringColumnProjection[T, *string] {
	return newCol066(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col066
	})
}

func (f forWideRecordsDto[T]) Col067() *ast.StringColumnProjection[T, *string] {
	return newCol067(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col067
	})
}

func (f forWideRecordsDto[T]) Col068() *ast.StringColumnProjection[T, *string] {
	return newCol068(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col068
	})
}

func (f forWideRecordsDto[T]) Col069() *ast.StringColumnProjection[T, *string] {
	return newCol069(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col069
	})
}

func (f forWideRecordsDto[T]) Col070() *ast.StringColumnProjection[T, *string] {
	return newCol070(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col070
	})
}

func (f forWideRecordsDto[T]) Col071() *ast.StringColumnProjection[T, *string] {
	return newCol071(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col071
	})
}

func (f forWideRecordsDto[T]) Col072() *ast.StringColumnProjection[T, *string] {
	return newCol072(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col072
	})
}

func (f forWideRecordsDto[T]) Col073() *ast.StringColumnProjection[T, *string] {
	return newCol073(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col073
	})
}

func (f forWideRecordsDto[T]) Col074() *ast.StringColumnProjection[T, *string] {
	return newCol074(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col074
	})
}

func (f forWideRecordsDto[T]) Col075() *ast.StringColumnProjection[T, *string] {
	return newCol075(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col075
	})
}

func (f forWideRecordsDto[T]) Col076() *ast.StringColumnProjection[T, *string] {
	return newCol076(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col076
	})
}

func (f forWideRecordsDto[T]) Col077() *ast.StringColumnProjection[T, *string] {
	return newCol077(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col077
	})
}

func (f forWideRecordsDto[T]) Col078() *ast.StringColumnProjection[T, *string] {
	return newCol078(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col078
	})
}

func (f forWideRecordsDto[T]) Col079() *ast.StringColumnProjection[T, *string] {
	return newCol079(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col079
	})
}

func (f forWideRecordsDto[T]) Col080() *ast.StringColumnProjection[T, *string] {
	return newCol080(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col080
	})
}

func (f forWideRecordsDto[T]) Col081() *ast.StringColumnProjection[T, *string] {
	return newCol081(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col081
	})
}

func (f forWideRecordsDto[T]) Col082() *ast.StringColumnProjection[T, *string] {
	return newCol082(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col082
	})
}

func (f forWideRecordsDto[T]) Col083() *ast.StringColumnProjection[T, *string] {
	return newCol083(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col083
	})
}

func (f forWideRecordsDto[T]) Col084() *ast.StringColumnProjection[T, *string] {
	return newCol084(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col084
	})
}

func (f forWideRecordsDto[T]) Col085() *ast.StringColumnProjection[T, *string] {
	return newCol085(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col085
	})
}

func (f forWideRecordsDto[T]) Col086() *ast.StringColumnProjection[T, *string] {
	return newCol086(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col086
	})
}

func (f forWideRecordsDto[T]) Col087() *ast.StringColumnProjection[T, *string] {
	return newCol087(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col087
	})
}

func (f forWideRecordsDto[T]) Col088() *ast.StringColumnProjection[T, *string] {
	return newCol088(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col088
	})
}

func (f forWideRecordsDto[T]) Col089() *ast.StringColumnProjection[T, *string] {
	return newCol089(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col089
	})
}

func (f forWideRecordsDto[T]) Col090() *ast.StringColumnProjection[T, *string] {
	return newCol090(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col090
	})
}

func (f forWideRecordsDto[T]) Col091() *ast.StringColumnProjection[T, *string] {
	return newCol091(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col091
	})
}

func (f forWideRecordsDto[T]) Col092() *ast.StringColumnProjection[T, *string] {
	return newCol092(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col092
	})
}

func (f forWideRecordsDto[T]) Col093() *ast.StringColumnProjection[T, *string] {
	return newCol093(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col093
	})
}

func (f forWideRecordsDto[T]) Col094() *ast.StringColumnProjection[T, *string] {
	return newCol094(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col094
	})
}

func (f forWideRecordsDto[T]) Col095() *ast.StringColumnProjection[T, *string] {
	return newCol095(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col095
	})
}

func (f forWideRecordsDto[T]) Col096() *ast.StringColumnProjection[T, *string] {
	return newCol096(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col096
	})
}

func (f forWideRecordsDto[T]) Col097() *ast.StringColumnProjection[T, *string] {
	return newCol097(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col097
	})
}

func (f forWideRecordsDto[T]) Col098() *ast.StringColumnProjection[T, *string] {
	return newCol098(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col098
	})
}

func (f forWideRecordsDto[T]) Col099() *ast.StringColumnProjection[T, *string] {
	return newCol099(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col099
	})
}

func (f forWideRecordsDto[T]) Col100() *ast.StringColumnProjection[T, *string] {
	return newCol100(`"public"."wide_records"`, func(t *T) *string {
		dto := f.dest(t)
		return &dto.Col100
	})
}

func (f forWideRecordsDto[T]) AllColumns() []ast.Projection[T] {
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

type Alias[T any] struct {
	*ast.TableAlias
	forWideRecordsDto[T]
}

func As(name string, columns ...ast.NamedExpression) Alias[models.WideRecordsDto] {
	return Alias[models.WideRecordsDto]{
		TableAlias: Table().As(name, columns...),
		forWideRecordsDto: forWideRecordsDto[models.WideRecordsDto]{
			dest: func(w *models.WideRecordsDto) *models.WideRecordsDto {
				return w
			},
		},
	}
}

func (a Alias[T]) Id() *ast.UUIDColumnProjection[T, **uuid.UUID] {
	projection := newId(a.TableAlias.GetName(), func(t *T) **uuid.UUID {
		dto := a.dest(t)
		return &dto.Id
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'id' in alias %s", a.TableAlias.GetName()))
	projection.UUIDColumnExpression = ast.NewUUIDColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col001() *ast.StringColumnProjection[T, *string] {
	projection := newCol001(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col001
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_001' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col002() *ast.StringColumnProjection[T, *string] {
	projection := newCol002(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col002
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_002' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col003() *ast.StringColumnProjection[T, *string] {
	projection := newCol003(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col003
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_003' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col004() *ast.StringColumnProjection[T, *string] {
	projection := newCol004(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col004
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_004' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col005() *ast.StringColumnProjection[T, *string] {
	projection := newCol005(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col005
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_005' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col006() *ast.StringColumnProjection[T, *string] {
	projection := newCol006(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col006
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_006' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col007() *ast.StringColumnProjection[T, *string] {
	projection := newCol007(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col007
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_007' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col008() *ast.StringColumnProjection[T, *string] {
	projection := newCol008(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col008
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_008' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col009() *ast.StringColumnProjection[T, *string] {
	projection := newCol009(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col009
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_009' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col010() *ast.StringColumnProjection[T, *string] {
	projection := newCol010(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col010
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_010' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col011() *ast.StringColumnProjection[T, *string] {
	projection := newCol011(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col011
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_011' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col012() *ast.StringColumnProjection[T, *string] {
	projection := newCol012(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col012
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_012' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col013() *ast.StringColumnProjection[T, *string] {
	projection := newCol013(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col013
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_013' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col014() *ast.StringColumnProjection[T, *string] {
	projection := newCol014(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col014
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_014' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col015() *ast.StringColumnProjection[T, *string] {
	projection := newCol015(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col015
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_015' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col016() *ast.StringColumnProjection[T, *string] {
	projection := newCol016(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col016
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_016' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col017() *ast.StringColumnProjection[T, *string] {
	projection := newCol017(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col017
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_017' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col018() *ast.StringColumnProjection[T, *string] {
	projection := newCol018(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col018
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_018' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col019() *ast.StringColumnProjection[T, *string] {
	projection := newCol019(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col019
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_019' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col020() *ast.StringColumnProjection[T, *string] {
	projection := newCol020(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col020
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_020' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col021() *ast.StringColumnProjection[T, *string] {
	projection := newCol021(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col021
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_021' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col022() *ast.StringColumnProjection[T, *string] {
	projection := newCol022(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col022
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_022' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col023() *ast.StringColumnProjection[T, *string] {
	projection := newCol023(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col023
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_023' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col024() *ast.StringColumnProjection[T, *string] {
	projection := newCol024(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col024
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_024' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col025() *ast.StringColumnProjection[T, *string] {
	projection := newCol025(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col025
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_025' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col026() *ast.StringColumnProjection[T, *string] {
	projection := newCol026(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col026
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_026' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col027() *ast.StringColumnProjection[T, *string] {
	projection := newCol027(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col027
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_027' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col028() *ast.StringColumnProjection[T, *string] {
	projection := newCol028(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col028
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_028' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col029() *ast.StringColumnProjection[T, *string] {
	projection := newCol029(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col029
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_029' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col030() *ast.StringColumnProjection[T, *string] {
	projection := newCol030(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col030
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_030' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col031() *ast.StringColumnProjection[T, *string] {
	projection := newCol031(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col031
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_031' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col032() *ast.StringColumnProjection[T, *string] {
	projection := newCol032(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col032
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_032' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col033() *ast.StringColumnProjection[T, *string] {
	projection := newCol033(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col033
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_033' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col034() *ast.StringColumnProjection[T, *string] {
	projection := newCol034(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col034
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_034' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col035() *ast.StringColumnProjection[T, *string] {
	projection := newCol035(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col035
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_035' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col036() *ast.StringColumnProjection[T, *string] {
	projection := newCol036(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col036
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_036' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col037() *ast.StringColumnProjection[T, *string] {
	projection := newCol037(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col037
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_037' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col038() *ast.StringColumnProjection[T, *string] {
	projection := newCol038(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col038
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_038' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col039() *ast.StringColumnProjection[T, *string] {
	projection := newCol039(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col039
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_039' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col040() *ast.StringColumnProjection[T, *string] {
	projection := newCol040(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col040
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_040' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col041() *ast.StringColumnProjection[T, *string] {
	projection := newCol041(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col041
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_041' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col042() *ast.StringColumnProjection[T, *string] {
	projection := newCol042(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col042
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_042' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col043() *ast.StringColumnProjection[T, *string] {
	projection := newCol043(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col043
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_043' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col044() *ast.StringColumnProjection[T, *string] {
	projection := newCol044(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col044
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_044' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col045() *ast.StringColumnProjection[T, *string] {
	projection := newCol045(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col045
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_045' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col046() *ast.StringColumnProjection[T, *string] {
	projection := newCol046(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col046
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_046' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col047() *ast.StringColumnProjection[T, *string] {
	projection := newCol047(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col047
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_047' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col048() *ast.StringColumnProjection[T, *string] {
	projection := newCol048(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col048
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_048' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col049() *ast.StringColumnProjection[T, *string] {
	projection := newCol049(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col049
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_049' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col050() *ast.StringColumnProjection[T, *string] {
	projection := newCol050(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col050
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_050' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col051() *ast.StringColumnProjection[T, *string] {
	projection := newCol051(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col051
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_051' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col052() *ast.StringColumnProjection[T, *string] {
	projection := newCol052(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col052
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_052' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col053() *ast.StringColumnProjection[T, *string] {
	projection := newCol053(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col053
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_053' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col054() *ast.StringColumnProjection[T, *string] {
	projection := newCol054(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col054
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_054' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col055() *ast.StringColumnProjection[T, *string] {
	projection := newCol055(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col055
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_055' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col056() *ast.StringColumnProjection[T, *string] {
	projection := newCol056(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col056
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_056' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col057() *ast.StringColumnProjection[T, *string] {
	projection := newCol057(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col057
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_057' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col058() *ast.StringColumnProjection[T, *string] {
	projection := newCol058(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col058
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_058' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col059() *ast.StringColumnProjection[T, *string] {
	projection := newCol059(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col059
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_059' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col060() *ast.StringColumnProjection[T, *string] {
	projection := newCol060(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col060
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_060' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col061() *ast.StringColumnProjection[T, *string] {
	projection := newCol061(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col061
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_061' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col062() *ast.StringColumnProjection[T, *string] {
	projection := newCol062(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col062
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_062' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col063() *ast.StringColumnProjection[T, *string] {
	projection := newCol063(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col063
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_063' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col064() *ast.StringColumnProjection[T, *string] {
	projection := newCol064(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col064
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_064' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col065() *ast.StringColumnProjection[T, *string] {
	projection := newCol065(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col065
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_065' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col066() *ast.StringColumnProjection[T, *string] {
	projection := newCol066(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col066
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_066' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col067() *ast.StringColumnProjection[T, *string] {
	projection := newCol067(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col067
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_067' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col068() *ast.StringColumnProjection[T, *string] {
	projection := newCol068(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col068
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_068' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col069() *ast.StringColumnProjection[T, *string] {
	projection := newCol069(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col069
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_069' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col070() *ast.StringColumnProjection[T, *string] {
	projection := newCol070(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col070
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_070' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col071() *ast.StringColumnProjection[T, *string] {
	projection := newCol071(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col071
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_071' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col072() *ast.StringColumnProjection[T, *string] {
	projection := newCol072(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col072
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_072' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col073() *ast.StringColumnProjection[T, *string] {
	projection := newCol073(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col073
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_073' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col074() *ast.StringColumnProjection[T, *string] {
	projection := newCol074(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col074
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_074' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col075() *ast.StringColumnProjection[T, *string] {
	projection := newCol075(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col075
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_075' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col076() *ast.StringColumnProjection[T, *string] {
	projection := newCol076(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col076
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_076' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col077() *ast.StringColumnProjection[T, *string] {
	projection := newCol077(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col077
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_077' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col078() *ast.StringColumnProjection[T, *string] {
	projection := newCol078(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col078
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_078' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col079() *ast.StringColumnProjection[T, *string] {
	projection := newCol079(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col079
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_079' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col080() *ast.StringColumnProjection[T, *string] {
	projection := newCol080(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col080
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_080' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col081() *ast.StringColumnProjection[T, *string] {
	projection := newCol081(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col081
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_081' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col082() *ast.StringColumnProjection[T, *string] {
	projection := newCol082(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col082
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_082' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col083() *ast.StringColumnProjection[T, *string] {
	projection := newCol083(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col083
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_083' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col084() *ast.StringColumnProjection[T, *string] {
	projection := newCol084(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col084
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_084' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col085() *ast.StringColumnProjection[T, *string] {
	projection := newCol085(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col085
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_085' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col086() *ast.StringColumnProjection[T, *string] {
	projection := newCol086(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col086
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_086' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col087() *ast.StringColumnProjection[T, *string] {
	projection := newCol087(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col087
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_087' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col088() *ast.StringColumnProjection[T, *string] {
	projection := newCol088(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col088
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_088' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col089() *ast.StringColumnProjection[T, *string] {
	projection := newCol089(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col089
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_089' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col090() *ast.StringColumnProjection[T, *string] {
	projection := newCol090(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col090
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_090' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col091() *ast.StringColumnProjection[T, *string] {
	projection := newCol091(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col091
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_091' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col092() *ast.StringColumnProjection[T, *string] {
	projection := newCol092(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col092
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_092' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col093() *ast.StringColumnProjection[T, *string] {
	projection := newCol093(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col093
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_093' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col094() *ast.StringColumnProjection[T, *string] {
	projection := newCol094(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col094
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_094' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col095() *ast.StringColumnProjection[T, *string] {
	projection := newCol095(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col095
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_095' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col096() *ast.StringColumnProjection[T, *string] {
	projection := newCol096(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col096
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_096' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col097() *ast.StringColumnProjection[T, *string] {
	projection := newCol097(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col097
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_097' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col098() *ast.StringColumnProjection[T, *string] {
	projection := newCol098(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col098
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_098' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col099() *ast.StringColumnProjection[T, *string] {
	projection := newCol099(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col099
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_099' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) Col100() *ast.StringColumnProjection[T, *string] {
	projection := newCol100(a.TableAlias.GetName(), func(t *T) *string {
		dto := a.dest(t)
		return &dto.Col100
	})
	if len(a.Columns()) == 0 || slices.ContainsFunc(a.Columns(), func(e ast.NamedExpression) bool {
		return e.GetName() == projection.GetName()
	}) {
		return projection
	}
	errExpr := ast.NewErrorExpression(fmt.Errorf("unknown column 'col_100' in alias %s", a.TableAlias.GetName()))
	projection.StringColumnExpression = ast.NewStringColumnExpressionFromExpr(projection.GetName(), errExpr)
	return projection
}

func (a Alias[T]) AllColumns() []ast.Projection[T] {
	return []ast.Projection[T]{
		a.Id(),
		a.Col001(),
		a.Col002(),
		a.Col003(),
		a.Col004(),
		a.Col005(),
		a.Col006(),
		a.Col007(),
		a.Col008(),
		a.Col009(),
		a.Col010(),
		a.Col011(),
		a.Col012(),
		a.Col013(),
		a.Col014(),
		a.Col015(),
		a.Col016(),
		a.Col017(),
		a.Col018(),
		a.Col019(),
		a.Col020(),
		a.Col021(),
		a.Col022(),
		a.Col023(),
		a.Col024(),
		a.Col025(),
		a.Col026(),
		a.Col027(),
		a.Col028(),
		a.Col029(),
		a.Col030(),
		a.Col031(),
		a.Col032(),
		a.Col033(),
		a.Col034(),
		a.Col035(),
		a.Col036(),
		a.Col037(),
		a.Col038(),
		a.Col039(),
		a.Col040(),
		a.Col041(),
		a.Col042(),
		a.Col043(),
		a.Col044(),
		a.Col045(),
		a.Col046(),
		a.Col047(),
		a.Col048(),
		a.Col049(),
		a.Col050(),
		a.Col051(),
		a.Col052(),
		a.Col053(),
		a.Col054(),
		a.Col055(),
		a.Col056(),
		a.Col057(),
		a.Col058(),
		a.Col059(),
		a.Col060(),
		a.Col061(),
		a.Col062(),
		a.Col063(),
		a.Col064(),
		a.Col065(),
		a.Col066(),
		a.Col067(),
		a.Col068(),
		a.Col069(),
		a.Col070(),
		a.Col071(),
		a.Col072(),
		a.Col073(),
		a.Col074(),
		a.Col075(),
		a.Col076(),
		a.Col077(),
		a.Col078(),
		a.Col079(),
		a.Col080(),
		a.Col081(),
		a.Col082(),
		a.Col083(),
		a.Col084(),
		a.Col085(),
		a.Col086(),
		a.Col087(),
		a.Col088(),
		a.Col089(),
		a.Col090(),
		a.Col091(),
		a.Col092(),
		a.Col093(),
		a.Col094(),
		a.Col095(),
		a.Col096(),
		a.Col097(),
		a.Col098(),
		a.Col099(),
		a.Col100(),
	}
}

func AliasFor[T any](alias *ast.TableAlias, dest func(*T) *models.WideRecordsDto) Alias[T] {
	return Alias[T]{
		TableAlias:        alias,
		forWideRecordsDto: forWideRecordsDto[T]{dest: dest},
	}
}

func newId[T any](table string, ref func(*T) **uuid.UUID) *ast.UUIDColumnProjection[T, **uuid.UUID] {
	return ast.NewUUIDColumnProjection(
		table,
		"id",
		ref,
	)
}

func newCol001[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_001",
		ref,
	)
}

func newCol002[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_002",
		ref,
	)
}

func newCol003[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_003",
		ref,
	)
}

func newCol004[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_004",
		ref,
	)
}

func newCol005[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_005",
		ref,
	)
}

func newCol006[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_006",
		ref,
	)
}

func newCol007[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_007",
		ref,
	)
}

func newCol008[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_008",
		ref,
	)
}

func newCol009[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_009",
		ref,
	)
}

func newCol010[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_010",
		ref,
	)
}

func newCol011[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_011",
		ref,
	)
}

func newCol012[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_012",
		ref,
	)
}

func newCol013[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_013",
		ref,
	)
}

func newCol014[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_014",
		ref,
	)
}

func newCol015[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_015",
		ref,
	)
}

func newCol016[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_016",
		ref,
	)
}

func newCol017[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_017",
		ref,
	)
}

func newCol018[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_018",
		ref,
	)
}

func newCol019[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_019",
		ref,
	)
}

func newCol020[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_020",
		ref,
	)
}

func newCol021[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_021",
		ref,
	)
}

func newCol022[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_022",
		ref,
	)
}

func newCol023[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_023",
		ref,
	)
}

func newCol024[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_024",
		ref,
	)
}

func newCol025[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_025",
		ref,
	)
}

func newCol026[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_026",
		ref,
	)
}

func newCol027[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_027",
		ref,
	)
}

func newCol028[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_028",
		ref,
	)
}

func newCol029[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_029",
		ref,
	)
}

func newCol030[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_030",
		ref,
	)
}

func newCol031[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_031",
		ref,
	)
}

func newCol032[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_032",
		ref,
	)
}

func newCol033[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_033",
		ref,
	)
}

func newCol034[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_034",
		ref,
	)
}

func newCol035[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_035",
		ref,
	)
}

func newCol036[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_036",
		ref,
	)
}

func newCol037[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_037",
		ref,
	)
}

func newCol038[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_038",
		ref,
	)
}

func newCol039[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_039",
		ref,
	)
}

func newCol040[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_040",
		ref,
	)
}

func newCol041[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_041",
		ref,
	)
}

func newCol042[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_042",
		ref,
	)
}

func newCol043[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_043",
		ref,
	)
}

func newCol044[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_044",
		ref,
	)
}

func newCol045[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_045",
		ref,
	)
}

func newCol046[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_046",
		ref,
	)
}

func newCol047[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_047",
		ref,
	)
}

func newCol048[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_048",
		ref,
	)
}

func newCol049[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_049",
		ref,
	)
}

func newCol050[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_050",
		ref,
	)
}

func newCol051[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_051",
		ref,
	)
}

func newCol052[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_052",
		ref,
	)
}

func newCol053[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_053",
		ref,
	)
}

func newCol054[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_054",
		ref,
	)
}

func newCol055[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_055",
		ref,
	)
}

func newCol056[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_056",
		ref,
	)
}

func newCol057[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_057",
		ref,
	)
}

func newCol058[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_058",
		ref,
	)
}

func newCol059[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_059",
		ref,
	)
}

func newCol060[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_060",
		ref,
	)
}

func newCol061[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_061",
		ref,
	)
}

func newCol062[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_062",
		ref,
	)
}

func newCol063[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_063",
		ref,
	)
}

func newCol064[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_064",
		ref,
	)
}

func newCol065[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_065",
		ref,
	)
}

func newCol066[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_066",
		ref,
	)
}

func newCol067[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_067",
		ref,
	)
}

func newCol068[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_068",
		ref,
	)
}

func newCol069[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_069",
		ref,
	)
}

func newCol070[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_070",
		ref,
	)
}

func newCol071[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_071",
		ref,
	)
}

func newCol072[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_072",
		ref,
	)
}

func newCol073[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_073",
		ref,
	)
}

func newCol074[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_074",
		ref,
	)
}

func newCol075[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_075",
		ref,
	)
}

func newCol076[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_076",
		ref,
	)
}

func newCol077[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_077",
		ref,
	)
}

func newCol078[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_078",
		ref,
	)
}

func newCol079[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_079",
		ref,
	)
}

func newCol080[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_080",
		ref,
	)
}

func newCol081[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_081",
		ref,
	)
}

func newCol082[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_082",
		ref,
	)
}

func newCol083[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_083",
		ref,
	)
}

func newCol084[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_084",
		ref,
	)
}

func newCol085[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_085",
		ref,
	)
}

func newCol086[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_086",
		ref,
	)
}

func newCol087[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_087",
		ref,
	)
}

func newCol088[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_088",
		ref,
	)
}

func newCol089[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_089",
		ref,
	)
}

func newCol090[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_090",
		ref,
	)
}

func newCol091[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_091",
		ref,
	)
}

func newCol092[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_092",
		ref,
	)
}

func newCol093[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_093",
		ref,
	)
}

func newCol094[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_094",
		ref,
	)
}

func newCol095[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_095",
		ref,
	)
}

func newCol096[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_096",
		ref,
	)
}

func newCol097[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_097",
		ref,
	)
}

func newCol098[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_098",
		ref,
	)
}

func newCol099[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_099",
		ref,
	)
}

func newCol100[T any](table string, ref func(*T) *string) *ast.StringColumnProjection[T, *string] {
	return ast.NewStringColumnProjection(
		table,
		"col_100",
		ref,
	)
}
