package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type PurchaseReturnHandler struct {
	purchaseReturnFirebase *models.PurchaseReturnFirebase
}

func NewPurchaseReturnHandler(purchaseReturnFirebase *models.PurchaseReturnFirebase) *PurchaseReturnHandler {
	return &PurchaseReturnHandler{
		purchaseReturnFirebase: purchaseReturnFirebase,
	}
}

func (h *PurchaseReturnHandler) List(c *gin.Context) {
	purchaseReturns, err := h.purchaseReturnFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchaseReturns)
}

func (h *PurchaseReturnHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Purchase Return ID is required"})
		return
	}

	purchaseReturn, err := h.purchaseReturnFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, purchaseReturn)
}

func (h *PurchaseReturnHandler) Create(c *gin.Context) {
	var purchaseReturn models.FirebasePurchaseReturn
	if err := c.ShouldBindJSON(&purchaseReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.purchaseReturnFirebase.Create(c.Request.Context(), &purchaseReturn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	purchaseReturn.ID = id
	c.JSON(http.StatusCreated, purchaseReturn)
}

func (h *PurchaseReturnHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Purchase Return ID is required"})
		return
	}

	var purchaseReturn models.FirebasePurchaseReturn
	if err := c.ShouldBindJSON(&purchaseReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	purchaseReturn.ID = id
	if err := h.purchaseReturnFirebase.Update(c.Request.Context(), id, &purchaseReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, purchaseReturn)
}

func (h *PurchaseReturnHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Purchase Return ID is required"})
		return
	}

	if err := h.purchaseReturnFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Purchase Return deleted successfully"})
}
