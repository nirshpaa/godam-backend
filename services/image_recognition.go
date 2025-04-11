package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/disintegration/imaging"
	"github.com/nirshpaa/godam-backend/types"
)

// ImageRecognitionService handles image processing and recognition
type ImageRecognitionService struct {
	Client             *firestore.Client
	BarcodeAPIEndpoint string
	CNNAPIEndpoint     string
}

// Product represents a product in the database
type Product struct {
	ID           string
	Name         string
	BarcodeValue string
}

// RecognitionResult represents the result of image processing
type RecognitionResult struct {
	RecognitionSuccess bool   `json:"recognition_success"`
	RecognitionData    string `json:"recognition_data"`
}

// NewImageRecognitionService creates a new image recognition service
func NewImageRecognitionService(client *firestore.Client, barcodeEndpoint, cnnEndpoint string) *ImageRecognitionService {
	return &ImageRecognitionService{
		Client:             client,
		BarcodeAPIEndpoint: barcodeEndpoint,
		CNNAPIEndpoint:     cnnEndpoint,
	}
}

// ProcessImage processes an image for product recognition
func (s *ImageRecognitionService) ProcessImage(imagePath string) (*types.ImageRecognitionResult, error) {
	// Try to scan barcode first
	barcode, err := s.scanBarcode(imagePath)
	if err == nil && barcode != "" {
		// If barcode is found, try to find the product
		productID, err := s.FindProductByBarcode(barcode)
		if err == nil && productID != "" {
			return &types.ImageRecognitionResult{
				Success: true,
				Data:    productID,
			}, nil
		}
	}

	// If barcode scanning fails, try CNN recognition
	recognitionData, err := s.recognizeImage(imagePath)
	if err != nil {
		return &types.ImageRecognitionResult{
			Success: false,
			Data:    "",
		}, fmt.Errorf("failed to recognize image: %v", err)
	}

	// Parse the recognition data
	var data struct {
		Class      string  `json:"class"`
		Confidence float64 `json:"confidence"`
	}
	if err := json.Unmarshal([]byte(recognitionData), &data); err != nil {
		return &types.ImageRecognitionResult{
			Success: false,
			Data:    recognitionData,
		}, nil
	}

	// If confidence is above 0.5, try to find the product
	if data.Confidence > 0.5 {
		productID, err := s.FindProductByName(data.Class)
		if err == nil && productID != "" {
			return &types.ImageRecognitionResult{
				Success: true,
				Data:    productID,
			}, nil
		}
	}

	// If no product found or confidence is low, return create product action
	createProductData := map[string]interface{}{
		"action":         "create_product",
		"suggested_name": data.Class,
		"confidence":     data.Confidence,
	}
	jsonData, _ := json.Marshal(createProductData)
	return &types.ImageRecognitionResult{
		Success: false,
		Data:    string(jsonData),
	}, nil
}

// scanBarcode attempts to scan a barcode from an image
func (s *ImageRecognitionService) scanBarcode(imagePath string) (string, error) {
	// Implementation of barcode scanning
	return "", fmt.Errorf("barcode not found")
}

// recognizeImage uses CNN to recognize the image
func (s *ImageRecognitionService) recognizeImage(imagePath string) (string, error) {
	result, err := s.recognizeWithCNN(imagePath)
	if err != nil {
		return "", err
	}
	return result.RecognitionData, nil
}

// ScanBarcode attempts to detect and read a barcode from an image
func (s *ImageRecognitionService) ScanBarcode(filePath string) (*RecognitionResult, error) {
	result := RecognitionResult{}

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Convert to grayscale for better barcode detection
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	// Detect barcode
	barcode, err := detectBarcode(gray)
	if err != nil {
		return nil, err
	}

	result.RecognitionSuccess = true
	result.RecognitionData = fmt.Sprintf(`{"barcode":"%s"}`, barcode)
	return &result, nil
}

