package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// AccessHandler handles HTTP requests for access
type AccessHandler struct {
	accessFirebase *models.AccessFirebase
}

// NewAccessHandler creates a new AccessHandler instance
func NewAccessHandler(accessFirebase *models.AccessFirebase) *AccessHandler {
	return &AccessHandler{
		accessFirebase: accessFirebase,
	}
}

// List handles GET requests to list all access
func (h *AccessHandler) List(c *gin.Context) {
	accesses, err := h.accessFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accesses)
}

// Get handles GET requests to get a specific access
func (h *AccessHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access ID is required"})
		return
	}

	access, err := h.accessFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, access)
}

// Create handles POST requests to create a new access
func (h *AccessHandler) Create(c *gin.Context) {
	var access models.AccessFirebaseModel
	if err := c.ShouldBindJSON(&access); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.accessFirebase.Create(c.Request.Context(), &access)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	access.ID = id
	c.JSON(http.StatusCreated, access)
}

// Update handles PUT requests to update a access
func (h *AccessHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access ID is required"})
		return
	}

	var access models.AccessFirebaseModel
	if err := c.ShouldBindJSON(&access); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access.ID = id
	if err := h.accessFirebase.Update(c.Request.Context(), id, &access); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, access)
}

// Delete handles DELETE requests to delete a access
func (h *AccessHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access ID is required"})
		return
	}

	if err := h.accessFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Access deleted successfully"})
}
