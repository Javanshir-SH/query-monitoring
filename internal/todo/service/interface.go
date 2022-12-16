package service

import stg "github.com/Javanshir-SH/query-monitoring/internal/todo/storage"

//go:generate moq -pkg mock -out ./mocks/storage.go . TaskRepository:MockTaskRepository
type TaskRepository interface {
	Create(t stg.Task) error
	Get(t stg.Task) (*stg.Task, error)
	Update(t stg.Task) error
	Delete(t stg.Task) error
}
