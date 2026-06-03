package router

import (
	"github.com/gin-gonic/gin"
)

// InitalizeRouter initializes all the routes for the application
func InitalizeRouter(router *gin.Engine) {
	initializerUserRouter(router)
}
