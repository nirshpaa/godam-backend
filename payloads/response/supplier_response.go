package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// SupplierResponse json
type SupplierResponse struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Transform Supplier models to Supplier response
func (u *SupplierResponse) Transform(c interface{}) {
	switch v := c.(type) {
	case *models.Supplier:
		u.ID = v.ID
		u.Code = v.Code
		u.Name = v.Name
		u.Address = v.Address
	}
}
