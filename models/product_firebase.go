package models

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/nirshpaa/godam-backend/interfaces"
	"github.com/nirshpaa/godam-backend/types"
)

// ProductFirebase represents a product in Firebase
type ProductFirebase struct {
	*FirebaseModel
}

// NewProductFirebase creates a new Firebase product model
func NewProductFirebase(client *firestore.Client) (*ProductFirebase, error) {
	return &ProductFirebase{
		FirebaseModel: NewFirebaseModel("products", client),
	}, nil
}

// FirebaseProduct represents a product in Firebase
type FirebaseProduct struct {
	Code                 string    `json:"code"`
	Name                 string    `json:"name"`
	PurchasePrice        float64   `json:"purchase_price"`
	SalePrice            float64   `json:"sale_price"`
	MinimumStock         float64   `json:"minimum_stock"`
	ImageURL             string    `json:"image_url"`
	BarcodeValue         string    `json:"barcode_value"`
	ImageRecognitionData string    `json:"image_recognition_data"`
	CompanyID            string    `json:"company_id"`
	BrandID              string    `json:"brand_id"`
	ProductCategoryID    string    `json:"product_category_id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

// UnmarshalJSON implements custom JSON unmarshaling for FirebaseProduct
func (p *FirebaseProduct) UnmarshalJSON(data []byte) error {
	type Alias FirebaseProduct
	aux := &struct {
		CreatedAt interface{} `json:"created_at"`
		UpdatedAt interface{} `json:"updated_at"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parse timestamps
	var err error
	if aux.CreatedAt != nil {
		switch v := aux.CreatedAt.(type) {
		case string:
			p.CreatedAt, err = time.Parse(time.RFC3339, v)
		case float64:
			p.CreatedAt = time.Unix(int64(v), 0)
		}
		if err != nil {
			return fmt.Errorf("failed to parse created_at: %v", err)
		}
	}

	if aux.UpdatedAt != nil {
		switch v := aux.UpdatedAt.(type) {
		case string:
			p.UpdatedAt, err = time.Parse(time.RFC3339, v)
		case float64:
			p.UpdatedAt = time.Unix(int64(v), 0)
		}
		if err != nil {
			return fmt.Errorf("failed to parse updated_at: %v", err)
		}
	}

	return nil
}

func toFloat64(v interface{}) float64 {
	switch val := v.(type) {
	case int64:
		return float64(val)
	case float64:
		return val
	default:
		return 0
	}
}

// RecognitionResult represents the result of image processing
type RecognitionResult struct {
	RecognitionSuccess bool   `json:"recognition_success"`
	RecognitionData    string `json:"recognition_data"`
}

// Get retrieves a product by code
func (p *ProductFirebase) Get(ctx context.Context, code string) (*FirebaseProduct, error) {
	// Query the products collection
	iter := p.client.Collection("products").
		Where("code", "==", code).
		Limit(1).
		Documents(ctx)

	// Get the first document
	doc, err := iter.Next()
	if err != nil {
		return nil, fmt.Errorf("no product found with code: %s", code)
	}

	// Map the document to a FirebaseProduct
	return mapFirebaseProduct(doc)
}

// mapFirebaseProduct maps a Firestore document to a FirebaseProduct struct
func mapFirebaseProduct(doc *firestore.DocumentSnapshot) (*FirebaseProduct, error) {
	data := doc.Data()
	if data == nil {
		return nil, fmt.Errorf("document data is nil")
	}

	// Create a new product and populate it
	product := &FirebaseProduct{
		Code:                 data["code"].(string),
		Name:                 data["name"].(string),
		PurchasePrice:        toFloat64(data["purchase_price"]),
		SalePrice:            toFloat64(data["sale_price"]),
		MinimumStock:         toFloat64(data["minimum_stock"]),
		ImageURL:             data["image_url"].(string),
		BarcodeValue:         data["barcode_value"].(string),
		ImageRecognitionData: data["image_recognition_data"].(string),
		CompanyID:            data["company_id"].(string),
		BrandID:              data["brand_id"].(string),
		ProductCategoryID:    data["product_category_id"].(string),
	}

	// Handle timestamps
	if createdAt, ok := data["created_at"].(time.Time); ok {
		product.CreatedAt = createdAt
	}
	if updatedAt, ok := data["updated_at"].(time.Time); ok {
		product.UpdatedAt = updatedAt
	}

	return product, nil
}

