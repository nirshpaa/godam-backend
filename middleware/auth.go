package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/services"
)

// contextKey is a custom type for context keys
type contextKey string

// UserIDKey is the key used to store the user ID in the context
const UserIDKey contextKey = "userID"

// CompanyIDKey is the key used to store the company ID in the context
const CompanyIDKey contextKey = "company_id"

// AuthMiddleware validates Firebase ID tokens and extracts company ID
func AuthMiddleware(firebaseService *services.FirebaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Skip auth for health check endpoint
		if c.Request.URL.Path == "/health" {
			c.Next()
			return
		}

		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check if the header is in the format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// Verify the Firebase ID token using a background context
		token := parts[1]
		ctx := context.Background()
		decodedToken, err := firebaseService.GetAuthClient().VerifyIDToken(ctx, token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		// Get the company ID from the header
		companyID := c.GetHeader("X-Company-ID")
		if companyID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
			c.Abort()
			return
		}

		// Set the user ID and company ID in the context
		c.Set("userID", decodedToken.UID)
		c.Set("company_id", companyID)
		c.Next()
	}
}

// SetUserID sets the user ID in the context
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

// GetUserID retrieves the user ID from the context
func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}

// SetCompanyID sets the company ID in the context
func SetCompanyID(ctx context.Context, companyID string) context.Context {
	return context.WithValue(ctx, CompanyIDKey, companyID)
}

// GetCompanyID retrieves the company ID from the context
func GetCompanyID(ctx context.Context) (string, bool) {
	companyID, ok := ctx.Value(CompanyIDKey).(string)
	return companyID, ok
}
