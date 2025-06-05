package tags

import "strings"

type (
	JsonOptions struct {
		Name        string
		IsOmitEmpty bool
		IsOmitZero  bool
		IsInline    bool
	}
	Json struct {
		JsonOptions
	}
)

func NewJson(options JsonOptions) Generable {
	return Json{JsonOptions: options}
}

func (c Json) Generate() (string, string) {
	sb := &strings.Builder{}

	if c.IsInline {
		sb.WriteString(",inline")
		return c.Name, sb.String()
	}
	sb.WriteString(c.Name)
	if c.IsOmitEmpty {
		sb.WriteString(",omitempty")
	}
	if c.IsOmitZero {
		sb.WriteString(",omitzero")
	}
	return c.Name, sb.String()
}
