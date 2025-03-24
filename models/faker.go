package models

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
)

const fakerPkgPath = "github.com/brianvoe/gofakeit/v7"

func FakerCallByType(t string) jen.Code {
	switch domain.SupportedType(t) {
	case domain.TypeString:
		return jen.Qual(fakerPkgPath, "Word").Call()
	case domain.TypeInt:
		return jen.Qual(fakerPkgPath, "Int").Call()
	case domain.TypeInt8:
		return jen.Qual(fakerPkgPath, "Int8").Call()
	case domain.TypeInt16:
		return jen.Qual(fakerPkgPath, "Int16").Call()
	case domain.TypeInt32:
		return jen.Qual(fakerPkgPath, "Int32").Call()
	case domain.TypeInt64:
		return jen.Qual(fakerPkgPath, "Int64").Call()
	case domain.TypeFloat32:
		return jen.Qual(fakerPkgPath, "Float32").Call()
	case domain.TypeFloat64:
		return jen.Qual(fakerPkgPath, "Float64").Call()
	case domain.TypeTime:
		return jen.Qual(fakerPkgPath, "Date").Call()
	case domain.TypeDecimal:
		return jen.Qual(domain.DecimalPkgPath, "MustFromFloat64").Call(
			jen.Qual(fakerPkgPath, "Float64").Call(),
		)
	case domain.TypeUuid:
		return jen.Qual(domain.UUIDPkgPath, "New").Call()
	default:
		panic(fmt.Sprintf("type %s can be faked", t))
	}
}
