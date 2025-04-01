package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/models"
)

type ProductHandler struct {
	productModel *models.ProductFirebase
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		productModel: models.NewProductFirebase(),
	}
}

// ListProducts handles GET /api/products
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.productModel.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct handles GET /api/products/:id
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productModel.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct handles POST /api/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.FirebaseProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.productModel.Create(c.Request.Context(), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.ID = id
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct handles PUT /api/products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.FirebaseProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productModel.Update(c.Request.Context(), id, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.ID = id
	c.JSON(http.StatusOK, product)
}

// DeleteProduct handles DELETE /api/products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.productModel.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// FindByBarcode handles GET /api/products/barcode/:barcode
func (h *ProductHandler) FindByBarcode(c *gin.Context) {
	barcode := c.Param("barcode")
	product, err := h.productModel.FindByBarcode(c.Request.Context(), barcode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// FindByCompany handles GET /api/products/company/:companyId
func (h *ProductHandler) FindByCompany(c *gin.Context) {
	companyID := c.Param("companyId")
	products, err := h.productModel.FindByCompany(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// UpdateProductImage handles PUT /api/products/:id/image
func (h *ProductHandler) UpdateProductImage(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		ImageURL             string `json:"image_url"`
		BarcodeValue         string `json:"barcode_value"`
		ImageRecognitionData string `json:"image_recognition_data"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productModel.UpdateImage(c.Request.Context(), id, request.ImageURL, request.BarcodeValue, request.ImageRecognitionData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product image updated successfully"})
}
