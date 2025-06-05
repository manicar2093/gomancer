package tags

import "strings"

const gormTagName = "gorm"

type (
	GormPKOptions struct {
		IsUuid bool
	}

	GormPK struct {
		GormPKOptions
	}
)

func NewGormPK(opts GormPKOptions) Generable {
	return &GormPK{GormPKOptions: opts}
}

func (g GormPK) Generate() (string, string) {
	sb := &strings.Builder{}

	if g.IsUuid {
		sb.WriteString("default:gen_random_uuid()")
		sb.WriteString(";")
	}
	sb.WriteString("primaryKey")

	return gormTagName, sb.String()
}
