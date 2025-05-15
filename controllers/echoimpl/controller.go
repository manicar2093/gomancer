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
	tplInput struct {
		domain.GenerateModelInput
		EchoDependency           string
		CoreDependency           string
		CoreCommonReqDependency  string
		InternalDomainModelsPath domain.Path
		InternalPackagePath      domain.Path
	}
)

func GenerateController(input domain.GenerateModelInput) error {
	log.Info("Generating echo controller...")
	var tpl = template.Must(template.
		New("controllers").
		Funcs(map[string]any{
			"GetByType": func() string {
				if input.IdAttribute.Type == string(domain.TypeUuid) {
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
		GenerateModelInput:       input,
		EchoDependency:           domain.EchoPkgPath,
		CoreDependency:           domain.GenerateCorePackage(input.ModuleInfo),
		CoreCommonReqDependency:  domain.GetCorePackage(input.ModuleInfo, domain.CoreCommonReqPkg),
		InternalDomainModelsPath: domain.InternalDomainModelsPackagePath,
		InternalPackagePath:      domain.InternalPackagePath,
	}); err != nil {
		return err
	}

	return nil
}
