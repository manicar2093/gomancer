package echoimpl

import (
	"embed"
	"github.com/charmbracelet/log"
	"github.com/manicar2093/gomancer/domain"
	"os"
	"path"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS

type (
	GenerateControllerInput struct {
		domain.GenerateModelInput
	}
	tplInput struct {
		domain.GenerateModelInput
		EchoDependency            string
		WinterDependency          string
		WinterCommonReqDependency string
		InternalDomainModelsPath  domain.Path
	}
)

func GenerateController(input domain.GenerateModelInput) error {
	log.Printf("Generating echo controller...")
	tpl := template.Must(template.ParseFS(templatesFS, "templates/*"))

	f, err := os.OpenFile(path.Join(string(domain.CmdApiControllersPackagePath), input.SnakeCase+".go"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := tpl.ExecuteTemplate(f, "controller_tmpl", tplInput{
		GenerateModelInput:        input,
		EchoDependency:            domain.EchoPkgPath,
		WinterDependency:          domain.WinterPkgPath,
		WinterCommonReqDependency: domain.WinterCommonReqPkgPath,
		InternalDomainModelsPath:  domain.InternalDomainModelsPackagePath,
	}); err != nil {
		return err
	}

	return nil
}
