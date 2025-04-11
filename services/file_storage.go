package services

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileStorageService handles local file storage operations
type FileStorageService struct {
	basePath string
}

// NewFileStorageService creates a new file storage service
func NewFileStorageService(basePath string) *FileStorageService {
	return &FileStorageService{
		basePath: basePath,
	}
}

// SaveBase64Image saves a base64 encoded image to the local filesystem
func (s *FileStorageService) SaveBase64Image(base64Data, folder string) (string, error) {
	// Remove the data URL prefix if present
	if len(base64Data) > 0 && base64Data[0] == 'd' {
		// Find the base64 part after the comma
		commaIndex := strings.Index(base64Data, ",")
		if commaIndex != -1 {
			base64Data = base64Data[commaIndex+1:]
		}
	}

	// Decode base64 data
	imageData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %v", err)
	}

	// Create a unique filename
	filename := fmt.Sprintf("%d-%s.jpg", time.Now().Unix(), generateUniqueString(8))

	// Create the full directory path
	dirPath := filepath.Join(s.basePath, folder)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Create the full file path
	fullPath := filepath.Join(dirPath, filename)

	// Write the file
	if err := os.WriteFile(fullPath, imageData, 0644); err != nil {
		return "", fmt.Errorf("failed to write image file: %v", err)
	}

	// Return the relative path that can be used in URLs
	relativePath := filepath.Join(folder, filename)
	return relativePath, nil
}

// SaveFile saves a file directly to the local filesystem
func (s *FileStorageService) SaveFile(file interface{}, folder string) (string, error) {
	// Type assert to *multipart.FileHeader
	fileHeader, ok := file.(*multipart.FileHeader)
	if !ok {
		return "", fmt.Errorf("invalid file type: expected *multipart.FileHeader")
	}

	// Create the full directory path
	dirPath := filepath.Join(s.basePath, folder)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %v", err)
	}

	// Create the full file path
	fullPath := filepath.Join(dirPath, filepath.Base(fileHeader.Filename))

	// Open the source file
	src, err := fileHeader.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open source file: %v", err)
	}
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(fullPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy the file contents
	if _, err := io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	// Return the relative path that can be used in URLs
	relativePath := filepath.Join(folder, filepath.Base(fileHeader.Filename))
	return relativePath, nil
}

// generateUniqueString generates a random string of the specified length
func generateUniqueString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}
