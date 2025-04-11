package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/types"
)

type SalesOrderHandler struct {
	salesOrderFirebase *models.SalesOrderFirebase
}

func NewSalesOrderHandler(salesOrderFirebase *models.SalesOrderFirebase) *SalesOrderHandler {
	return &SalesOrderHandler{
		salesOrderFirebase: salesOrderFirebase,
	}
}

func (h *SalesOrderHandler) List(c *gin.Context) {
	salesOrders, err := h.salesOrderFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salesOrders)
}

func (h *SalesOrderHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sales Order ID is required"})
		return
	}

	salesOrder, err := h.salesOrderFirebase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, salesOrder)
}

func (h *SalesOrderHandler) Create(c *gin.Context) {
	var salesOrder types.SalesOrder
	if err := c.ShouldBindJSON(&salesOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.salesOrderFirebase.Create(c.Request.Context(), salesOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, salesOrder)
}

func (h *SalesOrderHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sales Order ID is required"})
		return
	}

	var salesOrder types.SalesOrder
	if err := c.ShouldBindJSON(&salesOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesOrder.Code = id
	if err := h.salesOrderFirebase.Update(c.Request.Context(), id, salesOrder); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, salesOrder)
}

func (h *SalesOrderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sales Order ID is required"})
		return
	}

	if err := h.salesOrderFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales Order deleted successfully"})
}

func (h *SalesOrderHandler) Stats(c *gin.Context) {
	// Get all sales orders
	orders, err := h.salesOrderFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total sales and monthly income
	var totalSales float64
	var monthlyIncome float64
	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()

	for _, order := range orders {
		// Add to total sales
		totalSales += order.TotalAmount

		// Parse the order date
		orderDate, err := time.Parse(time.RFC3339, order.Date)
		if err != nil {
			continue // Skip orders with invalid dates
		}

		// Add to monthly income if the order is from the current month
		if orderDate.Month() == currentMonth && orderDate.Year() == currentYear {
			monthlyIncome += order.TotalAmount
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"total_sales":    totalSales,
		"monthly_income": monthlyIncome,
	})
}
