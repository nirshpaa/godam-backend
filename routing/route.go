package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/controllers"
	"github.com/nirshpaa/godam-backend/middleware"
)

// SetupRoutes configures all routes for the application
func SetupRoutes() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Health check
		c.GET("/health", controllers.HealthCheck)

		// Auth routes
		auth := controllers.NewAuthController()
		c.POST("/login", auth.Login)
		c.GET("/check/email", auth.CheckEmail)
		c.GET("/check/username", auth.CheckUsername)

		// Protected routes
		protected := c.Group("/api")
		protected.Use(middleware.Auth())
		{
			// Products routes
			products := controllers.NewProductController()
			protected.GET("/products", products.List)
			protected.GET("/products/:id", products.Get)
			protected.POST("/products", products.Create)
			protected.PUT("/products/:id", products.Update)
			protected.DELETE("/products/:id", products.Delete)
			protected.POST("/products/scan", products.ScanImage)

			// Companies routes
			companies := controllers.NewCompanyController()
			protected.GET("/companies", companies.List)
			protected.GET("/companies/:id", companies.Get)
			protected.POST("/companies", companies.Create)
			protected.PUT("/companies/:id", companies.Update)
			protected.DELETE("/companies/:id", companies.Delete)

			// Users routes
			users := controllers.NewUserController()
			protected.GET("/users", users.List)
			protected.GET("/users/:id", users.Get)
			protected.POST("/users", users.Create)
			protected.PUT("/users/:id", users.Update)
			protected.DELETE("/users/:id", users.Delete)
		}

		c.Next()
	}
}
