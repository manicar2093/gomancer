package prismaimpl

import (
	"embed"
	"fmt"
	"github.com/jinzhu/inflection"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"github.com/rjNemo/underscore"
	"strings"
	"text/template"
)

const (
	space   = " "
	empty   = ""
	newLine = "\n"
)

//go:embed templates/*
var templatesFS embed.FS
var funcMap = template.FuncMap{
	"CreateIdAttribute": func(attribute parser.Attribute) string {
		sb := strings.Builder{}

		writeString(&sb, attribute.SnakeCase, "     ", space)
		writeString(
			&sb,
			getMigrationType(attribute.Type, attribute),
			empty,
			space,
		)
		writeString(&sb, "@id", empty, space)
		idDefault := "@default(%s)"
		idGenerator := underscore.Ternary[string](
			types.SupportedType(attribute.Type) == types.TypeUuid,
			"dbgenerated(\"gen_random_uuid()\")",
			"autoincrement()",
		)
		writeString(&sb, fmt.Sprintf(idDefault, idGenerator), empty, space)
		typeAttribute := getTypeAttribute(attribute.Type)
		writeString(
			&sb,
			typeAttribute,
			underscore.Ternary(typeAttribute == "", empty, space),
			newLine,
		)
		return strings.TrimRight(sb.String(), " ")
	},
	"CreateAttribute": func(attribute parser.Attribute) string {
		sb := strings.Builder{}

		writeString(&sb, attribute.SnakeCase, "     ", space)
		writeString(
			&sb,
			getMigrationType(attribute.Type, attribute),
			empty,
			underscore.Ternary[string](attribute.IsOptional, "?", empty),
		)
		typeAttribute := getTypeAttribute(attribute.Type)
		writeString(
			&sb,
			typeAttribute,
			underscore.Ternary(typeAttribute == "", empty, space),
			newLine,
		)
		return sb.String()
	},
	"Pluralize": inflection.Plural,
	"CreateEnumValue": func(attribute parser.TransformedText) string {
		sb := strings.Builder{}

		writeString(&sb, attribute.SnakeCase, "     ", newLine)

		return sb.String()
	},
}

func writeString(sb *strings.Builder, content, prefix, suffix string) {
	sb.WriteString(fmt.Sprintf("%s%s%s", prefix, content, suffix))

}
