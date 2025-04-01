package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"firebase.google.com/go/v4/db"
	"github.com/nirshpaa/godam-backend/libraries/firebase"
)

// FirebaseModel provides common Firebase operations
type FirebaseModel struct {
	ref *db.Ref
}

// NewFirebaseModel creates a new Firebase model instance
func NewFirebaseModel(path string) *FirebaseModel {
	db := firebase.GetDB()
	return &FirebaseModel{
		ref: db.NewRef(path),
	}
}

// Create creates a new record in Firebase
func (m *FirebaseModel) Create(ctx context.Context, data interface{}) (string, error) {
	// Generate a new reference
	newRef, err := m.ref.Push(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create reference: %v", err)
	}

	// Add metadata
	metadata := map[string]interface{}{
		"created_at": time.Now().Unix(),
		"updated_at": time.Now().Unix(),
	}

	// Combine data with metadata
	record := map[string]interface{}{
		"data":     data,
		"metadata": metadata,
	}

	// Save to Firebase
	if err := newRef.Set(ctx, record); err != nil {
		return "", fmt.Errorf("failed to create record: %v", err)
	}

	return newRef.Key, nil
}

// Get retrieves a record by ID
func (m *FirebaseModel) Get(ctx context.Context, id string, result interface{}) error {
	ref := m.ref.Child(id)

	var record struct {
		Data     json.RawMessage `json:"data"`
		Metadata struct {
			CreatedAt int64 `json:"created_at"`
			UpdatedAt int64 `json:"updated_at"`
		} `json:"metadata"`
	}

	if err := ref.Get(ctx, &record); err != nil {
		return fmt.Errorf("failed to get record: %v", err)
	}

	return json.Unmarshal(record.Data, result)
}

// Update updates an existing record
func (m *FirebaseModel) Update(ctx context.Context, id string, data interface{}) error {
	ref := m.ref.Child(id)

	// Get existing record
	var record struct {
		Data     json.RawMessage `json:"data"`
		Metadata struct {
			CreatedAt int64 `json:"created_at"`
			UpdatedAt int64 `json:"updated_at"`
		} `json:"metadata"`
	}

	if err := ref.Get(ctx, &record); err != nil {
		return fmt.Errorf("failed to get existing record: %v", err)
	}

	// Update metadata
	record.Metadata.UpdatedAt = time.Now().Unix()

	// Create new record with updated data
	updatedRecord := map[string]interface{}{
		"data":     data,
		"metadata": record.Metadata,
	}

	return ref.Set(ctx, updatedRecord)
}

// Delete removes a record
func (m *FirebaseModel) Delete(ctx context.Context, id string) error {
	ref := m.ref.Child(id)
	return ref.Delete(ctx)
}

// List retrieves all records
func (m *FirebaseModel) List(ctx context.Context, result interface{}) error {
	var records map[string]struct {
		Data     json.RawMessage `json:"data"`
		Metadata struct {
			CreatedAt int64 `json:"created_at"`
			UpdatedAt int64 `json:"updated_at"`
		} `json:"metadata"`
	}

	if err := m.ref.Get(ctx, &records); err != nil {
		return fmt.Errorf("failed to list records: %v", err)
	}

	// Convert to slice of data
	var dataSlice []json.RawMessage
	for _, record := range records {
		dataSlice = append(dataSlice, record.Data)
	}

	// Marshal and unmarshal to convert to the desired type
	jsonData, err := json.Marshal(dataSlice)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	return json.Unmarshal(jsonData, result)
}

// Query retrieves records based on a query
func (m *FirebaseModel) Query(ctx context.Context, query *db.Query, result interface{}) error {
	var records map[string]struct {
		Data     json.RawMessage `json:"data"`
		Metadata struct {
			CreatedAt int64 `json:"created_at"`
			UpdatedAt int64 `json:"updated_at"`
		} `json:"metadata"`
	}

	if err := query.Get(ctx, &records); err != nil {
		return fmt.Errorf("failed to query records: %v", err)
	}

	// Convert to slice of data
	var dataSlice []json.RawMessage
	for _, record := range records {
		dataSlice = append(dataSlice, record.Data)
	}

	// Marshal and unmarshal to convert to the desired type
	jsonData, err := json.Marshal(dataSlice)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	return json.Unmarshal(jsonData, result)
}
