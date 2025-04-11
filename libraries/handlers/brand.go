package handlers

import (
	"net/http"
)

type brandHandler struct{}

func NewBrandHandler() BrandHandler {
	return &brandHandler{}
}

func (h *brandHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list brands
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *brandHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get brand
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *brandHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create brand
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *brandHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update brand
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *brandHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete brand
	w.WriteHeader(http.StatusNotImplemented)
}
