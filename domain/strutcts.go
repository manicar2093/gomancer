package domain

type (
	TransformedText struct {
		SnakeCase  string
		PascalCase string
		CamelCase  string
	}

	ModuleInfo struct {
		Name string
	}
)
