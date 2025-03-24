package prisma

import (
	"embed"
	"fmt"
	"github.com/manicar2093/gomancer/domain"
	"github.com/rjNemo/underscore"
	"strings"
	"text/template"
)

//go:embed templates/*
var templatesFS embed.FS
var funcMap = template.FuncMap{
	"CreateIdAttribute": func(attribute domain.Attribute) string {
		sb := strings.Builder{}

		writeString(&sb, attribute.SnakeCase, "     ", space)
		writeString(
			&sb,
			getMigrationType(attribute.Type),
			empty,
			space,
		)
		writeString(&sb, "@id", empty, space)
		idDefault := "@default(%s)"
		idGenerator := underscore.Ternary[string](
			domain.SupportedType(attribute.Type) == domain.TypeUuid,
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
	"CreateAttribute": func(attribute domain.Attribute) string {
		sb := strings.Builder{}

		writeString(&sb, attribute.SnakeCase, "     ", space)
		writeString(
			&sb,
			getMigrationType(attribute.Type),
			empty,
			underscore.Ternary(attribute.IsOptional, "?", empty),
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
}
