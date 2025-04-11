package models

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/nirshpaa/godam-backend/types"
)

// SalesOrderFirebase represents a sales order in Firebase
type SalesOrderFirebase struct {
	ID                string                   `json:"id"`
	Code              string                   `json:"code"`
	Date              string                   `json:"date"`
	CustomerID        string                   `json:"customer_id"`
	SalesmanID        string                   `json:"salesman_id"`
	TotalAmount       float64                  `json:"total_amount"`
	Discount          float64                  `json:"discount"`
	AdditionalDisc    float64                  `json:"additional_disc"`
	Status            string                   `json:"status"`
	SalesOrderDetails []types.SalesOrderDetail `json:"sales_order_details"`
	client            *firestore.Client
}

// NewSalesOrderFirebase creates a new Firebase sales order model
func NewSalesOrderFirebase(client *firestore.Client) *SalesOrderFirebase {
	return &SalesOrderFirebase{
		client: client,
	}
}

// List retrieves all sales orders
func (s *SalesOrderFirebase) List(ctx context.Context) ([]types.SalesOrder, error) {
	docs, err := s.client.Collection("sales_orders").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var orders []types.SalesOrder
	for _, doc := range docs {
		var order types.SalesOrder
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetByID retrieves a sales order by ID
func (s *SalesOrderFirebase) GetByID(ctx context.Context, id string) (*types.SalesOrder, error) {
	doc, err := s.client.Collection("sales_orders").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	var order types.SalesOrder
	if err := doc.DataTo(&order); err != nil {
		return nil, err
	}

	return &order, nil
}

// Create creates a new sales order
func (s *SalesOrderFirebase) Create(ctx context.Context, order types.SalesOrder) error {
	// Generate a unique ID if not provided
	if order.ID == "" {
		order.ID = uuid.New().String()
	}

	// Set the order date if not provided
	if order.Date == "" {
		order.Date = time.Now().Format(time.RFC3339)
	}

	// Validate the order
	if order.CustomerID == "" {
		return fmt.Errorf("customer ID is required")
	}
	if len(order.SalesOrderDetails) == 0 {
		return fmt.Errorf("order must have at least one item")
	}

	// Calculate total amount if not provided
	if order.TotalAmount == 0 {
		for _, detail := range order.SalesOrderDetails {
			order.TotalAmount += detail.TotalPrice
		}
	}

	// Set default status if not provided
	if order.Status == "" {
		order.Status = "pending"
	}

	_, err := s.client.Collection("sales_orders").Doc(order.ID).Set(ctx, order)
	return err
}

// Update updates an existing sales order
func (s *SalesOrderFirebase) Update(ctx context.Context, id string, order types.SalesOrder) error {
	// Get the existing order
	existingOrder, err := s.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get existing order: %v", err)
	}

	// If status is changing to completed, update stock
	if order.Status == "completed" && existingOrder.Status != "completed" {
		// TODO: Implement stock update logic here
	}

	// Update the order
	_, err = s.client.Collection("sales_orders").Doc(id).Set(ctx, order)
	return err
}

// Delete removes a sales order
func (s *SalesOrderFirebase) Delete(ctx context.Context, id string) error {
	_, err := s.client.Collection("sales_orders").Doc(id).Delete(ctx)
	return err
}

// FindByCompany retrieves all sales orders for a specific company
func (s *SalesOrderFirebase) FindByCompany(ctx context.Context, companyID string) ([]types.SalesOrder, error) {
	query := s.client.Collection("sales_orders").Where("company_id", "==", companyID)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var orders []types.SalesOrder
	for _, doc := range docs {
		var order types.SalesOrder
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetSalesByDateRange retrieves sales orders within a date range
func (s *SalesOrderFirebase) GetSalesByDateRange(ctx context.Context, companyID string, startDate, endDate time.Time) ([]types.SalesOrder, error) {
	query := s.client.Collection("sales_orders").Where("company_id", "==", companyID).Where("date", ">=", startDate).Where("date", "<=", endDate)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var orders []types.SalesOrder
	for _, doc := range docs {
		var order types.SalesOrder
		if err := doc.DataTo(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// GetProductTransactions retrieves all transactions for a specific product
func (s *SalesOrderFirebase) GetProductTransactions(ctx context.Context, productID string, startDate, endDate time.Time) ([]types.ProductTransaction, error) {
	// Get all sales orders in the date range
	orders, err := s.GetSalesByDateRange(ctx, "", startDate, endDate)
	if err != nil {
		return nil, err
	}

	// Filter transactions for the specific product
	transactions := make([]types.ProductTransaction, 0)
	for _, order := range orders {
		for _, detail := range order.SalesOrderDetails {
			if detail.ProductID == productID {
				transactions = append(transactions, types.ProductTransaction{
					ID:            order.ID,
					Type:          "sale",
					ProductCode:   detail.ProductCode,
					Quantity:      detail.Quantity,
					PreviousStock: 0, // This would need to be calculated based on stock history
					NewStock:      0, // This would need to be calculated based on stock history
					Timestamp:     time.Now(),
					Notes:         "Sales order: " + order.Code,
				})
			}
		}
	}

	return transactions, nil
}
