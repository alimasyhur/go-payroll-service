package driver

import (
	"fmt"

	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"

	"github.com/weanan/weanan-service/config"
)

func NewMySQLDatabase(cfg config.DBConfig) (db *gorm.DB, err error) {
	fmt.Println("Try NewDatabase ...")

	dsn := cfg.GetDSN()
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	if cfg.DebugMode {
		db = db.Debug()
	}

	return
}
