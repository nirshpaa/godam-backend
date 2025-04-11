package handlers

import (
	"net/http"
)

type userHandler struct{}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list users
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get user
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create user
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update user
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete user
	w.WriteHeader(http.StatusNotImplemented)
}
