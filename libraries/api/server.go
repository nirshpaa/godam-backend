package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/middleware"
	"github.com/nirshpaa/godam-backend/routing"
)

// Server represents the API server
type Server struct {
	router *gin.Engine
	logger *log.Logger
}

// NewServer creates a new API server instance
func NewServer() *Server {
	router := gin.Default()
	logger := log.New(log.Writer(), "[API] ", log.LstdFlags)

	// Add middleware
	router.Use(middleware.CORS())
	router.Use(middleware.Logger(logger))

	// Initialize routes
	router.Use(routing.SetupRoutes())

	return &Server{
		router: router,
		logger: logger,
	}
}

// Start starts the server
func (s *Server) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	s.logger.Printf("Starting server on port %s", port)
	return s.router.Run(":" + port)
}
