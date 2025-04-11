package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/services"
)

type UserHandler struct {
	firebaseService *services.FirebaseService
	userFirebase    *models.UserFirebase
}

func NewUserHandler(firebaseService *services.FirebaseService, userFirebase *models.UserFirebase) *UserHandler {
	return &UserHandler{
		firebaseService: firebaseService,
		userFirebase:    userFirebase,
	}
}

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

// ListUsers handles GET requests to list all users
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser handles GET requests to get a specific user
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := h.userFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.firebaseService.CreateUser(c.Request.Context(), req.Email, req.Password, req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser handles PUT requests to update a user
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	var user models.FirebaseUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = id
	if err := h.userFirebase.Update(c.Request.Context(), id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// Delete from Firebase Auth and Firestore
	if err := h.firebaseService.DeleteUser(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user from Firebase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) JoinCompany(c *gin.Context) {
	companyID := c.Param("id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	// Get user ID from context (assuming it's set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Update user's company ID in Firestore
	if err := h.firebaseService.UpdateUserCompany(c.Request.Context(), userID.(string), companyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to join company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully joined company"})
}
