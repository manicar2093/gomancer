package domain

import (
	"strings"
)

type (
	Tag         string
	Validations struct {
		Required bool
	}
)

const (
	JsonTag         Tag = "json"
	MapstructureTag     = "mapstructure"
	ParamTag            = "param"
	GormUuidTag         = "gorm|default:gen_random_uuid()"
)

func Tags(attribute Attribute, validations Validations, tags ...Tag) map[string]string {
	response := make(map[string]string)
	for _, tag := range tags {
		if tag == GormUuidTag {
			d := strings.Split(string(tag), "|")
			response[d[0]] = d[1]
			continue
		}
		response[string(tag)] = attribute.SnakeCase
	}

	validationsSB := strings.Builder{}
	if validations.Required {
		if attribute.Type == string(TypeUuid) {
			validationsSB.WriteString("required_uuid")
		} else {
			validationsSB.WriteString("required")
		}
	}
	if valString := validationsSB.String(); valString != "" {
		response["validate"] = valString
	}

	return response
}
