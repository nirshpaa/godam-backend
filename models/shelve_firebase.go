package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// ShelveFirebaseModel represents a shelve in the system for Firebase
type ShelveFirebaseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ShelveFirebase represents the Firestore client for shelve
type ShelveFirebase struct {
	client *firestore.Client
}

// NewShelveFirebase creates a new ShelveFirebase instance
func NewShelveFirebase(client *firestore.Client) *ShelveFirebase {
	return &ShelveFirebase{
		client: client,
	}
}

// List retrieves all shelve records
func (s *ShelveFirebase) List(ctx context.Context) ([]*ShelveFirebaseModel, error) {
	var shelves []*ShelveFirebaseModel
	docs, err := s.client.Collection("shelves").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting shelves: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var shelve ShelveFirebaseModel
		if err := doc.DataTo(&shelve); err != nil {
			log.Printf("Error converting shelve data: %v", err)
			continue
		}
		shelve.ID = doc.Ref.ID
		shelves = append(shelves, &shelve)
	}

	return shelves, nil
}

// Get retrieves a single shelve record by ID
func (s *ShelveFirebase) Get(ctx context.Context, id string) (*ShelveFirebaseModel, error) {
	doc, err := s.client.Collection("shelves").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting shelve: %v", err)
		return nil, err
	}

	var shelve ShelveFirebaseModel
	if err := doc.DataTo(&shelve); err != nil {
		log.Printf("Error converting shelve data: %v", err)
		return nil, err
	}
	shelve.ID = doc.Ref.ID

	return &shelve, nil
}

// Create creates a new shelve
func (s *ShelveFirebase) Create(ctx context.Context, shelve *ShelveFirebaseModel) (string, error) {
	docRef := s.client.Collection("shelves").NewDoc()
	_, err := docRef.Set(ctx, shelve)
	if err != nil {
		log.Printf("Error creating shelve: %v", err)
		return "", err
	}
	return docRef.ID, nil
}

// Update updates an existing shelve
func (s *ShelveFirebase) Update(ctx context.Context, id string, shelve *ShelveFirebaseModel) error {
	docRef := s.client.Collection("shelves").Doc(id)
	_, err := docRef.Set(ctx, shelve)
	if err != nil {
		log.Printf("Error updating shelve: %v", err)
		return err
	}
	return nil
}

// Delete removes a shelve
func (s *ShelveFirebase) Delete(ctx context.Context, id string) error {
	docRef := s.client.Collection("shelves").Doc(id)
	_, err := docRef.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting shelve: %v", err)
		return err
	}
	return nil
}

// FindByCompany retrieves all shelves for a specific company
func (s *ShelveFirebase) FindByCompany(ctx context.Context, companyID string) ([]*ShelveFirebaseModel, error) {
	var shelves []*ShelveFirebaseModel
	query := s.client.Collection("shelves").Where("company_id", "==", companyID)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting shelves: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var shelve ShelveFirebaseModel
		if err := doc.DataTo(&shelve); err != nil {
			log.Printf("Error converting shelve data: %v", err)
			continue
		}
		shelve.ID = doc.Ref.ID
		shelves = append(shelves, &shelve)
	}

	return shelves, nil
}
