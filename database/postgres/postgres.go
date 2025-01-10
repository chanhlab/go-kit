package postgres

import (
	"fmt"
	"time"

	"github.com/chanhlab/go-kit/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormzap "moul.io/zapgorm2"
)

// NewConnection create new postgres db connection
func NewConnection(
	host string, port int, database string,
	username string, password string,
	maxIdleConnection int, maxOpenConnection int, connMaxLifetime time.Duration,
) (*gorm.DB, error) {
	connectionStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, database,
	)

	config := &gorm.Config{}
	if logger.Log != nil {
		logger := gormzap.New(logger.Log)
		logger.SetAsDefault()
		config.Logger = logger
	}

	db, err := gorm.Open(postgres.Open(connectionStr), config)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(connMaxLifetime)
	sqlDB.SetMaxOpenConns(maxOpenConnection)
	sqlDB.SetMaxIdleConns(maxIdleConnection)
	return db, nil
}
