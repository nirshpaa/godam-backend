package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// AccessFirebaseModel represents an access in the system for Firebase
type AccessFirebaseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AccessFirebase represents the Firestore client for access
type AccessFirebase struct {
	client *firestore.Client
}

// NewAccessFirebase creates a new AccessFirebase instance
func NewAccessFirebase(client *firestore.Client) *AccessFirebase {
	return &AccessFirebase{
		client: client,
	}
}

// List retrieves all access records
func (a *AccessFirebase) List(ctx context.Context) ([]*AccessFirebaseModel, error) {
	var access []*AccessFirebaseModel
	docs, err := a.client.Collection("access").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting access: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var acc AccessFirebaseModel
		if err := doc.DataTo(&acc); err != nil {
			log.Printf("Error converting access data: %v", err)
			continue
		}
		acc.ID = doc.Ref.ID
		access = append(access, &acc)
	}

	return access, nil
}

// Get retrieves a single access record by ID
func (a *AccessFirebase) Get(ctx context.Context, id string) (*AccessFirebaseModel, error) {
	doc, err := a.client.Collection("access").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting access: %v", err)
		return nil, err
	}

	var access AccessFirebaseModel
	if err := doc.DataTo(&access); err != nil {
		log.Printf("Error converting access data: %v", err)
		return nil, err
	}
	access.ID = doc.Ref.ID

	return &access, nil
}

// Create creates a new access record
func (a *AccessFirebase) Create(ctx context.Context, access *AccessFirebaseModel) (string, error) {
	docRef, _, err := a.client.Collection("access").Add(ctx, access)
	if err != nil {
		log.Printf("Error creating access: %v", err)
		return "", err
	}
	return docRef.ID, nil
}

// Update updates an existing access record
func (a *AccessFirebase) Update(ctx context.Context, id string, access *AccessFirebaseModel) error {
	_, err := a.client.Collection("access").Doc(id).Set(ctx, access)
	if err != nil {
		log.Printf("Error updating access: %v", err)
		return err
	}
	return nil
}

// Delete deletes an access record
func (a *AccessFirebase) Delete(ctx context.Context, id string) error {
	_, err := a.client.Collection("access").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting access: %v", err)
		return err
	}
	return nil
}
