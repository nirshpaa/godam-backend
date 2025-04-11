package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// DeliveryReturnFirebase represents a delivery return in Firebase
type DeliveryReturnFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewDeliveryReturnFirebase creates a new Firebase delivery return model
func NewDeliveryReturnFirebase(client *firestore.Client) *DeliveryReturnFirebase {
	return &DeliveryReturnFirebase{
		FirebaseModel: NewFirebaseModel("delivery_returns", client),
		client:        client,
	}
}

// FirebaseDeliveryReturn represents a delivery return in Firebase
type FirebaseDeliveryReturn struct {
	ID                    string                         `json:"id"`
	Code                  string                         `json:"code"`
	Date                  time.Time                      `json:"date"`
	Remark                string                         `json:"remark"`
	DeliveryID            string                         `json:"delivery_id"`
	CompanyID             string                         `json:"company_id"`
	BranchID              string                         `json:"branch_id"`
	DeliveryReturnDetails []FirebaseDeliveryReturnDetail `json:"delivery_return_details"`
}

// FirebaseDeliveryReturnDetail represents a delivery return detail in Firebase
type FirebaseDeliveryReturnDetail struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Qty       uint            `json:"qty"`
	Code      string          `json:"code"`
	Product   FirebaseProduct `json:"product"`
}

// List retrieves all delivery returns
func (d *DeliveryReturnFirebase) List(ctx context.Context) ([]FirebaseDeliveryReturn, error) {
	var returns []FirebaseDeliveryReturn
	err := d.FirebaseModel.List(ctx, &returns)
	return returns, err
}

// Get retrieves a delivery return by ID
func (d *DeliveryReturnFirebase) Get(ctx context.Context, id string) (*FirebaseDeliveryReturn, error) {
	var ret FirebaseDeliveryReturn
	err := d.FirebaseModel.Get(ctx, id, &ret)
	if err != nil {
		return nil, err
	}
	ret.ID = id
	return &ret, nil
}

// Create creates a new delivery return
func (d *DeliveryReturnFirebase) Create(ctx context.Context, ret *FirebaseDeliveryReturn) (string, error) {
	return d.FirebaseModel.Create(ctx, ret)
}

// Update updates an existing delivery return
func (d *DeliveryReturnFirebase) Update(ctx context.Context, id string, ret *FirebaseDeliveryReturn) error {
	return d.FirebaseModel.Update(ctx, id, ret)
}

// Delete removes a delivery return
func (d *DeliveryReturnFirebase) Delete(ctx context.Context, id string) error {
	return d.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all delivery returns for a specific company
func (d *DeliveryReturnFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseDeliveryReturn, error) {
	query := d.ref.Where("company_id", "==", companyID)

	var returns []FirebaseDeliveryReturn
	err := d.FirebaseModel.Query(ctx, &query, &returns)
	return returns, err
}
