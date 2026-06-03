package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HandlerGetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get User",
	})
}

func HandlerCreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Create User",
		"user":    user,
	})
}

func HandlerUpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Update User",
	})
}

func HandlerDeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Delete User",
	})
}
