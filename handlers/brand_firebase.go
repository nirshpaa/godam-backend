package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/payloads/response"
)

// BrandHandler handles brand-related HTTP requests
type BrandHandler struct {
	model *models.BrandFirebase
}

// NewBrandHandler creates a new instance of BrandHandler
func NewBrandHandler(model *models.BrandFirebase) *BrandHandler {
	return &BrandHandler{
		model: model,
	}
}

// List returns all brands
func (h *BrandHandler) List(c *gin.Context) {
	brands, err := h.model.List(c.Request.Context())
	if err != nil {
		log.Printf("Error getting brands: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get brands"})
		return
	}

	var brandResponses []response.BrandResponse
	for _, brand := range brands {
		var res response.BrandResponse
		res.Transform(&brand)
		brandResponses = append(brandResponses, res)
	}

	c.JSON(http.StatusOK, brandResponses)
}

// Get returns a brand by ID
func (h *BrandHandler) Get(c *gin.Context) {
	id := c.Param("id")
	brand, err := h.model.Get(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error getting brand: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get brand"})
		return
	}

	var res response.BrandResponse
	res.Transform(brand)
	c.JSON(http.StatusOK, res)
}

// Create creates a new brand
func (h *BrandHandler) Create(c *gin.Context) {
	var brand models.BrandFirebaseModel
	if err := c.ShouldBindJSON(&brand); err != nil {
		log.Printf("Error binding brand: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand data"})
		return
	}

	if err := h.model.Create(c.Request.Context(), &brand); err != nil {
		log.Printf("Error creating brand: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create brand"})
		return
	}

	var res response.BrandResponse
	res.Transform(&brand)
	c.JSON(http.StatusCreated, res)
}

// Update updates an existing brand
func (h *BrandHandler) Update(c *gin.Context) {
	id := c.Param("id")
	brand, err := h.model.Get(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error getting brand: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get brand"})
		return
	}

	if err := c.ShouldBindJSON(brand); err != nil {
		log.Printf("Error binding brand: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid brand data"})
		return
	}

	if err := h.model.Update(c.Request.Context(), brand); err != nil {
		log.Printf("Error updating brand: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update brand"})
		return
	}

	var res response.BrandResponse
	res.Transform(brand)
	c.JSON(http.StatusOK, res)
}

// Delete deletes a brand
func (h *BrandHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.model.Delete(c.Request.Context(), id); err != nil {
		log.Printf("Error deleting brand: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete brand"})
		return
	}

	c.Status(http.StatusNoContent)
}
