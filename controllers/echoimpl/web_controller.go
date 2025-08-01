package echoimpl

import (
	"bytes"
	"cmp"
	"fmt"
	"github.com/a-h/templ/cmd/templ/fmtcmd"
	"github.com/charmbracelet/log"
	"github.com/coditory/go-errors"
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"maps"
	"os"
	"path"
	"slices"
	"text/template"
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

	tplData := tplInput{
		GenerateModelInput: input,
		GoDeps:             goDeps,
		InCreation:         inCreation,
	}

	if err := tpl.ExecuteTemplate(f, "web_controller_tmpl", tplData); err != nil {
		return err
	}
	if err := generateTemplates(input, tplData, tpl, goDeps); err != nil {
		return err
	}

	return nil
}

func generateTemplates(input parser.GenerateModelInput, tplData tplInput, tpl *template.Template, goDeps deps.Container) error {
	if err := os.Mkdir(path.Join(string(domain.CmdServiceControllersPackagePath), fmt.Sprintf("%spages", input.LowerNoSpaceCase)), 0755); err != nil {
		return errors.Wrap(err, "Error creating directory")
	}

	templsToGenerate := []struct {
		TemplateName, Filename string
		AdditionalComponents   []string
		GetTemplFormComponents bool
	}{
		{TemplateName: "web_form", Filename: fmt.Sprintf("%s_form.templ", input.SnakeCase), GetTemplFormComponents: true},
		{TemplateName: "web_table", Filename: fmt.Sprintf("%s_table.templ", input.SnakeCase)},
		{TemplateName: "web_show", Filename: fmt.Sprintf("%s_show.templ", input.SnakeCase)},
		{TemplateName: "web_register", Filename: fmt.Sprintf("%s_register.templ", input.SnakeCase)},
		{TemplateName: "web_index", Filename: fmt.Sprintf("%s_index.templ", input.SnakeCase)},
		{TemplateName: "web_edit", Filename: fmt.Sprintf("%s_edit.templ", input.SnakeCase)},
	}

	for _, item := range templsToGenerate {
		filePath := path.Join(
			string(domain.CmdServiceControllersPackagePath),
			fmt.Sprintf("%spages", input.LowerNoSpaceCase),
			item.Filename,
		)
		destinyFile, err := os.OpenFile(
			filePath,
			os.O_RDWR|os.O_CREATE,
			0644,
		)
		if err != nil {
			return errors.Wrap(err, "Error opening file")
		}

		if item.GetTemplFormComponents {
			tplData.ComponentsDeps = getFormComponentsToBeUsed(input.Attributes, goDeps)
		}

		bytesContent := bytes.NewBufferString("")
		if err := tpl.ExecuteTemplate(bytesContent, item.TemplateName, tplData); err != nil {
			return err
		}

		if err := fmtcmd.Run(nil, bytesContent, destinyFile, fmtcmd.Arguments{}); err != nil {
			return errors.Wrap(err, "Error formatting file")
		}
	}

	return nil
}

func getFormComponentsToBeUsed(attributes []parser.Attribute, goDeps deps.Container) []deps.Dependency {
	var componentsDetected = map[string]deps.Dependency{}
	for _, attr := range attributes {
		switch types.SupportedType(attr.Type) {
		case types.TypeInt, types.TypeInt8, types.TypeInt16, types.TypeInt32, types.TypeInt64, types.TypeFloat32, types.TypeFloat64:
			componentsDetected["input"] = goDeps.Project.Core.Web.Components.Input
			componentsDetected["strconv"] = goDeps.Std.StrConv
		case types.TypeDecimal, types.TypeString, types.TypeUuid:
			componentsDetected["input"] = goDeps.Project.Core.Web.Components.Input
		case types.TypeBool:
			componentsDetected["toggle"] = goDeps.Project.Core.Web.Components.Toggle
		case types.TypeTime:
			componentsDetected["datetime"] = goDeps.Project.Core.Web.Components.DateTime
		case types.TypeEnum:
			componentsDetected["selectbox"] = goDeps.Project.Core.Web.Components.SelectBox
		}
		componentsDetected["label"] = goDeps.Project.Core.Web.Components.Label
		componentsDetected["form"] = goDeps.Project.Core.Web.Components.Form
		componentsDetected["formerrors"] = goDeps.Project.Core.Web.Components.FormErrors
		componentsDetected["button"] = goDeps.Project.Core.Web.Components.Button
	}

	return slices.SortedFunc(
		maps.Values(componentsDetected),
		func(a deps.Dependency, b deps.Dependency) int {
			return cmp.Compare(a.Path, b.Path)
		})
}
