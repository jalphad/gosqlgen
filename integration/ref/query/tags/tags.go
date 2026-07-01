package tags

import (
	"github.com/jalphad/gosqlgen/integration/models.new"
	"github.com/jalphad/gosqlgen/integration/ref/ast"
)

var Into = intoTagsDto{}

type intoTagsDto struct{}

func Table() *ast.TableSource {
	return ast.NewTableSource("tags")
}

func Id() *ast.IntColumnExpression {
	return ast.NewIntColumnExpression("tags", "id")
}

func Name() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("tags", "name")
}

func Slug() *ast.StringColumnExpression {
	return ast.NewStringColumnExpression("tags", "slug")
}

func AllColumns() []ast.NamedExpression {
	return []ast.NamedExpression{
		Id(),
		Name(),
		Slug(),
	}
}

func (intoTagsDto) Id() ast.Projection[models.TagsDto] {
	return ast.NewProjection(Id(), func(t *models.TagsDto) **int64 {
		return &t.Id
	})
}

func (intoTagsDto) Name() ast.Projection[models.TagsDto] {
	return ast.NewProjection(Name(), func(t *models.TagsDto) *string {
		return &t.Name
	})
}

func (intoTagsDto) Slug() ast.Projection[models.TagsDto] {
	return ast.NewProjection(Slug(), func(t *models.TagsDto) *string {
		return &t.Slug
	})
}

func (intoTagsDto) AllColumns() []ast.Projection[models.TagsDto] {
	return []ast.Projection[models.TagsDto]{
		Into.Id(),
		Into.Name(),
		Into.Slug(),
	}
}
