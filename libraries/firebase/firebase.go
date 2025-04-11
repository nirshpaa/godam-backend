package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var (
	app      *firebase.App
	fsClient *firestore.Client
)

// Initialize Firebase
func Initialize() (*firebase.App, error) {
	opt := option.WithCredentialsFile("firebase-credentials.json")

	// Initialize Firebase app
	var err error
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing Firebase app: %v", err)
	}

	// Initialize Firestore client
	fsClient, err = app.Firestore(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing Firestore client: %v", err)
	}

	return app, nil
}

// GetFirestore returns the Firestore client
func GetFirestore() *firestore.Client {
	return fsClient
}

// GetApp returns the Firebase app instance
func GetApp() *firebase.App {
	return app
}
