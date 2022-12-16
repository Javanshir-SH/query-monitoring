package v1_test

import (
	"context"
	"errors"
	v1 "github.com/Javanshir-SH/query-monitoring/internal/query/delivery/http/v1"
	mock "github.com/Javanshir-SH/query-monitoring/internal/query/delivery/http/v1/mocks"
	"github.com/Javanshir-SH/query-monitoring/internal/query/service"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestQueryMonitoringHandler_List(t *testing.T) {
	l, _ := zap.NewProduction()

	tests := []struct {
		name       string
		handler    *v1.QueryMonitoringHandler
		route      string
		statusCode int
	}{
		{
			name:       "invalid type param",
			handler:    v1.NewQueryMonitoringHandler(l, &mock.MockService{}),
			route:      "/queries?type=selectu&sort=asc&page=1&per_page=2",
			statusCode: 400,
		},
		{
			name:       "invalid sort param",
			handler:    v1.NewQueryMonitoringHandler(l, &mock.MockService{}),
			route:      "/queries?type=select&sort=ascT&page=1&per_page=2",
			statusCode: 400,
		},
		{
			name: "service error",
			handler: v1.NewQueryMonitoringHandler(l, &mock.MockService{
				ListFunc: func(ctx context.Context, dto service.FilterOptionDto) (service.ListOfQueryStatementsDto, int, int, error) {
					return nil, 0, 0, errors.New("internal error")
				},
			}),
			route:      "/queries?type=select&sort=asc&page=1&per_page=2",
			statusCode: 500,
		},
		{
			name: "success",
			handler: v1.NewQueryMonitoringHandler(l, &mock.MockService{
				ListFunc: func(ctx context.Context, dto service.FilterOptionDto) (service.ListOfQueryStatementsDto, int, int, error) {
					r := service.ListOfQueryStatementsDto{
						{
							QueryID:           123,
							Query:             "test",
							MeanExecutionTime: 134.123,
							MaxExecutionTime:  12334.23,
						},
					}
					return r, 0, 0, nil
				},
			}),
			route:      "/queries?type=select&sort=asc&page=1&per_page=2",
			statusCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Get("/queries", tt.handler.List)
			resp, err := app.Test(httptest.NewRequest(
				http.MethodGet,
				tt.route,
				nil,
			))

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
