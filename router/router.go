package router

import (
	"github.com/gin-gonic/gin"
)

func InitalizeRouter(router *gin.Engine) {
	initializerUserRouter(router)
}
