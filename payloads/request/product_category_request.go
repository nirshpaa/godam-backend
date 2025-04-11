package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewProductCategoryRequest : format json request for new product category
type NewProductCategoryRequest struct {
	Name      string                       `json:"name" validate:"required"`
	CompanyID string                       `json:"company_id" validate:"required"`
	Category  models.CategoryFirebaseModel `json:"category" validate:"required"`
}

// Transform converts NewProductCategoryRequest to ProductCategoryFirebaseModel
func (p *NewProductCategoryRequest) Transform() *models.ProductCategoryFirebaseModel {
	return &models.ProductCategoryFirebaseModel{
		Name:      p.Name,
		CompanyID: p.CompanyID,
		Category:  p.Category,
	}
}

// ProductCategoryRequest : format json request for product category
type ProductCategoryRequest struct {
	ID        string                       `json:"id,omitempty" validate:"required"`
	Name      string                       `json:"name,omitempty" validate:"required"`
	CompanyID string                       `json:"company_id,omitempty" validate:"required"`
	Category  models.CategoryFirebaseModel `json:"category,omitempty" validate:"required"`
}

// Transform converts ProductCategoryRequest to ProductCategoryFirebaseModel
func (p *ProductCategoryRequest) Transform() *models.ProductCategoryFirebaseModel {
	return &models.ProductCategoryFirebaseModel{
		ID:        p.ID,
		Name:      p.Name,
		CompanyID: p.CompanyID,
		Category:  p.Category,
	}
}
