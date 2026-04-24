package parser

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gertd/go-pluralize"
	"github.com/manicar2093/gomancer/types"
	"github.com/rjNemo/underscore"
	strcase "github.com/stoewer/go-strcase"
)

const (
	notEnoughDataToContinue = "not enough data to continue. Remember syntax: attribute:type:optional"
	typeNotSupported        = "type '%s' is not supported"
	optionalNotDeclared     = "expected optional declaration, got '%s'"
	noEnumValuesFound       = "enum type must have at least one value"
	enumSeparator           = "/"
	argumentsSeparator      = ":"
	enumKeyword             = "enum"
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

	modelNameCamelCase := strcase.LowerCamelCase(modelName)
	lowerNoSpaceCase := strings.ToLower(modelNameCamelCase)
	packageName := pluralize.NewClient().Plural(lowerNoSpaceCase)
	response := GenerateModelInput{
		PackageEntityName: packageName,
		TransformedText: TransformedText{
			SnakeCase:        strcase.SnakeCase(modelName),
			PascalCase:       strcase.UpperCamelCase(modelName),
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

	for index, item := range attributesArgs {
		var (
			separated    = strings.Split(item, argumentsSeparator)
			separatedLen = len(separated)
		)
		index++
		log.Debug(item)

		if separatedLen <= 1 {
			parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
				Input: item,
				Err:   notEnoughDataToContinue,
				Index: index,
			})
			continue
		}

		var fieldName, fieldType, fieldOptional string
		var validateOptional, isFieldOptional bool

		switch separatedLen {
		case 2:
			fieldName = separated[0]
			fieldType = separated[1]
		case 3:
			fieldName = separated[0]
			fieldType = separated[1]
			fieldOptional = separated[2]
			validateOptional = true
		}

		fieldTypeNormalized, enumValues, isValidType, err := inferFieldType(fieldType)
		if err != nil {
			parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
				Input: item,
				Err:   err.Error(),
				Index: index,
			})
			continue
		}
		if !isValidType {
			parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
				Input: item,
				Err:   fmt.Sprintf(typeNotSupported, fieldType),
				Index: index,
			})
			continue
		}

		if validateOptional {
			if fieldOptional != "optional" {
				parseErrorsDetails = append(parseErrorsDetails, ParsingErrorDetail{
					Input: item,
					Err:   fmt.Sprintf(optionalNotDeclared, underscore.Ternary(fieldOptional == "", "<empty>", fieldOptional)),
					Index: index,
				})
				continue
			}
			isFieldOptional = true
		}

		tempAttrib := Attribute{
			TransformedText: TransformedText{},
			Type:            fieldTypeNormalized,
			IsOptional:      isFieldOptional,
		}

		if enumValues != nil {
			tempAttrib.EnumStrings = underscore.Map(enumValues, transformText)
		}

		tempAttrib.TransformedText = transformText(fieldName)
		response.Attributes = append(response.Attributes, tempAttrib)
	}

	return response, parseErrorsDetails, len(parseErrorsDetails) > 0
}

func inferFieldType(fieldType string) (string, []string, bool, error) {
	if fieldType == "" {
		return "", nil, false, fmt.Errorf(notEnoughDataToContinue)
	}
	if strings.HasPrefix(fieldType, enumKeyword) {
		argumentSeparated := strings.Split(fieldType, enumSeparator)
		if len(argumentSeparated) <= 1 {
			return string(types.TypeEnum), nil, false, fmt.Errorf(noEnumValuesFound)
		}
		return string(types.TypeEnum), argumentSeparated[1:len(argumentSeparated)], strings.ToLower(argumentSeparated[0]) == enumKeyword, nil
	}
	return fieldType, nil, types.IsValidType(fieldType), nil
}

func transformText(name string) TransformedText {
	return TransformedText{
		SnakeCase:        strcase.SnakeCase(name),
		PascalCase:       strcase.UpperCamelCase(name),
		CamelCase:        strcase.LowerCamelCase(name),
		LowerNoSpaceCase: strings.ToLower(strcase.LowerCamelCase(name)),
	}
}
