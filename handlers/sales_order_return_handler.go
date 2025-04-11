package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type SalesOrderReturnHandler struct {
	salesOrderReturnFirebase *models.SalesOrderReturnFirebase
}

func NewSalesOrderReturnHandler(salesOrderReturnFirebase *models.SalesOrderReturnFirebase) *SalesOrderReturnHandler {
	return &SalesOrderReturnHandler{
		salesOrderReturnFirebase: salesOrderReturnFirebase,
	}
}

func (h *SalesOrderReturnHandler) List(c *gin.Context) {
	salesOrderReturns, err := h.salesOrderReturnFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salesOrderReturns)
}

func (h *SalesOrderReturnHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sales Order Return ID is required"})
		return
	}

	salesOrderReturn, err := h.salesOrderReturnFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salesOrderReturn)
}

func (h *SalesOrderReturnHandler) Create(c *gin.Context) {
	var salesOrderReturn models.FirebaseSalesOrderReturn
	if err := c.ShouldBindJSON(&salesOrderReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.salesOrderReturnFirebase.Create(c.Request.Context(), &salesOrderReturn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	salesOrderReturn.ID = id
	c.JSON(http.StatusCreated, salesOrderReturn)
}

func (h *SalesOrderReturnHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sales Order Return ID is required"})
		return
	}

	var salesOrderReturn models.FirebaseSalesOrderReturn
	if err := c.ShouldBindJSON(&salesOrderReturn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesOrderReturn.ID = id
	if err := h.salesOrderReturnFirebase.Update(c.Request.Context(), id, &salesOrderReturn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salesOrderReturn)
}

func (h *SalesOrderReturnHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sales Order Return ID is required"})
		return
	}

	if err := h.salesOrderReturnFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales Order Return deleted successfully"})
}
