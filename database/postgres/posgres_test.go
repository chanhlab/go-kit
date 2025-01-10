package postgres

import (
	"testing"
	"time"

	"github.com/chanhlab/go-kit/logger"
	"github.com/stretchr/testify/assert"
)

func TestGetConnectionShouldReturnNil(t *testing.T) {
	logger.NewDefault()
	db, err := NewConnection("localhost", 5432, "db_service", "go_service", "go_service", 10, 10, time.Hour)
	assert.Nil(t, db)
	assert.NotNil(t, err)
}
