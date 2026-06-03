package main

import (
	"crud/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitalizeRouter(r)
	r.Run()
}
