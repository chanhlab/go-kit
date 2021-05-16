package mysql

import (
	"testing"

	"github.com/chanhteam/go-utils/logger"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionShouldReturnNil(t *testing.T) {
	logger.NewDefault()
	db := GetConnection("localhost", 9910, "db_service", "go_service", "go_service", 10, 10)
	assert.NotNil(t, db)
}
