package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// BranchFirebaseModel represents a branch in the system for Firebase
type BranchFirebaseModel struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
}

// BranchFirebase represents the Firestore client for branch
type BranchFirebase struct {
	client *firestore.Client
}

// NewBranchFirebase creates a new BranchFirebase instance
func NewBranchFirebase(client *firestore.Client) *BranchFirebase {
	return &BranchFirebase{
		client: client,
	}
}

// List retrieves all branch records
func (b *BranchFirebase) List(ctx context.Context) ([]*BranchFirebaseModel, error) {
	var branches []*BranchFirebaseModel
	docs, err := b.client.Collection("branches").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting branches: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var branch BranchFirebaseModel
		if err := doc.DataTo(&branch); err != nil {
			log.Printf("Error converting branch data: %v", err)
			continue
		}
		branch.ID = doc.Ref.ID
		branches = append(branches, &branch)
	}

	return branches, nil
}

// Get retrieves a single branch record by ID
func (b *BranchFirebase) Get(ctx context.Context, id string) (*BranchFirebaseModel, error) {
	doc, err := b.client.Collection("branches").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting branch: %v", err)
		return nil, err
	}

	var branch BranchFirebaseModel
	if err := doc.DataTo(&branch); err != nil {
		log.Printf("Error converting branch data: %v", err)
		return nil, err
	}
	branch.ID = doc.Ref.ID

	return &branch, nil
}

// Create creates a new branch record
func (b *BranchFirebase) Create(ctx context.Context, branch *BranchFirebaseModel) (string, error) {
	docRef, _, err := b.client.Collection("branches").Add(ctx, branch)
	if err != nil {
		log.Printf("Error creating branch: %v", err)
		return "", err
	}
	return docRef.ID, nil
}

// Update updates an existing branch record
func (b *BranchFirebase) Update(ctx context.Context, id string, branch *BranchFirebaseModel) error {
	_, err := b.client.Collection("branches").Doc(id).Set(ctx, branch)
	if err != nil {
		log.Printf("Error updating branch: %v", err)
		return err
	}
	return nil
}

// Delete deletes a branch record
func (b *BranchFirebase) Delete(ctx context.Context, id string) error {
	_, err := b.client.Collection("branches").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting branch: %v", err)
		return err
	}
	return nil
}
