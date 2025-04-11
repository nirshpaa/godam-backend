package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// ProductCategoryFirebase represents the product category model for Firebase
type ProductCategoryFirebase struct {
	client *firestore.Client
}

// NewProductCategoryFirebase creates a new instance of ProductCategoryFirebase
func NewProductCategoryFirebase(client *firestore.Client) *ProductCategoryFirebase {
	return &ProductCategoryFirebase{
		client: client,
	}
}

// ProductCategoryFirebaseModel represents a product category in the system for Firebase
type ProductCategoryFirebaseModel struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	CompanyID string                `json:"company_id"`
	Category  CategoryFirebaseModel `json:"category"`
}

// CategoryFirebaseModel represents a category in the system for Firebase
type CategoryFirebaseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// List returns all product categories
func (pc *ProductCategoryFirebase) List(ctx context.Context) ([]ProductCategoryFirebaseModel, error) {
	var categories []ProductCategoryFirebaseModel

	// Get all product categories from Firestore
	iter := pc.client.Collection("product_categories").Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		log.Printf("Error getting product categories: %v", err)
		return nil, err
	}

	// Convert Firestore documents to ProductCategoryFirebaseModel structs
	for _, doc := range docs {
		var category ProductCategoryFirebaseModel
		if err := doc.DataTo(&category); err != nil {
			log.Printf("Error converting product category document: %v", err)
			continue
		}
		category.ID = doc.Ref.ID
		categories = append(categories, category)
	}

	return categories, nil
}

// Get returns a product category by ID
func (pc *ProductCategoryFirebase) Get(ctx context.Context, id string) (*ProductCategoryFirebaseModel, error) {
	doc, err := pc.client.Collection("product_categories").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting product category: %v", err)
		return nil, err
	}

	var category ProductCategoryFirebaseModel
	if err := doc.DataTo(&category); err != nil {
		log.Printf("Error converting product category document: %v", err)
		return nil, err
	}
	category.ID = doc.Ref.ID

	return &category, nil
}

// Create creates a new product category
func (pc *ProductCategoryFirebase) Create(ctx context.Context, category *ProductCategoryFirebaseModel) error {
	docRef := pc.client.Collection("product_categories").NewDoc()
	category.ID = docRef.ID

	_, err := docRef.Set(ctx, category)
	if err != nil {
		log.Printf("Error creating product category: %v", err)
		return err
	}

	return nil
}

// Update updates an existing product category
func (pc *ProductCategoryFirebase) Update(ctx context.Context, category *ProductCategoryFirebaseModel) error {
	_, err := pc.client.Collection("product_categories").Doc(category.ID).Set(ctx, category)
	if err != nil {
		log.Printf("Error updating product category: %v", err)
		return err
	}

	return nil
}

// Delete deletes a product category
func (pc *ProductCategoryFirebase) Delete(ctx context.Context, id string) error {
	_, err := pc.client.Collection("product_categories").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting product category: %v", err)
		return err
	}

	return nil
}
