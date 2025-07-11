package parser

import (
	"github.com/manicar2093/gomancer/types"
	"strings"
)

type (
	TransformedText struct {
		SnakeCase        string
		PascalCase       string
		CamelCase        string
		LowerNoSpaceCase string
	}

	ModuleInfo struct {
		Name string
	}
	Attribute struct {
		TransformedText
		Type        string
		IsOptional  bool
		EnumStrings []TransformedText
	}
	GenerateModelInput struct {
		TransformedText
		PackageEntityName string
		ModuleInfo        ModuleInfo
		IdAttribute       Attribute
		Attributes        []Attribute
	}
)

func (c GenerateModelInput) ToStringCmd() string {
	sb := &strings.Builder{}
	sb.WriteString("gomancer gen ")
	if types.SupportedType(c.IdAttribute.Type) == types.TypeUuid {
		sb.WriteString("--pk-uuid ")
	}
	sb.WriteString(c.PascalCase)
	sb.WriteString(" ")
	for _, attr := range c.Attributes {
		sb.WriteString(attr.SnakeCase)
		sb.WriteString(":")
		sb.WriteString(attr.Type)
		if types.SupportedType(attr.Type) == types.TypeEnum {
			for _, enum := range attr.EnumStrings {
				sb.WriteString("/")
				sb.WriteString(enum.SnakeCase)
			}
		}
		if attr.IsOptional {
			sb.WriteString(":")
			sb.WriteString("optional")
		}
		sb.WriteString(" ")
	}

	return sb.String()
}
