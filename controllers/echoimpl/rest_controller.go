package echoimpl

import (
	"embed"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"os"
	"path"
)

//go:embed templates/*
var templatesFS embed.FS

func GenerateRestController(input parser.GenerateModelInput, goDeps deps.Container, inCreation deps.Dependency) error {
	log.Info("Generating rest controller...")
	var tpl = initTemplates(input)

	f, err := os.OpenFile(
		path.Join(
			string(domain.CmdServiceControllersPackagePath),
			fmt.Sprintf("%s_rest.go", inflection.Plural(input.SnakeCase)),
		),
		os.O_RDWR|os.O_CREATE,
		0755,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tpl.ExecuteTemplate(f, "rest_controller_tmpl", tplInput{
		GenerateModelInput: input,
		GoDeps:             goDeps,
		InCreation:         inCreation,
	}); err != nil {
		return err
	}

	return nil
}
