package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api")
	{
		books := main.Group("/books")
		{
			books.GET("/", nil)
		}
	}

	return router
}
