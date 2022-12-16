package v1_test

import (
	"bytes"
	"encoding/json"
	"errors"
	v1 "github.com/Javanshir-SH/query-monitoring/internal/todo/delivery/http/v1"
	mock "github.com/Javanshir-SH/query-monitoring/internal/todo/delivery/http/v1/mocks"
	"github.com/Javanshir-SH/query-monitoring/internal/todo/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoHandler_Create(t *testing.T) {
	l, _ := zap.NewProduction()

	tests := []struct {
		name       string
		handler    *v1.TodoHandler
		route      string
		body       entities.Task
		statusCode int
	}{
		{
			name: "success",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				CreateFunc: func(t entities.Task) error {
					return nil
				},
			}),
			route: "/todos",
			body: entities.Task{
				ID:    1,
				Title: "test",
			},
			statusCode: 201,
		},

		{
			name: "bad request error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				CreateFunc: func(t entities.Task) error {
					return nil
				},
			}),
			route:      "/todos",
			body:       entities.Task{},
			statusCode: 400,
		},

		{
			name: "internal error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				CreateFunc: func(t entities.Task) error {
					return errors.New("server error")
				},
			}),
			route: "/todos",
			body: entities.Task{
				ID:    1,
				Title: "test",
			},
			statusCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Post("/todos", tt.handler.Create)

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(tt.body)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(
				http.MethodPost,
				tt.route,
				&buf,
			)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)

			err = app.Shutdown()
			if err != nil {
				t.Fatal(err)
			}
		})
	}

}

func TestTodoHandler_Get(t *testing.T) {
	l, _ := zap.NewProduction()

	tests := []struct {
		name       string
		handler    *v1.TodoHandler
		route      string
		Id         int
		statusCode int
	}{
		{
			name: "success",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				GetFunc: func(id int) (*entities.Task, error) {
					return &entities.Task{
						ID:    1,
						Title: "test",
					}, nil
				},
			}),
			route:      "/todos/1",
			Id:         1,
			statusCode: 200,
		},

		{
			name: "bad request error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				GetFunc: func(id int) (*entities.Task, error) {
					return nil, nil
				},
			}),
			route:      "/todos/a",
			Id:         1,
			statusCode: 400,
		},

		{
			name: "internal error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				GetFunc: func(id int) (*entities.Task, error) {
					return nil, errors.New("server error")
				},
			}),
			route:      "/todos/1",
			Id:         1,
			statusCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/todos/:id", tt.handler.Get)

			req := httptest.NewRequest(
				http.MethodGet,
				tt.route,
				nil,
			)

			resp, err := app.Test(req)

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)

			err = app.Shutdown()
			if err != nil {
				t.Fatal(err)
			}
		})
	}

}

func TestTodoHandler_Update(t *testing.T) {
	l, _ := zap.NewProduction()

	tests := []struct {
		name       string
		handler    *v1.TodoHandler
		route      string
		body       entities.Task
		statusCode int
	}{
		{
			name: "success",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				UpdateFunc: func(t entities.Task) error {
					return nil
				},
			}),
			route: "/todos",
			body: entities.Task{
				ID:    1,
				Title: "test",
			},
			statusCode: 200,
		},

		{
			name: "bad request error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				UpdateFunc: func(t entities.Task) error {
					return nil
				},
			}),
			route:      "/todos",
			body:       entities.Task{},
			statusCode: 400,
		},

		{
			name: "internal error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				UpdateFunc: func(t entities.Task) error {
					return errors.New("server error")
				},
			}),
			route: "/todos",
			body: entities.Task{
				ID:    1,
				Title: "test",
			},
			statusCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Put("/todos", tt.handler.Update)

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(tt.body)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(
				http.MethodPut,
				tt.route,
				&buf,
			)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)

			err = app.Shutdown()
			if err != nil {
				t.Fatal(err)
			}
		})
	}

}

func TestTodoHandler_Delete(t *testing.T) {
	l, _ := zap.NewProduction()

	tests := []struct {
		name       string
		handler    *v1.TodoHandler
		route      string
		body       entities.Task
		statusCode int
	}{
		{
			name: "success",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				UpdateFunc: func(t entities.Task) error {
					return nil
				},
			}),
			route: "/todos",
			body: entities.Task{
				ID:    1,
				Title: "test",
			},
			statusCode: 200,
		},

		{
			name: "bad request error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				UpdateFunc: func(t entities.Task) error {
					return nil
				},
			}),
			route:      "/todos",
			body:       entities.Task{},
			statusCode: 400,
		},

		{
			name: "internal error",
			handler: v1.NewTodoHandler(l, &mock.MockService{
				UpdateFunc: func(t entities.Task) error {
					return errors.New("server error")
				},
			}),
			route: "/todos",
			body: entities.Task{
				ID:    1,
				Title: "test",
			},
			statusCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Put("/todos", tt.handler.Update)

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(tt.body)
			if err != nil {
				t.Fatal(err)
			}

			req := httptest.NewRequest(
				http.MethodPut,
				tt.route,
				&buf,
			)
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)

			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.statusCode, resp.StatusCode)

			err = app.Shutdown()
			if err != nil {
				t.Fatal(err)
			}
		})
	}

}
