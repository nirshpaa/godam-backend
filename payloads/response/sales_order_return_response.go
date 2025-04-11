package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// SalesOrderReturnResponse represents a sales order return response
type SalesOrderReturnResponse struct {
	ID                      string                           `json:"id"`
	Code                    string                           `json:"code"`
	Date                    time.Time                        `json:"date"`
	Price                   float64                          `json:"price"`
	Disc                    float64                          `json:"disc"`
	AdditionalDisc          float64                          `json:"additional_disc"`
	Total                   float64                          `json:"total"`
	SalesOrderID            string                           `json:"sales_order_id"`
	CompanyID               string                           `json:"company_id"`
	BranchID                string                           `json:"branch_id"`
	SalesOrderReturnDetails []SalesOrderReturnDetailResponse `json:"sales_order_return_details"`
}

// Transform converts a FirebaseSalesOrderReturn to SalesOrderReturnResponse
func (r *SalesOrderReturnResponse) Transform(ret *models.FirebaseSalesOrderReturn) *SalesOrderReturnResponse {
	r.ID = ret.ID
	r.Code = ret.Code
	r.Date = ret.Date
	r.Price = ret.Price
	r.Disc = ret.Disc
	r.AdditionalDisc = ret.AdditionalDisc
	r.Total = ret.Total
	r.SalesOrderID = ret.SalesOrderID
	r.CompanyID = ret.CompanyID
	r.BranchID = ret.BranchID
	r.SalesOrderReturnDetails = transformSalesOrderReturnDetails(ret.SalesOrderReturnDetails)
	return r
}

// SalesOrderReturnDetailResponse represents a sales order return detail response
type SalesOrderReturnDetailResponse struct {
	ID        string                 `json:"id"`
	ProductID string                 `json:"product_id"`
	Price     float64                `json:"price"`
	Disc      float64                `json:"disc"`
	Qty       uint                   `json:"qty"`
	Product   models.FirebaseProduct `json:"product"`
}

// Transform converts a FirebaseSalesOrderReturnDetail to SalesOrderReturnDetailResponse
func (r *SalesOrderReturnDetailResponse) Transform(detail *models.FirebaseSalesOrderReturnDetail) *SalesOrderReturnDetailResponse {
	r.ID = detail.ID
	r.ProductID = detail.ProductID
	r.Price = detail.Price
	r.Disc = detail.Disc
	r.Qty = detail.Qty
	r.Product = detail.Product
	return r
}

// SalesOrderReturnListResponse represents a list of sales order returns
type SalesOrderReturnListResponse struct {
	Data []SalesOrderReturnResponse `json:"data"`
}

// Transform converts a slice of FirebaseSalesOrderReturn to SalesOrderReturnListResponse
func (r *SalesOrderReturnListResponse) Transform(returns []models.FirebaseSalesOrderReturn) *SalesOrderReturnListResponse {
	r.Data = make([]SalesOrderReturnResponse, len(returns))
	for i, ret := range returns {
		r.Data[i] = *new(SalesOrderReturnResponse).Transform(&ret)
	}
	return r
}

// Helper function to transform sales order return details
func transformSalesOrderReturnDetails(details []models.FirebaseSalesOrderReturnDetail) []SalesOrderReturnDetailResponse {
	transformed := make([]SalesOrderReturnDetailResponse, len(details))
	for i, detail := range details {
		transformed[i] = *new(SalesOrderReturnDetailResponse).Transform(&detail)
	}
	return transformed
}
