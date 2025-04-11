package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// SalesmanFirebaseModel represents a salesman in the system for Firebase
type SalesmanFirebaseModel struct {
	ID        string `json:"id"`
	CompanyID string `json:"company_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
}

// SalesmanFirebase represents the Firestore client for salesman
type SalesmanFirebase struct {
	client *firestore.Client
}

// NewSalesmanFirebase creates a new SalesmanFirebase instance
func NewSalesmanFirebase(client *firestore.Client) *SalesmanFirebase {
	return &SalesmanFirebase{
		client: client,
	}
}

// List retrieves all salesman records
func (s *SalesmanFirebase) List(ctx context.Context) ([]*SalesmanFirebaseModel, error) {
	var salesmen []*SalesmanFirebaseModel
	docs, err := s.client.Collection("salesmen").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting salesmen: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var salesman SalesmanFirebaseModel
		if err := doc.DataTo(&salesman); err != nil {
			log.Printf("Error converting salesman data: %v", err)
			continue
		}
		salesman.ID = doc.Ref.ID
		salesmen = append(salesmen, &salesman)
	}

	return salesmen, nil
}

// Get retrieves a single salesman record by ID
func (s *SalesmanFirebase) Get(ctx context.Context, id string) (*SalesmanFirebaseModel, error) {
	doc, err := s.client.Collection("salesmen").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting salesman: %v", err)
		return nil, err
	}

	var salesman SalesmanFirebaseModel
	if err := doc.DataTo(&salesman); err != nil {
		log.Printf("Error converting salesman data: %v", err)
		return nil, err
	}
	salesman.ID = doc.Ref.ID

	return &salesman, nil
}

// Create creates a new salesman record
func (s *SalesmanFirebase) Create(ctx context.Context, salesman *SalesmanFirebaseModel) (string, error) {
	doc, _, err := s.client.Collection("salesmen").Add(ctx, salesman)
	if err != nil {
		log.Printf("Error creating salesman: %v", err)
		return "", err
	}

	return doc.ID, nil
}

// Update updates an existing salesman record
func (s *SalesmanFirebase) Update(ctx context.Context, id string, salesman *SalesmanFirebaseModel) error {
	_, err := s.client.Collection("salesmen").Doc(id).Set(ctx, salesman)
	if err != nil {
		log.Printf("Error updating salesman: %v", err)
		return err
	}

	return nil
}

// Delete deletes a salesman record
func (s *SalesmanFirebase) Delete(ctx context.Context, id string) error {
	_, err := s.client.Collection("salesmen").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting salesman: %v", err)
		return err
	}

	return nil
}
