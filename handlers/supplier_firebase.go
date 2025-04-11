package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/payloads/request"
	"github.com/nirshpaa/godam-backend/payloads/response"
)

// SupplierHandler represents the supplier handler
type SupplierHandler struct {
	supplier *models.SupplierFirebase
}

// NewSupplierHandler creates a new instance of SupplierHandler
func NewSupplierHandler(supplier *models.SupplierFirebase) *SupplierHandler {
	return &SupplierHandler{
		supplier: supplier,
	}
}

// List returns all suppliers
func (h *SupplierHandler) List(c *gin.Context) {
	suppliers, err := h.supplier.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res []response.SupplierResponse
	for _, supplier := range suppliers {
		var r response.SupplierResponse
		r.Transform(&supplier)
		res = append(res, r)
	}

	c.JSON(http.StatusOK, res)
}

// Get returns a supplier by ID
func (h *SupplierHandler) Get(c *gin.Context) {
	id := c.Param("id")
	supplier, err := h.supplier.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res response.SupplierResponse
	res.Transform(supplier)
	c.JSON(http.StatusOK, res)
}

// Create creates a new supplier
func (h *SupplierHandler) Create(c *gin.Context) {
	var req request.NewSupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := req.Transform()
	if err := h.supplier.Create(c.Request.Context(), &supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res response.SupplierResponse
	res.Transform(&supplier)
	c.JSON(http.StatusCreated, res)
}

// Update updates an existing supplier
func (h *SupplierHandler) Update(c *gin.Context) {
	var req request.SupplierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier, err := h.supplier.Get(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	supplier = req.Transform(supplier)
	if err := h.supplier.Update(c.Request.Context(), supplier); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var res response.SupplierResponse
	res.Transform(supplier)
	c.JSON(http.StatusOK, res)
}

// Delete deletes a supplier by ID
func (h *SupplierHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.supplier.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
