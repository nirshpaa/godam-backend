package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/nishanpandit/inventory/libraries/api"
	"github.com/nishanpandit/inventory/models"
	"github.com/nishanpandit/inventory/services"
	// "github.com/jacky-htg/inventory/services"
)

// ProductImageHandler processes image uploads for products
type ProductImageHandler struct {
	Db                   *sql.DB
	ImageRecognitionSvc  *services.ImageRecognitionService
	ImageUploadDirectory string
}

// Upload handles product image uploads
func (h *ProductImageHandler) Upload(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received image upload request")
	log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))
	log.Printf("Content-Length: %s", r.Header.Get("Content-Length"))

	// Extract product ID from the URL
	productIDStr := r.URL.Query().Get("product_id")
	if productIDStr == "" {
		api.ResponseError(w, fmt.Errorf("Missing product_id parameter", productIDStr))
		return
	}

	productID, err := strconv.ParseUint(productIDStr, 10, 64)
	if err != nil {
		api.ResponseError(w, fmt.Errorf("Invalid product_id format", err))
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to parse form: %v", err))
		return
	}

	// Get file from form
	file, header, err := r.FormFile("image")
	if err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to get image from form: %v", err))
		return
	}
	defer file.Close()

	// Validate file type
	if !isValidImageType(header.Filename) {
		api.ResponseError(w, fmt.Errorf("Invalid image type. Supported formats: jpg, jpeg, png", err))
		return
	}

	// Create unique filename
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("product_%d_%d%s", productID, time.Now().Unix(), ext)
	filePath := filepath.Join(h.ImageUploadDirectory, fileName)

	// Ensure upload directory exists
	if err := os.MkdirAll(h.ImageUploadDirectory, os.ModePerm); err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to create upload directory: ", err))
		return
	}

	// Create file on server
	dst, err := os.Create(filePath)
	if err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to create file: ", err))
		return
	}
	defer dst.Close()

	// Create a copy of the file for processing
	tempFile, err := os.Create(filepath.Join(h.ImageUploadDirectory, "temp_"+fileName))
	if err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to create temporary file: ", err))
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Create a TeeReader to write to both destinations
	teeReader := io.TeeReader(file, tempFile)

	// Copy the file content to the destination file
	if _, err = io.Copy(dst, teeReader); err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to save file: ", err))
		return
	}

	// Process the image for recognition
	tempFile.Seek(0, 0)
	recognitionResult, err := h.ImageRecognitionSvc.ProcessImage(tempFile.Name())
	if err != nil {
		// Log error but continue, as we still want to save the image
		fmt.Printf("Recognition error: %v\n", err)
	}

	// Parse recognition data to get barcode
	var recognitionData struct {
		Barcode string `json:"barcode"`
	}
	if err := json.Unmarshal([]byte(recognitionResult.RecognitionData), &recognitionData); err != nil {
		fmt.Printf("Error parsing recognition data: %v\n", err)
	}

	// Begin transaction
	ctx := context.Background()
	tx, err := h.Db.BeginTx(ctx, nil)
	if err != nil {
		api.ResponseError(w, fmt.Errorf("Database error:  %v", err))
		return
	}

	// Prepare product model for update
	product := models.Product{
		ID:                   productID,
		ImageURL:             fileName,
		BarcodeValue:         recognitionData.Barcode,
		ImageRecognitionData: recognitionResult.RecognitionData,
	}

	// Update the product record
	if err := product.UpdateImage(ctx, tx); err != nil {
		tx.Rollback()
		api.ResponseError(w, fmt.Errorf("Failed to update product: %v", err))
		return
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		api.ResponseError(w, fmt.Errorf("Failed to commit transaction: %v", err))
		return
	}

	// Return success response
	response := map[string]interface{}{
		"success":   true,
		"image_url": fileName,
		"recognition_result": map[string]interface{}{
			"barcode_detected": recognitionData.Barcode != "",
			"barcode_value":    recognitionData.Barcode,
			"cnn_recognition":  recognitionResult.RecognitionData != "",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ScanBarcode endpoint for scanning barcodes without attaching to products
func (h *ProductImageHandler) ScanBarcode(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting ScanBarcode handler")
	log.Printf("ImageRecognitionSvc state: %+v", h.ImageRecognitionSvc)

	// Log request details
	log.Printf("Request Method: %s", r.Method)
	log.Printf("Request URL: %s", r.URL.String())
	log.Printf("Request headers: %+v", r.Header)
	log.Printf("Content-Type: %s", r.Header.Get("Content-Type"))
	log.Printf("Content-Length: %s", r.Header.Get("Content-Length"))

	// Check if the request is multipart
	if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
		log.Printf("Invalid Content-Type: %s", r.Header.Get("Content-Type"))
		http.Error(w, "Content-Type must be multipart/form-data", http.StatusBadRequest)
		return
	}

	// Parse multipart form with a larger max memory size
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		log.Printf("Error parsing multipart form: %v", err)
		http.Error(w, fmt.Sprintf("Failed to parse form: %v", err), http.StatusBadRequest)
		return
	}

	// Log form data
	log.Printf("Form values: %+v", r.Form)
	if r.MultipartForm != nil {
		log.Printf("Form files: %+v", r.MultipartForm.File)
	}

	// Get the file from the form
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Printf("Error getting image file: %v", err)
		http.Error(w, fmt.Sprintf("Failed to get image file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	log.Printf("Received file: %s, size: %d", handler.Filename, handler.Size)

	// Get the file extension from the uploaded file
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".png" // Default to PNG if no extension
	}

	// Create a temporary file to store the uploaded image with the correct extension
	tempFile, err := os.CreateTemp("", "upload-*"+ext)
	if err != nil {
		log.Printf("Error creating temp file: %v", err)
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	// Copy the uploaded file to the temporary file
	if _, err := io.Copy(tempFile, file); err != nil {
		log.Printf("Error copying file: %v", err)
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	// Process the image
	result, err := h.ImageRecognitionSvc.ProcessImage(tempFile.Name())
	if err != nil {
		log.Printf("Error processing image: %v", err)
		http.Error(w, fmt.Sprintf("Failed to process image: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the result
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// isValidImageType checks if the file has a valid image extension
func isValidImageType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	validExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	return validExts[ext]
}
