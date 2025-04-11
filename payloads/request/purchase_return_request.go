package request

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// NewPurchaseReturnRequest represents a new purchase return request
type NewPurchaseReturnRequest struct {
	Date                  time.Time                        `json:"date" binding:"required"`
	Remark                string                           `json:"remark"`
	PurchaseReturnDetails []NewPurchaseReturnDetailRequest `json:"purchase_return_details" binding:"required"`
	PurchaseID            string                           `json:"purchase_id" binding:"required"`
}

// NewPurchaseReturnDetailRequest represents a new purchase return detail request
type NewPurchaseReturnDetailRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Qty       uint   `json:"qty" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// PurchaseReturnRequest represents an update purchase return request
type PurchaseReturnRequest struct {
	ID                    string                        `json:"id" binding:"required"`
	Date                  time.Time                     `json:"date" binding:"required"`
	Remark                string                        `json:"remark"`
	PurchaseReturnDetails []PurchaseReturnDetailRequest `json:"purchase_return_details" binding:"required"`
	PurchaseID            string                        `json:"purchase_id" binding:"required"`
}

// PurchaseReturnDetailRequest represents an update purchase return detail request
type PurchaseReturnDetailRequest struct {
	ID        string `json:"id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Qty       uint   `json:"qty" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// Transform transforms the request into a model
func (u *NewPurchaseReturnRequest) Transform() *models.FirebasePurchaseReturn {
	return &models.FirebasePurchaseReturn{
		Date:                  u.Date,
		Remark:                u.Remark,
		PurchaseID:            u.PurchaseID,
		PurchaseReturnDetails: transformNewPurchaseReturnDetails(u.PurchaseReturnDetails),
	}
}

// Transform transforms the request into a model
func (u *NewPurchaseReturnDetailRequest) Transform() models.FirebasePurchaseReturnDetail {
	return models.FirebasePurchaseReturnDetail{
		ProductID: u.ProductID,
		Qty:       u.Qty,
		Code:      u.Code,
	}
}

// Transform transforms the request into a model
func (u *PurchaseReturnRequest) Transform() *models.FirebasePurchaseReturn {
	return &models.FirebasePurchaseReturn{
		ID:                    u.ID,
		Date:                  u.Date,
		Remark:                u.Remark,
		PurchaseID:            u.PurchaseID,
		PurchaseReturnDetails: transformPurchaseReturnDetails(u.PurchaseReturnDetails),
	}
}

// Transform transforms the request into a model
func (u *PurchaseReturnDetailRequest) Transform() models.FirebasePurchaseReturnDetail {
	return models.FirebasePurchaseReturnDetail{
		ID:        u.ID,
		ProductID: u.ProductID,
		Qty:       u.Qty,
		Code:      u.Code,
	}
}

func transformNewPurchaseReturnDetails(details []NewPurchaseReturnDetailRequest) []models.FirebasePurchaseReturnDetail {
	var result []models.FirebasePurchaseReturnDetail
	for _, detail := range details {
		result = append(result, detail.Transform())
	}
	return result
}

func transformPurchaseReturnDetails(details []PurchaseReturnDetailRequest) []models.FirebasePurchaseReturnDetail {
	var result []models.FirebasePurchaseReturnDetail
	for _, detail := range details {
		result = append(result, detail.Transform())
	}
	return result
}

// Validate validates the request
func (u *NewPurchaseReturnRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}

// Validate validates the request
func (u *PurchaseReturnRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}
