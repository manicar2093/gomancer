package echoimpl

import (
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"text/template"
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
		}).
		ParseFS(templatesFS, "templates/*"))
}
