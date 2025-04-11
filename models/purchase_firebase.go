package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// PurchaseFirebase represents a purchase in Firebase
type PurchaseFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewPurchaseFirebase creates a new Firebase purchase model
func NewPurchaseFirebase(client *firestore.Client) *PurchaseFirebase {
	return &PurchaseFirebase{
		FirebaseModel: NewFirebaseModel("purchases", client),
		client:        client,
	}
}

// FirebasePurchase represents a purchase in Firebase
type FirebasePurchase struct {
	ID              string                   `json:"id"`
	Code            string                   `json:"code"`
	Date            time.Time                `json:"date"`
	AdditionalDisc  float64                  `json:"additional_disc"`
	SupplierID      string                   `json:"supplier_id"`
	CompanyID       string                   `json:"company_id"`
	BranchID        string                   `json:"branch_id"`
	PurchaseDetails []FirebasePurchaseDetail `json:"purchase_details"`
}

// FirebasePurchaseDetail represents a purchase detail in Firebase
type FirebasePurchaseDetail struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Price     float64         `json:"price"`
	Disc      float64         `json:"disc"`
	Qty       uint            `json:"qty"`
	Product   FirebaseProduct `json:"product"`
}

// List retrieves all purchases
func (p *PurchaseFirebase) List(ctx context.Context) ([]FirebasePurchase, error) {
	var purchases []FirebasePurchase
	err := p.FirebaseModel.List(ctx, &purchases)
	return purchases, err
}

// Get retrieves a purchase by ID
func (p *PurchaseFirebase) Get(ctx context.Context, id string) (*FirebasePurchase, error) {
	var purchase FirebasePurchase
	err := p.FirebaseModel.Get(ctx, id, &purchase)
	if err != nil {
		return nil, err
	}
	purchase.ID = id
	return &purchase, nil
}

// Create creates a new purchase
func (p *PurchaseFirebase) Create(ctx context.Context, purchase *FirebasePurchase) (string, error) {
	return p.FirebaseModel.Create(ctx, purchase)
}

// Update updates an existing purchase
func (p *PurchaseFirebase) Update(ctx context.Context, id string, purchase *FirebasePurchase) error {
	return p.FirebaseModel.Update(ctx, id, purchase)
}

// Delete removes a purchase
func (p *PurchaseFirebase) Delete(ctx context.Context, id string) error {
	return p.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all purchases for a specific company
func (p *PurchaseFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebasePurchase, error) {
	query := p.ref.Where("company_id", "==", companyID)

	var purchases []FirebasePurchase
	err := p.FirebaseModel.Query(ctx, &query, &purchases)
	return purchases, err
}
