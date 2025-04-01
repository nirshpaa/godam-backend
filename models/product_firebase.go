package models

import (
	"context"
	"fmt"
)

// ProductFirebase represents a product in Firebase
type ProductFirebase struct {
	*FirebaseModel
}

// NewProductFirebase creates a new Firebase product model
func NewProductFirebase() *ProductFirebase {
	return &ProductFirebase{
		FirebaseModel: NewFirebaseModel("products"),
	}
}

// FirebaseProduct represents a product in Firebase
type FirebaseProduct struct {
	ID                   string  `json:"id"`
	Code                 string  `json:"code"`
	Name                 string  `json:"name"`
	PurchasePrice        float64 `json:"purchase_price"`
	SalePrice            float64 `json:"sale_price"`
	MinimumStock         uint    `json:"minimum_stock"`
	ImageURL             string  `json:"image_url"`
	BarcodeValue         string  `json:"barcode_value"`
	ImageRecognitionData string  `json:"image_recognition_data"`
	CompanyID            string  `json:"company_id"`
	BrandID              string  `json:"brand_id"`
	ProductCategoryID    string  `json:"product_category_id"`
}

// List retrieves all products
func (p *ProductFirebase) List(ctx context.Context) ([]FirebaseProduct, error) {
	var products []FirebaseProduct
	err := p.FirebaseModel.List(ctx, &products)
	return products, err
}

// Get retrieves a product by ID
func (p *ProductFirebase) Get(ctx context.Context, id string) (*FirebaseProduct, error) {
	var product FirebaseProduct
	err := p.FirebaseModel.Get(ctx, id, &product)
	if err != nil {
		return nil, err
	}
	product.ID = id
	return &product, nil
}

// Create creates a new product
func (p *ProductFirebase) Create(ctx context.Context, product *FirebaseProduct) (string, error) {
	return p.FirebaseModel.Create(ctx, product)
}

// Update updates an existing product
func (p *ProductFirebase) Update(ctx context.Context, id string, product *FirebaseProduct) error {
	return p.FirebaseModel.Update(ctx, id, product)
}

// Delete removes a product
func (p *ProductFirebase) Delete(ctx context.Context, id string) error {
	return p.FirebaseModel.Delete(ctx, id)
}

// UpdateImage updates only the image-related fields of a product
func (p *ProductFirebase) UpdateImage(ctx context.Context, id string, imageURL, barcodeValue, recognitionData string) error {
	product, err := p.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get product: %v", err)
	}

	product.ImageURL = imageURL
	product.BarcodeValue = barcodeValue
	product.ImageRecognitionData = recognitionData

	return p.Update(ctx, id, product)
}

// FindByBarcode finds a product by its barcode
func (p *ProductFirebase) FindByBarcode(ctx context.Context, barcode string) (*FirebaseProduct, error) {
	query := p.ref.OrderByChild("barcode_value").EqualTo(barcode).LimitToFirst(1)

	var products []FirebaseProduct
	err := p.FirebaseModel.Query(ctx, query, &products)
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
	query := p.ref.OrderByChild("company_id").EqualTo(companyID)

	var products []FirebaseProduct
	err := p.FirebaseModel.Query(ctx, query, &products)
	return products, err
}
