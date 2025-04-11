package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// RegionFirebaseModel represents a region in the system for Firebase
type RegionFirebaseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RegionFirebase represents the Firestore client for region
type RegionFirebase struct {
	client *firestore.Client
}

// NewRegionFirebase creates a new RegionFirebase instance
func NewRegionFirebase(client *firestore.Client) *RegionFirebase {
	return &RegionFirebase{
		client: client,
	}
}

// List retrieves all region records
func (r *RegionFirebase) List(ctx context.Context) ([]*RegionFirebaseModel, error) {
	var regions []*RegionFirebaseModel
	docs, err := r.client.Collection("regions").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting regions: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var region RegionFirebaseModel
		if err := doc.DataTo(&region); err != nil {
			log.Printf("Error converting region data: %v", err)
			continue
		}
		region.ID = doc.Ref.ID
		regions = append(regions, &region)
	}

	return regions, nil
}

// Get retrieves a single region record by ID
func (r *RegionFirebase) Get(ctx context.Context, id string) (*RegionFirebaseModel, error) {
	doc, err := r.client.Collection("regions").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting region: %v", err)
		return nil, err
	}

	var region RegionFirebaseModel
	if err := doc.DataTo(&region); err != nil {
		log.Printf("Error converting region data: %v", err)
		return nil, err
	}
	region.ID = doc.Ref.ID

	return &region, nil
}

// Create creates a new region record
func (r *RegionFirebase) Create(ctx context.Context, region *RegionFirebaseModel) (string, error) {
	doc, _, err := r.client.Collection("regions").Add(ctx, region)
	if err != nil {
		log.Printf("Error creating region: %v", err)
		return "", err
	}

	return doc.ID, nil
}

// Update updates an existing region record
func (r *RegionFirebase) Update(ctx context.Context, id string, region *RegionFirebaseModel) error {
	_, err := r.client.Collection("regions").Doc(id).Set(ctx, region)
	if err != nil {
		log.Printf("Error updating region: %v", err)
		return err
	}

	return nil
}

// Delete deletes a region record
func (r *RegionFirebase) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection("regions").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting region: %v", err)
		return err
	}

	return nil
}
