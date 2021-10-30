package migrations

import (
	"github.com/zackwn/books-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{}, models.User{})
}
