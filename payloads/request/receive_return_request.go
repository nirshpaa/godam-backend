package request

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// NewReceiveReturnRequest represents a new receive return request
type NewReceiveReturnRequest struct {
	Date                 string                          `json:"date" validate:"required"`
	Remark               string                          `json:"remark"`
	ReceiveReturnDetails []NewReceiveReturnDetailRequest `json:"receive_return_details" validate:"required"`
	ReceiveID            string                          `json:"receive" validate:"required"`
}

// Transform converts NewReceiveReturnRequest to FirebaseReceiveReturn
func (u *NewReceiveReturnRequest) Transform() *models.FirebaseReceiveReturn {
	var p models.FirebaseReceiveReturn
	p.Date, _ = time.Parse("2006-01-02", u.Date)
	p.ReceiveID = u.ReceiveID
	p.Remark = u.Remark
	p.ReceiveReturnDetails = transformNewReceiveReturnDetails(u.ReceiveReturnDetails)
	return &p
}

// NewReceiveReturnDetailRequest represents a new receive return detail request
type NewReceiveReturnDetailRequest struct {
	ProductID string `json:"product" validate:"required"`
	Code      string `json:"code" validate:"required"`
}

// Transform converts NewReceiveReturnDetailRequest to FirebaseReceiveReturnDetail
func (u *NewReceiveReturnDetailRequest) Transform() models.FirebaseReceiveReturnDetail {
	return models.FirebaseReceiveReturnDetail{
		ProductID: u.ProductID,
		Code:      u.Code,
		Qty:       1,
	}
}

// ReceiveReturnRequest represents a receive return request
type ReceiveReturnRequest struct {
	ID                   string                       `json:"id" validate:"required"`
	Date                 string                       `json:"date"`
	Remark               string                       `json:"remark"`
	ReceiveReturnDetails []ReceiveReturnDetailRequest `json:"receive_return_details"`
	ReceiveID            string                       `json:"receive"`
}

// Transform converts ReceiveReturnRequest to FirebaseReceiveReturn
func (u *ReceiveReturnRequest) Transform() *models.FirebaseReceiveReturn {
	var p models.FirebaseReceiveReturn
	p.ID = u.ID
	p.Date, _ = time.Parse("2006-01-02", u.Date)
	p.ReceiveID = u.ReceiveID
	p.Remark = u.Remark
	p.ReceiveReturnDetails = transformReceiveReturnDetails(u.ReceiveReturnDetails)
	return &p
}

// ReceiveReturnDetailRequest represents a receive return detail request
type ReceiveReturnDetailRequest struct {
	ID        string `json:"id"`
	ProductID string `json:"product"`
	Code      string `json:"code"`
}

// Transform converts ReceiveReturnDetailRequest to FirebaseReceiveReturnDetail
func (u *ReceiveReturnDetailRequest) Transform() models.FirebaseReceiveReturnDetail {
	return models.FirebaseReceiveReturnDetail{
		ID:        u.ID,
		ProductID: u.ProductID,
		Code:      u.Code,
		Qty:       1,
	}
}

// Helper function to transform new receive return details
func transformNewReceiveReturnDetails(details []NewReceiveReturnDetailRequest) []models.FirebaseReceiveReturnDetail {
	transformed := make([]models.FirebaseReceiveReturnDetail, len(details))
	for i, detail := range details {
		transformed[i] = detail.Transform()
	}
	return transformed
}

// Helper function to transform receive return details
func transformReceiveReturnDetails(details []ReceiveReturnDetailRequest) []models.FirebaseReceiveReturnDetail {
	transformed := make([]models.FirebaseReceiveReturnDetail, len(details))
	for i, detail := range details {
		transformed[i] = detail.Transform()
	}
	return transformed
}
