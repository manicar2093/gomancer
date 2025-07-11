package commonreq

import "github.com/google/uuid"

type (
	GetById struct {
		Id int `header:"id" param:"id" form:"id" query:"id" json:"id" xml:"id" validate:"required"`
	}

	GetByIdUUID struct {
		Id uuid.UUID `header:"id" param:"id" form:"id" query:"id" json:"id" xml:"id" validate:"required_uuid"`
	}

	PageData struct {
		PageNumber int `header:"page_number" param:"page_number" form:"page_number" query:"page_number" json:"page_number" xml:"page_number" validate:"greaterThan:0"`
		PageSize   int `header:"page_size" param:"page_size" form:"page_size" query:"page_size" json:"page_size" xml:"page_size" validate:"greaterThan:0"`
	}
)
