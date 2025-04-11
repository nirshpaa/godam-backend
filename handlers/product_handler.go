package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/interfaces"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/services"
)

// ProductHandler handles product-related HTTP requests
type ProductHandler struct {
	productModel         *models.ProductFirebase
	fileStorage          interfaces.FileStorage
	imageRecognition     interfaces.ImageRecognition
	imageTrainingService *services.ImageTrainingService
}

// NewProductHandler creates a new product handler
func NewProductHandler(
	productModel *models.ProductFirebase,
	fileStorage interfaces.FileStorage,
	imageRecognition interfaces.ImageRecognition,
	imageTrainingService *services.ImageTrainingService,
) *ProductHandler {
	return &ProductHandler{
		productModel:         productModel,
		fileStorage:          fileStorage,
		imageRecognition:     imageRecognition,
		imageTrainingService: imageTrainingService,
	}
}

// List handles GET /products
func (h *ProductHandler) List(c *gin.Context) {
	products, err := h.productModel.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// Get handles GET /products/:code
func (h *ProductHandler) Get(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	product, err := h.productModel.Get(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Create handles POST /products
func (h *ProductHandler) Create(c *gin.Context) {
	var product models.FirebaseProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for duplicate code before creating
	existingProduct, err := h.productModel.Get(c.Request.Context(), product.Code)
	if err == nil && existingProduct != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("product with code %s already exists", product.Code)})
		return
	}

	// Create the product
	id, err := h.productModel.Create(c.Request.Context(), &product, h.fileStorage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// Update handles PUT /products/:code
func (h *ProductHandler) Update(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code is required"})
		return
	}

	var product models.FirebaseProduct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure the product code matches the URL parameter
	product.Code = code

	// Check if the product exists
	existingProduct, err := h.productModel.Get(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// If no new image URL is provided, keep the existing one
	if product.ImageURL == "" {
		product.ImageURL = existingProduct.ImageURL
	}

	// Update the product
	if err := h.productModel.Update(c.Request.Context(), code, &product, h.fileStorage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the updated product to return
	updatedProduct, err := h.productModel.Get(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get updated product"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// Delete handles DELETE /products/:code
func (h *ProductHandler) Delete(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product code is required"})
		return
	}

	// Check if the product exists
	_, err := h.productModel.Get(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Delete the product
	err = h.productModel.Delete(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// FindByBarcode handles GET /products/barcode/:barcode
func (h *ProductHandler) FindByBarcode(c *gin.Context) {
	barcode := c.Param("barcode")
	product, err := h.productModel.FindByBarcode(c.Request.Context(), barcode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// FindByCompany handles GET /products/company/:companyId
func (h *ProductHandler) FindByCompany(c *gin.Context) {
	companyID := c.Param("companyId")
	products, err := h.productModel.FindByCompany(c.Request.Context(), companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// ProcessImage handles image processing for product recognition
func (h *ProductHandler) ProcessImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the uploaded file temporarily
	tempPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer os.Remove(tempPath)

	// Process the image
	result, err := h.productModel.ProcessImage(tempPath, h.imageRecognition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result.Data)
}

// UpdateImage handles PUT /products/:code/image
func (h *ProductHandler) UpdateImage(c *gin.Context) {
	code := c.Param("code")
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
		return
	}

	// Save the uploaded file temporarily
	imagePath := "/tmp/" + file.Filename
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	result, err := h.productModel.ProcessImage(imagePath, h.imageRecognition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the image URL from the form data
	imageURL := c.PostForm("image_url")
	if imageURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image URL is required"})
		return
	}

	err = h.productModel.UpdateImage(c.Request.Context(), code, imageURL, result.Data, result.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}

// ScanProduct handles POST /products/scan
func (h *ProductHandler) ScanProduct(c *gin.Context) {
	// Parse multipart form with a larger size limit
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil { // 32MB max
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Failed to parse form: %v", err),
			"data":    nil,
		})
		return
	}

	// Get the file from the form
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   fmt.Sprintf("No image file provided: %v", err),
			"data":    nil,
		})
		return
	}
	defer file.Close()

	// Validate file type
	contentType := fileHeader.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid file type. Only images are allowed",
			"data":    nil,
		})
		return
	}

	// Create a temporary file with the correct extension
	ext := filepath.Ext(fileHeader.Filename)
	if ext == "" {
		ext = ".jpg" // Default extension
	}
	tempFile, err := os.CreateTemp("", "scan-*"+ext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Failed to create temporary file: %v", err),
			"data":    nil,
		})
		return
	}
	defer os.Remove(tempFile.Name())

	// Save the uploaded file
	if _, err := io.Copy(tempFile, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Failed to save uploaded file: %v", err),
			"data":    nil,
		})
		return
	}

	// Close the file to ensure all data is written
	if err := tempFile.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Failed to close temporary file: %v", err),
			"data":    nil,
		})
		return
	}

	// Process the image
	result, err := h.productModel.ProcessImage(tempFile.Name(), h.imageRecognition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   fmt.Sprintf("Failed to process image: %v", err),
			"data":    nil,
		})
		return
	}

	// Log the recognition result
	fmt.Printf("Recognition result: %+v\n", result)

	response := gin.H{
		"success": true,
		"error":   nil,
		"data":    result.Data,
	}

	// Log the response
	fmt.Printf("Sending response: %+v\n", response)

	c.JSON(http.StatusOK, response)
}

// UploadImage handles POST /products/upload
func (h *ProductHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate a unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	path := "products/" + filename

	// Save the file directly
	imageURL, err := h.fileStorage.SaveFile(file, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": imageURL})
}

// Upload handles POST /products/upload
func (h *ProductHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
		return
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	imagePath := filepath.Join("assets", "products", filename)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(imagePath), 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create image directory"})
		return
	}

	// Save the uploaded file
	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Return the image URL
	c.JSON(http.StatusOK, gin.H{
		"success":   true,
		"image_url": "assets/products/" + filename,
	})
}
