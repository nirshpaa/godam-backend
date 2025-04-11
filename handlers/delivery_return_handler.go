package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type DeliveryReturnHandler struct {
	deliveryReturnFirebase *models.DeliveryReturnFirebase
}

func NewDeliveryReturnHandler(deliveryReturnFirebase *models.DeliveryReturnFirebase) *DeliveryReturnHandler {
	return &DeliveryReturnHandler{
		deliveryReturnFirebase: deliveryReturnFirebase,
	}
}

func (h *DeliveryReturnHandler) List(c *gin.Context) {
	deliveryReturns, err := h.deliveryReturnFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deliveryReturns)
}

func (h *DeliveryReturnHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery Return ID is required"})
		return
	}

	deliveryReturn, err := h.deliveryReturnFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deliveryReturn)
}

func (h *DeliveryReturnHandler) Create(c *gin.Context) {
	var deliveryReturn models.FirebaseDeliveryReturn
	if err := c.ShouldBindJSON(&deliveryReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.deliveryReturnFirebase.Create(c.Request.Context(), &deliveryReturn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deliveryReturn.ID = id
	c.JSON(http.StatusCreated, deliveryReturn)
}

func (h *DeliveryReturnHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery Return ID is required"})
		return
	}

	var deliveryReturn models.FirebaseDeliveryReturn
	if err := c.ShouldBindJSON(&deliveryReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deliveryReturn.ID = id
	if err := h.deliveryReturnFirebase.Update(c.Request.Context(), id, &deliveryReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deliveryReturn)
}

func (h *DeliveryReturnHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery Return ID is required"})
		return
	}

	if err := h.deliveryReturnFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery Return deleted successfully"})
}
