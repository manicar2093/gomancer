package types

import (
	"github.com/charmbracelet/log"
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/deps"
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
	TypeEnum    SupportedType = "enum"

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
		TypeEnum:    true,
	}
)

func QualifiersByType(t string, goDeps deps.Container) *jen.Statement {
	switch t {
	case string(TypeUuid):
		return jen.Qual(goDeps.Uuid.Path, "UUID")
	case string(TypeDecimal):
		return jen.Qual(goDeps.UDecimal.Path, "Decimal")
	case string(TypeTime):
		return jen.Qual(goDeps.Std.Time.Path, "Time")
	default:
		return jen.Id(t)
	}
}

func OptionalQualifier(goDeps deps.Container) jen.Code {
	return jen.Qual(goDeps.Goption.Path, "Optional")
}

func IsValidType(t string) bool {
	_, ok := ValidTypesMap[SupportedType(t)]
	log.Debug("validating type", "type", t, "ok", ok)
	return ok
}
