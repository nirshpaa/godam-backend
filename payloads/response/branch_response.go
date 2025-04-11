package response

import (
	"github.com/nirshpaa/godam-backend/models"
)

// BranchResponse json
type BranchResponse struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
}

// Transform Branch models to Branch response
func (u *BranchResponse) Transform(branch *models.BranchFirebaseModel) {
	u.ID = branch.ID
	u.Code = branch.Code
	u.Name = branch.Name
	u.Type = branch.Type
	u.Address = branch.Address
}
