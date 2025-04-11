package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// CustomerHandler handles HTTP requests for customers
type CustomerHandler struct {
	customerFirebase *models.CustomerFirebase
}

// NewCustomerHandler creates a new CustomerHandler instance
func NewCustomerHandler(customerFirebase *models.CustomerFirebase) *CustomerHandler {
	return &CustomerHandler{
		customerFirebase: customerFirebase,
	}
}

// List handles GET requests to list all customers
func (h *CustomerHandler) List(c *gin.Context) {
	customers, err := h.customerFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

// Get handles GET requests to get a specific customer
func (h *CustomerHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	customer, err := h.customerFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Create handles POST requests to create a new customer
func (h *CustomerHandler) Create(c *gin.Context) {
	var customer models.FirebaseCustomer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.customerFirebase.Create(c.Request.Context(), &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	customer.ID = id
	c.JSON(http.StatusCreated, customer)
}

// Update handles PUT requests to update a customer
func (h *CustomerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	var customer models.FirebaseCustomer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer.ID = id
	if err := h.customerFirebase.Update(c.Request.Context(), id, &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

// Delete handles DELETE requests to delete a customer
func (h *CustomerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer ID is required"})
		return
	}

	if err := h.customerFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
