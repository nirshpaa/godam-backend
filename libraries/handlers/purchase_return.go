package handlers

import (
	"net/http"
)

type purchaseReturnHandler struct{}

func NewPurchaseReturnHandler() PurchaseReturnHandler {
	return &purchaseReturnHandler{}
}

func (h *purchaseReturnHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list purchase returns
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseReturnHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get purchase return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseReturnHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create purchase return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseReturnHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update purchase return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseReturnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete purchase return
	w.WriteHeader(http.StatusNotImplemented)
}
