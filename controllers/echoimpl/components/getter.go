package components

import (
	"fmt"
	"github.com/manicar2093/gomancer/parser"
	"github.com/manicar2093/gomancer/types"
)

var ErrNoComponentFound = fmt.Errorf("no component found")

func GetInputByType(attribute parser.Attribute, modelBeingGenerated parser.TransformedText) (string, error) {
	input := InputGenerationData{
		Attribute:            attribute,
		ModelTransformedText: modelBeingGenerated,
	}
	switch types.SupportedType(attribute.Type) {
	case types.TypeInt, types.TypeInt8, types.TypeInt16, types.TypeInt32, types.TypeInt64:
		return InputNumber(input)
	case types.TypeFloat32, types.TypeFloat64, types.TypeDecimal:
		return InputNumberFloat(input)
	case types.TypeString, types.TypeUuid:
		return InputText(input)
	case types.TypeBool:
		return InputToggle(input)
	case types.TypeTime:
		return InputDateTime(input)
	case types.TypeEnum:
		return InputSelectBox(input)
	default:
		return "", ErrNoComponentFound
	}
}
