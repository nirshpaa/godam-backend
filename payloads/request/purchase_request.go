package request

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// NewPurchaseRequest represents a new purchase request
type NewPurchaseRequest struct {
	Date            time.Time                  `json:"date" binding:"required"`
	AdditionalDisc  float64                    `json:"additional_disc"`
	PurchaseDetails []NewPurchaseDetailRequest `json:"purchase_details" binding:"required"`
	SupplierID      string                     `json:"supplier_id" binding:"required"`
}

// NewPurchaseDetailRequest represents a new purchase detail request
type NewPurchaseDetailRequest struct {
	Price     float64 `json:"price" binding:"required"`
	Disc      float64 `json:"disc"`
	Qty       uint    `json:"qty" binding:"required"`
	ProductID string  `json:"product_id" binding:"required"`
}

// PurchaseRequest represents an update purchase request
type PurchaseRequest struct {
	ID              string                  `json:"id" binding:"required"`
	Date            time.Time               `json:"date" binding:"required"`
	AdditionalDisc  float64                 `json:"additional_disc"`
	PurchaseDetails []PurchaseDetailRequest `json:"purchase_details" binding:"required"`
	SupplierID      string                  `json:"supplier_id" binding:"required"`
}

// PurchaseDetailRequest represents an update purchase detail request
type PurchaseDetailRequest struct {
	ID        string  `json:"id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	Disc      float64 `json:"disc"`
	Qty       uint    `json:"qty" binding:"required"`
	ProductID string  `json:"product_id" binding:"required"`
}

// Transform transforms the request into a model
func (u *NewPurchaseRequest) Transform() *models.FirebasePurchase {
	return &models.FirebasePurchase{
		Date:            u.Date,
		AdditionalDisc:  u.AdditionalDisc,
		SupplierID:      u.SupplierID,
		PurchaseDetails: transformNewPurchaseDetails(u.PurchaseDetails),
	}
}

// Transform transforms the request into a model
func (u *NewPurchaseDetailRequest) Transform() models.FirebasePurchaseDetail {
	return models.FirebasePurchaseDetail{
		Price:     u.Price,
		Disc:      u.Disc,
		Qty:       u.Qty,
		ProductID: u.ProductID,
	}
}

// Transform transforms the request into a model
func (u *PurchaseRequest) Transform() *models.FirebasePurchase {
	return &models.FirebasePurchase{
		ID:              u.ID,
		Date:            u.Date,
		AdditionalDisc:  u.AdditionalDisc,
		SupplierID:      u.SupplierID,
		PurchaseDetails: transformPurchaseDetails(u.PurchaseDetails),
	}
}

// Transform transforms the request into a model
func (u *PurchaseDetailRequest) Transform() models.FirebasePurchaseDetail {
	return models.FirebasePurchaseDetail{
		ID:        u.ID,
		Price:     u.Price,
		Disc:      u.Disc,
		Qty:       u.Qty,
		ProductID: u.ProductID,
	}
}

func transformNewPurchaseDetails(details []NewPurchaseDetailRequest) []models.FirebasePurchaseDetail {
	var result []models.FirebasePurchaseDetail
	for _, detail := range details {
		result = append(result, detail.Transform())
	}
	return result
}

func transformPurchaseDetails(details []PurchaseDetailRequest) []models.FirebasePurchaseDetail {
	var result []models.FirebasePurchaseDetail
	for _, detail := range details {
		result = append(result, detail.Transform())
	}
	return result
}

// Validate validates the request
func (u *NewPurchaseRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}

// Validate validates the request
func (u *PurchaseRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}
