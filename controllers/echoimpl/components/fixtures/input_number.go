package fixtures

const (
	OptionalInputNumberInt = `{{ anOptionalIntKey := "an_optional_int" }}
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
		Value:    userData.AnOptionalInt,
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
`
	RequiredInputNumberInt = `{{ anIntKey := "an_int" }}
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
		Value:    userData.AnInt,
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
`
)
