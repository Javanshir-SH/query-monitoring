package config_test

import (
	"github.com/Javanshir-SH/query-monitoring/internal/pkg/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTodoHandler_Create(t *testing.T) {

	tests := []struct {
		name          string
		key           string
		value         string
		expectedValue string
	}{
		{
			name:          "get from os env",
			key:           "DB_HOST",
			value:         "test",
			expectedValue: "test",
		},
		{
			name:          "default value if no value",
			key:           "DB_HOST",
			expectedValue: "localhost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv(tt.key, tt.value)
			if err != nil {
				t.Fatal(err.Error())
			}
			cnf := config.NewConfig()
			assert.Equal(t, tt.expectedValue, cnf.PgDB.Host)
		})
	}
}
