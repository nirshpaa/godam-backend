package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewProductRequest : format json request for new product
type NewProductRequest struct {
	Code              string  `json:"code" validate:"required"`
	Name              string  `json:"name" validate:"required"`
	SalePrice         float64 `json:"price" validate:"required"`
	MinimumStock      float64 `json:"minimum_stock" validate:"required"`
	BrandID           string  `json:"brand" validate:"required"`
	ProductCategoryID string  `json:"product_category" validate:"required"`
}

// Transform NewProductRequest to FirebaseProduct
func (u *NewProductRequest) Transform() *models.FirebaseProduct {
	product := &models.FirebaseProduct{
		Code:              u.Code,
		Name:              u.Name,
		SalePrice:         u.SalePrice,
		BrandID:           u.BrandID,
		ProductCategoryID: u.ProductCategoryID,
	}

	product.MinimumStock = u.MinimumStock

	return product
}

// ProductRequest : format json request for product
type ProductRequest struct {
	ID                string  `json:"id,omitempty"`
	Code              string  `json:"code,omitempty"`
	Name              string  `json:"name,omitempty"`
	SalePrice         float64 `json:"price,omitempty"`
	MinimumStock      float64 `json:"minimum_stock,omitempty"`
	BrandID           string  `json:"brand"`
	ProductCategoryID string  `json:"product_category"`
}

// Transform ProductRequest to FirebaseProduct
func (u *ProductRequest) Transform(product *models.FirebaseProduct) *models.FirebaseProduct {
	if u.Code == product.Code {
		if len(u.Code) > 0 {
			product.Code = u.Code
		}

		if len(u.Name) > 0 {
			product.Name = u.Name
		}

		if u.SalePrice > 0 {
			product.SalePrice = u.SalePrice
		}

		if u.MinimumStock > 0 {
			product.MinimumStock = u.MinimumStock
		}

		if len(u.BrandID) > 0 {
			product.BrandID = u.BrandID
		}

		if len(u.ProductCategoryID) > 0 {
			product.ProductCategoryID = u.ProductCategoryID
		}
	}
	return product
}
