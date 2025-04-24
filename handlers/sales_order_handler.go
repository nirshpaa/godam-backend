package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/types"
)

type SalesOrderHandler struct {
	salesOrderModel *models.SalesOrderFirebase
	productModel    *models.ProductFirebase
}

func NewSalesOrderHandler(salesOrderModel *models.SalesOrderFirebase, productModel *models.ProductFirebase) *SalesOrderHandler {
	return &SalesOrderHandler{
		salesOrderModel: salesOrderModel,
		productModel:    productModel,
	}
}

func (h *SalesOrderHandler) List(c *gin.Context) {
	// Get company ID from context
	companyID := c.GetString("company_id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	// Get sales orders for the company
	salesOrders, err := h.salesOrderModel.FindByCompany(c.Request.Context(), companyID)
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

	salesOrder, err := h.salesOrderModel.GetByID(c.Request.Context(), id)
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

	// Set the company ID
	salesOrder.CompanyID = companyID

	err := h.salesOrderModel.Create(c.Request.Context(), salesOrder)
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
	if err := h.salesOrderModel.Update(c.Request.Context(), id, salesOrder); err != nil {
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

	if err := h.salesOrderModel.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales Order deleted successfully"})
}

func (h *SalesOrderHandler) Stats(c *gin.Context) {
	// Get company ID from context
	companyID := c.GetString("company_id")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	// Get all sales orders for the company
	orders, err := h.salesOrderModel.FindByCompany(c.Request.Context(), companyID)
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

// UpdateProductStock handles POST /sales-orders/update-stock
func (h *SalesOrderHandler) UpdateProductStock(c *gin.Context) {
	var request struct {
		ProductID string  `json:"productId" binding:"required"`
		Quantity  float64 `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the product
	product, err := h.productModel.Get(c.Request.Context(), request.ProductID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Check if there's enough stock
	if product.MinimumStock < request.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Update the stock
	product.MinimumStock -= request.Quantity

	// Save the updated product
	if err := h.productModel.UpdateStock(c.Request.Context(), request.ProductID, product.MinimumStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock updated successfully"})
}
