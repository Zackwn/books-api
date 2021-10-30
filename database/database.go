package database

import (
	"log"
	"time"

	"github.com/zackwn/books-api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Start() {
	dsn := "host=localhost port=25432 user=postgres dbname=books sslmode=disable password=123456"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	config, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

func Get() *gorm.DB {
	return db
}
