package entities

import validation "github.com/go-ozzo/ozzo-validation"

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func (p Task) Validate() error {

	return validation.ValidateStruct(&p,
		validation.Field(&p.ID, validation.Required, validation.Min(1)),
		validation.Field(&p.Title, validation.Required, validation.Length(1, 100)),
	)
}
