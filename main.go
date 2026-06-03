package main

import (
	"crud/database"
	"crud/router"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	router.InitalizeRouter(r)
	r.Run()
}
