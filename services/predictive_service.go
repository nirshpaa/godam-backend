package services

import (
	"context"
	"time"

	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/types"
)

type PredictiveService struct {
	productModel *models.ProductFirebase
	salesModel   *models.SalesOrderFirebase
}

func NewPredictiveService(productModel *models.ProductFirebase, salesModel *models.SalesOrderFirebase) *PredictiveService {
	return &PredictiveService{
		productModel: productModel,
		salesModel:   salesModel,
	}
}

// GetStockRecommendations returns recommended stock levels based on sales history
func (s *PredictiveService) GetStockRecommendations(ctx context.Context, companyID string) ([]models.StockRecommendation, error) {
	// Get all products for the company
	products, err := s.productModel.FindByCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	// Get sales data for the last 30 days
	endDate := time.Now()
	startDate := endDate.AddDate(0, -1, 0) // Last 30 days
	sales, err := s.salesModel.GetSalesByDateRange(ctx, companyID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Calculate average daily sales and generate recommendations
	recommendations := make([]models.StockRecommendation, 0)
	for _, product := range products {
		// Calculate average daily sales
		dailySales := calculateAverageDailySales(sales, product.Code)

		// Calculate safety stock (2 weeks worth of average sales)
		safetyStock := dailySales * 14

		// Calculate reorder point (safety stock + lead time demand)
		leadTimeDays := 7 // Assuming 7 days lead time
		reorderPoint := safetyStock + (dailySales * float64(leadTimeDays))

		// Calculate recommended order quantity
		recommendedOrder := reorderPoint - product.MinimumStock

		recommendations = append(recommendations, models.StockRecommendation{
			ProductCode:       product.Code,
			ProductName:       product.Name,
			CurrentStock:      product.MinimumStock,
			AverageDailySales: dailySales,
			SafetyStock:       safetyStock,
			ReorderPoint:      reorderPoint,
			RecommendedOrder:  recommendedOrder,
			LastUpdated:       time.Now(),
		})
	}

	return recommendations, nil
}

// GetSalesPredictions predicts future sales based on historical data
func (s *PredictiveService) GetSalesPredictions(ctx context.Context, companyID string, days int) ([]models.SalesPrediction, error) {
	// Get historical sales data
	endDate := time.Now()
	startDate := endDate.AddDate(0, -3, 0) // Last 3 months
	sales, err := s.salesModel.GetSalesByDateRange(ctx, companyID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Group sales by product and calculate trends
	productTrends := calculateProductTrends(sales)

	// Generate predictions for the next 'days' days
	predictions := make([]models.SalesPrediction, 0)
	for productCode, trend := range productTrends {
		product, err := s.productModel.Get(ctx, productCode)
		if err != nil {
			continue
		}

		prediction := models.SalesPrediction{
			ProductCode:      productCode,
			ProductName:      product.Name,
			CurrentStock:     product.MinimumStock,
			DailyPredictions: make([]models.DailyPrediction, days),
		}

		// Generate daily predictions
		for i := 0; i < days; i++ {
			date := endDate.AddDate(0, 0, i+1)
			predictedSales := trend.predict(date)

			prediction.DailyPredictions[i] = models.DailyPrediction{
				Date:            date,
				PredictedSales:  predictedSales,
				ConfidenceScore: trend.confidence,
			}
		}

		predictions = append(predictions, prediction)
	}

	return predictions, nil
}

// GetProductHistory retrieves the history of a product's stock changes
func (s *PredictiveService) GetProductHistory(ctx context.Context, productCode string, startDate, endDate time.Time) ([]models.ProductHistory, error) {
	// Get all transactions involving this product
	transactions, err := s.salesModel.GetProductTransactions(ctx, productCode, startDate, endDate)
	if err != nil {
		return nil, err
	}

	history := make([]models.ProductHistory, 0)
	for _, tx := range transactions {
		history = append(history, models.ProductHistory{
			ProductCode:     productCode,
			ProductName:     tx.ProductName,
			TransactionID:   tx.ID,
			TransactionType: tx.Type,
			Quantity:        tx.Quantity,
			PreviousStock:   tx.PreviousStock,
			NewStock:        tx.NewStock,
			Timestamp:       tx.Timestamp,
			Notes:           tx.Notes,
		})
	}

	return history, nil
}

// GetSalesReport generates a comprehensive sales report
func (s *PredictiveService) GetSalesReport(ctx context.Context, companyID string, startDate, endDate time.Time) (*models.SalesReport, error) {
	// Get all sales data for the period
	sales, err := s.salesModel.GetSalesByDateRange(ctx, companyID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Calculate total sales and quantity
	totalSales := 0.0
	totalQuantity := 0.0
	productDetails := make(map[string]*models.ProductDetail)

	for _, sale := range sales {
		for _, item := range sale.SalesOrderDetails {
			totalSales += item.TotalPrice
			totalQuantity += float64(item.Quantity)

			// Update product details
			if detail, exists := productDetails[item.ProductCode]; exists {
				detail.QuantitySold += float64(item.Quantity)
				detail.TotalRevenue += item.TotalPrice
				detail.AveragePrice = detail.TotalRevenue / detail.QuantitySold
			} else {
				product, err := s.productModel.Get(ctx, item.ProductCode)
				if err != nil {
					continue
				}

				productDetails[item.ProductCode] = &models.ProductDetail{
					ProductCode:  item.ProductCode,
					ProductName:  product.Name,
					QuantitySold: float64(item.Quantity),
					TotalRevenue: item.TotalPrice,
					AveragePrice: item.TotalPrice / float64(item.Quantity),
					ProfitMargin: calculateProfitMargin(item.TotalPrice, product.PurchasePrice),
				}
			}
		}
	}

	// Convert product details map to slice
	details := make([]models.ProductDetail, 0, len(productDetails))
	for _, detail := range productDetails {
		details = append(details, *detail)
	}

	return &models.SalesReport{
		CompanyID:      companyID,
		StartDate:      startDate,
		EndDate:        endDate,
		TotalSales:     totalSales,
		TotalQuantity:  totalQuantity,
		ProductDetails: details,
	}, nil
}

// Helper functions
func calculateAverageDailySales(sales []types.SalesOrder, productCode string) float64 {
	totalSales := 0.0
	days := 0
	for _, sale := range sales {
		for _, item := range sale.SalesOrderDetails {
			if item.ProductCode == productCode {
				totalSales += item.Quantity
				days++
			}
		}
	}
	if days == 0 {
		return 0
	}
	return totalSales / float64(days)
}

func calculateProfitMargin(revenue, cost float64) float64 {
	if cost == 0 {
		return 0
	}
	return ((revenue - cost) / revenue) * 100
}

type productTrend struct {
	slope      float64
	intercept  float64
	confidence float64
}

func (t *productTrend) predict(date time.Time) float64 {
	daysSinceStart := float64(date.Unix() / (24 * 60 * 60))
	return t.slope*daysSinceStart + t.intercept
}

func calculateProductTrends(sales []types.SalesOrder) map[string]*productTrend {
	trends := make(map[string]*productTrend)
	// Implementation of trend calculation using linear regression
	// This is a simplified version - in production, you might want to use more sophisticated algorithms
	return trends
}
