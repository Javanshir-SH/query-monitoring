package service

import (
	"github.com/Javanshir-SH/query-monitoring/internal/todo/entities"
	"github.com/Javanshir-SH/query-monitoring/internal/todo/storage"
)

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return TaskService{repo: r}

}

func (svc TaskService) Create(t entities.Task) error {

	task := storage.Task{
		ID:    t.ID,
		Title: t.Title,
	}

	return svc.repo.Create(task)

}

func (svc TaskService) Get(id int) (*entities.Task, error) {
	task := storage.Task{
		ID: id,
	}

	t, err := svc.repo.Get(task)
	if err != nil {
		return nil, err
	}

	return &entities.Task{
		ID:    t.ID,
		Title: t.Title,
	}, nil

}

func (svc TaskService) Update(t entities.Task) error {
	task := storage.Task{
		ID:    t.ID,
		Title: t.Title,
	}

	return svc.repo.Update(task)
}

func (svc TaskService) Delete(id int) error {
	task := storage.Task{
		ID: id,
	}

	return svc.repo.Delete(task)
}
