package mysql

import (
	"fmt"
	"sync"
	"time"

	"github.com/chanhlab/go-utils/logger"

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
func Init(host string, port int, database string, username string, password string, maxIdleConnection int, maxOpenConnection int) {
	strConnect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	logger := gormzap.New(logger.Log)
	logger.SetAsDefault()
	db, _ := gorm.Open(mysql.Open(strConnect), &gorm.Config{Logger: logger})

	if db == nil {
		return
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetConnMaxLifetime(time.Hour)
		sqlDB.SetMaxOpenConns(maxOpenConnection)
		sqlDB.SetMaxIdleConns(maxIdleConnection)
	}
	connect = db
}

// GetConnection gets a MySQL connection
func GetConnection(host string, port int, database string, username string, password string, maxIdleConnection int, maxOpenConnection int) *gorm.DB {
	onceInit.Do(func() {
		Init(host, port, database, username, password, maxIdleConnection, maxOpenConnection)
	})
	return connect
}
