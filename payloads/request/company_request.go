package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// CompanyRequest represents a company request
type CompanyRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

// Transform transforms a CompanyRequest to FirebaseCompany
func (r *CompanyRequest) Transform(company *models.FirebaseCompany) *models.FirebaseCompany {
	company.ID = r.ID
	company.Name = r.Name
	company.Address = r.Address
	company.Phone = r.Phone
	company.Email = r.Email
	company.Description = r.Description
	return company
}

// NewCompanyRequest represents a new company request
type NewCompanyRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

// Transform transforms a NewCompanyRequest to FirebaseCompany
func (r *NewCompanyRequest) Transform() *models.FirebaseCompany {
	return &models.FirebaseCompany{
		Name:        r.Name,
		Address:     r.Address,
		Phone:       r.Phone,
		Email:       r.Email,
		Description: r.Description,
	}
}
