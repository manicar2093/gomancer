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
	DescribeTable("InputNumber", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputNumber(attr)
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
		result, err := components.InputNumberFloat(attr)

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
		}, fixtures.OptionalInputNumberFloat),
		Entry("not optional float", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_float_32",
				PascalCase:       "AnFloat32",
				CamelCase:        "anFloat32",
				LowerNoSpaceCase: "anfloat32",
			},
			Type: "float32",
		}, fixtures.RequiredInputNumberFloat),
	)

	DescribeTable("InputText", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputText(attr)

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
		Value:    postTest.OptionalString,
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
		Value:    postTest.String,
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
		result, err := components.InputToggle(attr)
		fmt.Println(result)
		fmt.Println(expectedResult)

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
		Checked: postTest.OptionalBool,
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
		Checked: postTest.Bool,
	})
	@form.Description() {
		Check for Bool
	}
}
`),
	)

	//	Describe("InputDateTime", func() {
	//		It("should generate HTML for a datetime input", func() {
	//			// Use test attribute from testfixtures
	//			attr := testfixtures.ModelBinaryIdSuccess.Attributes[21] // Time (non-optional time)
	//			// Expected output
	//			expected := `{{ decimalKey := "decimal" }}
	//{{ hasDecimalErrors := errors.HasField(decimalKey) }}
	//@form.Item(form.ItemProps{}) {
	//        @label.Label(label.Props{
	//                For: decimalKey,
	//        }) {
	//                Decimal
	//        }
	//        @datetime.Datetime(datetime.DatetimeProps{
	//                ID:       decimalKey,
	//                Name:     decimalKey,
	//                Value:    postTest.Decimal.UTC(),
	//                HasError: hasDecimalErrors,
	//                Required: true,
	//        })
	//        @form.Description() {
	//                Enter Decimal
	//        }
	//        if hasDecimalErrors {
	//                for _,value := range errors.Field(decimalKey) {
	//                        @form.Message(form.MessageProps{
	//                                Variant: form.MessageVariantError,
	//                        }) {
	//                                { value }
	//                        }
	//                }
	//        }
	//}%`
	//
	//			// Create the InputDateTime component
	//			component := InputDateTime{
	//				Attribute: attr,
	//			}
	//
	//			// Generate the HTML
	//			result, err := component.Generate()
	//
	//			Expect(err).ToNot(HaveOccurred())
	//			Expect(result).To(Equal(expected))
	//		})
	//	})
	//
	//	Describe("InputToggle", func() {
	//		It("should generate HTML for a toggle input", func() {
	//			// Use test attribute from testfixtures
	//			attr := testfixtures.ModelBinaryIdSuccess.Attributes[18] // OptionalBool (optional bool)
	//
	//			// Create the InputToggle component
	//			component := InputToggle{
	//				Attribute: attr,
	//			}
	//			// Expected output
	//			expected := `{{ optionalTimeKey := "optional_time" }}
	//@label.Label(label.Props{
	//        For: optionalTimeKey,
	//}) {
	//        OptionalTime
	//}
	//@form.ItemFlex() {
	//        @toggle.Toggle(toggle.Props{
	//                ID:      optionalTimeKey,
	//                Name:    optionalTimeKey,
	//                Checked: postTest.OptionalTime,
	//        })
	//        @form.Description() {
	//                Check for OptionalTime
	//        }
	//}%`
	//
	//			// Generate the HTML
	//			// Generate the HTML
	//			result, err := component.Generate()
	//
	//			Expect(err).ToNot(HaveOccurred())
	//			Expect(result).To(Equal(expected))
	//		})
	//	})
	//
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
	//					Selected: postTest.Enum == "enum_1",
	//					Disabled: postTest.Enum == "enum_1",
	//				}) {
	//					Enum1
	//				}
	//				@selectbox.Item(selectbox.ItemProps{
	//					Value:    "enum_2",
	//					Selected: postTest.Enum == "enum_2",
	//					Disabled: postTest.Enum == "enum_2",
	//				}) {
	//					Enum2
	//				}
	//				@selectbox.Item(selectbox.ItemProps{
	//					Value:    "enum_3",
	//					Selected: postTest.Enum == "enum_3",
	//					Disabled: postTest.Enum == "enum_3",
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
