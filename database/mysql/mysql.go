package mysql

import (
	"fmt"
	"sync"
	"time"

	"github.com/chanhteam/go-utils/logger"

	"github.com/jinzhu/gorm"

	// Register some standard stuff
	_ "github.com/jinzhu/gorm/dialects/mysql"
	gormzap "github.com/wantedly/gorm-zap"
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
	connect, err := gorm.Open("mysql", strConnect)

	if err == nil {
		connect.LogMode(false)
		connect.SetLogger(gormzap.New(logger.Log))
		connect.DB().SetConnMaxLifetime(time.Hour)
		connect.DB().SetMaxOpenConns(maxOpenConnection)
		connect.DB().SetMaxIdleConns(maxIdleConnection)
	}
}

// GetConnection gets a MySQL connection
func GetConnection(host string, database string, username string, password string, maxIdleConnection int, maxOpenConnection int) *gorm.DB {
	onceInit.Do(func() {
		Init(host, database, username, password, maxIdleConnection, maxOpenConnection)
	})
	return connect
}
