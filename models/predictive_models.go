package models

import "time"

// StockRecommendation represents a recommendation for stock levels
type StockRecommendation struct {
	ProductCode       string    `json:"product_code"`
	ProductName       string    `json:"product_name"`
	CurrentStock      float64   `json:"current_stock"`
	AverageDailySales float64   `json:"average_daily_sales"`
	SafetyStock       float64   `json:"safety_stock"`
	ReorderPoint      float64   `json:"reorder_point"`
	RecommendedOrder  float64   `json:"recommended_order"`
	LastUpdated       time.Time `json:"last_updated"`
}

// SalesPrediction represents predicted sales for a product
type SalesPrediction struct {
	ProductCode      string            `json:"product_code"`
	ProductName      string            `json:"product_name"`
	CurrentStock     float64           `json:"current_stock"`
	DailyPredictions []DailyPrediction `json:"daily_predictions"`
}

// DailyPrediction represents predicted sales for a specific day
type DailyPrediction struct {
	Date            time.Time `json:"date"`
	PredictedSales  float64   `json:"predicted_sales"`
	ConfidenceScore float64   `json:"confidence_score"`
}

// ProductHistory represents the history of a product's stock changes
type ProductHistory struct {
	ProductCode     string    `json:"product_code"`
	ProductName     string    `json:"product_name"`
	TransactionID   string    `json:"transaction_id"`
	TransactionType string    `json:"transaction_type"` // "purchase", "sale", "return", etc.
	Quantity        float64   `json:"quantity"`
	PreviousStock   float64   `json:"previous_stock"`
	NewStock        float64   `json:"new_stock"`
	Timestamp       time.Time `json:"timestamp"`
	Notes           string    `json:"notes"`
}

// SalesReport represents a sales report for a specific period
type SalesReport struct {
	CompanyID      string          `json:"company_id"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	TotalSales     float64         `json:"total_sales"`
	TotalQuantity  float64         `json:"total_quantity"`
	ProductDetails []ProductDetail `json:"product_details"`
}

// ProductDetail represents detailed sales information for a product
type ProductDetail struct {
	ProductCode  string  `json:"product_code"`
	ProductName  string  `json:"product_name"`
	QuantitySold float64 `json:"quantity_sold"`
	TotalRevenue float64 `json:"total_revenue"`
	AveragePrice float64 `json:"average_price"`
	ProfitMargin float64 `json:"profit_margin"`
}
