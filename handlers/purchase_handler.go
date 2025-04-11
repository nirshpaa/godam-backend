package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// PurchaseHandler handles HTTP requests for purchases
type PurchaseHandler struct {
	purchaseFirebase *models.PurchaseFirebase
}

// NewPurchaseHandler creates a new PurchaseHandler instance
func NewPurchaseHandler(purchaseFirebase *models.PurchaseFirebase) *PurchaseHandler {
	return &PurchaseHandler{
		purchaseFirebase: purchaseFirebase,
	}
}

// List handles GET requests to list all purchases
func (h *PurchaseHandler) List(c *gin.Context) {
	purchases, err := h.purchaseFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchases)
}

// Get handles GET requests to get a specific purchase
func (h *PurchaseHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Purchase ID is required"})
		return
	}

	purchase, err := h.purchaseFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchase)
}

// Create handles POST requests to create a new purchase
func (h *PurchaseHandler) Create(c *gin.Context) {
	var purchase models.FirebasePurchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.purchaseFirebase.Create(c.Request.Context(), &purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	purchase.ID = id
	c.JSON(http.StatusCreated, purchase)
}

// Update handles PUT requests to update a purchase
func (h *PurchaseHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Purchase ID is required"})
		return
	}

	var purchase models.FirebasePurchase
	if err := c.ShouldBindJSON(&purchase); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchase.ID = id
	if err := h.purchaseFirebase.Update(c.Request.Context(), id, &purchase); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchase)
}

// Delete handles DELETE requests to delete a purchase
func (h *PurchaseHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Purchase ID is required"})
		return
	}

	if err := h.purchaseFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase deleted successfully"})
}
