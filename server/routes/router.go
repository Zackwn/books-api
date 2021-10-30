package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api")
	{
		users := main.Group("users")
		{
			users.POST("/create", controllers.CreateUser)
		}
		books := main.Group("books")
		{
			books.GET("/:id", controllers.ShowBook)
			books.DELETE("/:id", controllers.DeleteBook)
			books.POST("/create", controllers.CreateBook)
			books.PUT("/edit", controllers.EditBook)
			books.GET("/list", controllers.ListBooks)
		}
	}
	return router
}
