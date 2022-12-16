package storage

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

const (
	defaultSorting = "DESC"
	defaultLimit   = 10
)

type QueryMonitoring struct {
	db *gorm.DB
}

// NewQueryMonitoringRepo constructor for Query
func NewQueryMonitoringRepo(db *gorm.DB) QueryMonitoring {
	return QueryMonitoring{db: db}
}

func (s QueryMonitoring) List(ctx context.Context, option FilterOption) (ListOfQuery, error) {
	result := ListOfQuery{}
	rows := make([]*Query, 0)

	q := s.db.WithContext(ctx)

	// if option limit is 0 apply default value
	if option.Limit == 0 {
		option.Limit = defaultLimit
	}

	// if option sort is empty apply default value
	if option.Sort == "" {
		option.Sort = defaultSorting
	}

	// apply option command filter if it is not 0
	if option.Command != "" {
		q = q.Where("starts_with(lower(query), lower(?))", option.Command)
	}

	// order by sort value if sort value not provided in the option order will work by default value
	q.Order(fmt.Sprintf("max_exec_time %s", option.Sort))

	// total count
	allRows := make([]*Query, 0)
	if err := q.Find(&allRows).Error; err != nil {
		return result, err
	}

	result.TotalCount = len(allRows)

	// apply limit and offset option
	q = q.Limit(option.Limit).Offset(option.Offset)

	if err := q.Find(&rows).Error; err != nil {
		return result, err
	}

	result.Rows = rows

	return result, nil
}
