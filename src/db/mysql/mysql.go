package mysql

import (
	"context"
	"fmt"
	"github.com/HuckOps/notify/pkg/restful"
	"github.com/HuckOps/notify/src/config"
	mysqlModel "github.com/HuckOps/notify/src/model/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var MySQL *sqlEngine

type sqlEngine struct {
	context context.Context
	db      *gorm.DB
}

func init() {
	MySQL = &sqlEngine{context: context.Background()}
}

func (s *sqlEngine) Load() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DB.MySQL.User, config.Config.DB.MySQL.Password, config.Config.DB.MySQL.Host, config.Config.DB.MySQL.Port, config.Config.DB.MySQL.DB)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        512,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&mysqlModel.User{}, &mysqlModel.Active{})
	sqldb, _ := db.DB()
	sqldb.SetConnMaxLifetime(time.Hour)
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	s.db = db
}

func (s *sqlEngine) DB() *gorm.DB {
	return s.db
}

func SearchByPagination[T interface{}](db *gorm.DB, skip int, limit int, v *restful.ItemResponse[T]) *gorm.DB {
	return db.Count(&v.Total).Limit(limit).Offset(skip).Scan(&v.Items)
}
