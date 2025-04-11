package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/payloads/response"
)

// ProductCategoryHandler handles product category-related HTTP requests
type ProductCategoryHandler struct {
	model *models.ProductCategoryFirebase
}

// NewProductCategoryHandler creates a new instance of ProductCategoryHandler
func NewProductCategoryHandler(model *models.ProductCategoryFirebase) *ProductCategoryHandler {
	return &ProductCategoryHandler{
		model: model,
	}
}

// List returns all product categories
func (h *ProductCategoryHandler) List(c *gin.Context) {
	categories, err := h.model.List(c.Request.Context())
	if err != nil {
		log.Printf("Error getting product categories: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product categories"})
		return
	}

	var categoryResponses []response.ProductCategoryResponse
	for _, category := range categories {
		var res response.ProductCategoryResponse
		res.Transform(&category)
		categoryResponses = append(categoryResponses, res)
	}

	c.JSON(http.StatusOK, categoryResponses)
}

// Get returns a product category by ID
func (h *ProductCategoryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	category, err := h.model.Get(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error getting product category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product category"})
		return
	}

	var res response.ProductCategoryResponse
	res.Transform(category)
	c.JSON(http.StatusOK, res)
}

// Create creates a new product category
func (h *ProductCategoryHandler) Create(c *gin.Context) {
	var category models.ProductCategoryFirebaseModel
	if err := c.ShouldBindJSON(&category); err != nil {
		log.Printf("Error binding product category: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category data"})
		return
	}

	if err := h.model.Create(c.Request.Context(), &category); err != nil {
		log.Printf("Error creating product category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product category"})
		return
	}

	var res response.ProductCategoryResponse
	res.Transform(&category)
	c.JSON(http.StatusCreated, res)
}

// Update updates an existing product category
func (h *ProductCategoryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	category, err := h.model.Get(c.Request.Context(), id)
	if err != nil {
		log.Printf("Error getting product category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product category"})
		return
	}

	if err := c.ShouldBindJSON(category); err != nil {
		log.Printf("Error binding product category: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product category data"})
		return
	}

	if err := h.model.Update(c.Request.Context(), category); err != nil {
		log.Printf("Error updating product category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product category"})
		return
	}

	var res response.ProductCategoryResponse
	res.Transform(category)
	c.JSON(http.StatusOK, res)
}

// Delete deletes a product category
func (h *ProductCategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.model.Delete(c.Request.Context(), id); err != nil {
		log.Printf("Error deleting product category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product category"})
		return
	}

	c.Status(http.StatusNoContent)
}
