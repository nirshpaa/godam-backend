package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// ProductResponse represents a product response
type ProductResponse struct {
	ID                string  `json:"id"`
	Code              string  `json:"code"`
	Name              string  `json:"name"`
	SalePrice         float64 `json:"price"`
	MinimumStock      float64 `json:"minimum_stock"`
	CompanyID         string  `json:"company_id"`
	BrandID           string  `json:"brand_id"`
	ProductCategoryID string  `json:"product_category_id"`
}

// Transform transforms the model into a response
func (r ProductResponse) Transform(p *models.FirebaseProduct) ProductResponse {
	return ProductResponse{
		Code:              p.Code,
		Name:              p.Name,
		SalePrice:         p.SalePrice,
		MinimumStock:      p.MinimumStock,
		CompanyID:         p.CompanyID,
		BrandID:           p.BrandID,
		ProductCategoryID: p.ProductCategoryID,
	}
}

// ProductListResponse represents a list of products
type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}

// Transform transforms the model into a response
func (r ProductListResponse) Transform(products []models.FirebaseProduct) ProductListResponse {
	for _, p := range products {
		r.Products = append(r.Products, ProductResponse{}.Transform(&p))
	}
	return r
}
