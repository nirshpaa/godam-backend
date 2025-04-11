package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

// RoleFirebaseModel represents a role in the system for Firebase
type RoleFirebaseModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// RoleFirebase represents the Firestore client for role
type RoleFirebase struct {
	client *firestore.Client
}

// NewRoleFirebase creates a new RoleFirebase instance
func NewRoleFirebase(client *firestore.Client) *RoleFirebase {
	return &RoleFirebase{
		client: client,
	}
}

// List retrieves all role records
func (r *RoleFirebase) List(ctx context.Context) ([]*RoleFirebaseModel, error) {
	var roles []*RoleFirebaseModel
	docs, err := r.client.Collection("roles").Documents(ctx).GetAll()
	if err != nil {
		log.Printf("Error getting roles: %v", err)
		return nil, err
	}

	for _, doc := range docs {
		var role RoleFirebaseModel
		if err := doc.DataTo(&role); err != nil {
			log.Printf("Error converting role data: %v", err)
			continue
		}
		role.ID = doc.Ref.ID
		roles = append(roles, &role)
	}

	return roles, nil
}

// Get retrieves a single role record by ID
func (r *RoleFirebase) Get(ctx context.Context, id string) (*RoleFirebaseModel, error) {
	doc, err := r.client.Collection("roles").Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Error getting role: %v", err)
		return nil, err
	}

	var role RoleFirebaseModel
	if err := doc.DataTo(&role); err != nil {
		log.Printf("Error converting role data: %v", err)
		return nil, err
	}
	role.ID = doc.Ref.ID

	return &role, nil
}

// Create creates a new role record
func (r *RoleFirebase) Create(ctx context.Context, role *RoleFirebaseModel) (string, error) {
	doc, _, err := r.client.Collection("roles").Add(ctx, role)
	if err != nil {
		log.Printf("Error creating role: %v", err)
		return "", err
	}

	return doc.ID, nil
}

// Update updates an existing role record
func (r *RoleFirebase) Update(ctx context.Context, id string, role *RoleFirebaseModel) error {
	_, err := r.client.Collection("roles").Doc(id).Set(ctx, role)
	if err != nil {
		log.Printf("Error updating role: %v", err)
		return err
	}

	return nil
}

// Delete deletes a role record
func (r *RoleFirebase) Delete(ctx context.Context, id string) error {
	_, err := r.client.Collection("roles").Doc(id).Delete(ctx)
	if err != nil {
		log.Printf("Error deleting role: %v", err)
		return err
	}

	return nil
}
