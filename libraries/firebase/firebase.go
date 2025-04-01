package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var (
	app      *firebase.App
	dbClient *db.Client
)

// Initialize Firebase
func Initialize() error {
	ctx := context.Background()

	// Initialize Firebase with credentials
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_FILE"))
	var err error
	app, err = firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return err
	}

	// Initialize Realtime Database
	dbClient, err = app.Database(ctx)
	if err != nil {
		return err
	}

	log.Println("Firebase initialized successfully")
	return nil
}

// GetDB returns the Firebase database client
func GetDB() *db.Client {
	return dbClient
}

// GetApp returns the Firebase app instance
func GetApp() *firebase.App {
	return app
}
