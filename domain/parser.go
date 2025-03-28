package domain

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/iancoleman/strcase"
	"github.com/rjNemo/underscore"
	"strings"
)

const (
	notEnoughDataToContinue = "not enough data to continue. Remember syntax: attribute:type:optional"
	typeNotSupported        = "type '%s' is not supported"
	optionalNotDeclared     = "expected optional declaration, got '%s'"
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
		var (
			separated        = strings.Split(item, ":")
			separatedLen     = len(separated)
			attribName       string
			attribType       string
			isOptionalString string
		)
		index++
		log.Debug(item)

		switch separatedLen {
		case 1:
			parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
				Input: item,
				Err:   notEnoughDataToContinue,
				Index: index,
			})
			continue ParsingFor
		case 2:
			attribName = separated[0]
			attribType = separated[1]
			if attribType == "" {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   notEnoughDataToContinue,
					Index: index,
				})
				continue ParsingFor
			}
			if !isValidType(attribType) {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   fmt.Sprintf(typeNotSupported, attribType),
					Index: index,
				})
			}
		case 3:
			attribName = separated[0]
			attribType = separated[1]
			isOptionalString = separated[2]
			if !isValidType(attribType) {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   fmt.Sprintf(typeNotSupported, attribType),
					Index: index,
				})
				continue ParsingFor
			}
			if isOptionalString != "optional" {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err: fmt.Sprintf(
						optionalNotDeclared,
						underscore.Ternary(isOptionalString == "", "<empty>", isOptionalString),
					),
					Index: index,
				})
				continue ParsingFor
			}
		}

		switch separatedLen {
		case 2:
			response.Attributes = append(response.Attributes, Attribute{
				TransformedText: TransformedText{
					SnakeCase:  strcase.ToSnake(attribName),
					PascalCase: strcase.ToCamel(attribName),
					CamelCase:  strcase.ToLowerCamel(attribName),
				},
				Type: attribType,
			})
		case 3:
			isOptional := strings.ToLower(isOptionalString) == "optional"

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

	return response, parseErrorsDetails, len(parseErrorsDetails) > 0
}
