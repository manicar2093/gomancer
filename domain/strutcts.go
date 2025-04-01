package domain

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
)
