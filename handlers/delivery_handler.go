package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type DeliveryHandler struct {
	deliveryFirebase *models.DeliveryFirebase
}

func NewDeliveryHandler(deliveryFirebase *models.DeliveryFirebase) *DeliveryHandler {
	return &DeliveryHandler{
		deliveryFirebase: deliveryFirebase,
	}
}

func (h *DeliveryHandler) List(c *gin.Context) {
	deliveries, err := h.deliveryFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, deliveries)
}

func (h *DeliveryHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery ID is required"})
		return
	}

	delivery, err := h.deliveryFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, delivery)
}

func (h *DeliveryHandler) Create(c *gin.Context) {
	var delivery models.FirebaseDelivery
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.deliveryFirebase.Create(c.Request.Context(), &delivery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	delivery.ID = id
	c.JSON(http.StatusCreated, delivery)
}

func (h *DeliveryHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery ID is required"})
		return
	}

	var delivery models.FirebaseDelivery
	if err := c.ShouldBindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	delivery.ID = id
	if err := h.deliveryFirebase.Update(c.Request.Context(), id, &delivery); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, delivery)
}

func (h *DeliveryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Delivery ID is required"})
		return
	}

	if err := h.deliveryFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery deleted successfully"})
}
