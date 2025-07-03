package fixtures

const (
	OptionalInputNumberFloat = `{{ anOptionalFloat64Key := "an_optional_float_64" }}
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
		Value:    postTest.AnOptionalFloat64,
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
`
	RequiredInputNumberFloat = `{{ anFloat32Key := "an_float_32" }}
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
		Value:    postTest.AnFloat32,
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
`
)
