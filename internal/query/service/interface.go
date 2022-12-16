package service

import (
	"context"
	"github.com/Javanshir-SH/query-monitoring/internal/query/storage"
)

//go:generate moq -pkg mock -out ./mocks/repo.go . QueryStatementsRepository:MockQueryStatementsRepository
type QueryStatementsRepository interface {
	List(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error)
}

//go:generate moq -pkg mock -out ./mocks/cache.go . CacheRepo:MockCacheRepo
type CacheRepo interface {
	Set(ctx context.Context, key string, sts storage.ListOfQuery) error
	Get(ctx context.Context, key string) (storage.ListOfQuery, error)
}
