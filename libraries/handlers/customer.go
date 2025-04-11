package handlers

import (
	"net/http"
)

type customerHandler struct{}

func NewCustomerHandler() CustomerHandler {
	return &customerHandler{}
}

func (h *customerHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list customers
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *customerHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get customer
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *customerHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create customer
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *customerHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update customer
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *customerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete customer
	w.WriteHeader(http.StatusNotImplemented)
}
