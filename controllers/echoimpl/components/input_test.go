package components_test

import (
	"github.com/manicar2093/gomancer/controllers/echoimpl/components"
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
		Value:    strconv.Itoa(userData.AnOptionalInt.GetValue()),
		HasError: hasAnOptionalIntErrors,
	})
	@form.Description() {
		Enter AnOptionalInt
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasAnOptionalIntErrors,
		Errors: errors,
		FormItemKey: anOptionalIntKey,
	})
}
`),
		Entry("not optional int", parser.Attribute{
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
		Value:    strconv.Itoa(userData.AnInt),
		HasError: hasAnIntErrors,
		Required: true,
	})
	@form.Description() {
		Enter AnInt
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasAnIntErrors,
		Errors: errors,
		FormItemKey: anIntKey,
	})
}
`),
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
		Value:    strconv.FormatFloat(userData.AnOptionalFloat64.GetValue(), 'f', 2, 64),
		HasError: hasAnOptionalFloat64Errors,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter AnOptionalFloat64
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasAnOptionalFloat64Errors,
		Errors: errors,
		FormItemKey: anOptionalFloat64Key,
	})
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
		Value:    strconv.FormatFloat(float64(userData.AnFloat32), 'f', 2, 64),
		HasError: hasAnFloat32Errors,
		Required: true,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter AnFloat32
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasAnFloat32Errors,
		Errors: errors,
		FormItemKey: anFloat32Key,
	})
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
		Value:    userData.OptionalDecimal.GetValue().StringFixed(2),
		HasError: hasOptionalDecimalErrors,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter OptionalDecimal
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasOptionalDecimalErrors,
		Errors: errors,
		FormItemKey: optionalDecimalKey,
	})
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
		Value:    userData.Decimal.StringFixed(2),
		HasError: hasDecimalErrors,
		Required: true,
		Attributes: map[string]any{
			"step": "0.01",
		},
	})
	@form.Description() {
		Enter Decimal
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasDecimalErrors,
		Errors: errors,
		FormItemKey: decimalKey,
	})
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
		Value:    userData.OptionalString.GetValue(),
		HasError: hasOptionalStringErrors,
	})
	@form.Description() {
		Enter OptionalString
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasOptionalStringErrors,
		Errors: errors,
		FormItemKey: optionalStringKey,
	})
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
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasStringErrors,
		Errors: errors,
		FormItemKey: stringKey,
	})
}
`),
	)

	DescribeTable("InputTextUuid", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputTextUuid(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},
		Entry("optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "optional_uuid",
				PascalCase:       "OptionalUuid",
				CamelCase:        "optionalUuid",
				LowerNoSpaceCase: "optionaluuid",
			},
			Type:       "uuid",
			IsOptional: true,
		}, `{{ optionalUuidKey := "optional_uuid" }}
{{ hasOptionalUuidErrors := errors.HasField(optionalUuidKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: optionalUuidKey,
	}) {
		OptionalUuid
	}
	@input.Input(input.Props{
		ID:       optionalUuidKey,
		Name:     optionalUuidKey,
		Type:     input.TypeText,
		Value:    userData.OptionalUuid.GetValue().String(),
		HasError: hasOptionalUuidErrors,
		Attributes: templ.Attributes{
			"pattern": utils.UuidV4Pattern,
			"title": i18n.T(ctx, "uuid_v4_html_input_title"),
		},
	})
	@form.Description() {
		Enter OptionalUuid
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasOptionalUuidErrors,
		Errors: errors,
		FormItemKey: optionalUuidKey,
	})
}
`),
		Entry("not optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "uuid",
				PascalCase:       "Uuid",
				CamelCase:        "uuid",
				LowerNoSpaceCase: "uuid",
			},
			Type: "uuid",
		}, `{{ uuidKey := "uuid" }}
{{ hasUuidErrors := errors.HasField(uuidKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: uuidKey,
	}) {
		Uuid
	}
	@input.Input(input.Props{
		ID:       uuidKey,
		Name:     uuidKey,
		Type:     input.TypeText,
		Value:    userData.Uuid.String(),
		HasError: hasUuidErrors,
		Required: true,
		Attributes: templ.Attributes{
			"pattern": utils.UuidV4Pattern,
			"title": i18n.T(ctx, "uuid_v4_html_input_title"),
		},
	})
	@form.Description() {
		Enter Uuid
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasUuidErrors,
		Errors: errors,
		FormItemKey: uuidKey,
	})
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
		Checked: userData.OptionalBool.GetValue(),
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
		Value:    userData.OptionalTime.GetValue().UTC(),
		HasError: hasOptionalTimeErrors,
	})
	@form.Description() {
		Enter OptionalTime
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasOptionalTimeErrors,
		Errors: errors,
		FormItemKey: optionalTimeKey,
	})
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
		Value:    userData.Time.UTC(),
		HasError: hasTimeErrors,
		Required: true,
	})
	@form.Description() {
		Enter Time
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasTimeErrors,
		Errors: errors,
		FormItemKey: timeKey,
	})
}
`),
	)

	DescribeTable("InputSelectBox", func(attr parser.Attribute, expectedResult string) {
		result, err := components.InputSelectBox(components.InputGenerationData{
			Attribute:            attr,
			ModelTransformedText: generateModelTransformedText,
		})

		Expect(err).ToNot(HaveOccurred())
		Expect(result).To(Equal(expectedResult))

	},

		Entry("optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "optional_enum",
				PascalCase:       "OptionalEnum",
				CamelCase:        "optionalEnum",
				LowerNoSpaceCase: "optionalenum",
			},
			Type:       "enum",
			IsOptional: true,
			EnumStrings: []parser.TransformedText{
				{
					SnakeCase:        "optional_enum_1",
					PascalCase:       "OptionalEnum1",
					CamelCase:        "optionalEnum1",
					LowerNoSpaceCase: "optionalenum1",
				},
				{
					SnakeCase:        "optional_enum_2",
					PascalCase:       "OptionalEnum2",
					CamelCase:        "optionalEnum2",
					LowerNoSpaceCase: "optionalenum2",
				},
				{
					SnakeCase:        "optional_enum_3",
					PascalCase:       "OptionalEnum3",
					CamelCase:        "optionalEnum3",
					LowerNoSpaceCase: "optionalenum3",
				},
			},
		}, `{{ optionalEnumKey := "optional_enum" }}
{{ hasOptionalEnumErrors := errors.HasField(optionalEnumKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: optionalEnumKey,
	}) {
		OptionalEnum
	}
	@selectbox.SelectBox() {
		@selectbox.Trigger(selectbox.TriggerProps{
			ID:       optionalEnumKey,
			Name:     optionalEnumKey,
			HasError: hasOptionalEnumErrors,
		}) {
			@selectbox.Value(selectbox.ValueProps{
				Placeholder: i18n.T(ctx, "select_placeholder"),
			})
		}
		@selectbox.Content() {
			@selectbox.Group() {
				@selectbox.Item(selectbox.ItemProps{
					Value:    "optional_enum_1",
					Selected: userData.OptionalEnum.GetValue() == "optional_enum_1",
					Disabled: userData.OptionalEnum.GetValue() == "optional_enum_1",
				}) {
					OptionalEnum1
				}
				@selectbox.Item(selectbox.ItemProps{
					Value:    "optional_enum_2",
					Selected: userData.OptionalEnum.GetValue() == "optional_enum_2",
					Disabled: userData.OptionalEnum.GetValue() == "optional_enum_2",
				}) {
					OptionalEnum2
				}
				@selectbox.Item(selectbox.ItemProps{
					Value:    "optional_enum_3",
					Selected: userData.OptionalEnum.GetValue() == "optional_enum_3",
					Disabled: userData.OptionalEnum.GetValue() == "optional_enum_3",
				}) {
					OptionalEnum3
				}
			}
		}
	}
	@form.Description() {
		Select OptionalEnum
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasOptionalEnumErrors,
		Errors: errors,
		FormItemKey: optionalEnumKey,
	})
}
`),
		Entry("not optional", parser.Attribute{
			TransformedText: parser.TransformedText{
				SnakeCase:        "enum",
				PascalCase:       "Enum",
				CamelCase:        "enum",
				LowerNoSpaceCase: "enum",
			},
			Type: "enum",
			EnumStrings: []parser.TransformedText{
				{
					SnakeCase:        "enum_1",
					PascalCase:       "Enum1",
					CamelCase:        "enum1",
					LowerNoSpaceCase: "enum1",
				},
				{
					SnakeCase:        "enum_2",
					PascalCase:       "Enum2",
					CamelCase:        "enum2",
					LowerNoSpaceCase: "enum2",
				},
				{
					SnakeCase:        "enum_3",
					PascalCase:       "Enum3",
					CamelCase:        "enum3",
					LowerNoSpaceCase: "enum3",
				},
			},
		}, `{{ enumKey := "enum" }}
{{ hasEnumErrors := errors.HasField(enumKey) }}
@form.Item(form.ItemProps{}) {
	@label.Label(label.Props{
		For: enumKey,
	}) {
		Enum
	}
	@selectbox.SelectBox() {
		@selectbox.Trigger(selectbox.TriggerProps{
			ID:       enumKey,
			Name:     enumKey,
			HasError: hasEnumErrors,
			Required: true,
		}) {
			@selectbox.Value(selectbox.ValueProps{
				Placeholder: i18n.T(ctx, "select_placeholder"),
			})
		}
		@selectbox.Content() {
			@selectbox.Group() {
				@selectbox.Item(selectbox.ItemProps{
					Value:    "enum_1",
					Selected: userData.Enum == "enum_1",
					Disabled: userData.Enum == "enum_1",
				}) {
					Enum1
				}
				@selectbox.Item(selectbox.ItemProps{
					Value:    "enum_2",
					Selected: userData.Enum == "enum_2",
					Disabled: userData.Enum == "enum_2",
				}) {
					Enum2
				}
				@selectbox.Item(selectbox.ItemProps{
					Value:    "enum_3",
					Selected: userData.Enum == "enum_3",
					Disabled: userData.Enum == "enum_3",
				}) {
					Enum3
				}
			}
		}
	}
	@form.Description() {
		Select Enum
	}
	@formerrors.FormErrors(formerrors.FormErrorsProps{
		HasErrors: hasEnumErrors,
		Errors: errors,
		FormItemKey: enumKey,
	})
}
`),
	)

})
