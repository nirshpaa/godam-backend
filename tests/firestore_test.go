package tests

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/nirshpaa/godam-backend/libraries/firebase"
)

// NewFirestoreTest creates a new Firestore client for testing.
// It returns the client and a cleanup function.
func NewFirestoreTest(t *testing.T) (*firestore.Client, func()) {
	t.Helper()

	// Initialize Firebase
	app, err := firebase.Initialize()
	if err != nil {
		t.Fatalf("initializing Firebase: %v", err)
	}

	// Get Firestore client
	client, err := app.Firestore(context.Background())
	if err != nil {
		t.Fatalf("getting Firestore client: %v", err)
	}

	// Create a test collection
	testCollection := client.Collection("test_products")

	// Cleanup function
	cleanup := func() {
		// Delete all documents in the test collection
		docs, err := testCollection.Documents(context.Background()).GetAll()
		if err != nil {
			t.Logf("warning: failed to get test documents for cleanup: %v", err)
			return
		}

		for _, doc := range docs {
			_, err := doc.Ref.Delete(context.Background())
			if err != nil {
				t.Logf("warning: failed to delete test document: %v", err)
			}
		}

		// Close the client
		client.Close()
	}

	return client, cleanup
}
