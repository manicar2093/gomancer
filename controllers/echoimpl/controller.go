package echoimpl

import (
	"embed"
	"github.com/charmbracelet/log"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"os"
	"path"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

type (
	tplInput struct {
		parser.GenerateModelInput
		GoDeps     deps.Container
		InCreation deps.Dependency
	}
)

func GenerateController(input parser.GenerateModelInput, goDeps deps.Container, inCreation deps.Dependency) error {
	log.Info("Generating echo controller...")
	var tpl = template.Must(template.
		New("controllers").
		Funcs(map[string]any{
			"GetByType": func() string {
				if input.IdAttribute.Type == string(types.TypeUuid) {
					return "GetByIdUUID"
				}

				return "GetById"
			},
		}).
		ParseFS(templatesFS, "templates/*"))

	f, err := os.OpenFile(path.Join(string(domain.CmdApiControllersPackagePath), input.SnakeCase+".go"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tpl.ExecuteTemplate(f, "controller_tmpl", tplInput{
		GenerateModelInput: input,
		GoDeps:             goDeps,
		InCreation:         inCreation,
	}); err != nil {
		return err
	}

	return nil
}
