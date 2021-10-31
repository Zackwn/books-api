package migrations

import (
	"log"

	"github.com/zackwn/books-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(models.Book{}, models.User{})
	if err != nil {
		log.Fatal(err)
	}
}
