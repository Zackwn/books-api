package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/database"
	"github.com/zackwn/books-api/models"
	"github.com/zackwn/books-api/services"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok := user.Validate()
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user",
		})
		return
	}

	db := database.Get()

	err = db.Where("email = ?", user.Email).First(&models.User{}).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already in use",
		})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password, err = services.Hash(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// hide password
	user.Password = ""

	c.JSON(http.StatusOK, user)
}
