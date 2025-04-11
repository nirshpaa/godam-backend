package request

import (
	"github.com/nirshpaa/godam-backend/models"
)

// NewShelveRequest : format json request for new shelve
type NewShelveRequest struct {
	Name string `json:"name" validate:"required"`
}

// Transform converts NewShelveRequest to ShelveFirebaseModel
func (s *NewShelveRequest) Transform() *models.ShelveFirebaseModel {
	return &models.ShelveFirebaseModel{
		Name: s.Name,
	}
}

// ShelveRequest : format json request for shelve
type ShelveRequest struct {
	ID   string `json:"id,omitempty" validate:"required"`
	Name string `json:"name,omitempty" validate:"required"`
}

// Transform converts ShelveRequest to ShelveFirebaseModel
func (s *ShelveRequest) Transform() *models.ShelveFirebaseModel {
	return &models.ShelveFirebaseModel{
		ID:   s.ID,
		Name: s.Name,
	}
}
