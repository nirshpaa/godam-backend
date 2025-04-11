package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// PurchaseReturnResponse : format json response for purchase return
type PurchaseReturnResponse struct {
	ID                    string                         `json:"id"`
	Code                  string                         `json:"code"`
	Date                  time.Time                      `json:"date"`
	Remark                string                         `json:"remark"`
	PurchaseID            string                         `json:"purchase_id"`
	CompanyID             string                         `json:"company_id"`
	BranchID              string                         `json:"branch_id"`
	PurchaseReturnDetails []PurchaseReturnDetailResponse `json:"purchase_return_details"`
}

// Transform from FirebasePurchaseReturn model to Purchase Return response
func (u *PurchaseReturnResponse) Transform(purchaseReturn *models.FirebasePurchaseReturn) *PurchaseReturnResponse {
	u.ID = purchaseReturn.ID
	u.Code = purchaseReturn.Code
	u.Date = purchaseReturn.Date
	u.Remark = purchaseReturn.Remark
	u.PurchaseID = purchaseReturn.PurchaseID
	u.CompanyID = purchaseReturn.CompanyID
	u.BranchID = purchaseReturn.BranchID

	u.PurchaseReturnDetails = make([]PurchaseReturnDetailResponse, len(purchaseReturn.PurchaseReturnDetails))
	for i, d := range purchaseReturn.PurchaseReturnDetails {
		u.PurchaseReturnDetails[i] = *(&PurchaseReturnDetailResponse{}).Transform(&d)
	}
	return u
}

// PurchaseReturnListResponse : format json response for purchase return list
type PurchaseReturnListResponse struct {
	Data []PurchaseReturnResponse `json:"data"`
}

// Transform from FirebasePurchaseReturn model to PurchaseReturn List response
func (u *PurchaseReturnListResponse) Transform(returns []*models.FirebasePurchaseReturn) *PurchaseReturnListResponse {
	u.Data = make([]PurchaseReturnResponse, len(returns))
	for i, ret := range returns {
		u.Data[i] = *(&PurchaseReturnResponse{}).Transform(ret)
	}
	return u
}

// PurchaseReturnDetailResponse : format json response for purchase return detail
type PurchaseReturnDetailResponse struct {
	ID        string                 `json:"id"`
	ProductID string                 `json:"product_id"`
	Qty       uint                   `json:"qty"`
	Code      string                 `json:"code"`
	Product   models.FirebaseProduct `json:"product"`
}

// Transform from FirebasePurchaseReturnDetail model to PurchaseReturnDetail response
func (u *PurchaseReturnDetailResponse) Transform(pd *models.FirebasePurchaseReturnDetail) *PurchaseReturnDetailResponse {
	u.ID = pd.ID
	u.ProductID = pd.ProductID
	u.Qty = pd.Qty
	u.Code = pd.Code
	u.Product = pd.Product
	return u
}
