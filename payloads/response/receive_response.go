package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// ReceiveResponse represents a receive response
type ReceiveResponse struct {
	ID             string                  `json:"id"`
	Code           string                  `json:"code"`
	Date           time.Time               `json:"date"`
	Remark         string                  `json:"remark"`
	PurchaseID     string                  `json:"purchase_id"`
	CompanyID      string                  `json:"company_id"`
	BranchID       string                  `json:"branch_id"`
	ReceiveDetails []ReceiveDetailResponse `json:"receive_details"`
}

// Transform converts a FirebaseReceive to a ReceiveResponse
func (u *ReceiveResponse) Transform(receive *models.FirebaseReceive) {
	u.ID = receive.ID
	u.Code = receive.Code
	u.Date = receive.Date
	u.Remark = receive.Remark
	u.PurchaseID = receive.PurchaseID
	u.CompanyID = receive.CompanyID
	u.BranchID = receive.BranchID

	for _, d := range receive.ReceiveDetails {
		var p ReceiveDetailResponse
		p.Transform(&d)
		u.ReceiveDetails = append(u.ReceiveDetails, p)
	}
}

// ReceiveListResponse represents a list of receives
type ReceiveListResponse struct {
	Data []ReceiveResponse `json:"data"`
}

// Transform converts a slice of FirebaseReceive to a ReceiveListResponse
func (u *ReceiveListResponse) Transform(receives []models.FirebaseReceive) {
	u.Data = make([]ReceiveResponse, len(receives))
	for i, receive := range receives {
		u.Data[i].Transform(&receive)
	}
}

// ReceiveDetailResponse represents a receive detail response
type ReceiveDetailResponse struct {
	ID        string                     `json:"id"`
	ProductID string                     `json:"product_id"`
	Qty       uint                       `json:"qty"`
	Code      string                     `json:"code"`
	ShelveID  string                     `json:"shelve_id"`
	Product   models.FirebaseProduct     `json:"product"`
	Shelve    models.ShelveFirebaseModel `json:"shelve"`
}

// Transform converts a FirebaseReceiveDetail to a ReceiveDetailResponse
func (u *ReceiveDetailResponse) Transform(pd *models.FirebaseReceiveDetail) {
	u.ID = pd.ID
	u.Qty = pd.Qty
	u.ProductID = pd.ProductID
	u.Code = pd.Code
	u.ShelveID = pd.ShelveID
	u.Product = pd.Product
	u.Shelve = pd.Shelve
}
