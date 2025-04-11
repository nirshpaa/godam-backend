package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// ReceiveReturnResponse represents a receive return response
type ReceiveReturnResponse struct {
	ID                   string                        `json:"id"`
	Code                 string                        `json:"code"`
	Date                 time.Time                     `json:"date"`
	Remark               string                        `json:"remark"`
	ReceiveID            string                        `json:"receive_id"`
	CompanyID            string                        `json:"company_id"`
	BranchID             string                        `json:"branch_id"`
	ReceiveReturnDetails []ReceiveReturnDetailResponse `json:"receive_return_details"`
}

// Transform converts a FirebaseReceiveReturn to a ReceiveReturnResponse
func (u *ReceiveReturnResponse) Transform(receiveReturn *models.FirebaseReceiveReturn) {
	u.ID = receiveReturn.ID
	u.Code = receiveReturn.Code
	u.Date = receiveReturn.Date
	u.Remark = receiveReturn.Remark
	u.ReceiveID = receiveReturn.ReceiveID
	u.CompanyID = receiveReturn.CompanyID
	u.BranchID = receiveReturn.BranchID

	for _, d := range receiveReturn.ReceiveReturnDetails {
		var p ReceiveReturnDetailResponse
		p.Transform(&d)
		u.ReceiveReturnDetails = append(u.ReceiveReturnDetails, p)
	}
}

// ReceiveReturnListResponse represents a list of receive returns
type ReceiveReturnListResponse struct {
	Data []ReceiveReturnResponse `json:"data"`
}

// Transform converts a slice of FirebaseReceiveReturn to a ReceiveReturnListResponse
func (u *ReceiveReturnListResponse) Transform(returns []models.FirebaseReceiveReturn) {
	u.Data = make([]ReceiveReturnResponse, len(returns))
	for i, ret := range returns {
		u.Data[i].Transform(&ret)
	}
}

// ReceiveReturnDetailResponse represents a receive return detail response
type ReceiveReturnDetailResponse struct {
	ID        string                 `json:"id"`
	ProductID string                 `json:"product_id"`
	Qty       uint                   `json:"qty"`
	Code      string                 `json:"code"`
	Product   models.FirebaseProduct `json:"product"`
}

// Transform converts a FirebaseReceiveReturnDetail to a ReceiveReturnDetailResponse
func (u *ReceiveReturnDetailResponse) Transform(pd *models.FirebaseReceiveReturnDetail) {
	u.ID = pd.ID
	u.ProductID = pd.ProductID
	u.Qty = pd.Qty
	u.Code = pd.Code
	u.Product = pd.Product
}
