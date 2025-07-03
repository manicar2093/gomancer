package components

import (
	"bytes"
	"fmt"
	"github.com/manicar2093/gomancer/parser"
	"path/filepath"
	"text/template"
)

// executeTemplate loads and executes a template with the given data
func executeTemplate(templateName string, data interface{}) (string, error) {
	// Try different paths to find the template
	possiblePaths := []string{
		filepath.Join("templates", templateName),                                          // When running from components dir
		filepath.Join("controllers", "echoimpl", "components", "templates", templateName), // When running from project root
	}

	var tmpl *template.Template
	var err error
	var foundPath string

	// Try each path until we find the template
	for _, path := range possiblePaths {
		tmpl, err = template.ParseFiles(path)
		if err == nil {
			foundPath = path
			break
		}
	}

	if err != nil {
		return "", fmt.Errorf("error parsing template %s: could not find template in any of the expected locations", templateName)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error executing template %s (found at %s): %w", templateName, foundPath, err)
	}

	return buf.String(), nil
}

func InputNumber(attribute parser.Attribute) (string, error) {
	result, err := executeTemplate("input_number.tmpl", attribute)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputNumberFloat(attribute parser.Attribute) (string, error) {
	result, err := executeTemplate("input_number_float.tmpl", attribute)
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputText(attribute parser.Attribute) (string, error) {
	result, err := executeTemplate("input_text.tmpl", attribute)
	if err != nil {
		return "", err
	}
	return result, nil
}

type InputDateTime struct {
	parser.Attribute
}

func (c InputDateTime) Generate() (string, error) {
	result, err := executeTemplate("input_datetime.tmpl", c)
	if err != nil {
		return "", err
	}
	return result, nil
}

type InputToggle struct {
	parser.Attribute
}

func (c InputToggle) Generate() (string, error) {
	result, err := executeTemplate("input_toggle.tmpl", c)
	if err != nil {
		return "", err
	}
	return result, nil
}

type InputSelectBox struct {
	parser.Attribute
}

func (c InputSelectBox) Generate() (string, error) {
	result, err := executeTemplate("input_selectbox.tmpl", c)
	if err != nil {
		return "", err
	}
	return result, nil
}
