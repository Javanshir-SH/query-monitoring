package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RequestBody struct {
	ID    int    `json:"ID"`
	Title string `json:"title"`
}

func (req RequestBody) Validate() error {

	return validation.ValidateStruct(&req,
		validation.Field(&req.ID, validation.Required),
		validation.Field(&req.ID, validation.Required),
	)
}
