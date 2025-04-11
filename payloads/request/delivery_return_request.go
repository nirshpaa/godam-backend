package request

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

// NewDeliveryReturnRequest represents a new delivery return request
type NewDeliveryReturnRequest struct {
	Date                  time.Time                        `json:"date" binding:"required"`
	Remark                string                           `json:"remark"`
	DeliveryReturnDetails []NewDeliveryReturnDetailRequest `json:"delivery_return_details" binding:"required"`
	DeliveryID            string                           `json:"delivery_id" binding:"required"`
}

// NewDeliveryReturnDetailRequest represents a new delivery return detail request
type NewDeliveryReturnDetailRequest struct {
	ProductID string `json:"product_id" binding:"required"`
	Qty       uint   `json:"qty" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// DeliveryReturnRequest represents an update delivery return request
type DeliveryReturnRequest struct {
	ID                    string                        `json:"id" binding:"required"`
	Date                  time.Time                     `json:"date" binding:"required"`
	Remark                string                        `json:"remark"`
	DeliveryReturnDetails []DeliveryReturnDetailRequest `json:"delivery_return_details" binding:"required"`
	DeliveryID            string                        `json:"delivery_id" binding:"required"`
}

// DeliveryReturnDetailRequest represents an update delivery return detail request
type DeliveryReturnDetailRequest struct {
	ID        string `json:"id" binding:"required"`
	ProductID string `json:"product_id" binding:"required"`
	Qty       uint   `json:"qty" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// Transform transforms the request into a model
func (u *NewDeliveryReturnRequest) Transform() *models.FirebaseDeliveryReturn {
	return &models.FirebaseDeliveryReturn{
		Date:                  u.Date,
		Remark:                u.Remark,
		DeliveryID:            u.DeliveryID,
		DeliveryReturnDetails: transformNewDeliveryReturnDetails(u.DeliveryReturnDetails),
	}
}

// Transform transforms the request into a model
func (u *NewDeliveryReturnDetailRequest) Transform() models.FirebaseDeliveryReturnDetail {
	return models.FirebaseDeliveryReturnDetail{
		ProductID: u.ProductID,
		Qty:       u.Qty,
		Code:      u.Code,
	}
}

// Transform transforms the request into a model
func (u *DeliveryReturnRequest) Transform() *models.FirebaseDeliveryReturn {
	return &models.FirebaseDeliveryReturn{
		ID:                    u.ID,
		Date:                  u.Date,
		Remark:                u.Remark,
		DeliveryID:            u.DeliveryID,
		DeliveryReturnDetails: transformDeliveryReturnDetails(u.DeliveryReturnDetails),
	}
}

// Transform transforms the request into a model
func (u *DeliveryReturnDetailRequest) Transform() models.FirebaseDeliveryReturnDetail {
	return models.FirebaseDeliveryReturnDetail{
		ID:        u.ID,
		ProductID: u.ProductID,
		Qty:       u.Qty,
		Code:      u.Code,
	}
}

func transformNewDeliveryReturnDetails(details []NewDeliveryReturnDetailRequest) []models.FirebaseDeliveryReturnDetail {
	var result []models.FirebaseDeliveryReturnDetail
	for _, detail := range details {
		result = append(result, detail.Transform())
	}
	return result
}

func transformDeliveryReturnDetails(details []DeliveryReturnDetailRequest) []models.FirebaseDeliveryReturnDetail {
	var result []models.FirebaseDeliveryReturnDetail
	for _, detail := range details {
		result = append(result, detail.Transform())
	}
	return result
}

// Validate validates the request
func (u *NewDeliveryReturnRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}

// Validate validates the request
func (u *DeliveryReturnRequest) Validate(c *gin.Context) error {
	return c.ShouldBindJSON(u)
}
