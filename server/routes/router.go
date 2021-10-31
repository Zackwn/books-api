package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/controllers"
	"github.com/zackwn/books-api/services"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	jwt := services.NewJWT()

	main := router.Group("api")
	{
		users := main.Group("users")
		{
			users.POST("/create", controllers.CreateUser)
			users.POST("/login", controllers.NewLogin(jwt))
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
