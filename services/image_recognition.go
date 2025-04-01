package services

import (
	"bytes"
	"database/sql"
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
	"strings"

	"github.com/disintegration/imaging"
)

// ImageRecognitionService handles image analysis operations
type ImageRecognitionService struct {
	Db                 *sql.DB
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
	RecognitionSuccess bool
	RecognitionData    string
}

// NewImageRecognitionService creates a new image recognition service
func NewImageRecognitionService(db *sql.DB, barcodeEndpoint, cnnEndpoint string) *ImageRecognitionService {
	return &ImageRecognitionService{
		Db:                 db,
		BarcodeAPIEndpoint: barcodeEndpoint,
		CNNAPIEndpoint:     cnnEndpoint,
	}
}

// ProcessImage processes an image file for barcode scanning and CNN recognition
func (s *ImageRecognitionService) ProcessImage(imagePath string) (*RecognitionResult, error) {
	// Try barcode scanning first
	barcode, err := s.scanBarcode(imagePath)
	if err != nil {
		// If barcode scanning fails, try CNN recognition
		recognitionData, err := s.recognizeImage(imagePath)
		if err != nil {
			return &RecognitionResult{
				RecognitionSuccess: false,
				RecognitionData:    "",
			}, err
		}
		return &RecognitionResult{
			RecognitionSuccess: true,
			RecognitionData:    recognitionData,
		}, nil
	}

	// If barcode is found, look up product
	productJSON, err := s.FindProductByBarcode(barcode)
	if err != nil {
		// If product not found, return create product suggestion
		return &RecognitionResult{
			RecognitionSuccess: true,
			RecognitionData:    fmt.Sprintf(`{"action":"create_product","barcode":"%s"}`, barcode),
		}, nil
	}

	// Return existing product data
	return &RecognitionResult{
		RecognitionSuccess: true,
		RecognitionData:    productJSON,
	}, nil
}

// scanBarcode attempts to scan a barcode from an image
func (s *ImageRecognitionService) scanBarcode(imagePath string) (string, error) {
	// Implementation of barcode scanning
	return "", fmt.Errorf("barcode not found")
}

// recognizeImage uses CNN to recognize the image
func (s *ImageRecognitionService) recognizeImage(imagePath string) (string, error) {
	// Implementation of CNN recognition
	return `{"action":"create_product","suggested_name":"Unknown Product","confidence":0.0}`, nil
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

// FindProductByBarcode attempts to find a product by its barcode
func (s *ImageRecognitionService) FindProductByBarcode(barcode string) (string, error) {
	if strings.TrimSpace(barcode) == "" {
		return "", errors.New("empty barcode value")
	}

	// Query your database for the product with this barcode
	var product struct {
		ID           uint64  `json:"id"`
		Name         string  `json:"name"`
		Code         string  `json:"code"`
		BarcodeValue string  `json:"barcode_value"`
		Price        float64 `json:"price"`
	}

	query := `
		SELECT id, name, code, barcode_value, price 
		FROM products 
		WHERE barcode_value = ?
	`

	err := s.Db.QueryRow(query, barcode).Scan(
		&product.ID,
		&product.Name,
		&product.Code,
		&product.BarcodeValue,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no product found with barcode: %s", barcode)
		}
		return "", fmt.Errorf("database error: %v", err)
	}

	// Convert to JSON string
	productJSON, err := json.Marshal(product)
	if err != nil {
		return "", fmt.Errorf("error marshaling product: %v", err)
	}

	return string(productJSON), nil
}

// FindProductByRecognitionData attempts to find a product by its image recognition data
func (s *ImageRecognitionService) FindProductByRecognitionData(recognitionData string) (string, error) {
	if strings.TrimSpace(recognitionData) == "" {
		return "", errors.New("empty recognition data")
	}

	// Parse the recognition data
	var result struct {
		Class      string  `json:"class"`
		Confidence float64 `json:"confidence"`
	}
	if err := json.Unmarshal([]byte(recognitionData), &result); err != nil {
		return "", fmt.Errorf("failed to parse recognition data: %v", err)
	}

	// Query your database for products with matching recognition data
	var product struct {
		ID                   uint64  `json:"id"`
		Name                 string  `json:"name"`
		Code                 string  `json:"code"`
		ImageRecognitionData string  `json:"image_recognition_data"`
		Price                float64 `json:"price"`
		Confidence           float64 `json:"confidence"`
	}

	query := `
		SELECT id, name, code, image_recognition_data, sale_price 
		FROM products 
		WHERE image_recognition_data LIKE ?
		ORDER BY sale_price DESC
		LIMIT 1
	`

	pattern := "%" + result.Class + "%"
	err := s.Db.QueryRow(query, pattern).Scan(
		&product.ID,
		&product.Name,
		&product.Code,
		&product.ImageRecognitionData,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no product found matching class: %s", result.Class)
		}
		return "", fmt.Errorf("database error: %v", err)
	}

	product.Confidence = result.Confidence

	// Convert to JSON string
	productJSON, err := json.Marshal(product)
	if err != nil {
		return "", fmt.Errorf("error marshaling product: %v", err)
	}

	return string(productJSON), nil
}
