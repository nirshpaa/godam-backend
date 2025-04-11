package handlers

import (
	"net/http"
)

type purchaseHandler struct{}

func NewPurchaseHandler() PurchaseHandler {
	return &purchaseHandler{}
}

func (h *purchaseHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list purchases
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get purchase
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create purchase
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update purchase
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *purchaseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete purchase
	w.WriteHeader(http.StatusNotImplemented)
}