// Helper function to detect barcode in grayscale image
func detectBarcode(img *image.Gray) (string, error) {
	// Implement barcode detection algorithm here
	// This could be using edge detection, pattern matching, etc.
	// For now, we'll return a placeholder
	return "", errors.New("barcode detection not implemented")
}

// recognizeWithCNN uses a CNN-based model to recognize the product in the image
func (s *ImageRecognitionService) recognizeWithCNN(filePath string) (*RecognitionResult, error) {
	result := RecognitionResult{}

	// Open the image file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %v", err)
	}

	// Resize image to 224x224
	resized := imaging.Resize(img, 224, 224, imaging.Lanczos)

	// Convert NRGBA to RGB
	bounds := resized.Bounds()
	rgb := image.NewRGBA(bounds)

	// Copy and convert NRGBA to RGB
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := resized.At(x, y).RGBA()
			rgb.Set(x, y, color.RGBA{
				uint8(r >> 8),
				uint8(g >> 8),
				uint8(b >> 8),
				255,
			})
		}
	}

	// Create a new buffer for the file content
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)

	// Create form file
	part, err := writer.CreateFormFile("image", "processed_image.png")
	if err != nil {
		return nil, err
	}

	// Encode the processed image as PNG
	if err := png.Encode(part, rgb); err != nil {
		return nil, fmt.Errorf("failed to encode processed image: %v", err)
	}

	// Close the writer
	writer.Close()

	// Create request to Python API
	req, err := http.NewRequest("POST", s.CNNAPIEndpoint, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("CNN API returned non-200 status: %d", resp.StatusCode)
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse response from Python API
	var cnnResp struct {
		Success bool            `json:"success"`
		Results json.RawMessage `json:"results"`
		Error   string          `json:"error"`
	}

	if err := json.Unmarshal(respBody, &cnnResp); err != nil {
		return nil, err
	}

	if !cnnResp.Success {
		return nil, errors.New(cnnResp.Error)
	}

	// Return the JSON string of recognition results
	result.RecognitionData = string(cnnResp.Results)
	result.RecognitionSuccess = true
	return &result, nil
}

// FindProductByBarcode finds a product by its barcode in Firestore
func (s *ImageRecognitionService) FindProductByBarcode(barcode string) (string, error) {
	ctx := context.Background()

	// Query Firestore for products with matching barcode
	iter := s.Client.Collection("products").Where("barcode_value", "==", barcode).Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		return "", fmt.Errorf("failed to query products: %v", err)
	}

	if len(docs) == 0 {
		return "", fmt.Errorf("no product found with barcode: %s", barcode)
	}

	// Get the first matching product
	doc := docs[0]
	var product Product
	if err := doc.DataTo(&product); err != nil {
		return "", fmt.Errorf("failed to parse product data: %v", err)
	}

	// Return product info as JSON
	productInfo := map[string]interface{}{
		"id":            product.ID,
		"name":          product.Name,
		"barcode_value": product.BarcodeValue,
	}

	jsonData, err := json.Marshal(productInfo)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product info: %v", err)
	}

	return string(jsonData), nil
}

// FindProductByName finds a product by its name in Firestore
func (s *ImageRecognitionService) FindProductByName(name string) (string, error) {
	ctx := context.Background()

	// Query Firestore for products with matching name
	iter := s.Client.Collection("products").Where("name", "==", name).Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		return "", fmt.Errorf("failed to query products: %v", err)
	}

	if len(docs) == 0 {
		return "", fmt.Errorf("no product found with name: %s", name)
	}

	// Get the first matching product
	doc := docs[0]
	var product Product
	if err := doc.DataTo(&product); err != nil {
		return "", fmt.Errorf("failed to parse product data: %v", err)
	}

	// Return product info as JSON
	productInfo := map[string]interface{}{
		"product_code": product.ID,
		"name":         product.Name,
	}

	jsonData, err := json.Marshal(productInfo)
	if err != nil {
		return "", fmt.Errorf("failed to marshal product info: %v", err)
	}

	return string(jsonData), nil
}
