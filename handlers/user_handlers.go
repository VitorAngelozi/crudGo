package handlers

import (
	"crud/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// HandlerGetUser retrieves a user from the database
func HandlerGetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Get User",
	})
}

// HandlerCreateUser creates a new user in the database
func HandlerCreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Insert the user into the database and return the created user with its ID
	query := `
			INSERT INTO users (name, email)
			VALUES ($1, $2)
			RETURNING id
		`
	// Execute the query and scan the returned ID into the user struct
	err := database.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)

}

// HandlerUpdateUser updates an existing user in the database
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
