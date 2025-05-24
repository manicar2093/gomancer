package models

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
	"github.com/manicar2093/gomancer/types"
)

const fakerPkgPath = "github.com/brianvoe/gofakeit/v7"

func FakerCallByType(t string) jen.Code {
	switch types.SupportedType(t) {
	case types.TypeString:
		return jen.Qual(fakerPkgPath, "Word").Call()
	case types.TypeInt:
		return jen.Qual(fakerPkgPath, "Int").Call()
	case types.TypeInt8:
		return jen.Qual(fakerPkgPath, "Int8").Call()
	case types.TypeInt16:
		return jen.Qual(fakerPkgPath, "Int16").Call()
	case types.TypeInt32:
		return jen.Qual(fakerPkgPath, "Int32").Call()
	case types.TypeInt64:
		return jen.Qual(fakerPkgPath, "Int64").Call()
	case types.TypeFloat32:
		return jen.Qual(fakerPkgPath, "Float32").Call()
	case types.TypeFloat64:
		return jen.Qual(fakerPkgPath, "Float64").Call()
	case types.TypeTime:
		return jen.Qual(fakerPkgPath, "Date").Call()
	case types.TypeDecimal:
		return jen.Qual(domain.DecimalPkgPath, "MustFromFloat64").Call(
			jen.Qual(fakerPkgPath, "Float64").Call(),
		)
	case types.TypeUuid:
		return jen.Qual(domain.UUIDPkgPath, "New").Call()
	default:
		panic(fmt.Sprintf("type %s can be faked", t))
	}
}
