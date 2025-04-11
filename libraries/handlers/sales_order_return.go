package handlers

import (
	"net/http"
)

type salesOrderReturnHandler struct{}

func NewSalesOrderReturnHandler() SalesOrderReturnHandler {
	return &salesOrderReturnHandler{}
}

func (h *salesOrderReturnHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list sales order returns
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderReturnHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get sales order return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderReturnHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create sales order return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderReturnHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update sales order return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *salesOrderReturnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete sales order return
	w.WriteHeader(http.StatusNotImplemented)
}
