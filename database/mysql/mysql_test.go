package mysql

import (
	"testing"

	"github.com/chanhteam/go-utils/logger"

	"github.com/stretchr/testify/assert"
)

func TestGetConnectionShouldReturnNil(t *testing.T) {
	logger.NewDefault()
	db := GetConnection("127.0.0.1", "test", "root", "", 10, 10)
	assert.Nil(t, db)
}
