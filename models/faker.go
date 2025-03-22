package models

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/domain"
)

const fakerPkgPath = "github.com/brianvoe/gofakeit/v7"

func FakerCallByType(t string) jen.Code {
	switch t {
	case "string":
		return jen.Qual(fakerPkgPath, "Word").Call()
	case "int":
		return jen.Qual(fakerPkgPath, "Int").Call()
	case "int8":
		return jen.Qual(fakerPkgPath, "Int8").Call()
	case "int16":
		return jen.Qual(fakerPkgPath, "Int16").Call()
	case "int32":
		return jen.Qual(fakerPkgPath, "Int32").Call()
	case "int64":
		return jen.Qual(fakerPkgPath, "Int64").Call()
	case "uint":
		return jen.Qual(fakerPkgPath, "UInt").Call()
	case "uint8":
		return jen.Qual(fakerPkgPath, "UInt8").Call()
	case "uint16":
		return jen.Qual(fakerPkgPath, "UInt16").Call()
	case "uint32":
		return jen.Qual(fakerPkgPath, "UInt32").Call()
	case "float32":
		return jen.Qual(fakerPkgPath, "Float32").Call()
	case "float64":
		return jen.Qual(fakerPkgPath, "Float64").Call()
	case "time":
		return jen.Qual(fakerPkgPath, "Date").Call()
	case "decimal":
		return jen.Qual(domain.DecimalPkgPath, "MustFromFloat64").Call(
			jen.Qual(fakerPkgPath, "Float64").Call(),
		)
	case "uuid":
		return jen.Qual(domain.UUIDPkgPath, "New").Call()
	default:
		panic(fmt.Sprintf("type %s can be faked", t))
	}
}
