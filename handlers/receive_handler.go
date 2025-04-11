package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type ReceiveHandler struct {
	receiveFirebase *models.ReceiveFirebase
}

func NewReceiveHandler(receiveFirebase *models.ReceiveFirebase) *ReceiveHandler {
	return &ReceiveHandler{
		receiveFirebase: receiveFirebase,
	}
}

func (h *ReceiveHandler) List(c *gin.Context) {
	receives, err := h.receiveFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receives)
}

func (h *ReceiveHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receive ID is required"})
		return
	}

	receive, err := h.receiveFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receive)
}

func (h *ReceiveHandler) Create(c *gin.Context) {
	var receive models.FirebaseReceive
	if err := c.ShouldBindJSON(&receive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.receiveFirebase.Create(c.Request.Context(), &receive)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	receive.ID = id
	c.JSON(http.StatusCreated, receive)
}

func (h *ReceiveHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receive ID is required"})
		return
	}

	var receive models.FirebaseReceive
	if err := c.ShouldBindJSON(&receive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receive.ID = id
	if err := h.receiveFirebase.Update(c.Request.Context(), id, &receive); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receive)
}

func (h *ReceiveHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receive ID is required"})
		return
	}

	if err := h.receiveFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Receive deleted successfully"})
}
