package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/nirshpaa/godam-backend/cmd/server/setup"
	"github.com/nirshpaa/godam-backend/services"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Create a background context
	ctx := context.Background()

	// Initialize Firebase service
	firebaseService, err := services.NewFirebaseService(ctx, "firebase-credentials.json")
	if err != nil {
		log.Fatal("Failed to initialize Firebase service:", err)
	}

	// Create and start the server
	srv, err := setup.NewServer()
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	// Setup routes
	setup.SetupRoutes(srv.GetRouter(), firebaseService)

	// Start the server
	if err := srv.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
