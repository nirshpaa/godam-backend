package services

import "time"

// FirebaseSalesOrderDetail represents a sales order detail in Firebase
type FirebaseSalesOrderDetail struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Disc      float64 `json:"disc"`
	Qty       uint    `json:"qty"`
}

// ProductTransaction represents a transaction involving a product
type ProductTransaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	ProductName   string    `json:"product_name"`
	Quantity      float64   `json:"quantity"`
	PreviousStock float64   `json:"previous_stock"`
	NewStock      float64   `json:"new_stock"`
	Timestamp     time.Time `json:"timestamp"`
	Notes         string    `json:"notes"`
}