// List retrieves all products
func (p *ProductFirebase) List(ctx context.Context) ([]FirebaseProduct, error) {
	var products []FirebaseProduct
	err := p.FirebaseModel.List(ctx, &products)
	return products, err
}

// Create creates a new product
func (p *ProductFirebase) Create(ctx context.Context, product *FirebaseProduct, fileStorage interfaces.FileStorage) (string, error) {
	// Check for duplicate code
	iter := p.client.Collection("products").
		Where("code", "==", product.Code).
		Limit(1).
		Documents(ctx)

	_, err := iter.Next()
	if err == nil {
		return "", fmt.Errorf("product with code %s already exists", product.Code)
	}

	return p.FirebaseModel.Create(ctx, product)
}

// Update updates an existing product
func (p *ProductFirebase) Update(ctx context.Context, code string, product *FirebaseProduct, fileStorage interfaces.FileStorage) error {
	// Query the products collection
	iter := p.client.Collection("products").
		Where("code", "==", code).
		Limit(1).
		Documents(ctx)

	// Get the first document
	doc, err := iter.Next()
	if err != nil {
		return fmt.Errorf("product not found with code: %s", code)
	}

	// Update the product using the document ID
	return p.FirebaseModel.Update(ctx, doc.Ref.ID, product)
}

// Delete deletes a product by code
func (p *ProductFirebase) Delete(ctx context.Context, code string) error {
	// Query the products collection
	iter := p.client.Collection("products").
		Where("code", "==", code).
		Limit(1).
		Documents(ctx)

	// Get the first document
	doc, err := iter.Next()
	if err != nil {
		return fmt.Errorf("no product found with code: %s", code)
	}

	// Delete the document
	_, err = doc.Ref.Delete(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	return nil
}

// UpdateImage updates only the image-related fields of a product
func (p *ProductFirebase) UpdateImage(ctx context.Context, code string, imageURL, barcodeValue, recognitionData string) error {
	// First get the product to ensure it exists
	product, err := p.Get(ctx, code)
	if err != nil {
		return fmt.Errorf("failed to get product with code %s: %v", code, err)
	}

	// Update the image-related fields
	product.ImageURL = imageURL
	product.BarcodeValue = barcodeValue
	product.ImageRecognitionData = recognitionData

	// Get the document reference
	docRef := p.client.Collection("products").Where("code", "==", code).Limit(1)
	docs, err := docRef.Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("failed to get document reference: %v", err)
	}
	if len(docs) == 0 {
		return fmt.Errorf("no document found for product with code: %s", code)
	}

	// Update the document
	_, err = docs[0].Ref.Update(ctx, []firestore.Update{
		{Path: "image_url", Value: imageURL},
		{Path: "barcode_value", Value: barcodeValue},
		{Path: "image_recognition_data", Value: recognitionData},
		{Path: "updated_at", Value: time.Now()},
	})
	if err != nil {
		return fmt.Errorf("failed to update product image: %v", err)
	}

	return nil
}

// FindByBarcode finds a product by its barcode
func (p *ProductFirebase) FindByBarcode(ctx context.Context, barcode string) (*FirebaseProduct, error) {
	query := p.ref.Where("data.barcode_value", "==", barcode).Limit(1)

	var products []FirebaseProduct
	err := p.FirebaseModel.Query(ctx, &query, &products)
	if err != nil {
		return nil, err
	}

	if len(products) == 0 {
		return nil, fmt.Errorf("no product found with barcode: %s", barcode)
	}

	return &products[0], nil
}

// FindByCompany retrieves all products for a specific company
func (p *ProductFirebase) FindByCompany(ctx context.Context, companyID string) ([]FirebaseProduct, error) {
	query := p.ref.Where("data.company_id", "==", companyID)

	var products []FirebaseProduct
	err := p.FirebaseModel.Query(ctx, &query, &products)
	return products, err
}

// ProcessImage processes an image for product recognition
func (p *ProductFirebase) ProcessImage(imagePath string, imageRecognition interfaces.ImageRecognition) (*types.ImageRecognitionResult, error) {
	// Check if the image file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("image file not found: %s", imagePath)
	}

	// Process the image
	result, err := imageRecognition.ProcessImage(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to process image: %v", err)
	}

	// Validate the result
	if result == nil {
		return nil, fmt.Errorf("image recognition returned nil result")
	}

	return result, nil
}
