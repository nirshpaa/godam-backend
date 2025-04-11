package types

import "time"

// SalesOrderDetail represents a single item in a sales order
type SalesOrderDetail struct {
	ProductID   string  `json:"product_id"`
	ProductCode string  `json:"product_code"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
	Discount    float64 `json:"discount"`
}

// SalesOrder represents a sales order
type SalesOrder struct {
	ID                string             `json:"id"`
	Code              string             `json:"code"`
	Date              string             `json:"date"`
	CustomerID        string             `json:"customer_id"`
	SalesmanID        string             `json:"salesman_id"`
	TotalAmount       float64            `json:"total_amount"`
	Discount          float64            `json:"discount"`
	AdditionalDisc    float64            `json:"additional_disc"`
	Status            string             `json:"status"`
	SalesOrderDetails []SalesOrderDetail `json:"sales_order_details"`
}

// ProductTransaction represents a transaction involving a product
type ProductTransaction struct {
	ID            string    `json:"id"`
	Type          string    `json:"type"`
	ProductName   string    `json:"product_name"`
	ProductCode   string    `json:"product_code"`
	Quantity      float64   `json:"quantity"`
	PreviousStock float64   `json:"previous_stock"`
	NewStock      float64   `json:"new_stock"`
	Timestamp     time.Time `json:"timestamp"`
	Notes         string    `json:"notes"`
}
