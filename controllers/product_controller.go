package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/nirshpaa/godam-backend/services"
)

// ProductController handles product-related HTTP requests
type ProductController struct {
	imageRecognitionService *services.ImageRecognitionService
}

// NewProductController creates a new ProductController instance
func NewProductController(imageRecognitionService *services.ImageRecognitionService) *ProductController {
	return &ProductController{
		imageRecognitionService: imageRecognitionService,
	}
}

// ScanProduct handles product scanning requests
func (c *ProductController) ScanProduct(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get file from form
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "No image file provided: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded image
	tempFile, err := os.CreateTemp("", "scan-*.png")
	if err != nil {
		http.Error(w, "Failed to create temporary file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	// Copy the uploaded file to the temporary file
	if _, err := tempFile.ReadFrom(file); err != nil {
		http.Error(w, "Failed to save uploaded file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Process the image
	result, err := c.imageRecognitionService.ProcessImage(tempFile.Name())
	if err != nil {
		http.Error(w, "Failed to process image: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the recognition result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": result.RecognitionSuccess,
		"data":    result.RecognitionData,
	})
}
