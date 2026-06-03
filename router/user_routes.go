package router

import (
	"crud/handlers"

	"github.com/gin-gonic/gin"
)

// initializerUserRouter initializes the user routes for the application
func initializerUserRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	users := v1.Group("/users")
	{
		users.GET("/", handlers.HandlerGetUser)
		users.POST("/", handlers.HandlerCreateUser)
		users.PUT("/:id", handlers.HandlerUpdateUser)
		users.DELETE("/:id", handlers.HandlerDeleteUser)
	}
}
