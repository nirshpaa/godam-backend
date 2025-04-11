package handlers

import (
	"net/http"
)

type authHandler struct{}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement login
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *authHandler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement register
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *authHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement refresh token
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *authHandler) Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement logout
	w.WriteHeader(http.StatusNotImplemented)
}
