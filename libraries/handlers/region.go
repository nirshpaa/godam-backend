package handlers

import (
	"net/http"
)

type regionHandler struct{}

func NewRegionHandler() RegionHandler {
	return &regionHandler{}
}

func (h *regionHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list regions
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *regionHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get region
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *regionHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create region
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *regionHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update region
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *regionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete region
	w.WriteHeader(http.StatusNotImplemented)
}
