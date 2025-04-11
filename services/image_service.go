package services

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
)

type ImageService struct {
	storageClient *storage.Client
	bucketName    string
}

func NewImageService(ctx context.Context, storageClient *storage.Client) (*ImageService, error) {
	bucketName := os.Getenv("STORAGE_BUCKET")
	if bucketName == "" {
		return nil, fmt.Errorf("storage bucket name not configured")
	}

	return &ImageService{
		storageClient: storageClient,
		bucketName:    bucketName,
	}, nil
}

// UploadImage handles image upload to Firebase Storage
func (s *ImageService) UploadImage(ctx context.Context, base64Data string, path string) (string, error) {
	// Remove the data URL prefix if present
	base64Data = strings.TrimPrefix(base64Data, "data:image/jpeg;base64,")
	base64Data = strings.TrimPrefix(base64Data, "data:image/png;base64,")

	// Decode base64 data
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %v", err)
	}

	// Create a unique filename
	filename := fmt.Sprintf("%s/%d.jpg", path, time.Now().UnixNano())

	// Get bucket handle
	bucket := s.storageClient.Bucket(s.bucketName)
	obj := bucket.Object(filename)

	// Upload the image
	wc := obj.NewWriter(ctx)
	if _, err := wc.Write(imageData); err != nil {
		return "", fmt.Errorf("failed to write image: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Get the public URL
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get object attributes: %v", err)
	}

	return attrs.MediaLink, nil
}

// DeleteImage handles image deletion from Firebase Storage
func (s *ImageService) DeleteImage(ctx context.Context, imageURL string) error {
	// Extract the path from the URL
	path := strings.Split(imageURL, "/o/")[1]
	path = strings.Split(path, "?")[0]
	decodedPath, err := base64.StdEncoding.DecodeString(path)
	if err != nil {
		return fmt.Errorf("failed to decode path: %v", err)
	}

	// Get bucket handle
	bucket := s.storageClient.Bucket(s.bucketName)
	obj := bucket.Object(string(decodedPath))

	// Delete the object
	if err := obj.Delete(ctx); err != nil {
		return fmt.Errorf("failed to delete image: %v", err)
	}

	return nil
}
