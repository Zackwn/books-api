package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/database"
	"github.com/zackwn/books-api/models"
	"github.com/zackwn/books-api/services"
)

type loginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewLogin(jwtService services.JWT) func(c *gin.Context) {
	return func(c *gin.Context) {
		var dto loginDTO
		err := c.ShouldBindJSON(&dto)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := database.Get()

		var user models.User
		err = db.Where("email = ?", dto.Email).First(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = services.Compare(dto.Password, user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := jwtService.Sign(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
