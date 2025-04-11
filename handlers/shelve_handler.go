package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// ShelveHandler handles HTTP requests for shelves
type ShelveHandler struct {
	shelveFirebase *models.ShelveFirebase
}

// NewShelveHandler creates a new ShelveHandler instance
func NewShelveHandler(shelveFirebase *models.ShelveFirebase) *ShelveHandler {
	return &ShelveHandler{
		shelveFirebase: shelveFirebase,
	}
}

// List handles GET requests to list all shelves
func (h *ShelveHandler) List(c *gin.Context) {
	shelves, err := h.shelveFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelves)
}

// Get handles GET requests to get a specific shelve
func (h *ShelveHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Shelve ID is required"})
		return
	}

	shelve, err := h.shelveFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelve)
}

// Create handles POST requests to create a new shelve
func (h *ShelveHandler) Create(c *gin.Context) {
	var shelve models.ShelveFirebaseModel
	if err := c.ShouldBindJSON(&shelve); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.shelveFirebase.Create(c.Request.Context(), &shelve)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	shelve.ID = id
	c.JSON(http.StatusCreated, shelve)
}

// Update handles PUT requests to update a shelve
func (h *ShelveHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Shelve ID is required"})
		return
	}

	var shelve models.ShelveFirebaseModel
	if err := c.ShouldBindJSON(&shelve); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shelve.ID = id
	if err := h.shelveFirebase.Update(c.Request.Context(), id, &shelve); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shelve)
}

// Delete handles DELETE requests to delete a shelve
func (h *ShelveHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Shelve ID is required"})
		return
	}

	if err := h.shelveFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shelve deleted successfully"})
}
