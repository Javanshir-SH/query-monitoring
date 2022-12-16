package v1

import (
	"context"
	"github.com/Javanshir-SH/query-monitoring/internal/query/service"
)

//go:generate moq -pkg mock -out ./mocks/services.go . Service:MockService
type Service interface {
	List(ctx context.Context, dto service.FilterOptionDto) (
		res service.ListOfQueryStatementsDto,
		totalCount int,
		pageCount int,
		err error)
}
