package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// RegionResponse json
type RegionResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Transform Region models to Region response
func (u *RegionResponse) Transform(region *models.RegionFirebaseModel) {
	u.ID = region.ID
	u.Name = region.Name
}
