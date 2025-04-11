package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewRoleRequest : format json request for new role
type NewRoleRequest struct {
	Name string `json:"name" validate:"required"`
}

// Transform converts NewRoleRequest to RoleFirebaseModel
func (r *NewRoleRequest) Transform() *models.RoleFirebaseModel {
	return &models.RoleFirebaseModel{
		Name: r.Name,
	}
}

// RoleRequest : format json request for role
type RoleRequest struct {
	ID   string `json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}

// Transform converts RoleRequest to RoleFirebaseModel
func (r *RoleRequest) Transform() *models.RoleFirebaseModel {
	return &models.RoleFirebaseModel{
		ID:   r.ID,
		Name: r.Name,
	}
}
