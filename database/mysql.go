package database

import (
	"e-assets/config"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type mysqlDatabase struct {
	Db *gorm.DB
}

func NewMysqlDatabase(cfg *config.Config) Database {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Pass, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect mysql.")
	}

	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic("Failed to get *sql.DB")
		}
		if err := sqlDB.Close(); err != nil {
			panic("Failed to close database.")
		}
	}()

	return &mysqlDatabase{Db: db}
}

func (mysql *mysqlDatabase) GetDb() *gorm.DB {
	return mysql.Db
}
