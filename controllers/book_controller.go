package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/database"
	"github.com/zackwn/books-api/models"
)

func ShowBook(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}

	db := database.Get()
	var book models.Book
	err = db.First(&book, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	// err := json.NewDecoder(c.Request.Body).Decode(&book)
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok := book.Validate()
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid book",
		})
		return
	}

	db := database.Get()
	err = db.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}

func ListBooks(c *gin.Context) {
	db := database.Get()
	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func EditBook(c *gin.Context) {
	var bookEdited models.Book

	err := c.ShouldBindJSON(&bookEdited)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok := bookEdited.Validate()
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid book",
		})
		return
	}

	db := database.Get()
	err = db.Save(&bookEdited).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, bookEdited)
}

func DeleteBook(c *gin.Context) {
	strID := c.Param("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := database.Get()
	err = db.Delete(&models.Book{}, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
