package request

import (
	"strconv"

	"github.com/nirshpaa/godam-backend/models"
)

// NewUserRequest : format json request for new user
type NewUserRequest struct {
	Username   string                     `json:"username" validate:"required"`
	Email      string                     `json:"email" validate:"required"`
	Password   string                     `json:"password" validate:"required"`
	RePassword string                     `json:"re_password" validate:"required"`
	IsActive   bool                       `json:"is_active"`
	Roles      []models.RoleFirebaseModel `json:"roles"`
	RegionID   uint32                     `json:"region,omitempty"`
	BranchID   uint32                     `json:"branch,omitempty"`
}

// Transform converts NewUserRequest to FirebaseUser
func (u *NewUserRequest) Transform() *models.FirebaseUser {
	return &models.FirebaseUser{
		Email:     u.Email,
		Name:      u.Username,
		Role:      u.Roles[0].Name,
		CompanyID: strconv.FormatUint(uint64(u.RegionID), 10),
	}
}

// UserRequest : format json request for user
type UserRequest struct {
	ID         uint64                     `json:"id,omitempty" validate:"required"`
	Username   string                     `json:"username,omitempty" validate:"required"`
	Email      string                     `json:"email,omitempty" validate:"required"`
	Password   string                     `json:"password,omitempty" validate:"required"`
	RePassword string                     `json:"re_password,omitempty" validate:"required"`
	IsActive   bool                       `json:"is_active,omitempty"`
	Roles      []models.RoleFirebaseModel `json:"roles,omitempty"`
	RegionID   uint32                     `json:"region,omitempty"`
	BranchID   uint32                     `json:"branch,omitempty"`
}

// Transform converts UserRequest to FirebaseUser
func (u *UserRequest) Transform() *models.FirebaseUser {
	return &models.FirebaseUser{
		ID:        strconv.FormatUint(u.ID, 10),
		Email:     u.Email,
		Name:      u.Username,
		Role:      u.Roles[0].Name,
		CompanyID: strconv.FormatUint(uint64(u.RegionID), 10),
	}
}
