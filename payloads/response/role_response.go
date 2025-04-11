package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// RoleResponse json
type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Transform Role models to Role response
func (u *RoleResponse) Transform(role *models.RoleFirebaseModel) {
	u.ID = role.ID
	u.Name = role.Name
}
