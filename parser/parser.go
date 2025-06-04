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
			separated    = strings.Split(item, ":")
			separatedLen = len(separated)
			attribName   string
		)
		index++
		log.Debug(item)

		if separatedLen <= 1 {
			parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
				Input: item,
				Err:   notEnoughDataToContinue,
				Index: index,
			})
			continue ParsingFor
		}

		tempAttrib := Attribute{}

		for attribIndex, attrib := range separated {
			if attribIndex == 0 {
				continue
			}

			if attribIndex == 1 {
				enumData, isEnumType := func(argument string) ([]string, bool) {
					if !strings.Contains(argument, "enum") {
						return nil, false
					}
					argumentSeparated := strings.Split(argument, "/")
					return argumentSeparated[1:len(argumentSeparated)], strings.ToLower(argumentSeparated[0]) == "enum"
				}(attrib)
				if isEnumType {
					attrib = string(types.TypeEnum)
				}

				isValidType := types.IsValidType(attrib)
				if isValidType {
					tempAttrib.Type = attrib
					if isEnumType {
						tempAttrib.EnumStrings = underscore.Map(enumData, func(enumItem string) TransformedText {
							return TransformedText{
								SnakeCase:        strcase.ToSnake(enumItem),
								PascalCase:       strcase.ToCamel(enumItem),
								CamelCase:        enumItem,
								LowerNoSpaceCase: strings.ToLower(strcase.ToLowerCamel(enumItem)),
							}
						})
					}
				}
				if !isValidType {
					parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
						Input: item,
						Err:   fmt.Sprintf(typeNotSupported, attrib),
						Index: index,
					})
					continue ParsingFor

				}
			}

			if attribIndex == 2 {
				isOptional := strings.ToLower(attrib) == "optional"
				if isOptional {
					tempAttrib.IsOptional = true
				}
				if !isOptional {
					parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
						Input: item,
						Err: fmt.Sprintf(
							optionalNotDeclared,
							underscore.Ternary(attrib == "", "<empty>", attrib),
						),
						Index: index,
					})
					continue ParsingFor
				}
			}
		}

		attribName = separated[0]

		attributeNameCamelCase := strcase.ToLowerCamel(attribName)
		tempAttrib.TransformedText = TransformedText{
			SnakeCase:        strcase.ToSnake(attribName),
			PascalCase:       strcase.ToCamel(attribName),
			CamelCase:        attributeNameCamelCase,
			LowerNoSpaceCase: strings.ToLower(attributeNameCamelCase),
		}
		response.Attributes = append(response.Attributes, tempAttrib)
	}

	return response, parseErrorsDetails, len(parseErrorsDetails) > 0
}
