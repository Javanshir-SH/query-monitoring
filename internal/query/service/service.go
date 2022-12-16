package service

import (
	"context"
	"fmt"
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/utils"
	"github.com/Javanshir-SH/query-monitoring/internal/query/storage"
)

type QueryMonitoringService struct {
	repo  QueryStatementsRepository
	cache CacheRepo
}

func NewQueryMonitoringService(repo QueryStatementsRepository, cache CacheRepo) QueryMonitoringService {
	return QueryMonitoringService{
		repo:  repo,
		cache: cache,
	}
}

func (s QueryMonitoringService) List(ctx context.Context, dto FilterOptionDto) (
	res ListOfQueryStatementsDto,
	totalCount int,
	pageCount int,
	err error) {
	option := storage.FilterOption{
		Command: dto.Command,
		Sort:    dto.Sort,
		Limit:   dto.PerPage,
		Offset:  (dto.Page - 1) * dto.PerPage,
	}

	// Get data from cache
	var results storage.ListOfQuery
	cacheKey := fmt.Sprintf("%s-%s-%d-%d", dto.Command, dto.Sort, dto.Page, dto.PerPage)

	results, err = s.cache.Get(ctx, cacheKey)
	if err != nil {

		// Get from db
		results, err = s.repo.List(ctx, option)
		if err != nil {
			return ListOfQueryStatementsDto{}, 0, 0, err
		}

		//Add the latest page to cache
		err = s.cache.Set(ctx, cacheKey, results)
		if err != nil {
			return ListOfQueryStatementsDto{}, 0, 0, err
		}

	}

	// Set total count
	totalCount = results.TotalCount

	output := make(ListOfQueryStatementsDto, 0, len(results.Rows))
	for _, result := range results.Rows {
		output = append(output, &QueryStatement{
			QueryID:           result.QueryID,
			Query:             result.Query,
			MaxExecutionTime:  result.MaxExecutionTime,
			MeanExecutionTime: result.MeanExecutionTime,
		})
	}

	return output, totalCount, utils.PageCount(totalCount, dto.PerPage), nil
}
