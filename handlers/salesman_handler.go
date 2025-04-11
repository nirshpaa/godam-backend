package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// SalesmanHandler handles HTTP requests for salesmen
type SalesmanHandler struct {
	salesmanFirebase *models.SalesmanFirebase
}

// NewSalesmanHandler creates a new SalesmanHandler instance
func NewSalesmanHandler(salesmanFirebase *models.SalesmanFirebase) *SalesmanHandler {
	return &SalesmanHandler{
		salesmanFirebase: salesmanFirebase,
	}
}

// List handles GET requests to list all salesmen
func (h *SalesmanHandler) List(c *gin.Context) {
	salesmen, err := h.salesmanFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salesmen)
}

// Get handles GET requests to get a specific salesman
func (h *SalesmanHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Salesman ID is required"})
		return
	}

	salesman, err := h.salesmanFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salesman)
}

// Create handles POST requests to create a new salesman
func (h *SalesmanHandler) Create(c *gin.Context) {
	var salesman models.SalesmanFirebaseModel
	if err := c.ShouldBindJSON(&salesman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.salesmanFirebase.Create(c.Request.Context(), &salesman)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	salesman.ID = id
	c.JSON(http.StatusCreated, salesman)
}

// Update handles PUT requests to update a salesman
func (h *SalesmanHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Salesman ID is required"})
		return
	}

	var salesman models.SalesmanFirebaseModel
	if err := c.ShouldBindJSON(&salesman); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesman.ID = id
	if err := h.salesmanFirebase.Update(c.Request.Context(), id, &salesman); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salesman)
}

// Delete handles DELETE requests to delete a salesman
func (h *SalesmanHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Salesman ID is required"})
		return
	}

	if err := h.salesmanFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Salesman deleted successfully"})
}
