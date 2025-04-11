package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// BrandResponse json
type BrandResponse struct {
	ID        string          `json:"id"`
	CompanyID string          `json:"company_id"`
	Company   CompanyResponse `json:"company"`
	Code      string          `json:"code"`
	Name      string          `json:"name"`
}

// Transform Brand models to Brand response
func (u *BrandResponse) Transform(brand *models.BrandFirebaseModel) {
	u.ID = brand.ID
	u.Code = brand.Code
	u.Name = brand.Name
	u.CompanyID = brand.CompanyID
}
