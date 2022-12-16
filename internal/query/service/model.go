package service

import validation "github.com/go-ozzo/ozzo-validation"

const (
	SELECT = "select"
	UPDATE = "update"
	INSERT = "insert"
	DELETE = "delete"

	ASC  = "asc"
	DESC = "desc"
)

type FilterOptionDto struct {
	Command string
	Sort    string
	Page    int
	PerPage int
}

func (dto FilterOptionDto) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Command, validation.Required, validation.In(SELECT, UPDATE, INSERT, DELETE)),
		validation.Field(&dto.Sort, validation.Required, validation.In(ASC, DESC)),
		validation.Field(&dto.Page, validation.Required),
		validation.Field(&dto.PerPage, validation.Required),
	)
}

type QueryStatement struct {
	QueryID           int64
	Query             string
	MaxExecutionTime  float64
	MeanExecutionTime float64
}

type ListOfQueryStatementsDto []*QueryStatement
