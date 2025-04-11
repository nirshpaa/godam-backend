package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// ProductCategoryResponse json
type ProductCategoryResponse struct {
	ID        string           `json:"id"`
	CompanyID string           `json:"company_id"`
	Company   CompanyResponse  `json:"company"`
	Name      string           `json:"name"`
	Category  CategoryResponse `json:"category"`
}

// Transform ProductCategory models to ProductCategory response
func (u *ProductCategoryResponse) Transform(category *models.ProductCategoryFirebaseModel) {
	u.ID = category.ID
	u.Name = category.Name
	u.CompanyID = category.CompanyID
	u.Category.Transform(&category.Category)
}

// CategoryResponse json
type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Transform Category models to Category response
func (u *CategoryResponse) Transform(category *models.CategoryFirebaseModel) {
	u.ID = category.ID
	u.Name = category.Name
}
