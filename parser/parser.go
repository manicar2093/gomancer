package parser

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/manicar2093/gomancer/types"
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

	modelNameCamelCase := strcase.ToLowerCamel(modelName)
	lowerNoSpaceCase := strings.ToLower(modelNameCamelCase)
	packageName := pluralize.NewClient().Plural(lowerNoSpaceCase)
	response := GenerateModelInput{
		PackageEntityName: packageName,
		TransformedText: TransformedText{
			SnakeCase:        strcase.ToSnake(modelName),
			PascalCase:       strcase.ToCamel(modelName),
			CamelCase:        modelNameCamelCase,
			LowerNoSpaceCase: lowerNoSpaceCase,
		},
		ModuleInfo: ModuleInfo{
			Name: moduleName,
		},
		IdAttribute: Attribute{
			TransformedText: TransformedText{
				SnakeCase:        "id",
				PascalCase:       "Id",
				CamelCase:        "id",
				LowerNoSpaceCase: "id",
			},
			Type: underscore.Ternary(isPkUuid, string(types.TypeUuid), string(types.TypeInt)),
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
			if !types.IsValidType(attribType) {
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
			if !types.IsValidType(attribType) {
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

		attributeNameCamelCase := strcase.ToLowerCamel(attribName)
		switch separatedLen {
		case 2:
			response.Attributes = append(response.Attributes, Attribute{
				TransformedText: TransformedText{
					SnakeCase:        strcase.ToSnake(attribName),
					PascalCase:       strcase.ToCamel(attribName),
					CamelCase:        attributeNameCamelCase,
					LowerNoSpaceCase: strings.ToLower(attributeNameCamelCase),
				},
				Type: attribType,
			})
		case 3:
			isOptional := strings.ToLower(isOptionalString) == "optional"

			response.Attributes = append(response.Attributes, Attribute{
				TransformedText: TransformedText{
					SnakeCase:        strcase.ToSnake(attribName),
					PascalCase:       strcase.ToCamel(attribName),
					CamelCase:        strcase.ToLowerCamel(attribName),
					LowerNoSpaceCase: strings.ToLower(attributeNameCamelCase),
				},
				Type:       attribType,
				IsOptional: isOptional,
			})
		}
	}

	return response, parseErrorsDetails, len(parseErrorsDetails) > 0
}
