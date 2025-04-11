package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// ReceiveReturnFirebase represents a receive return in Firebase
type ReceiveReturnFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewReceiveReturnFirebase creates a new Firebase receive return model
func NewReceiveReturnFirebase(client *firestore.Client) *ReceiveReturnFirebase {
	return &ReceiveReturnFirebase{
		FirebaseModel: NewFirebaseModel("receive_returns", client),
		client:        client,
	}
}

// FirebaseReceiveReturn represents a receive return in Firebase
type FirebaseReceiveReturn struct {
	ID                   string                        `json:"id"`
	Code                 string                        `json:"code"`
	Date                 time.Time                     `json:"date"`
	Remark               string                        `json:"remark"`
	ReceiveID            string                        `json:"receive_id"`
	CompanyID            string                        `json:"company_id"`
	BranchID             string                        `json:"branch_id"`
	ReceiveReturnDetails []FirebaseReceiveReturnDetail `json:"receive_return_details"`
}

// FirebaseReceiveReturnDetail represents a receive return detail in Firebase
type FirebaseReceiveReturnDetail struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Qty       uint            `json:"qty"`
	Code      string          `json:"code"`
	Product   FirebaseProduct `json:"product"`
}

// List retrieves all receive returns
func (r *ReceiveReturnFirebase) List(ctx context.Context) ([]FirebaseReceiveReturn, error) {
	var returns []FirebaseReceiveReturn
	err := r.FirebaseModel.List(ctx, &returns)
	return returns, err
}

// Get retrieves a receive return by ID
func (r *ReceiveReturnFirebase) Get(ctx context.Context, id string) (*FirebaseReceiveReturn, error) {
	var ret FirebaseReceiveReturn
	err := r.FirebaseModel.Get(ctx, id, &ret)
	if err != nil {
		return nil, err
	}
	ret.ID = id
	return &ret, nil
}

// Create creates a new receive return
func (r *ReceiveReturnFirebase) Create(ctx context.Context, ret *FirebaseReceiveReturn) (string, error) {
	return r.FirebaseModel.Create(ctx, ret)
}

// Update updates an existing receive return
func (r *ReceiveReturnFirebase) Update(ctx context.Context, id string, ret *FirebaseReceiveReturn) error {
	return r.FirebaseModel.Update(ctx, id, ret)
}

// Delete removes a receive return
func (r *ReceiveReturnFirebase) Delete(ctx context.Context, id string) error {
	return r.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all receive returns for a specific company
func (r *ReceiveReturnFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseReceiveReturn, error) {
	query := r.ref.Where("company_id", "==", companyID)

	var returns []FirebaseReceiveReturn
	err := r.FirebaseModel.Query(ctx, &query, &returns)
	return returns, err
}
