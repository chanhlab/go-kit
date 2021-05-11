package mysql

import (
	"fmt"
	"sync"
	"time"

	"github.com/chanhteam/go-utils/logger"

	"gorm.io/gorm"

	// Register some standard stuff
	"gorm.io/driver/mysql"
	gormzap "moul.io/zapgorm2"
)

var (
	// Connect is connection global
	connect *gorm.DB

	// onceInit guarantee initialize connection only once
	onceInit sync.Once
)

// Init initialize mysql connection
func Init(host string, database string, username string, password string, maxIdleConnection int, maxOpenConnection int) {
	strConnect := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, database)
	logger := gormzap.New(logger.Log)
	logger.SetAsDefault()
	connect, err := gorm.Open(mysql.Open(strConnect), &gorm.Config{Logger: logger})
	if err != nil {
		return
	}

	sqlDB, err := connect.DB()
	if err == nil {
		sqlDB.SetConnMaxLifetime(time.Hour)
		sqlDB.SetMaxOpenConns(maxOpenConnection)
		sqlDB.SetMaxIdleConns(maxIdleConnection)
	}
}

// GetConnection gets a MySQL connection
func GetConnection(host string, database string, username string, password string, maxIdleConnection int, maxOpenConnection int) *gorm.DB {
	onceInit.Do(func() {
		Init(host, database, username, password, maxIdleConnection, maxOpenConnection)
	})
	return connect
}
