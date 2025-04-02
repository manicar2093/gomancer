package prismaimpl

import (
	"fmt"
	"github.com/manicar2093/gomancer/domain"
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

func GenerateMigration(input domain.GenerateModelInput) error {
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
	switch domain.SupportedType(typ) {
	case domain.TypeInt, domain.TypeInt8, domain.TypeInt16, domain.TypeInt32:
		return "Int"
	case domain.TypeInt64:
		return "BigInt"
	case domain.TypeFloat32, domain.TypeFloat64:
		return "Float"
	case domain.TypeString:
		return "String"
	case domain.TypeBool:
		return "Boolean"
	case domain.TypeTime:
		return "DateTime"
	case domain.TypeDecimal:
		return "Decimal"
	case domain.TypeUuid:
		return "String"
	default:
		panic(fmt.Sprintf("unsupported type %s for prisma migrations", typ))
	}
}

func getTypeAttribute(t string) string {
	switch domain.SupportedType(t) {
	case domain.TypeDecimal:
		return "@db.Decimal(10,2)"
	case domain.TypeUuid:
		return "@db.Uuid"
	default:
		return empty
	}

}
