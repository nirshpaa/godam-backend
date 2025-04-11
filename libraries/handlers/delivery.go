package handlers

import (
	"net/http"
)

type deliveryHandler struct{}

func NewDeliveryHandler() DeliveryHandler {
	return &deliveryHandler{}
}

func (h *deliveryHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list deliveries
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get delivery
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create delivery
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update delivery
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *deliveryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete delivery
	w.WriteHeader(http.StatusNotImplemented)
}
