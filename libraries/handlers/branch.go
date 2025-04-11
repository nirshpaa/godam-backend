package handlers

import (
	"net/http"
)

type branchHandler struct{}

func NewBranchHandler() BranchHandler {
	return &branchHandler{}
}

func (h *branchHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list branches
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *branchHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get branch
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *branchHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create branch
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *branchHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update branch
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *branchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete branch
	w.WriteHeader(http.StatusNotImplemented)
}
