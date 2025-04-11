package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// CompanyHandler handles HTTP requests for companies
type CompanyHandler struct {
	companyFirebase *models.CompanyFirebase
}

// NewCompanyHandler creates a new CompanyHandler instance
func NewCompanyHandler(companyFirebase *models.CompanyFirebase) *CompanyHandler {
	return &CompanyHandler{
		companyFirebase: companyFirebase,
	}
}

// List handles GET requests to list all companies
func (h *CompanyHandler) List(c *gin.Context) {
	companies, err := h.companyFirebase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}

// Get handles GET requests to get a specific company
func (h *CompanyHandler) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	company, err := h.companyFirebase.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

// Create handles POST requests to create a new company
func (h *CompanyHandler) Create(c *gin.Context) {
	var company models.FirebaseCompany
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.companyFirebase.Create(c.Request.Context(), &company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	company.ID = id
	c.JSON(http.StatusCreated, company)
}

// Update handles PUT requests to update a company
func (h *CompanyHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	var company models.FirebaseCompany
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company.ID = id
	if err := h.companyFirebase.Update(c.Request.Context(), id, &company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}

// Delete handles DELETE requests to delete a company
func (h *CompanyHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	if err := h.companyFirebase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
