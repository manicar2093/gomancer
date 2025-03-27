package domain

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/iancoleman/strcase"
	"github.com/rjNemo/underscore"
	"strings"
)

type (
	ParsingErrorDetail struct {
		Input string
		Err   string
		Index int
	}
)

func (p ParsingErrorDetail) String() string {
	return fmt.Sprintf("Error: %s. Input: [%d] %s", p.Err, p.Index, p.Input)
}

func ParseArgs(args []string, moduleName string, isPkUuid bool) (GenerateModelInput, []ParsingErrorDetail, bool) {
	var (
		modelName          = args[0]
		attributesArgs     = args[1:]
		parseErrorsDetails []ParsingErrorDetail
		hasParsedError     bool
	)

	response := GenerateModelInput{
		TransformedText: TransformedText{
			SnakeCase:  strcase.ToSnake(modelName),
			PascalCase: strcase.ToCamel(modelName),
			CamelCase:  strcase.ToLowerCamel(modelName),
		},
		ModuleInfo: ModuleInfo{
			Name: moduleName,
		},
		IdAttribute: Attribute{
			TransformedText: TransformedText{
				SnakeCase:  "id",
				PascalCase: "Id",
				CamelCase:  "id",
			},
			Type: underscore.Ternary(isPkUuid, string(TypeUuid), string(TypeInt)),
		},
		Attributes: []Attribute{},
	}

ParsingFor:
	for index, item := range attributesArgs {
		index++
		log.Debug(item)
		separated := strings.Split(item, ":")

		switch len(separated) {
		case 1:
			parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
				Input: item,
				Err:   "not enough data to continue",
			})
			hasParsedError = true
		case 2:
			attribName := separated[0]
			attribType := separated[1]

			if !isValidType(attribType) {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   fmt.Sprintf("%s is not a valid type", attribType),
					Index: index,
				})
				hasParsedError = true
			}

			response.Attributes = append(response.Attributes, Attribute{
				TransformedText: TransformedText{
					SnakeCase:  strcase.ToSnake(attribName),
					PascalCase: strcase.ToCamel(attribName),
					CamelCase:  strcase.ToLowerCamel(attribName),
				},
				Type: attribType,
			})
		case 3:
			attribName := separated[0]
			attribType := separated[1]
			optionalString := separated[2]
			if optionalString != "optional" {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   fmt.Sprintf("expected optional declaration, got '%s'", optionalString),
					Index: index,
				})
				hasParsedError = true
				continue ParsingFor
			}
			isOptional := strings.ToLower(separated[2]) == "optional"

			if !isValidType(attribType) {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   fmt.Sprintf("%s is not a valid type", attribType),
					Index: index,
				})
			}

			response.Attributes = append(response.Attributes, Attribute{
				TransformedText: TransformedText{
					SnakeCase:  strcase.ToSnake(attribName),
					PascalCase: strcase.ToCamel(attribName),
					CamelCase:  strcase.ToLowerCamel(attribName),
				},
				Type:       attribType,
				IsOptional: isOptional,
			})
		}
	}

	return response, parseErrorsDetails, hasParsedError
}
