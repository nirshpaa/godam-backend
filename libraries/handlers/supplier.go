package handlers

import (
	"net/http"
)

type supplierHandler struct{}

func NewSupplierHandler() SupplierHandler {
	return &supplierHandler{}
}

func (h *supplierHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list suppliers
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *supplierHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get supplier
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *supplierHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create supplier
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *supplierHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update supplier
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *supplierHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete supplier
	w.WriteHeader(http.StatusNotImplemented)
}
