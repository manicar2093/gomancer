package echoimpl

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/coditory/go-errors"
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"os"
	"path"
)

func GenerateWebController(input parser.GenerateModelInput, goDeps deps.Container, inCreation deps.Dependency) error {
	log.Info("Generating echo web controller...")
	var tpl = initTemplates(input)

	f, err := os.OpenFile(
		path.Join(
			string(domain.CmdServiceControllersPackagePath),
			fmt.Sprintf("%s_web.go", inflection.Plural(input.SnakeCase)),
		),
		os.O_RDWR|os.O_CREATE,
		0755,
	)
	if err != nil {
		return errors.Wrap(err, "Error opening file")
	}
	defer f.Close()

	if err := tpl.ExecuteTemplate(f, "web_controller_tmpl", tplInput{
		GenerateModelInput: input,
		GoDeps:             goDeps,
		InCreation:         inCreation,
	}); err != nil {
		return err
	}

	return nil
}
