package domain

type (
	Attribute struct {
		TransformedText
		Type       string
		IsOptional bool
	}
	GenerateModelInput struct {
		TransformedText
		PackageEntityName string
		ModuleInfo        ModuleInfo
		IdAttribute       Attribute
		Attributes        []Attribute
	}
)
