package request

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// NewSalesOrderReturnRequest represents a new sales order return request
type NewSalesOrderReturnRequest struct {
	Date                    string                             `json:"date" validate:"required"`
	AdditionalDisc          float64                            `json:"additional_disc"`
	SalesOrderReturnDetails []NewSalesOrderReturnDetailRequest `json:"sales_order_return_details" validate:"required"`
	SalesOrderID            string                             `json:"sales_order" validate:"required"`
}

// Transform converts NewSalesOrderReturnRequest to FirebaseSalesOrderReturn
func (u *NewSalesOrderReturnRequest) Transform() *models.FirebaseSalesOrderReturn {
	var p models.FirebaseSalesOrderReturn
	p.Date, _ = time.Parse("2006-01-02", u.Date)
	p.SalesOrderID = u.SalesOrderID
	p.AdditionalDisc = u.AdditionalDisc
	p.SalesOrderReturnDetails = transformNewSalesOrderReturnDetails(u.SalesOrderReturnDetails)
	return &p
}

// NewSalesOrderReturnDetailRequest represents a new sales order return detail request
type NewSalesOrderReturnDetailRequest struct {
	Price     float64 `json:"price"`
	Disc      float64 `json:"disc"`
	Qty       uint    `json:"qty" validate:"required"`
	ProductID string  `json:"product"`
}

// Transform converts NewSalesOrderReturnDetailRequest to FirebaseSalesOrderReturnDetail
func (u *NewSalesOrderReturnDetailRequest) Transform() models.FirebaseSalesOrderReturnDetail {
	if u.Qty < 1 {
		u.Qty = 1
	}
	return models.FirebaseSalesOrderReturnDetail{
		Price:     u.Price,
		Disc:      u.Disc,
		Qty:       u.Qty,
		ProductID: u.ProductID,
	}
}

// SalesOrderReturnRequest represents a sales order return request
type SalesOrderReturnRequest struct {
	ID                      string                          `json:"id" validate:"required"`
	Date                    string                          `json:"date"`
	AdditionalDisc          float64                         `json:"additional_disc"`
	SalesOrderReturnDetails []SalesOrderReturnDetailRequest `json:"sales_order_return_details"`
	SalesOrderID            string                          `json:"sales_order"`
}

// Transform converts SalesOrderReturnRequest to FirebaseSalesOrderReturn
func (u *SalesOrderReturnRequest) Transform() *models.FirebaseSalesOrderReturn {
	var p models.FirebaseSalesOrderReturn
	p.ID = u.ID
	p.Date, _ = time.Parse("2006-01-02", u.Date)
	p.SalesOrderID = u.SalesOrderID
	p.AdditionalDisc = u.AdditionalDisc
	p.SalesOrderReturnDetails = transformSalesOrderReturnDetails(u.SalesOrderReturnDetails)
	return &p
}

// SalesOrderReturnDetailRequest represents a sales order return detail request
type SalesOrderReturnDetailRequest struct {
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	Disc      float64 `json:"disc"`
	Qty       uint    `json:"qty"`
	ProductID string  `json:"product"`
}

// Transform converts SalesOrderReturnDetailRequest to FirebaseSalesOrderReturnDetail
func (u *SalesOrderReturnDetailRequest) Transform() models.FirebaseSalesOrderReturnDetail {
	if u.Qty < 1 {
		u.Qty = 1
	}
	return models.FirebaseSalesOrderReturnDetail{
		ID:        u.ID,
		Price:     u.Price,
		Disc:      u.Disc,
		Qty:       u.Qty,
		ProductID: u.ProductID,
	}
}

// Helper function to transform new sales order return details
func transformNewSalesOrderReturnDetails(details []NewSalesOrderReturnDetailRequest) []models.FirebaseSalesOrderReturnDetail {
	transformed := make([]models.FirebaseSalesOrderReturnDetail, len(details))
	for i, detail := range details {
		transformed[i] = detail.Transform()
	}
	return transformed
}

// Helper function to transform sales order return details
func transformSalesOrderReturnDetails(details []SalesOrderReturnDetailRequest) []models.FirebaseSalesOrderReturnDetail {
	transformed := make([]models.FirebaseSalesOrderReturnDetail, len(details))
	for i, detail := range details {
		transformed[i] = detail.Transform()
	}
	return transformed
}
