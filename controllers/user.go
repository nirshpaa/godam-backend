package controllers

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
)

// UserController handles user-related operations
type UserController struct {
	dbRef *db.Ref
}

// NewUserController creates a new user controller
func NewUserController() *UserController {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	dbClient, err := app.Database(context.Background())
	if err != nil {
		panic(err)
	}

	return &UserController{
		dbRef: dbClient.NewRef("users"),
	}
}

// List retrieves all users
func (u *UserController) List(ctx *gin.Context) {
	var users []map[string]interface{}
	if err := u.dbRef.Get(context.Background(), &users); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// Get retrieves a single user by ID
func (u *UserController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var user map[string]interface{}
	if err := u.dbRef.Child(id).Get(context.Background(), &user); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// Create creates a new user
func (u *UserController) Create(ctx *gin.Context) {
	var user map[string]interface{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRef, err := u.dbRef.Push(context.Background(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user["id"] = newRef.Key
	ctx.JSON(http.StatusCreated, user)
}

// Update updates an existing user
func (u *UserController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var user map[string]interface{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.dbRef.Child(id).Set(context.Background(), user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	user["id"] = id
	ctx.JSON(http.StatusOK, user)
}

// Delete removes a user
func (u *UserController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	if err := u.dbRef.Child(id).Delete(context.Background()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
