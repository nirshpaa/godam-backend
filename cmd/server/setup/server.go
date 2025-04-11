package setup

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/middleware"
	"github.com/nirshpaa/godam-backend/services"
)

// Server represents the API server
type Server struct {
	router          *gin.Engine
	logger          *log.Logger
	firebaseService *services.FirebaseService
	ctx             context.Context
	cancel          context.CancelFunc
}

// NewServer creates a new API server instance
func NewServer() (*Server, error) {
	logger := log.New(os.Stdout, "[API] ", log.LstdFlags|log.Lshortfile)

	// Create a background context that won't be cancelled
	ctx := context.Background()

	// Get the absolute path to the credentials file
	credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_FILE")
	if credentialsPath == "" {
		credentialsPath = "firebase-credentials.json"
	}

	absPath, err := filepath.Abs(credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path to credentials file: %v", err)
	}
	logger.Printf("Using credentials file: %s", absPath)

	// Create Firebase service first
	firebaseService, err := services.NewFirebaseService(ctx, absPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create Firebase service: %v", err)
	}

	// Verify the Firestore client is not nil
	if firebaseService.GetFirestore() == nil {
		return nil, fmt.Errorf("Firestore client is nil after initialization")
	}

	// Create new router
	router := gin.Default()

	// Apply middleware in correct order
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.Logger(logger))

	// Serve static files from the assets directory
	router.Static("/assets", "./assets")
	router.Static("/images", "./public/images")

	// Apply auth middleware last, but exclude static file routes
	router.Use(middleware.AuthMiddleware(firebaseService))

	return &Server{
		router:          router,
		logger:          logger,
		firebaseService: firebaseService,
		ctx:             ctx,
		cancel:          nil, // No cancel function needed
	}, nil
}

// Start starts the server
func (s *Server) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// Start a goroutine to monitor the context
	go func() {
		<-s.ctx.Done()
		s.logger.Printf("Server context cancelled, shutting down...")
	}()

	return s.router.Run(":" + port)
}

// Close closes the server and its resources
func (s *Server) Close() error {
	if s.cancel != nil {
		s.cancel()
	}
	if s.firebaseService != nil {
		return s.firebaseService.Close()
	}
	return nil
}

// GetRouter returns the server's router
func (s *Server) GetRouter() *gin.Engine {
	return s.router
}
