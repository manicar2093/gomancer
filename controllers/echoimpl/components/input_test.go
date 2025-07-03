package components_test

import (
	"fmt"
	"github.com/manicar2093/gomancer/controllers/echoimpl/components"
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
		Entry("not optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_int",
				PascalCase:       "AnInt",
				CamelCase:        "anInt",
				LowerNoSpaceCase: "anint",
			},
			Type: "int",
		}, `{{ anIntKey := "an_int" }}
{{ hasAnIntErrors := errors.HasField(anIntKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: anIntKey,
	}) {
		AnInt
	}
	@input.Input(input.Props{
		ID:       anIntKey,
		Name:     anIntKey,
		Type:     input.TypeNumber,
		Value:    postTest.AnInt,
		HasError: hasAnIntErrors,
		Required: true,
	})
	@form.Description() {
		Enter AnInt
	}
	if hasAnIntErrors {
		for _,value := range errors.Field(anIntKey) {
			@form.Message(form.MessageProps{
				Variant: form.MessageVariantError,
			}) {
				{ value }
			}
		}
	}
}
`),
		Entry("optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "an_optional_int",
				PascalCase:       "AnOptionalInt",
				CamelCase:        "anOptionalInt",
				LowerNoSpaceCase: "anoptionalint",
			},
			Type:       "int",
			IsOptional: true,
		}, `{{ anOptionalIntKey := "an_optional_int" }}
{{ hasAnOptionalIntErrors := errors.HasField(anOptionalIntKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: anOptionalIntKey,
	}) {
		AnOptionalInt
	}
	@input.Input(input.Props{
		ID:       anOptionalIntKey,
		Name:     anOptionalIntKey,
		Type:     input.TypeNumber,
		Value:    postTest.AnOptionalInt,
		HasError: hasAnOptionalIntErrors,
	})
	@form.Description() {
		Enter AnOptionalInt
	}
	if hasAnOptionalIntErrors {
		for _,value := range errors.Field(anOptionalIntKey) {
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

	//	Describe("InputText", func() {
	//		It("should generate HTML for a text input", func() {
	//			// Use test attribute from testfixtures
	//			attr := testfixtures.ModelBinaryIdSuccess.Attributes[16] // OptionalString (optional string)
	//
	//			// Create the InputText component
	//			component := InputText{
	//				Attribute: attr,
	//			}
	//
	//			// Generate the HTML
	//			result, err := component.Generate()
	//
	//			Expect(err).ToNot(HaveOccurred())
	//			Expect(result).To(Equal(""))
	//
	//		})
	//	})
	//
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
