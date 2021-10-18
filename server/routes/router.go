package routes

import (
	"books-api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")

	{
		session := main.Group("/session")
		{
			session.POST("/", controllers.SendLogin)
		}
		users := main.Group("/users")
		{
			users.POST("/", controllers.CreateUser)
		}
		books := main.Group("/books")
		{
			books.GET("/", controllers.FindBooks)
			books.GET("/:id", controllers.FindBookByID)
			books.POST("/", controllers.CreateBook)
			books.PUT("/:id", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
