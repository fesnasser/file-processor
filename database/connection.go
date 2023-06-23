package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/fesnasser/file-processor/model"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "file-processor"
)

var (
	once = sync.Once{}
	db   *gorm.DB
)

func GetCon() *gorm.DB {
	once.Do(func() {
		var err error

		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:                 logger.Default.LogMode(logger.Silent),
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
		if err != nil {
			panic(err.Error())
		}

		db.AutoMigrate(model.Line{})

		sqlDB, err := db.DB()
		if err != nil {
			panic(err.Error())
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	})
	return db
}
