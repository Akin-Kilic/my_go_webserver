package db

import (
	"fmt"
	"my_go_webserver/pkg/config"
	"my_go_webserver/pkg/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func GetConnection(db config.Database) {
	if DBClient != nil {
		return
	}
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", db.Host, db.Port, db.User, db.Password, db.Name)
	DBClient, err = gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		))
	if err != nil {
		panic(err)
	}
	sqlDB, err := DBClient.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	if err != nil {
		panic(err)
	}

	err = DBClient.AutoMigrate(&models.Post{})
	if err != nil {
		panic(err)
	}
}

func Client() *gorm.DB {
	return DBClient
}
