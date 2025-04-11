package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewRegionRequest : format json request for new region
type NewRegionRequest struct {
	Name string `json:"name" validate:"required"`
}

// Transform converts NewRegionRequest to RegionFirebaseModel
func (r *NewRegionRequest) Transform() *models.RegionFirebaseModel {
	return &models.RegionFirebaseModel{
		Name: r.Name,
	}
}

// RegionRequest : format json request for region
type RegionRequest struct {
	ID   string `json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}

// Transform converts RegionRequest to RegionFirebaseModel
func (r *RegionRequest) Transform() *models.RegionFirebaseModel {
	return &models.RegionFirebaseModel{
		ID:   r.ID,
		Name: r.Name,
	}
}
