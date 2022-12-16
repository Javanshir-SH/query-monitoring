package v1

import (
	"github.com/Javanshir-SH/query-monitoring/internal/todo/entities"
)

//go:generate moq -pkg mock -out ./mocks/service.go . Service:MockService
type Service interface {
	Create(t entities.Task) error
	Get(id int) (*entities.Task, error)
	Update(t entities.Task) error
	Delete(id int) error
}
