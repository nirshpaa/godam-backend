package controllers

import (
	"net/http"

	fb "firebase.google.com/go/v4"
	"github.com/nirshpaa/godam-backend/libraries/api"
	"github.com/nirshpaa/godam-backend/libraries/firebase"
)

// HealthController : struct for set Health Dependency Injection
type HealthController struct {
	app *fb.App
}

// NewHealthController : constructor for HealthController
func NewHealthController(app *fb.App) *HealthController {
	return &HealthController{
		app: app,
	}
}

// Check : http handler for health checking
func (h *HealthController) Check(w http.ResponseWriter, r *http.Request) {
	var health struct {
		Status string `json:"status"`
	}

	// Get Firestore client
	client := firebase.GetFirestore()
	if client == nil {
		api.ResponseError(w, api.ErrInternal(nil, "Firestore client not initialized"))
		return
	}

	// Check if Firestore is ready by attempting to get a document
	_, err := client.Collection("health").Doc("check").Get(r.Context())
	if err != nil {
		api.ResponseError(w, api.ErrInternal(err, "Firestore not ready"))
		return
	}

	health.Status = "ok"
	api.ResponseOK(w, health, http.StatusOK)
}
