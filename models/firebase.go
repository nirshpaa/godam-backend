package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

// FirebaseModel provides common Firestore operations
type FirebaseModel struct {
	ref    *firestore.CollectionRef
	client *firestore.Client
}

// NewFirebaseModel creates a new Firestore model instance
func NewFirebaseModel(path string, client *firestore.Client) *FirebaseModel {
	if client == nil {
		log.Fatal("Firestore client is nil")
	}
	return &FirebaseModel{
		ref:    client.Collection(path),
		client: client,
	}
}

// Create creates a new record in Firestore
func (m *FirebaseModel) Create(ctx context.Context, data interface{}) (string, error) {
	// Convert data to map if it's not already
	var dataMap map[string]interface{}
	switch v := data.(type) {
	case map[string]interface{}:
		dataMap = v
	default:
		// Convert to JSON and back to map to handle structs
		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", fmt.Errorf("failed to marshal data: %v", err)
		}
		if err := json.Unmarshal(jsonData, &dataMap); err != nil {
			return "", fmt.Errorf("failed to unmarshal data: %v", err)
		}
	}

	// Add timestamps in RFC3339 format
	now := time.Now().Format(time.RFC3339)
	dataMap["created_at"] = now
	dataMap["updated_at"] = now

	// Create new document
	docRef, _, err := m.ref.Add(ctx, dataMap)
	if err != nil {
		return "", fmt.Errorf("failed to create document: %v", err)
	}

	return docRef.ID, nil
}

// Get retrieves a record by ID
func (m *FirebaseModel) Get(ctx context.Context, id string, result interface{}) error {
	doc, err := m.ref.Doc(id).Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to get record: %v", err)
	}

	// Get the data directly without the data/metadata wrapper
	data := doc.Data()
	if data == nil {
		return fmt.Errorf("no data found for record: %s", id)
	}

	// Convert the data to JSON and then to the result type
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	return json.Unmarshal(jsonData, result)
}

// Update updates an existing record
func (m *FirebaseModel) Update(ctx context.Context, id string, data interface{}) error {
	docRef := m.ref.Doc(id)

	// Convert data to map
	var dataMap map[string]interface{}

	// Try to convert struct to map
	if product, ok := data.(*FirebaseProduct); ok {
		// Convert struct to map
		jsonData, err := json.Marshal(product)
		if err != nil {
			return fmt.Errorf("failed to marshal product: %v", err)
		}

		if err := json.Unmarshal(jsonData, &dataMap); err != nil {
			return fmt.Errorf("failed to unmarshal product: %v", err)
		}
	} else if mapData, ok := data.(map[string]interface{}); ok {
		dataMap = mapData
	} else {
		return fmt.Errorf("data must be either a FirebaseProduct struct or a map")
	}

	// Add updated timestamp in RFC3339 format
	dataMap["updated_at"] = time.Now().Format(time.RFC3339)

	// Convert map to updates
	updates := make([]firestore.Update, 0, len(dataMap))
	for k, v := range dataMap {
		updates = append(updates, firestore.Update{
			Path:  k,
			Value: v,
		})
	}

	// Use Update instead of Set to prevent document duplication
	_, err := docRef.Update(ctx, updates)
	return err
}

// Delete removes a record
func (m *FirebaseModel) Delete(ctx context.Context, id string) error {
	docRef := m.ref.Doc(id)
	_, err := docRef.Delete(ctx)
	return err
}

// List retrieves all records
func (m *FirebaseModel) List(ctx context.Context, result interface{}) error {
	docs, err := m.ref.Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("failed to list records: %v", err)
	}

	log.Printf("Found %d documents in collection", len(docs))

	// Create a slice to hold the results
	results := make([]map[string]interface{}, 0, len(docs))

	for _, doc := range docs {
		log.Printf("Processing document %s", doc.Ref.ID)

		// Get raw data from document
		data := doc.Data()
		log.Printf("Raw document data: %+v", data)

		// Check if data exists
		if data == nil {
			log.Printf("Document %s has no data", doc.Ref.ID)
			continue
		}

		// Add the document ID
		data["id"] = doc.Ref.ID

		// Ensure all required fields exist with default values
		if data["code"] == nil {
			data["code"] = ""
		}
		if data["name"] == nil {
			data["name"] = ""
		}
		if data["purchase_price"] == nil {
			data["purchase_price"] = 0.0
		}
		if data["sale_price"] == nil {
			data["sale_price"] = 0.0
		}
		if data["minimum_stock"] == nil {
			data["minimum_stock"] = 0
		}
		if data["image_url"] == nil {
			data["image_url"] = ""
		}
		if data["company_id"] == nil {
			data["company_id"] = ""
		}
		if data["brand_id"] == nil {
			data["brand_id"] = ""
		}
		if data["product_category_id"] == nil {
			data["product_category_id"] = ""
		}
		if data["barcode_value"] == nil {
			data["barcode_value"] = ""
		}
		if data["image_recognition_data"] == nil {
			data["image_recognition_data"] = ""
		}

		// Handle timestamps
		now := time.Now().Format(time.RFC3339)

		// Handle created_at
		switch v := data["created_at"].(type) {
		case time.Time:
			data["created_at"] = v.Format(time.RFC3339)
		case nil:
			data["created_at"] = now
		}

		// Handle updated_at
		switch v := data["updated_at"].(type) {
		case time.Time:
			data["updated_at"] = v.Format(time.RFC3339)
		case nil:
			data["updated_at"] = now
		}

		log.Printf("Processed product: %+v", data)
		results = append(results, data)
	}

	log.Printf("Processed %d products successfully", len(results))

	// Convert the results to JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		return fmt.Errorf("failed to marshal results: %v", err)
	}

	// Unmarshal into the provided result interface
	return json.Unmarshal(jsonData, result)
}

// Query retrieves records based on a query
func (m *FirebaseModel) Query(ctx context.Context, query *firestore.Query, result interface{}) error {
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		return fmt.Errorf("failed to query records: %v", err)
	}

	var results []map[string]interface{}
	for _, doc := range docs {
		data := doc.Data()
		if data == nil {
			continue
		}
		results = append(results, data)
	}

	// Convert to JSON and then to the result type
	jsonData, err := json.Marshal(results)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	return json.Unmarshal(jsonData, result)
}

// NewFirebaseClient creates a new Firebase client
func NewFirebaseClient() (*firestore.Client, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "godam-inventory")
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %v", err)
	}
	return client, nil
}
