package handlers

import (
	"net/http"
)

type roleHandler struct{}

func NewRoleHandler() RoleHandler {
	return &roleHandler{}
}

func (h *roleHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list roles
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *roleHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get role
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *roleHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create role
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *roleHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update role
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *roleHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete role
	w.WriteHeader(http.StatusNotImplemented)
}
