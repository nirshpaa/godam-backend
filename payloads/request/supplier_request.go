package request

import "github.com/nirshpaa/godam-backend/models"

// NewSupplierRequest is json request for new supplier and validation
type NewSupplierRequest struct {
	Code    string `json:"code" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address,omitempty"`
}

// Transform NewSupplierRequest to Supplier model
func (u *NewSupplierRequest) Transform() models.Supplier {
	var c models.Supplier
	c.Code = u.Code
	c.Name = u.Name
	c.Address = u.Address
	return c
}

// SupplierRequest is json request for update supplier and validation
type SupplierRequest struct {
	ID      string `json:"id" validate:"required"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Transform SupplierRequest to Supplier model
func (u *SupplierRequest) Transform(c *models.Supplier) *models.Supplier {
	if c.ID == u.ID {
		if len(u.Name) > 0 {
			c.Name = u.Name
		}
		if len(u.Address) > 0 {
			c.Address = u.Address
		}
	}
	return c
}
