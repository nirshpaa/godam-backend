package models

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
)

// PurchaseReturnFirebase represents a purchase return in Firebase
type PurchaseReturnFirebase struct {
	*FirebaseModel
	client *firestore.Client
}

// NewPurchaseReturnFirebase creates a new Firebase purchase return model
func NewPurchaseReturnFirebase(client *firestore.Client) *PurchaseReturnFirebase {
	return &PurchaseReturnFirebase{
		FirebaseModel: NewFirebaseModel("purchase_returns", client),
		client:        client,
	}
}

// FirebasePurchaseReturn represents a purchase return in Firebase
type FirebasePurchaseReturn struct {
	ID                    string                         `json:"id"`
	Code                  string                         `json:"code"`
	Date                  time.Time                      `json:"date"`
	Remark                string                         `json:"remark"`
	PurchaseID            string                         `json:"purchase_id"`
	CompanyID             string                         `json:"company_id"`
	BranchID              string                         `json:"branch_id"`
	PurchaseReturnDetails []FirebasePurchaseReturnDetail `json:"purchase_return_details"`
}

// FirebasePurchaseReturnDetail represents a purchase return detail in Firebase
type FirebasePurchaseReturnDetail struct {
	ID        string          `json:"id"`
	ProductID string          `json:"product_id"`
	Qty       uint            `json:"qty"`
	Code      string          `json:"code"`
	Product   FirebaseProduct `json:"product"`
}

// List retrieves all purchase returns
func (p *PurchaseReturnFirebase) List(ctx context.Context) ([]FirebasePurchaseReturn, error) {
	var returns []FirebasePurchaseReturn
	err := p.FirebaseModel.List(ctx, &returns)
	return returns, err
}

// Get retrieves a purchase return by ID
func (p *PurchaseReturnFirebase) Get(ctx context.Context, id string) (*FirebasePurchaseReturn, error) {
	var ret FirebasePurchaseReturn
	err := p.FirebaseModel.Get(ctx, id, &ret)
	if err != nil {
		return nil, err
	}
	ret.ID = id
	return &ret, nil
}

// Create creates a new purchase return
func (p *PurchaseReturnFirebase) Create(ctx context.Context, ret *FirebasePurchaseReturn) (string, error) {
	return p.FirebaseModel.Create(ctx, ret)
}

// Update updates an existing purchase return
func (p *PurchaseReturnFirebase) Update(ctx context.Context, id string, ret *FirebasePurchaseReturn) error {
	return p.FirebaseModel.Update(ctx, id, ret)
}

// Delete removes a purchase return
func (p *PurchaseReturnFirebase) Delete(ctx context.Context, id string) error {
	return p.FirebaseModel.Delete(ctx, id)
}

// FindByCompany retrieves all purchase returns for a specific company
func (p *PurchaseReturnFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebasePurchaseReturn, error) {
	query := p.ref.Where("company_id", "==", companyID)

	var returns []FirebasePurchaseReturn
	err := p.FirebaseModel.Query(ctx, &query, &returns)
	return returns, err
}
