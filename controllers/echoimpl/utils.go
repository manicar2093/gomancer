package echoimpl

import (
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/controllers/echoimpl/components"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"text/template"
)

type (
	tplInput struct {
		parser.GenerateModelInput
		GoDeps         deps.Container
		InCreation     deps.Dependency
		ComponentsDeps []deps.Dependency
	}
)

func initTemplates(input parser.GenerateModelInput) *template.Template {
	return template.Must(template.
		New("controllers").
		Funcs(map[string]any{
			"GetByType": func() string {
				if input.IdAttribute.Type == string(types.TypeUuid) {
					return "GetByIdUUID"
				}

				return "GetById"
			},
			"Pluralize": inflection.Plural,
			"GenerateFormItemInput": func(a parser.Attribute) string {
				comp, err := components.GetInputByType(a, input.TransformedText)
				if err != nil {
					panic(err)
				}
				return comp
			},
		}).
		ParseFS(templatesFS, "templates/*"))
}
