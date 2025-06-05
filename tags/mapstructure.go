package tags

import "strings"

const (
	mapstructureTagName = "mapstructure"
)

type (
	MapstructureOptions struct {
		Name        string
		IsSquash    bool
		IsRemain    bool
		IsOmitEmpty bool
	}

	Mapstructure struct {
		MapstructureOptions
	}
)

func NewMapstructure(options MapstructureOptions) Generable {
	return Mapstructure{MapstructureOptions: options}
}

func (c Mapstructure) Generate() (string, string) {
	sb := &strings.Builder{}

	if c.IsSquash {
		sb.WriteString(",squash")
		return mapstructureTagName, sb.String()
	}

	if c.IsRemain {
		sb.WriteString(",remain")
		return mapstructureTagName, sb.String()
	}

	sb.WriteString(c.Name)
	if c.IsOmitEmpty {
		sb.WriteString(",omitempty")
	}

	return mapstructureTagName, sb.String()
}
