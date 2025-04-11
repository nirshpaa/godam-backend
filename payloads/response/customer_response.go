package response

import "github.com/nirshpaa/godam-backend/models"

// CustomerResponse is json response for customer
type CustomerResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	CompanyID string `json:"company_id"`
}

// Transform FirebaseCustomer to CustomerResponse
func (r *CustomerResponse) Transform(c *models.FirebaseCustomer) *CustomerResponse {
	r.ID = c.ID
	r.Name = c.Name
	r.Email = c.Email
	r.Address = c.Address
	r.Phone = c.Phone
	r.CompanyID = c.CompanyID
	return r
}

// CustomerListResponse is json response for list of customers
type CustomerListResponse struct {
	Customers []CustomerResponse `json:"customers"`
}

// Transform []FirebaseCustomer to CustomerListResponse
func (r *CustomerListResponse) Transform(customers []*models.FirebaseCustomer) *CustomerListResponse {
	r.Customers = make([]CustomerResponse, len(customers))
	for i, c := range customers {
		r.Customers[i] = *(&CustomerResponse{}).Transform(c)
	}
	return r
}
