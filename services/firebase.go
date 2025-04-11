package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	fbstorage "firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

type User struct {
	CompanyID string `firestore:"company_id"`
	CreatedAt int64  `firestore:"created_at"`
	Email     string `firestore:"email"`
	IsActive  bool   `firestore:"is_active"`
	UpdatedAt int64  `firestore:"updated_at"`
	Username  string `firestore:"username"`
}

type FirebaseService struct {
	app        *firebase.App
	authClient *auth.Client
	db         *firestore.Client
	storage    *fbstorage.Client
	mu         sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewFirebaseService(parentCtx context.Context, credentialsPath string) (*FirebaseService, error) {
	log.Printf("Initializing Firebase service with credentials: %s", credentialsPath)

	// Create our own context that won't be cancelled
	ctx, cancel := context.WithCancel(context.Background())

	// Verify credentials file exists and is readable
	if _, err := os.Stat(credentialsPath); os.IsNotExist(err) {
		cancel()
		return nil, fmt.Errorf("credentials file does not exist: %s", credentialsPath)
	}

	// Get absolute path to credentials
	absPath, err := filepath.Abs(credentialsPath)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to get absolute path for credentials: %v", err)
	}

	// Create Firebase app with credentials
	opt := option.WithCredentialsFile(absPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create Firebase app: %v", err)
	}
	log.Printf("Firebase app initialized successfully")

	// Initialize Auth client
	authClient, err := app.Auth(ctx)
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create Auth client: %v", err)
	}
	log.Printf("Auth client initialized successfully")

	// Initialize Firestore client with retry logic
	var db *firestore.Client
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		db, err = app.Firestore(ctx)
		if err == nil {
			break
		}
		log.Printf("Attempt %d to initialize Firestore failed: %v", i+1, err)
		if i < maxRetries-1 {
			time.Sleep(time.Second * 2)
		}
	}
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create Firestore client after %d attempts: %v", maxRetries, err)
	}
	log.Printf("Firestore client initialized successfully")

	// Verify clients are not nil
	if db == nil {
		cancel()
		return nil, fmt.Errorf("Firestore client is nil after initialization")
	}
	if authClient == nil {
		cancel()
		return nil, fmt.Errorf("Auth client is nil after initialization")
	}

	service := &FirebaseService{
		app:        app,
		authClient: authClient,
		db:         db,
		ctx:        ctx,
		cancel:     cancel,
	}

	// Test Firestore connection
	testCtx, testCancel := context.WithTimeout(ctx, 5*time.Second)
	defer testCancel()

	// Try to create a test document to verify write access
	_, err = db.Collection("test").Doc("test").Set(testCtx, map[string]interface{}{
		"timestamp": time.Now().Unix(),
	})
	if err != nil {
		log.Printf("Warning: Firestore connection test failed: %v", err)
		// Don't fail initialization if test fails, just log warning
	}

	// Clean up test document
	_, err = db.Collection("test").Doc("test").Delete(testCtx)
	if err != nil {
		log.Printf("Warning: Failed to clean up test document: %v", err)
	}

	// Start a goroutine to monitor the parent context
	go func() {
		select {
		case <-parentCtx.Done():
			log.Printf("Parent context cancelled, cleaning up Firebase service")
			service.Close()
		case <-ctx.Done():
			// Our own context was cancelled, nothing to do
		}
	}()

	return service, nil
}

// GetFirestore returns the Firestore client
func (s *FirebaseService) GetFirestore() *firestore.Client {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s == nil {
		log.Printf("Warning: GetFirestore called on nil FirebaseService")
		return nil
	}
	if s.db == nil {
		log.Printf("Warning: Firestore client is nil in FirebaseService")
		return nil
	}

	return s.db
}

// GetAuthClient returns the Firebase Auth client
func (s *FirebaseService) GetAuthClient() *auth.Client {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s == nil {
		log.Printf("Warning: GetAuthClient called on nil FirebaseService")
		return nil
	}
	if s.authClient == nil {
		log.Printf("Warning: Auth client is nil in FirebaseService")
		return nil
	}
	return s.authClient
}

