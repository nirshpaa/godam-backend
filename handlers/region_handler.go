package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// RegionHandler handles HTTP requests for regions
type RegionHandler struct {
	regionFirebase *models.RegionFirebase
}

// NewRegionHandler creates a new RegionHandler instance
func NewRegionHandler(regionFirebase *models.RegionFirebase) *RegionHandler {
	return &RegionHandler{
		regionFirebase: regionFirebase,
	}
}

// List handles GET requests to list all regions
func (h *RegionHandler) List(c *gin.Context) {
	regions, err := h.regionFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, regions)
}

// Get handles GET requests to get a specific region
func (h *RegionHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region ID is required"})
		return
	}

	region, err := h.regionFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, region)
}

// Create handles POST requests to create a new region
func (h *RegionHandler) Create(c *gin.Context) {
	var region models.RegionFirebaseModel
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.regionFirebase.Create(c.Request.Context(), &region)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	region.ID = id
	c.JSON(http.StatusCreated, region)
}

// Update handles PUT requests to update a region
func (h *RegionHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region ID is required"})
		return
	}

	var region models.RegionFirebaseModel
	if err := c.ShouldBindJSON(&region); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	region.ID = id
	if err := h.regionFirebase.Update(c.Request.Context(), id, &region); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, region)
}

// Delete handles DELETE requests to delete a region
func (h *RegionHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Region ID is required"})
		return
	}

	if err := h.regionFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Region deleted successfully"})
}
