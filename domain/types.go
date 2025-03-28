package domain

import (
	"github.com/charmbracelet/log"
	"github.com/dave/jennifer/jen"
)

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

	ValidTypesMap = map[SupportedType]bool{
		TypeInt:     true,
		TypeInt8:    true,
		TypeInt16:   true,
		TypeInt32:   true,
		TypeInt64:   true,
		TypeFloat32: true,
		TypeFloat64: true,
		TypeString:  true,
		TypeBool:    true,
		TypeTime:    true,
		TypeDecimal: true,
		TypeUuid:    true,
	}
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
	_, ok := ValidTypesMap[SupportedType(t)]
	log.Debug("validating type", "type", t, "ok", ok)
	return ok
}
