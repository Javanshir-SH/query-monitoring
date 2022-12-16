package service_test

import (
	"context"
	"errors"
	"github.com/Javanshir-SH/query-monitoring/internal/query/service"
	mock "github.com/Javanshir-SH/query-monitoring/internal/query/service/mocks"
	"github.com/Javanshir-SH/query-monitoring/internal/query/storage"
	"testing"
)

func TestNewQueryMonitoringService_List(t *testing.T) {
	tests := []struct {
		name    string
		repo    mock.MockQueryStatementsRepository
		cache   mock.MockCacheRepo
		wantErr bool
	}{
		{
			name: "add cache error",
			repo: mock.MockQueryStatementsRepository{
				ListFunc: func(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error) {
					return storage.ListOfQuery{}, nil
				},
			},
			cache: mock.MockCacheRepo{
				SetFunc: func(ctx context.Context, key string, sts storage.ListOfQuery) error {
					return errors.New("add cache error")
				},
				GetFunc: func(ctx context.Context, key string) (storage.ListOfQuery, error) {
					return storage.ListOfQuery{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "error storage list",
			repo: mock.MockQueryStatementsRepository{
				ListFunc: func(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error) {
					return storage.ListOfQuery{}, errors.New("storage error")
				},
			},
			cache: mock.MockCacheRepo{
				SetFunc: func(ctx context.Context, key string, sts storage.ListOfQuery) error {
					return nil
				},
				GetFunc: func(ctx context.Context, key string) (storage.ListOfQuery, error) {
					return storage.ListOfQuery{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "success",
			repo: mock.MockQueryStatementsRepository{
				ListFunc: func(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error) {

					return storage.ListOfQuery{
						TotalCount: 123,
						Rows: []*storage.Query{
							{
								QueryID:           123,
								Query:             "test",
								MaxExecutionTime:  12334.33,
								MeanExecutionTime: 134.55,
							},
						},
					}, nil
				},
			},
			cache: mock.MockCacheRepo{
				SetFunc: func(ctx context.Context, key string, sts storage.ListOfQuery) error {
					return nil
				},
				GetFunc: func(ctx context.Context, key string) (storage.ListOfQuery, error) {
					return storage.ListOfQuery{}, nil
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := service.NewQueryMonitoringService(&tt.repo, &tt.cache)

			_, _, _, err := svc.List(context.Background(), service.FilterOptionDto{
				Command: "test",
				Sort:    "test",
				Page:    1,
				PerPage: 3,
			})
			if (err != nil) && !tt.wantErr {
				t.Fatal("unexpected error")
			}

		})
	}
}
