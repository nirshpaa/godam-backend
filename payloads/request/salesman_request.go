package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewSalesmanRequest : format json request for new salesman
type NewSalesmanRequest struct {
	Code      string `json:"code" validate:"required"`
	Name      string `json:"name" validate:"required"`
	CompanyID string `json:"company_id" validate:"required"`
}

// Transform converts NewSalesmanRequest to SalesmanFirebaseModel
func (s *NewSalesmanRequest) Transform() *models.SalesmanFirebaseModel {
	return &models.SalesmanFirebaseModel{
		Code:      s.Code,
		Name:      s.Name,
		CompanyID: s.CompanyID,
	}
}

// SalesmanRequest : format json request for salesman
type SalesmanRequest struct {
	ID        string `json:"id,omitempty" validate:"required"`
	Code      string `json:"code,omitempty" validate:"required"`
	Name      string `json:"name,omitempty" validate:"required"`
	CompanyID string `json:"company_id,omitempty" validate:"required"`
}

// Transform converts SalesmanRequest to SalesmanFirebaseModel
func (s *SalesmanRequest) Transform() *models.SalesmanFirebaseModel {
	return &models.SalesmanFirebaseModel{
		ID:        s.ID,
		Code:      s.Code,
		Name:      s.Name,
		CompanyID: s.CompanyID,
	}
}
