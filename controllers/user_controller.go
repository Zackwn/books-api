package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/database"
	"github.com/zackwn/books-api/models"
	"github.com/zackwn/books-api/services"
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

	user.Password, err = services.Hash(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := database.Get()
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
