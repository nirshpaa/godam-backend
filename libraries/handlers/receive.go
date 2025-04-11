package handlers

import (
	"net/http"
)

type receiveHandler struct{}

func NewReceiveHandler() ReceiveHandler {
	return &receiveHandler{}
}

func (h *receiveHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list receives
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get receive
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create receive
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update receive
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete receive
	w.WriteHeader(http.StatusNotImplemented)
}
