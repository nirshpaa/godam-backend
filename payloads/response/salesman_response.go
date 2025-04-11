package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// SalesmanResponse json
type SalesmanResponse struct {
	ID        string          `json:"id"`
	CompanyID string          `json:"company_id"`
	Company   CompanyResponse `json:"company"`
	Code      string          `json:"code"`
	Name      string          `json:"name"`
}

// Transform Salesman models to Salesman response
func (u *SalesmanResponse) Transform(salesman *models.SalesmanFirebaseModel) {
	u.ID = salesman.ID
	u.Code = salesman.Code
	u.Name = salesman.Name
	u.CompanyID = salesman.CompanyID
}
