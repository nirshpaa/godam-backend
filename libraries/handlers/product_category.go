package handlers

import (
	"net/http"
)

type productCategoryHandler struct{}

func NewProductCategoryHandler() ProductCategoryHandler {
	return &productCategoryHandler{}
}

func (h *productCategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement list product categories
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productCategoryHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get product category
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productCategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create product category
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productCategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update product category
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productCategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete product category
	w.WriteHeader(http.StatusNotImplemented)
}
