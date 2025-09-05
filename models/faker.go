package models

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/manicar2093/gomancer/deps"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
)

func FakerCallByType(t string, goDeps deps.Container, attribute parser.Attribute) jen.Code {
	switch types.SupportedType(t) {
	case types.TypeString:
		return jen.Qual(goDeps.GoFakeIt.Path, "Word").Call()
	case types.TypeInt:
		return jen.Qual(goDeps.GoFakeIt.Path, "Int").Call()
	case types.TypeInt8:
		return jen.Qual(goDeps.GoFakeIt.Path, "Int8").Call()
	case types.TypeInt16:
		return jen.Qual(goDeps.GoFakeIt.Path, "Int16").Call()
	case types.TypeInt32:
		return jen.Qual(goDeps.GoFakeIt.Path, "Int32").Call()
	case types.TypeInt64:
		return jen.Qual(goDeps.GoFakeIt.Path, "Int64").Call()
	case types.TypeFloat32:
		return jen.Qual(goDeps.GoFakeIt.Path, "Float32").Call()
	case types.TypeFloat64:
		return jen.Qual(goDeps.GoFakeIt.Path, "Float64").Call()
	case types.TypeTime:
		return jen.Qual(goDeps.GoFakeIt.Path, "Date").Call()
	case types.TypeDecimal:
		return jen.Qual(goDeps.UDecimal.Path, "MustFromFloat64").Call(
			jen.Qual(goDeps.GoFakeIt.Path, "Float64").Call(),
		)
	case types.TypeUuid:
		return jen.Qual(goDeps.Uuid.Path, "New").Call()
	case types.TypeBool:
		return jen.False()
	case types.TypeEnum:
		return jen.Qual(goDeps.Internal.Models.Path, attribute.EnumStrings[0].PascalCase)
	default:
		panic(fmt.Sprintf("type %s can be faked", t))
	}
}
