package domain

import "github.com/dave/jennifer/jen"

type SupportedType string

var (
	TypeInt     SupportedType = "int"
	TypeInt8    SupportedType = "int8"
	TypeInt16   SupportedType = "int16"
	TypeInt32   SupportedType = "int32"
	TypeInt64   SupportedType = "int64"
	TypeFloat32 SupportedType = "float32"
	TypeFloat64 SupportedType = "float64"
	TypeString  SupportedType = "string"
	TypeBool    SupportedType = "bool"
	TypeTime    SupportedType = "time"
	TypeDecimal SupportedType = "decimal"
	TypeUuid    SupportedType = "uuid"
)

func QualifiersByType(t string) *jen.Statement {
	switch t {
	case "uuid":
		return jen.Qual(UUIDPkgPath, "UUID")
	case "decimal":
		return jen.Qual(DecimalPkgPath, "Decimal")
	default:
		return jen.Id(t)
	}
}

func OptionalQualifier() jen.Code {
	return jen.Qual(GoptionPkgPath, "Optional")
}

func isValidType(t string) bool {
	return true
}
