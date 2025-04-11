package handlers

import (
	"net/http"
)

type salesmanHandler struct{}

func NewSalesmanHandler() SalesmanHandler {
	return &salesmanHandler{}
}

func (h *salesmanHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list salesmen
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesmanHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get salesman
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesmanHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create salesman
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesmanHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update salesman
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesmanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete salesman
	w.WriteHeader(http.StatusNotImplemented)
}
