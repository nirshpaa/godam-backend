package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// UserListResponse : format json response for list of users
type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

// Transform from User model to UserList response
func (u *UserListResponse) Transform(users []*models.FirebaseUser) {
	u.Users = make([]UserResponse, len(users))
	for i, user := range users {
		var userResponse UserResponse
		userResponse.Transform(user)
		u.Users[i] = userResponse
	}
}

// UserResponse : format json response for user
type UserResponse struct {
	ID       string          `json:"id"`
	Username string          `json:"username"`
	IsActive bool            `json:"is_active"`
	Role     string          `json:"role"`
	Company  CompanyResponse `json:"company"`
	Region   *RegionResponse `json:"region,omitempty"`
	Branch   *BranchResponse `json:"branch,omitempty"`
}

// Transform from FirebaseUser model to User response
func (u *UserResponse) Transform(user *models.FirebaseUser) {
	u.ID = user.ID
	u.Username = user.Name
	u.IsActive = true
	u.Role = user.Role
	u.Company = CompanyResponse{ID: user.CompanyID}
}
