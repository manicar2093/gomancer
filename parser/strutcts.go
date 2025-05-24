package parser

type (
	TransformedText struct {
		SnakeCase        string
		PascalCase       string
		CamelCase        string
		LowerNoSpaceCase string
	}

	ModuleInfo struct {
		Name string
	}
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
