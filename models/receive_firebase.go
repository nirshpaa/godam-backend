package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// ReceiveFirebase represents a receive in Firebase
type ReceiveFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewReceiveFirebase creates a new Firebase receive model
func NewReceiveFirebase(client *firestore.Client) *ReceiveFirebase {
	return &ReceiveFirebase{
		FirebaseModel: NewFirebaseModel("receives", client),
		client:        client,
	}
}

// FirebaseReceive represents a receive in Firebase
type FirebaseReceive struct {
	ID             string                  `json:"id"`
	Code           string                  `json:"code"`
	Date           time.Time               `json:"date"`
	Remark         string                  `json:"remark"`
	PurchaseID     string                  `json:"purchase_id"`
	CompanyID      string                  `json:"company_id"`
	BranchID       string                  `json:"branch_id"`
	ReceiveDetails []FirebaseReceiveDetail `json:"receive_details"`
}

// FirebaseReceiveDetail represents a receive detail in Firebase
type FirebaseReceiveDetail struct {
	ID        string              `json:"id"`
	ProductID string              `json:"product_id"`
	Qty       uint                `json:"qty"`
	Code      string              `json:"code"`
	ShelveID  string              `json:"shelve_id"`
	Product   FirebaseProduct     `json:"product"`
	Shelve    ShelveFirebaseModel `json:"shelve"`
}

// List retrieves all receives
func (r *ReceiveFirebase) List(ctx context.Context) ([]FirebaseReceive, error) {
	var receives []FirebaseReceive
	err := r.FirebaseModel.List(ctx, &receives)
	return receives, err
}

// Get retrieves a receive by ID
func (r *ReceiveFirebase) Get(ctx context.Context, id string) (*FirebaseReceive, error) {
	var receive FirebaseReceive
	err := r.FirebaseModel.Get(ctx, id, &receive)
	if err != nil {
		return nil, err
	}
	receive.ID = id
	return &receive, nil
}

// Create creates a new receive
func (r *ReceiveFirebase) Create(ctx context.Context, receive *FirebaseReceive) (string, error) {
	return r.FirebaseModel.Create(ctx, receive)
}

// Update updates an existing receive
func (r *ReceiveFirebase) Update(ctx context.Context, id string, receive *FirebaseReceive) error {
	return r.FirebaseModel.Update(ctx, id, receive)
}

// Delete removes a receive
func (r *ReceiveFirebase) Delete(ctx context.Context, id string) error {
	return r.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all receives for a specific company
func (r *ReceiveFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseReceive, error) {
	query := r.ref.Where("company_id", "==", companyID)

	var receives []FirebaseReceive
	err := r.FirebaseModel.Query(ctx, &query, &receives)
	return receives, err
}
