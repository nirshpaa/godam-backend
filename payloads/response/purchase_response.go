package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// PurchaseResponse represents a purchase response
type PurchaseResponse struct {
	ID              string                   `json:"id"`
	Code            string                   `json:"code"`
	Date            time.Time                `json:"date"`
	AdditionalDisc  float64                  `json:"additional_disc"`
	SupplierID      string                   `json:"supplier_id"`
	CompanyID       string                   `json:"company_id"`
	BranchID        string                   `json:"branch_id"`
	PurchaseDetails []PurchaseDetailResponse `json:"purchase_details"`
}

// PurchaseDetailResponse represents a purchase detail response
type PurchaseDetailResponse struct {
	ID        string                 `json:"id"`
	ProductID string                 `json:"product_id"`
	Price     float64                `json:"price"`
	Disc      float64                `json:"disc"`
	Qty       uint                   `json:"qty"`
	Product   models.FirebaseProduct `json:"product"`
}

// Transform transforms the model into a response
func (r PurchaseResponse) Transform(p *models.FirebasePurchase) PurchaseResponse {
	return PurchaseResponse{
		ID:              p.ID,
		Code:            p.Code,
		Date:            p.Date,
		AdditionalDisc:  p.AdditionalDisc,
		SupplierID:      p.SupplierID,
		CompanyID:       p.CompanyID,
		BranchID:        p.BranchID,
		PurchaseDetails: transformPurchaseDetails(p.PurchaseDetails),
	}
}

// Transform transforms the model into a response
func (r PurchaseDetailResponse) Transform(p models.FirebasePurchaseDetail) PurchaseDetailResponse {
	return PurchaseDetailResponse{
		ID:        p.ID,
		ProductID: p.ProductID,
		Price:     p.Price,
		Disc:      p.Disc,
		Qty:       p.Qty,
		Product:   p.Product,
	}
}

func transformPurchaseDetails(details []models.FirebasePurchaseDetail) []PurchaseDetailResponse {
	var result []PurchaseDetailResponse
	for _, detail := range details {
		result = append(result, PurchaseDetailResponse{}.Transform(detail))
	}
	return result
}

// PurchaseListResponse represents a list of purchases
type PurchaseListResponse struct {
	Purchases []PurchaseResponse `json:"purchases"`
}

// Transform transforms the model into a response
func (r PurchaseListResponse) Transform(purchases []models.FirebasePurchase) PurchaseListResponse {
	for _, p := range purchases {
		r.Purchases = append(r.Purchases, PurchaseResponse{}.Transform(&p))
	}
	return r
}
