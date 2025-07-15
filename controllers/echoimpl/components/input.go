package components

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"path"
	"text/template"
)

type InputGenerationData struct {
	parser.Attribute
	ModelTransformedText parser.TransformedText
}

//go:embed templates/*
var templatesFS embed.FS

func InputNumber(input InputGenerationData) (string, error) {
	result, err := executeTemplate("input_number.tmpl", input)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputNumberFloat(input InputGenerationData) (string, error) {
	result, err := executeTemplate("input_number_float.tmpl", input)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputText(input InputGenerationData) (string, error) {
	result, err := executeTemplate("input_text.tmpl", input)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputToggle(input InputGenerationData) (string, error) {
	result, err := executeTemplate("input_toggle.tmpl", input)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputDateTime(input InputGenerationData) (string, error) {
	result, err := executeTemplate("input_datetime.tmpl", input)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputSelectBox(input InputGenerationData) (string, error) {
	result, err := executeTemplate("input_selectbox.tmpl", input)
	if err != nil {
		return "", err
	}
	return result, nil
}

func executeTemplate(templateName string, data InputGenerationData) (string, error) {
	tplStarted := template.Must(
		template.New(templateName).
			Funcs(map[string]any{
				"GenerateFormValue": func(input InputGenerationData) string {
					gen := ""

					if input.IsOptional {
						gen = fmt.Sprintf("%s.%s.GetValue()", input.ModelTransformedText.CamelCase, input.PascalCase)
					} else {
						gen = fmt.Sprintf("%s.%s", input.ModelTransformedText.CamelCase, input.PascalCase)
					}

					switch types.SupportedType(input.Type) {
					case types.TypeInt:
						gen = fmt.Sprintf("strconv.Itoa(%s)", gen)
					case types.TypeInt8, types.TypeInt16, types.TypeInt32, types.TypeInt64:
						gen = fmt.Sprintf("strconv.Itoa(int(%s))", gen)
					case types.TypeFloat32:
						gen = fmt.Sprintf("strconv.FormatFloat(float64(%s), 'f', 2, 64)", gen)
					case types.TypeFloat64:
						gen = fmt.Sprintf("strconv.FormatFloat(%s, 'f', 2, 64)", gen)
					}
					return gen
				},
			}).
			ParseFS(templatesFS, path.Join("templates", templateName)),
	)
	var buf bytes.Buffer
	if err := tplStarted.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error executing template %s: %w", templateName, err)
	}

	return buf.String(), nil
}
