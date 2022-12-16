package v1

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

const (
	SELECT = "select"
	UPDATE = "update"
	INSERT = "insert"
	DELETE = "delete"

	ASC            = "asc"
	DESC           = "desc"
	pageDefault    = 1
	perPageDefault = 10
)

type FilterQueryParam struct {
	Type    string `json:"type"`
	Sort    string `json:"sort"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

func (p *FilterQueryParam) checkAndSetDefaults() {
	if p.Page <= 0 {
		p.Page = pageDefault
	}

	if p.PerPage <= 0 {
		p.PerPage = perPageDefault
	}

}

func (p FilterQueryParam) Validate() error {

	return validation.ValidateStruct(&p,
		validation.Field(&p.Type, validation.In(SELECT, UPDATE, INSERT, DELETE)),
		validation.Field(&p.Sort, validation.In(ASC, DESC)),
	)
}
