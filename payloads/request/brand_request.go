package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewBrandRequest : format json request for new brand
type NewBrandRequest struct {
	Code      string `json:"code" validate:"required"`
	Name      string `json:"name" validate:"required"`
	CompanyID string `json:"company_id" validate:"required"`
}

// Transform converts NewBrandRequest to BrandFirebaseModel
func (b *NewBrandRequest) Transform() *models.BrandFirebaseModel {
	return &models.BrandFirebaseModel{
		Code:      b.Code,
		Name:      b.Name,
		CompanyID: b.CompanyID,
	}
}

// BrandRequest : format json request for brand
type BrandRequest struct {
	ID        string `json:"id,omitempty" validate:"required"`
	Code      string `json:"code,omitempty" validate:"required"`
	Name      string `json:"name,omitempty" validate:"required"`
	CompanyID string `json:"company_id,omitempty" validate:"required"`
}

// Transform converts BrandRequest to BrandFirebaseModel
func (b *BrandRequest) Transform() *models.BrandFirebaseModel {
	return &models.BrandFirebaseModel{
		ID:        b.ID,
		Code:      b.Code,
		Name:      b.Name,
		CompanyID: b.CompanyID,
	}
}
