package components_test

import (
	"fmt"
	"github.com/manicar2093/gomancer/controllers/echoimpl/components"
	"github.com/manicar2093/gomancer/controllers/echoimpl/components/fixtures"
	"github.com/manicar2093/gomancer/parser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Input Components", func() {

	var generateModelTransformedText = parser.TransformedText{
		SnakeCase:        "user_data",
		PascalCase:       "UserData",
		CamelCase:        "userData",
		LowerNoSpaceCase: "userdata",
	}

	DescribeTable("InputNumber", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputNumber(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})
		fmt.Println(result)

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},
		Entry("optional int", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_optional_int",
				PascalCase:       "AnOptionalInt",
				CamelCase:        "anOptionalInt",
				LowerNoSpaceCase: "anoptionalint",
			},
			Type:       "int",
			IsOptional: true,
		}, fixtures.OptionalInputNumberInt),
		Entry("not optional int", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_int",
				PascalCase:       "AnInt",
				CamelCase:        "anInt",
				LowerNoSpaceCase: "anint",
			},
			Type: "int",
		}, fixtures.RequiredInputNumberInt),
	)

	DescribeTable("InputNumberFloat", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputNumberFloat(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},
		Entry("optional float", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_optional_float_64",
				PascalCase:       "AnOptionalFloat64",
				CamelCase:        "anOptionalFloat64",
				LowerNoSpaceCase: "anoptionalfloat64",
			},
			Type:       "float64",
			IsOptional: true,
		}, `{{ anOptionalFloat64Key := "an_optional_float_64" }}
{{ hasAnOptionalFloat64Errors := errors.HasField(anOptionalFloat64Key) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: anOptionalFloat64Key,
	}) {
		AnOptionalFloat64
	}
	@input.Input(input.Props{
		ID:       anOptionalFloat64Key,
		Name:     anOptionalFloat64Key,
		Type:     input.TypeNumber,
		Value:    userData.AnOptionalFloat64,
		HasError: hasAnOptionalFloat64Errors,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter AnOptionalFloat64
	}
	if hasAnOptionalFloat64Errors {
		for _,value := range errors.Field(anOptionalFloat64Key) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
		Entry("not optional float", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_float_32",
				PascalCase:       "AnFloat32",
				CamelCase:        "anFloat32",
				LowerNoSpaceCase: "anfloat32",
			},
			Type: "float32",
		}, `{{ anFloat32Key := "an_float_32" }}
{{ hasAnFloat32Errors := errors.HasField(anFloat32Key) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: anFloat32Key,
	}) {
		AnFloat32
	}
	@input.Input(input.Props{
		ID:       anFloat32Key,
		Name:     anFloat32Key,
		Type:     input.TypeNumber,
		Value:    userData.AnFloat32,
		HasError: hasAnFloat32Errors,
		Required: true,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter AnFloat32
	}
	if hasAnFloat32Errors {
		for _,value := range errors.Field(anFloat32Key) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
		Entry("optional decimal", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "optional_decimal",
				PascalCase:       "OptionalDecimal",
				CamelCase:        "optionalDecimal",
				LowerNoSpaceCase: "optionaldecimal",
			},
			Type:       "decimal",
			IsOptional: true,
		}, `{{ optionalDecimalKey := "optional_decimal" }}
{{ hasOptionalDecimalErrors := errors.HasField(optionalDecimalKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: optionalDecimalKey,
	}) {
		OptionalDecimal
	}
	@input.Input(input.Props{
		ID:       optionalDecimalKey,
		Name:     optionalDecimalKey,
		Type:     input.TypeNumber,
		Value:    userData.OptionalDecimal,
		HasError: hasOptionalDecimalErrors,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter OptionalDecimal
	}
	if hasOptionalDecimalErrors {
		for _,value := range errors.Field(optionalDecimalKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
		Entry("not optional decimal", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "decimal",
				PascalCase:       "Decimal",
				CamelCase:        "decimal",
				LowerNoSpaceCase: "decimal",
			},
			Type: "decimal",
		}, `{{ decimalKey := "decimal" }}
{{ hasDecimalErrors := errors.HasField(decimalKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: decimalKey,
	}) {
		Decimal
	}
	@input.Input(input.Props{
		ID:       decimalKey,
		Name:     decimalKey,
		Type:     input.TypeNumber,
		Value:    userData.Decimal,
		HasError: hasDecimalErrors,
		Required: true,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter Decimal
	}
	if hasDecimalErrors {
		for _,value := range errors.Field(decimalKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
	)

	DescribeTable("InputText", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputText(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},

		Entry("optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "optional_string",
				PascalCase:       "OptionalString",
				CamelCase:        "optionalString",
				LowerNoSpaceCase: "optionalstring",
			},
			Type:       "string",
			IsOptional: true,
		}, `{{ optionalStringKey := "optional_string" }}
{{ hasOptionalStringErrors := errors.HasField(optionalStringKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: optionalStringKey,
	}) {
		OptionalString
	}
	@input.Input(input.Props{
		ID:       optionalStringKey,
		Name:     optionalStringKey,
		Type:     input.TypeText,
		Value:    userData.OptionalString,
		HasError: hasOptionalStringErrors,
	})
	@form.Description() {
		Enter OptionalString
	}
	if hasOptionalStringErrors {
		for _,value := range errors.Field(optionalStringKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
		Entry("not optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "string",
				PascalCase:       "String",
				CamelCase:        "string",
				LowerNoSpaceCase: "string",
			},
			Type: "string",
		}, `{{ stringKey := "string" }}
{{ hasStringErrors := errors.HasField(stringKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: stringKey,
	}) {
		String
	}
	@input.Input(input.Props{
		ID:       stringKey,
		Name:     stringKey,
		Type:     input.TypeText,
		Value:    userData.String,
		HasError: hasStringErrors,
		Required: true,
	})
	@form.Description() {
		Enter String
	}
	if hasStringErrors {
		for _,value := range errors.Field(stringKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
	)

	DescribeTable("InputToggle", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputToggle(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},

		Entry("optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "optional_bool",
				PascalCase:       "OptionalBool",
				CamelCase:        "optionalBool",
				LowerNoSpaceCase: "optionalbool",
			},
			Type:       "bool",
			IsOptional: true,
		}, `{{ optionalBoolKey := "optional_bool" }}
@label.Label(label.Props{
	For: optionalBoolKey,
}) {
	OptionalBool
}
@form.ItemFlex() {
	@toggle.Toggle(toggle.Props{
		ID:      optionalBoolKey,
		Name:    optionalBoolKey,
		Checked: userData.OptionalBool,
	})
	@form.Description() {
		Check for OptionalBool
	}
}
`),
		Entry("not optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "bool",
				PascalCase:       "Bool",
				CamelCase:        "bool",
				LowerNoSpaceCase: "bool",
			},
			Type: "bool",
		}, `{{ boolKey := "bool" }}
@label.Label(label.Props{
	For: boolKey,
}) {
	Bool
}
@form.ItemFlex() {
	@toggle.Toggle(toggle.Props{
		ID:      boolKey,
		Name:    boolKey,
		Checked: userData.Bool,
	})
	@form.Description() {
		Check for Bool
	}
}
`),
	)

	DescribeTable("InputDateTime", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputDateTime(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},

		Entry("optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "optional_time",
				PascalCase:       "OptionalTime",
				CamelCase:        "optionalTime",
				LowerNoSpaceCase: "optionaltime",
			},
			Type:       "time",
			IsOptional: true,
		}, `{{ optionalTimeKey := "optional_time" }}
{{ hasOptionalTimeErrors := errors.HasField(optionalTimeKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: optionalTimeKey,
	}) {
		OptionalTime
	}
	@datetime.Datetime(datetime.DatetimeProps{
		ID:       optionalTimeKey,
		Name:     optionalTimeKey,
		Value:    postTest.OptionalTime.UTC(),
		HasError: hasOptionalTimeErrors,
	})
	@form.Description() {
		Enter OptionalTime
	}
	if hasOptionalTimeErrors {
		for _,value := range errors.Field(optionalTimeKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
		Entry("not optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "time",
				PascalCase:       "Time",
				CamelCase:        "time",
				LowerNoSpaceCase: "time",
			},
			Type: "time",
		}, `{{ timeKey := "time" }}
{{ hasTimeErrors := errors.HasField(timeKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: timeKey,
	}) {
		Time
	}
	@datetime.Datetime(datetime.DatetimeProps{
		ID:       timeKey,
		Name:     timeKey,
		Value:    postTest.Time.UTC(),
		HasError: hasTimeErrors,
		Required: true,
	})
	@form.Description() {
		Enter Time
	}
	if hasTimeErrors {
		for _,value := range errors.Field(timeKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
	)

	//	Describe("InputSelectBox", func() {
	//		It("should generate HTML for a select box with enum values", func() {
	//			// Use test attribute from testfixtures
	//			attr := testfixtures.ModelBinaryIdSuccess.Attributes[25] // Enum (non-optional enum)
	//
	//			// Create the InputSelectBox component
	//			component := InputSelectBox{
	//				Attribute: attr,
	//			}
	//			expected := `{{ enumKey := "enum" }}
	//{{ hasEnumErrors := errors.HasField(enumKey) }}
	//@form.Item(form.ItemProps{}) {
	//	@label.Label(label.Props{
	//		For: enumKey,
	//	}) {
	//		Enum
	//	}
	//	@selectbox.SelectBox() {
	//		@selectbox.Trigger(selectbox.TriggerProps{
	//			ID:       enumKey,
	//			Name:     enumKey,
	//			HasError: hasEnumErrors,
	//			Required: true,
	//		}) {
	//			@selectbox.Value(selectbox.ValueProps{
	//				Placeholder: i18n.T(ctx, "select_placeholder"),
	//			})
	//		}
	//		@selectbox.Content() {
	//			@selectbox.Group() {
	//				@selectbox.Item(selectbox.ItemProps{
	//					Value:    "enum_1",
	//					Selected: userData.Enum == "enum_1",
	//					Disabled: userData.Enum == "enum_1",
	//				}) {
	//					Enum1
	//				}
	//				@selectbox.Item(selectbox.ItemProps{
	//					Value:    "enum_2",
	//					Selected: userData.Enum == "enum_2",
	//					Disabled: userData.Enum == "enum_2",
	//				}) {
	//					Enum2
	//				}
	//				@selectbox.Item(selectbox.ItemProps{
	//					Value:    "enum_3",
	//					Selected: userData.Enum == "enum_3",
	//					Disabled: userData.Enum == "enum_3",
	//				}) {
	//					Enum3
	//				}
	//			}
	//		}
	//	}
	//	@form.Description() {
	//		Select Enum
	//	}
	//	if hasEnumErrors {
	//		for _,value := range errors.Field(enumKey) {
	//			@form.Message(form.MessageProps{
	//				Variant: form.MessageVariantError,
	//			}) {
	//				{ value }
	//			}
	//		}
	//	}
	//}`
	//
	//			result, err := component.Generate()
	//
	//			Expect(err).ToNot(HaveOccurred())
	//			Expect(result).To(Equal(expected))
	//		})
	//	})
})
