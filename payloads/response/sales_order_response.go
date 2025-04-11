package response

import (
	"github.com/nirshpaa/godam-backend/types"
)

// SalesOrderResponse represents a sales order response
type SalesOrderResponse struct {
	ID                string                   `json:"id"`
	Code              string                   `json:"code"`
	Date              string                   `json:"date"`
	CustomerID        string                   `json:"customer_id"`
	SalesmanID        string                   `json:"salesman_id"`
	TotalAmount       float64                  `json:"total_amount"`
	Discount          float64                  `json:"discount"`
	AdditionalDisc    float64                  `json:"additional_disc"`
	Status            string                   `json:"status"`
	SalesOrderDetails []types.SalesOrderDetail `json:"sales_order_details"`
}

func NewSalesOrderResponse(order *types.SalesOrder) *SalesOrderResponse {
	return &SalesOrderResponse{
		ID:                order.ID,
		Code:              order.Code,
		Date:              order.Date,
		CustomerID:        order.CustomerID,
		SalesmanID:        order.SalesmanID,
		TotalAmount:       order.TotalAmount,
		Discount:          order.Discount,
		AdditionalDisc:    order.AdditionalDisc,
		Status:            order.Status,
		SalesOrderDetails: order.SalesOrderDetails,
	}
}

// SalesOrderListResponse represents a list of sales orders
type SalesOrderListResponse struct {
	Data []SalesOrderResponse `json:"data"`
}

// Transform converts a slice of FirebaseSalesOrder to a SalesOrderListResponse
func (u *SalesOrderListResponse) Transform(orders []types.SalesOrder) {
	u.Data = make([]SalesOrderResponse, len(orders))
	for i, order := range orders {
		u.Data[i] = *NewSalesOrderResponse(&order)
	}
}

// SalesOrderDetailResponse represents a sales order detail response
type SalesOrderDetailResponse struct {
	ProductID   string  `json:"product_id"`
	ProductCode string  `json:"product_code"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
	Discount    float64 `json:"discount"`
}

// Transform converts a SalesOrderDetail to a SalesOrderDetailResponse
func (u *SalesOrderDetailResponse) Transform(sod *types.SalesOrderDetail) {
	u.ProductID = sod.ProductID
	u.ProductCode = sod.ProductCode
	u.Quantity = sod.Quantity
	u.UnitPrice = sod.UnitPrice
	u.TotalPrice = sod.TotalPrice
	u.Discount = sod.Discount
}
