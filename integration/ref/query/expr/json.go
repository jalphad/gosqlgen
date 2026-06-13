package expr

import (
	"encoding/json"

	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

func JsonAggObject(alias string, filter ast.OfType[bool], columns ...ast.NamedExpression) ast.AliasedFunction[json.RawMessage] {
	aggregation := JsonAgg(Distinct(JsonbBuildObject(columns...)))

	var aggregated ast.OfType[json.RawMessage] = aggregation
	if filter != nil {
		aggregated = ast.SetType[json.RawMessage](aggregation.Filter(filter))
	}

	emptyJSON := Cast(ast.SetType[string](InlineStringLiteral("[]"))).AsJson()
	return Coalesce(aggregated, emptyJSON).As(ast.NewAlias(alias))
}
