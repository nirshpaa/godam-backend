package controllers

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related operations
type AuthController struct {
	authClient *auth.Client
}

// NewAuthController creates a new auth controller
func NewAuthController() *AuthController {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		panic(err)
	}

	return &AuthController{
		authClient: authClient,
	}
}

// Login handles user login
func (a *AuthController) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Sign in with Firebase
	user, err := a.authClient.SignInWithEmailAndPassword(context.Background(), credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate custom token
	token, err := a.authClient.CustomToken(context.Background(), user.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// CheckEmail checks if an email is available
func (a *AuthController) CheckEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	// Check if email exists in Firebase
	_, err := a.authClient.GetUserByEmail(context.Background(), email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"available": true})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available": false})
}

// CheckUsername checks if a username is available
func (a *AuthController) CheckUsername(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
		return
	}

	// TODO: Implement username check against Firebase
	// This would require storing usernames in a separate collection
	c.JSON(http.StatusOK, gin.H{"available": true})
}
