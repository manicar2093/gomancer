package domain

import "github.com/dave/jennifer/jen"

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
