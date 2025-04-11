package handlers

import (
	"net/http"
)

type inventoryHandler struct{}

func NewInventoryHandler() InventoryHandler {
	return &inventoryHandler{}
}

func (h *inventoryHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list inventories
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *inventoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get inventory
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *inventoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create inventory
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *inventoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update inventory
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *inventoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete inventory
	w.WriteHeader(http.StatusNotImplemented)
}
