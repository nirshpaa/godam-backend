package handlers

import (
	"net/http"
)

type receiveReturnHandler struct{}

func NewReceiveReturnHandler() ReceiveReturnHandler {
	return &receiveReturnHandler{}
}

func (h *receiveReturnHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list receive returns
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveReturnHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get receive return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveReturnHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create receive return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveReturnHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update receive return
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *receiveReturnHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete receive return
	w.WriteHeader(http.StatusNotImplemented)
}
