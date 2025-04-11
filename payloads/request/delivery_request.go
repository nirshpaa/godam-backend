package request

import (
	"time"

	"github.com/nirshpaa/godam-backend/models"
)

// NewDeliveryRequest : format json request for new Delivery
type NewDeliveryRequest struct {
	Date            string                     `json:"date" validate:"required"`
	Remark          string                     `json:"remark"`
	DeliveryDetails []NewDeliveryDetailRequest `json:"delivery_details" validate:"required"`
	SalesOrderID    string                     `json:"sales_order" validate:"required"`
}

// Transform NewDeliveryRequest to FirebaseDelivery
func (u *NewDeliveryRequest) Transform() *models.FirebaseDelivery {
	delivery := &models.FirebaseDelivery{
		Date:         time.Now(),
		Remark:       u.Remark,
		SalesOrderID: u.SalesOrderID,
	}

	if u.Date != "" {
		delivery.Date, _ = time.Parse("2006-01-02", u.Date)
	}

	for _, pd := range u.DeliveryDetails {
		delivery.DeliveryDetails = append(delivery.DeliveryDetails, pd.Transform())
	}

	return delivery
}

// NewDeliveryDetailRequest : format json request for Delivery detail
type NewDeliveryDetailRequest struct {
	ProductID string `json:"product" validate:"required"`
	ShelveID  string `json:"shelve" validate:"required"`
}

// Transform NewDeliveryDetailRequest to FirebaseDeliveryDetail
func (u *NewDeliveryDetailRequest) Transform() models.FirebaseDeliveryDetail {
	return models.FirebaseDeliveryDetail{
		Qty:       1,
		ProductID: u.ProductID,
		ShelveID:  u.ShelveID,
	}
}

// DeliveryRequest : format json request for Delivery
type DeliveryRequest struct {
	ID              string                  `json:"id" validate:"required"`
	Date            string                  `json:"date"`
	Remark          string                  `json:"remark"`
	DeliveryDetails []DeliveryDetailRequest `json:"delivery_details"`
	SalesOrderID    string                  `json:"sales_order"`
}

// Transform DeliveryRequest to FirebaseDelivery
func (u *DeliveryRequest) Transform(delivery *models.FirebaseDelivery) *models.FirebaseDelivery {
	if u.ID == delivery.ID {
		if u.Date != "" {
			delivery.Date, _ = time.Parse("2006-01-02", u.Date)
		}
		if u.SalesOrderID != "" {
			delivery.SalesOrderID = u.SalesOrderID
		}
		if u.Remark != "" {
			delivery.Remark = u.Remark
		}

		var details []models.FirebaseDeliveryDetail
		for _, pd := range u.DeliveryDetails {
			details = append(details, pd.Transform())
		}

		delivery.DeliveryDetails = details
	}
	return delivery
}

// DeliveryDetailRequest : format json request for Delivery detail
type DeliveryDetailRequest struct {
	ID        string `json:"id"`
	ProductID string `json:"product"`
	ShelveID  string `json:"shelve"`
}

// Transform DeliveryDetailRequest to FirebaseDeliveryDetail
func (u *DeliveryDetailRequest) Transform() models.FirebaseDeliveryDetail {
	return models.FirebaseDeliveryDetail{
		ID:        u.ID,
		Qty:       1,
		ProductID: u.ProductID,
		ShelveID:  u.ShelveID,
	}
}
