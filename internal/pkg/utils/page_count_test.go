package utils_test

import (
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoHandler_Create(t *testing.T) {

	tests := []struct {
		name              string
		totalCount        int
		perPage           int
		expectedPageCount int
	}{
		{
			name:              "last page have last page not full",
			totalCount:        100,
			perPage:           30,
			expectedPageCount: 4,
		},
		{
			name:              "invalid total count and per page",
			totalCount:        100,
			perPage:           30,
			expectedPageCount: 4,
		},
		{
			name:              "perPage 0",
			totalCount:        100,
			perPage:           0,
			expectedPageCount: 0,
		},
		{
			name:              "totalCount 0",
			totalCount:        0,
			perPage:           7,
			expectedPageCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pageCount := utils.PageCount(tt.totalCount, tt.perPage)
			assert.Equal(t, tt.expectedPageCount, pageCount)
		})
	}
}
