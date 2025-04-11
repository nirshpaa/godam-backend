package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/services"
)

type PredictiveHandler struct {
	predictiveService *services.PredictiveService
}

func NewPredictiveHandler(predictiveService *services.PredictiveService) *PredictiveHandler {
	return &PredictiveHandler{
		predictiveService: predictiveService,
	}
}

// GetStockRecommendations handles GET /api/predictive/stock-recommendations
func (h *PredictiveHandler) GetStockRecommendations(c *gin.Context) {
	companyID := c.Param("companyId")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	recommendations, err := h.predictiveService.GetStockRecommendations(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}

// GetSalesPredictions handles GET /api/predictive/sales-predictions
func (h *PredictiveHandler) GetSalesPredictions(c *gin.Context) {
	companyID := c.Param("companyId")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	daysStr := c.DefaultQuery("days", "7")
	days, err := strconv.Atoi(daysStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid days parameter"})
		return
	}

	predictions, err := h.predictiveService.GetSalesPredictions(c.Request.Context(), companyID, days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, predictions)
}

// GetProductHistory handles GET /api/predictive/product-history/:productCode
func (h *PredictiveHandler) GetProductHistory(c *gin.Context) {
	productCode := c.Param("productCode")
	if productCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code is required"})
		return
	}

	// Parse date range parameters
	startDateStr := c.DefaultQuery("start_date", time.Now().AddDate(0, -1, 0).Format(time.RFC3339))
	endDateStr := c.DefaultQuery("end_date", time.Now().Format(time.RFC3339))

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	history, err := h.predictiveService.GetProductHistory(c.Request.Context(), productCode, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, history)
}

// GetSalesReport handles GET /api/predictive/sales-report
func (h *PredictiveHandler) GetSalesReport(c *gin.Context) {
	companyID := c.Param("companyId")
	if companyID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	// Parse date range parameters
	startDateStr := c.DefaultQuery("start_date", time.Now().AddDate(0, -1, 0).Format(time.RFC3339))
	endDateStr := c.DefaultQuery("end_date", time.Now().Format(time.RFC3339))

	startDate, err := time.Parse(time.RFC3339, startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date format"})
		return
	}

	endDate, err := time.Parse(time.RFC3339, endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date format"})
		return
	}

	report, err := h.predictiveService.GetSalesReport(c.Request.Context(), companyID, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
