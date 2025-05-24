package prismaimpl

import (
	"fmt"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
	"os"
	"path"
	"strings"
	"text/template"
)

const (
	space   = " "
	empty   = ""
	newLine = "\n"
)

func GenerateMigration(input parser.GenerateModelInput) error {
	var tpl = template.Must(template.
		New("migrations").
		Funcs(funcMap).
		ParseFS(templatesFS, "templates/*"))

	migrationFilePath := path.Join(string(domain.PrismaSchemaPackagePath), fmt.Sprintf("%s.prisma", input.SnakeCase))
	f, err := os.OpenFile(migrationFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	if err := tpl.ExecuteTemplate(f, "migration", input); err != nil {
		return err
	}
	return nil
}

func writeString(sb *strings.Builder, content, prefix, suffix string) {
	sb.WriteString(fmt.Sprintf("%s%s%s", prefix, content, suffix))

}

func getMigrationType(typ string) string {
	switch types.SupportedType(typ) {
	case types.TypeInt, types.TypeInt8, types.TypeInt16, types.TypeInt32:
		return "Int"
	case types.TypeInt64:
		return "BigInt"
	case types.TypeFloat32, types.TypeFloat64:
		return "Float"
	case types.TypeString:
		return "String"
	case types.TypeBool:
		return "Boolean"
	case types.TypeTime:
		return "DateTime"
	case types.TypeDecimal:
		return "Decimal"
	case types.TypeUuid:
		return "String"
	default:
		panic(fmt.Sprintf("unsupported type %s for prisma migrations", typ))
	}
}

func getTypeAttribute(t string) string {
	switch types.SupportedType(t) {
	case types.TypeDecimal:
		return "@db.Decimal(10,2)"
	case types.TypeUuid:
		return "@db.Uuid"
	default:
		return empty
	}

}
