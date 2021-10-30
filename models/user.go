package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password,omitempty"`
}

func (user *User) Validate() bool {
	return !(user.Name == "" ||
		user.Email == "" ||
		user.Password == "")
}
