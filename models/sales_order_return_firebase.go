package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// SalesOrderReturnFirebase represents a sales order return in Firebase
type SalesOrderReturnFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewSalesOrderReturnFirebase creates a new Firebase sales order return model
func NewSalesOrderReturnFirebase(client *firestore.Client) *SalesOrderReturnFirebase {
	return &SalesOrderReturnFirebase{
		FirebaseModel: NewFirebaseModel("sales_order_returns", client),
		client:        client,
	}
}

// FirebaseSalesOrderReturn represents a sales order return in Firebase
type FirebaseSalesOrderReturn struct {
	ID                      string                           `json:"id"`
	Code                    string                           `json:"code"`
	Date                    time.Time                        `json:"date"`
	Price                   float64                          `json:"price"`
	Disc                    float64                          `json:"disc"`
	AdditionalDisc          float64                          `json:"additional_disc"`
	Total                   float64                          `json:"total"`
	SalesOrderID            string                           `json:"sales_order_id"`
	CompanyID               string                           `json:"company_id"`
	BranchID                string                           `json:"branch_id"`
	SalesOrderReturnDetails []FirebaseSalesOrderReturnDetail `json:"sales_order_return_details"`
}

// FirebaseSalesOrderReturnDetail represents a sales order return detail in Firebase
type FirebaseSalesOrderReturnDetail struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Price     float64         `json:"price"`
	Disc      float64         `json:"disc"`
	Qty       uint            `json:"qty"`
	Product   FirebaseProduct `json:"product"`
}

// List retrieves all sales order returns
func (s *SalesOrderReturnFirebase) List(ctx context.Context) ([]FirebaseSalesOrderReturn, error) {
	var returns []FirebaseSalesOrderReturn
	err := s.FirebaseModel.List(ctx, &returns)
	return returns, err
}

// Get retrieves a sales order return by ID
func (s *SalesOrderReturnFirebase) Get(ctx context.Context, id string) (*FirebaseSalesOrderReturn, error) {
	var ret FirebaseSalesOrderReturn
	err := s.FirebaseModel.Get(ctx, id, &ret)
	if err != nil {
		return nil, err
	}
	ret.ID = id
	return &ret, nil
}

// Create creates a new sales order return
func (s *SalesOrderReturnFirebase) Create(ctx context.Context, ret *FirebaseSalesOrderReturn) (string, error) {
	return s.FirebaseModel.Create(ctx, ret)
}

// Update updates an existing sales order return
func (s *SalesOrderReturnFirebase) Update(ctx context.Context, id string, ret *FirebaseSalesOrderReturn) error {
	return s.FirebaseModel.Update(ctx, id, ret)
}

// Delete removes a sales order return
func (s *SalesOrderReturnFirebase) Delete(ctx context.Context, id string) error {
	return s.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all sales order returns for a specific company
func (s *SalesOrderReturnFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseSalesOrderReturn, error) {
	query := s.ref.Where("company_id", "==", companyID)

	var returns []FirebaseSalesOrderReturn
	err := s.FirebaseModel.Query(ctx, &query, &returns)
	return returns, err
}
