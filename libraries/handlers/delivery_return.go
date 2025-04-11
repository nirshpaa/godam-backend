package handlers

import (
	"net/http"
)

type deliveryReturnHandler struct{}

func NewDeliveryReturnHandler() DeliveryReturnHandler {
	return &deliveryReturnHandler{}
}

func (h *deliveryReturnHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list delivery returns
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryReturnHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get delivery return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryReturnHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create delivery return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryReturnHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update delivery return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryReturnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete delivery return
	w.WriteHeader(http.StatusNotImplemented)
}
