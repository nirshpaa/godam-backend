package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// CompanyResponse represents a company response
type CompanyResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

// CompanyListResponse represents a list of companies response
type CompanyListResponse struct {
	Companies []CompanyResponse `json:"companies"`
}

// Transform transforms a FirebaseCompany to CompanyResponse
func (r *CompanyResponse) Transform(company *models.FirebaseCompany) *CompanyResponse {
	r.ID = company.ID
	r.Name = company.Name
	r.Address = company.Address
	r.Phone = company.Phone
	r.Email = company.Email
	r.Description = company.Description
	return r
}

// Transform transforms a slice of FirebaseCompany to CompanyListResponse
func (r *CompanyListResponse) Transform(companies []models.FirebaseCompany) *CompanyListResponse {
	r.Companies = make([]CompanyResponse, len(companies))
	for i, company := range companies {
		r.Companies[i] = *(&CompanyResponse{}).Transform(&company)
	}
	return r
}
