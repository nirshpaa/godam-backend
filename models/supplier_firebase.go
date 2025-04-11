package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// SupplierFirebase represents the supplier model for Firebase
type SupplierFirebase struct {
	client *firestore.Client
}

// NewSupplierFirebase creates a new instance of SupplierFirebase
func NewSupplierFirebase(client *firestore.Client) *SupplierFirebase {
	return &SupplierFirebase{
		client: client,
	}
}

// Supplier represents a supplier in the system
type Supplier struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// List returns all suppliers
func (s *SupplierFirebase) List(ctx context.Context) ([]Supplier, error) {
	var suppliers []Supplier

	// Get all suppliers from Firestore
	iter := s.client.Collection("suppliers").Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		log.Printf("Error getting suppliers: %v", err)
		return nil, err
	}

	// Convert Firestore documents to Supplier structs
	for _, doc := range docs {
		var supplier Supplier
		if err := doc.DataTo(&supplier); err != nil {
			log.Printf("Error converting supplier document: %v", err)
			continue
		}
		supplier.ID = doc.Ref.ID
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

// Get returns a supplier by ID
func (s *SupplierFirebase) Get(ctx context.Context, id string) (*Supplier, error) {
	doc, err := s.client.Collection("suppliers").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting supplier: %v", err)
		return nil, err
	}

	var supplier Supplier
	if err := doc.DataTo(&supplier); err != nil {
		log.Printf("Error converting supplier document: %v", err)
		return nil, err
	}
	supplier.ID = doc.Ref.ID

	return &supplier, nil
}

// Create creates a new supplier
func (s *SupplierFirebase) Create(ctx context.Context, supplier *Supplier) error {
	docRef, _, err := s.client.Collection("suppliers").Add(ctx, supplier)
	if err != nil {
		log.Printf("Error creating supplier: %v", err)
		return err
	}
	supplier.ID = docRef.ID
	return nil
}

// Update updates an existing supplier
func (s *SupplierFirebase) Update(ctx context.Context, supplier *Supplier) error {
	_, err := s.client.Collection("suppliers").Doc(supplier.ID).Set(ctx, supplier)
	if err != nil {
		log.Printf("Error updating supplier: %v", err)
		return err
	}
	return nil
}

// Delete deletes a supplier by ID
func (s *SupplierFirebase) Delete(ctx context.Context, id string) error {
	_, err := s.client.Collection("suppliers").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting supplier: %v", err)
		return err
	}
	return nil
}