// GetApp returns the Firebase app
func (s *FirebaseService) GetApp() *firebase.App {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s == nil {
		log.Printf("Warning: GetApp called on nil FirebaseService")
		return nil
	}
	return s.app
}

// Close closes the Firebase service and its clients
func (s *FirebaseService) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s == nil {
		return fmt.Errorf("cannot close nil FirebaseService")
	}
	if s.cancel != nil {
		s.cancel()
	}
	if s.db == nil {
		return fmt.Errorf("cannot close nil Firestore client")
	}
	return s.db.Close()
}

func (s *FirebaseService) CreateUser(ctx context.Context, email, password, username string) (*User, error) {
	// Create user in Firebase Auth
	userRecord, err := s.authClient.CreateUser(ctx, (&auth.UserToCreate{}).
		Email(email).
		Password(password))
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	// Create user document in Firestore
	user := &User{
		CompanyID: "",
		CreatedAt: now,
		Email:     email,
		IsActive:  true,
		UpdatedAt: now,
		Username:  username,
	}

	_, err = s.db.Collection("users").Doc(userRecord.UID).Set(ctx, user)
	if err != nil {
		// If Firestore creation fails, delete the Auth user
		if deleteErr := s.authClient.DeleteUser(ctx, userRecord.UID); deleteErr != nil {
			log.Printf("Failed to delete Auth user after Firestore creation failed: %v", deleteErr)
		}
		return nil, err
	}

	return user, nil
}

func (s *FirebaseService) DeleteUser(ctx context.Context, userID string) error {
	// Delete from Firebase Auth
	if err := s.authClient.DeleteUser(ctx, userID); err != nil {
		return err
	}

	// Delete from Firestore
	_, err := s.db.Collection("users").Doc(userID).Delete(ctx)
	if err != nil {
		log.Printf("Warning: Failed to delete user from Firestore: %v", err)
		// Don't return error here as Auth deletion was successful
	}

	return nil
}

func (s *FirebaseService) UpdateUserCompany(ctx context.Context, userID, companyID string) error {
	_, err := s.db.Collection("users").Doc(userID).Update(ctx, []firestore.Update{
		{
			Path:  "company_id",
			Value: companyID,
		},
		{
			Path:  "updated_at",
			Value: firestore.ServerTimestamp,
		},
	})
	return err
}

// UploadBase64Image uploads a base64 encoded image to Firebase Storage
func (s *FirebaseService) UploadBase64Image(ctx context.Context, base64Data, folder, filename string) (string, error) {
	if s.storage == nil {
		var err error
		s.storage, err = s.app.Storage(ctx)
		if err != nil {
			return "", fmt.Errorf("failed to initialize storage client: %v", err)
		}
	}

	// Decode base64 data
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %v", err)
	}

	// Create a unique filename if none provided
	if filename == "" {
		filename = fmt.Sprintf("%d-%s.jpg", time.Now().Unix(), generateRandomString(8))
	}

	// Create the full path including folder
	fullPath := fmt.Sprintf("%s/%s", folder, filename)

	// Get bucket handle
	bucket, err := s.storage.DefaultBucket()
	if err != nil {
		return "", fmt.Errorf("failed to get default bucket: %v", err)
	}

	// Create the file in Firebase Storage
	obj := bucket.Object(fullPath)
	writer := obj.NewWriter(ctx)

	// Set content type
	writer.ContentType = "image/jpeg"

	// Write the decoded image data
	if _, err := writer.Write(imageData); err != nil {
		return "", fmt.Errorf("failed to write image to storage: %v", err)
	}

	// Close the writer
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Make the file public
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", fmt.Errorf("failed to make file public: %v", err)
	}

	// Get the public URL
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get object attributes: %v", err)
	}

	return attrs.MediaLink, nil
}

// generateRandomString generates a random string of the specified length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
