package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// BrandFirebase represents the brand model for Firebase
type BrandFirebase struct {
	client *firestore.Client
}

// NewBrandFirebase creates a new instance of BrandFirebase
func NewBrandFirebase(client *firestore.Client) *BrandFirebase {
	return &BrandFirebase{
		client: client,
	}
}

// BrandFirebaseModel represents a brand in the system for Firebase
type BrandFirebaseModel struct {
	ID        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CompanyID string `json:"company_id"`
}

// List returns all brands
func (b *BrandFirebase) List(ctx context.Context) ([]BrandFirebaseModel, error) {
	var brands []BrandFirebaseModel

	// Get all brands from Firestore
	iter := b.client.Collection("brands").Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		log.Printf("Error getting brands: %v", err)
		return nil, err
	}

	// Convert Firestore documents to BrandFirebaseModel structs
	for _, doc := range docs {
		var brand BrandFirebaseModel
		if err := doc.DataTo(&brand); err != nil {
			log.Printf("Error converting brand document: %v", err)
			continue
		}
		brand.ID = doc.Ref.ID
		brands = append(brands, brand)
	}

	return brands, nil
}

// Get returns a brand by ID
func (b *BrandFirebase) Get(ctx context.Context, id string) (*BrandFirebaseModel, error) {
	doc, err := b.client.Collection("brands").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting brand: %v", err)
		return nil, err
	}

	var brand BrandFirebaseModel
	if err := doc.DataTo(&brand); err != nil {
		log.Printf("Error converting brand document: %v", err)
		return nil, err
	}
	brand.ID = doc.Ref.ID

	return &brand, nil
}

// Create creates a new brand
func (b *BrandFirebase) Create(ctx context.Context, brand *BrandFirebaseModel) error {
	docRef := b.client.Collection("brands").NewDoc()
	brand.ID = docRef.ID

	_, err := docRef.Set(ctx, brand)
	if err != nil {
		log.Printf("Error creating brand: %v", err)
		return err
	}

	return nil
}

// Update updates an existing brand
func (b *BrandFirebase) Update(ctx context.Context, brand *BrandFirebaseModel) error {
	_, err := b.client.Collection("brands").Doc(brand.ID).Set(ctx, brand)
	if err != nil {
		log.Printf("Error updating brand: %v", err)
		return err
	}

	return nil
}

// Delete deletes a brand
func (b *BrandFirebase) Delete(ctx context.Context, id string) error {
	_, err := b.client.Collection("brands").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting brand: %v", err)
		return err
	}

	return nil
}
