package response

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// DeliveryResponse : format json response for Delivery
type DeliveryResponse struct {
	ID              string                   `json:"id"`
	Code            string                   `json:"code"`
	Date            time.Time                `json:"date"`
	Remark          string                   `json:"remark"`
	SalesOrderID    string                   `json:"sales_order_id"`
	CompanyID       string                   `json:"company_id"`
	BranchID        string                   `json:"branch_id"`
	DeliveryDetails []DeliveryDetailResponse `json:"delivery_details"`
}

// Transform from FirebaseDelivery model to Delivery response
func (u *DeliveryResponse) Transform(delivery *models.FirebaseDelivery) {
	u.ID = delivery.ID
	u.Code = delivery.Code
	u.Date = delivery.Date
	u.Remark = delivery.Remark
	u.SalesOrderID = delivery.SalesOrderID
	u.CompanyID = delivery.CompanyID
	u.BranchID = delivery.BranchID

	for _, d := range delivery.DeliveryDetails {
		var p DeliveryDetailResponse
		p.Transform(&d)
		u.DeliveryDetails = append(u.DeliveryDetails, p)
	}
}

// DeliveryListResponse : format json response for Delivery list
type DeliveryListResponse struct {
	ID           string    `json:"id"`
	Code         string    `json:"code"`
	Date         time.Time `json:"date"`
	Remark       string    `json:"remark"`
	SalesOrderID string    `json:"sales_order_id"`
	CompanyID    string    `json:"company_id"`
	BranchID     string    `json:"branch_id"`
}

// Transform from FirebaseDelivery model to Delivery List response
func (u *DeliveryListResponse) Transform(delivery *models.FirebaseDelivery) {
	u.ID = delivery.ID
	u.Code = delivery.Code
	u.Date = delivery.Date
	u.Remark = delivery.Remark
	u.SalesOrderID = delivery.SalesOrderID
	u.CompanyID = delivery.CompanyID
	u.BranchID = delivery.BranchID
}

// DeliveryDetailResponse : format json response for Delivery detail
type DeliveryDetailResponse struct {
	ID        string          `json:"id"`
	Qty       uint            `json:"qty"`
	ProductID string          `json:"product_id"`
	Code      string          `json:"code"`
	ShelveID  string          `json:"shelve_id"`
	Product   ProductResponse `json:"product"`
}

// Transform from FirebaseDeliveryDetail model to DeliveryDetail response
func (u *DeliveryDetailResponse) Transform(pd *models.FirebaseDeliveryDetail) {
	u.ID = pd.ID
	u.Qty = pd.Qty
	u.ProductID = pd.ProductID
	u.Code = pd.Code
	u.ShelveID = pd.ShelveID
	u.Product.Transform(&pd.Product)
}
