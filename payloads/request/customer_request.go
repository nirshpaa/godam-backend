package request

import "github.com/nirshpaa/godam-backend/models"

// NewCustomerRequest is json request for new customer and validation
type NewCustomerRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
	CompanyID string `json:"company_id" validate:"required"`
}

// Transform NewCustomerRequest to FirebaseCustomer model
func (u *NewCustomerRequest) Transform() *models.FirebaseCustomer {
	return &models.FirebaseCustomer{
		Name:      u.Name,
		Email:     u.Email,
		Address:   u.Address,
		Phone:     u.Phone,
		CompanyID: u.CompanyID,
	}
}

// CustomerRequest is json request for update customer and validation
type CustomerRequest struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	CompanyID string `json:"company_id"`
}

// Transform CustomerRequest to FirebaseCustomer model
func (u *CustomerRequest) Transform() *models.FirebaseCustomer {
	return &models.FirebaseCustomer{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Address:   u.Address,
		Phone:     u.Phone,
		CompanyID: u.CompanyID,
	}
}
