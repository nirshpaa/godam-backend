package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type ReceiveReturnHandler struct {
	receiveReturnFirebase *models.ReceiveReturnFirebase
}

func NewReceiveReturnHandler(receiveReturnFirebase *models.ReceiveReturnFirebase) *ReceiveReturnHandler {
	return &ReceiveReturnHandler{
		receiveReturnFirebase: receiveReturnFirebase,
	}
}

func (h *ReceiveReturnHandler) List(c *gin.Context) {
	receiveReturns, err := h.receiveReturnFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receiveReturns)
}

func (h *ReceiveReturnHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receive Return ID is required"})
		return
	}

	receiveReturn, err := h.receiveReturnFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receiveReturn)
}

func (h *ReceiveReturnHandler) Create(c *gin.Context) {
	var receiveReturn models.FirebaseReceiveReturn
	if err := c.ShouldBindJSON(&receiveReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.receiveReturnFirebase.Create(c.Request.Context(), &receiveReturn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	receiveReturn.ID = id
	c.JSON(http.StatusCreated, receiveReturn)
}

func (h *ReceiveReturnHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receive Return ID is required"})
		return
	}

	var receiveReturn models.FirebaseReceiveReturn
	if err := c.ShouldBindJSON(&receiveReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	receiveReturn.ID = id
	if err := h.receiveReturnFirebase.Update(c.Request.Context(), id, &receiveReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, receiveReturn)
}

func (h *ReceiveReturnHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receive Return ID is required"})
		return
	}

	if err := h.receiveReturnFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Receive Return deleted successfully"})
}
