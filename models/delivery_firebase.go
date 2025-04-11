package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// DeliveryFirebase represents a delivery in Firebase
type DeliveryFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewDeliveryFirebase creates a new Firebase delivery model
func NewDeliveryFirebase(client *firestore.Client) *DeliveryFirebase {
	return &DeliveryFirebase{
		FirebaseModel: NewFirebaseModel("deliveries", client),
		client:        client,
	}
}

// FirebaseDelivery represents a delivery in Firebase
type FirebaseDelivery struct {
	ID              string                   `json:"id"`
	Code            string                   `json:"code"`
	Date            time.Time                `json:"date"`
	Remark          string                   `json:"remark"`
	SalesOrderID    string                   `json:"sales_order_id"`
	CompanyID       string                   `json:"company_id"`
	BranchID        string                   `json:"branch_id"`
	DeliveryDetails []FirebaseDeliveryDetail `json:"delivery_details"`
}

// FirebaseDeliveryDetail represents a delivery detail in Firebase
type FirebaseDeliveryDetail struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Qty       uint            `json:"qty"`
	Code      string          `json:"code"`
	ShelveID  string          `json:"shelve_id"`
	Product   FirebaseProduct `json:"product"`
}

// List retrieves all deliveries
func (d *DeliveryFirebase) List(ctx context.Context) ([]FirebaseDelivery, error) {
	var deliveries []FirebaseDelivery
	err := d.FirebaseModel.List(ctx, &deliveries)
	return deliveries, err
}

// Get retrieves a delivery by ID
func (d *DeliveryFirebase) Get(ctx context.Context, id string) (*FirebaseDelivery, error) {
	var delivery FirebaseDelivery
	err := d.FirebaseModel.Get(ctx, id, &delivery)
	if err != nil {
		return nil, err
	}
	delivery.ID = id
	return &delivery, nil
}

// Create creates a new delivery
func (d *DeliveryFirebase) Create(ctx context.Context, delivery *FirebaseDelivery) (string, error) {
	return d.FirebaseModel.Create(ctx, delivery)
}

// Update updates an existing delivery
func (d *DeliveryFirebase) Update(ctx context.Context, id string, delivery *FirebaseDelivery) error {
	return d.FirebaseModel.Update(ctx, id, delivery)
}

// Delete removes a delivery
func (d *DeliveryFirebase) Delete(ctx context.Context, id string) error {
	return d.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all deliveries for a specific company
func (d *DeliveryFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseDelivery, error) {
	query := d.ref.Where("company_id", "==", companyID)

	var deliveries []FirebaseDelivery
	err := d.FirebaseModel.Query(ctx, &query, &deliveries)
	return deliveries, err
}
