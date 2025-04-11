package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// ShelveResponse json
type ShelveResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Transform Shelve models to Shelve response
func (u *ShelveResponse) Transform(shelve *models.ShelveFirebaseModel) {
	u.ID = shelve.ID
	u.Name = shelve.Name
}
