package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// RoleHandler handles HTTP requests for roles
type RoleHandler struct {
	roleFirebase *models.RoleFirebase
}

// NewRoleHandler creates a new RoleHandler instance
func NewRoleHandler(roleFirebase *models.RoleFirebase) *RoleHandler {
	return &RoleHandler{
		roleFirebase: roleFirebase,
	}
}

// List handles GET requests to list all roles
func (h *RoleHandler) List(c *gin.Context) {
	roles, err := h.roleFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// Get handles GET requests to get a specific role
func (h *RoleHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	role, err := h.roleFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// Create handles POST requests to create a new role
func (h *RoleHandler) Create(c *gin.Context) {
	var role models.RoleFirebaseModel
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.roleFirebase.Create(c.Request.Context(), &role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	role.ID = id
	c.JSON(http.StatusCreated, role)
}

// Update handles PUT requests to update a role
func (h *RoleHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	var role models.RoleFirebaseModel
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role.ID = id
	if err := h.roleFirebase.Update(c.Request.Context(), id, &role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// Delete handles DELETE requests to delete a role
func (h *RoleHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role ID is required"})
		return
	}

	if err := h.roleFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}
