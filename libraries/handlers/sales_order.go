package handlers

import (
	"net/http"
)

type salesOrderHandler struct{}

func NewSalesOrderHandler() SalesOrderHandler {
	return &salesOrderHandler{}
}

func (h *salesOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list sales orders
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get sales order
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create sales order
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update sales order
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete sales order
	w.WriteHeader(http.StatusNotImplemented)
}
