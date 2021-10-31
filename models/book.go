package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	MediumPrice float32        `json:"medium_price" gorm:"not null"`
	Author      string         `json:"author" gorm:"not null"`
	ImageURL    string         `json:"image_url" gorm:"not null"`
}

func (book *Book) Validate() bool {
	return !(book.Name == "" ||
		book.Author == "" ||
		book.Description == "" ||
		book.MediumPrice == 0 ||
		book.ImageURL == "")
}
