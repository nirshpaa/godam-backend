package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// AccessResponse json
type AccessResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Transform Access models to Access response
func (u *AccessResponse) Transform(access *models.AccessFirebaseModel) {
	u.ID = access.ID
	u.Name = access.Name
}
