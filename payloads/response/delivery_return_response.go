package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// DeliveryReturnResponse represents a delivery return response
type DeliveryReturnResponse struct {
	ID                    string                         `json:"id"`
	Code                  string                         `json:"code"`
	Date                  time.Time                      `json:"date"`
	Remark                string                         `json:"remark"`
	DeliveryID            string                         `json:"delivery_id"`
	CompanyID             string                         `json:"company_id"`
	BranchID              string                         `json:"branch_id"`
	DeliveryReturnDetails []DeliveryReturnDetailResponse `json:"delivery_return_details"`
}

// DeliveryReturnDetailResponse represents a delivery return detail response
type DeliveryReturnDetailResponse struct {
	ID        string                 `json:"id"`
	ProductID string                 `json:"product_id"`
	Qty       uint                   `json:"qty"`
	Code      string                 `json:"code"`
	Product   models.FirebaseProduct `json:"product"`
}

// Transform transforms the model into a response
func (r DeliveryReturnResponse) Transform(p *models.FirebaseDeliveryReturn) DeliveryReturnResponse {
	r.ID = p.ID
	r.Code = p.Code
	r.Date = p.Date
	r.Remark = p.Remark
	r.DeliveryID = p.DeliveryID
	r.CompanyID = p.CompanyID
	r.BranchID = p.BranchID
	r.DeliveryReturnDetails = transformDeliveryReturnDetails(p.DeliveryReturnDetails)
	return r
}

// Transform transforms the model into a response
func (r DeliveryReturnDetailResponse) Transform(p models.FirebaseDeliveryReturnDetail) DeliveryReturnDetailResponse {
	return DeliveryReturnDetailResponse{
		ID:        p.ID,
		ProductID: p.ProductID,
		Qty:       p.Qty,
		Code:      p.Code,
		Product:   p.Product,
	}
}

func transformDeliveryReturnDetails(details []models.FirebaseDeliveryReturnDetail) []DeliveryReturnDetailResponse {
	var result []DeliveryReturnDetailResponse
	for _, detail := range details {
		result = append(result, DeliveryReturnDetailResponse{}.Transform(detail))
	}
	return result
}

// DeliveryReturnListResponse represents a list of delivery returns
type DeliveryReturnListResponse struct {
	DeliveryReturns []DeliveryReturnResponse `json:"delivery_returns"`
}

// Transform transforms the model into a response
func (r DeliveryReturnListResponse) Transform(returns []models.FirebaseDeliveryReturn) DeliveryReturnListResponse {
	for _, ret := range returns {
		r.DeliveryReturns = append(r.DeliveryReturns, DeliveryReturnResponse{}.Transform(&ret))
	}
	return r
}
