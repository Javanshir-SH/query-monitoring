package service_test

import (
	"errors"
	"github.com/Javanshir-SH/query-monitoring/internal/todo/entities"
	"github.com/Javanshir-SH/query-monitoring/internal/todo/service"
	mock "github.com/Javanshir-SH/query-monitoring/internal/todo/service/mocks"
	"github.com/Javanshir-SH/query-monitoring/internal/todo/storage"

	"testing"
)

func TestTaskService_Create(t *testing.T) {
	taskId := 1
	tests := []struct {
		name    string
		task    entities.Task
		repo    mock.MockTaskRepository
		wantErr bool
	}{
		{
			name: "success",
			task: entities.Task{
				ID:    taskId,
				Title: "test",
			},
			repo: mock.MockTaskRepository{
				CreateFunc: func(t storage.Task) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "storage error",
			task: entities.Task{
				ID:    taskId,
				Title: "test",
			},
			repo: mock.MockTaskRepository{
				CreateFunc: func(t storage.Task) error {
					return errors.New("storage error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewTaskService(&tt.repo)

			err := svc.Create(tt.task)
			if (err != nil) && !tt.wantErr {
				t.Fatal("unexpected error")
			}
		})
	}
}

func TestTaskService_Get(t *testing.T) {
	taskId := 1
	tests := []struct {
		name    string
		Id      int
		repo    mock.MockTaskRepository
		wantErr bool
	}{
		{
			name: "success",
			Id:   taskId,
			repo: mock.MockTaskRepository{
				GetFunc: func(t storage.Task) (*storage.Task, error) {
					return &storage.Task{
						ID:    taskId,
						Title: "test",
					}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "storage error",
			Id:   taskId,
			repo: mock.MockTaskRepository{
				GetFunc: func(t storage.Task) (*storage.Task, error) {
					return nil, errors.New("todo storage get todo error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewTaskService(&tt.repo)

			_, err := svc.Get(tt.Id)
			if (err != nil) && !tt.wantErr {
				t.Fatal("unexpected error")
			}
		})
	}

}

func TestTaskService_Update(t *testing.T) {
	taskId := 1
	tests := []struct {
		name    string
		task    entities.Task
		repo    mock.MockTaskRepository
		wantErr bool
	}{
		{
			name: "success",
			task: entities.Task{
				ID:    taskId,
				Title: "test",
			},
			repo: mock.MockTaskRepository{
				UpdateFunc: func(cn storage.Task) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "storage error",
			task: entities.Task{
				ID:    taskId,
				Title: "test",
			},
			repo: mock.MockTaskRepository{
				UpdateFunc: func(cn storage.Task) error {
					return errors.New("storage error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewTaskService(&tt.repo)

			err := svc.Update(tt.task)
			if (err != nil) && !tt.wantErr {
				t.Fatal("unexpected error")
			}
		})
	}

}

func TestTaskService_Delete(t *testing.T) {
	taskId := 1
	tests := []struct {
		name    string
		Id      int
		repo    mock.MockTaskRepository
		wantErr bool
	}{
		{
			name: "success",
			Id:   taskId,
			repo: mock.MockTaskRepository{
				DeleteFunc: func(cn storage.Task) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "storage error",
			Id:   taskId,
			repo: mock.MockTaskRepository{
				DeleteFunc: func(cn storage.Task) error {
					return errors.New("storage error")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewTaskService(&tt.repo)

			err := svc.Delete(tt.Id)
			if (err != nil) && !tt.wantErr {
				t.Fatal("unexpected error")
			}
		})
	}
}
