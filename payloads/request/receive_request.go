package request

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// NewReceiveRequest represents a new receive request
type NewReceiveRequest struct {
	Date           string                    `json:"date" validate:"required"`
	Remark         string                    `json:"remark"`
	ReceiveDetails []NewReceiveDetailRequest `json:"receive_details" validate:"required"`
	PurchaseID     string                    `json:"purchase" validate:"required"`
}

// Transform converts NewReceiveRequest to FirebaseReceive
func (u *NewReceiveRequest) Transform() *models.FirebaseReceive {
	var p models.FirebaseReceive
	p.Date, _ = time.Parse("2006-01-02", u.Date)
	p.PurchaseID = u.PurchaseID
	p.Remark = u.Remark
	p.ReceiveDetails = transformNewReceiveDetails(u.ReceiveDetails)
	return &p
}

// NewReceiveDetailRequest represents a new receive detail request
type NewReceiveDetailRequest struct {
	ProductID string `json:"product" validate:"required"`
	ShelveID  string `json:"shelve" validate:"required"`
}

// Transform converts NewReceiveDetailRequest to FirebaseReceiveDetail
func (u *NewReceiveDetailRequest) Transform() models.FirebaseReceiveDetail {
	return models.FirebaseReceiveDetail{
		ProductID: u.ProductID,
		ShelveID:  u.ShelveID,
		Qty:       1,
	}
}

// ReceiveRequest represents a receive request
type ReceiveRequest struct {
	ID             string                 `json:"id" validate:"required"`
	Date           string                 `json:"date"`
	Remark         string                 `json:"remark"`
	ReceiveDetails []ReceiveDetailRequest `json:"receive_details"`
	PurchaseID     string                 `json:"purchase"`
}

// Transform converts ReceiveRequest to FirebaseReceive
func (u *ReceiveRequest) Transform() *models.FirebaseReceive {
	var p models.FirebaseReceive
	p.ID = u.ID
	p.Date, _ = time.Parse("2006-01-02", u.Date)
	p.PurchaseID = u.PurchaseID
	p.Remark = u.Remark
	p.ReceiveDetails = transformReceiveDetails(u.ReceiveDetails)
	return &p
}

// ReceiveDetailRequest represents a receive detail request
type ReceiveDetailRequest struct {
	ID        string `json:"id"`
	ProductID string `json:"product"`
	ShelveID  string `json:"shelve"`
}

// Transform converts ReceiveDetailRequest to FirebaseReceiveDetail
func (u *ReceiveDetailRequest) Transform() models.FirebaseReceiveDetail {
	return models.FirebaseReceiveDetail{
		ID:        u.ID,
		ProductID: u.ProductID,
		ShelveID:  u.ShelveID,
		Qty:       1,
	}
}

// Helper function to transform new receive details
func transformNewReceiveDetails(details []NewReceiveDetailRequest) []models.FirebaseReceiveDetail {
	transformed := make([]models.FirebaseReceiveDetail, len(details))
	for i, detail := range details {
		transformed[i] = detail.Transform()
	}
	return transformed
}

// Helper function to transform receive details
func transformReceiveDetails(details []ReceiveDetailRequest) []models.FirebaseReceiveDetail {
	transformed := make([]models.FirebaseReceiveDetail, len(details))
	for i, detail := range details {
		transformed[i] = detail.Transform()
	}
	return transformed
}
