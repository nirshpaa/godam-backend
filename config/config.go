package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	StorageBucket string
	UploadPath    string
	BarcodeAPI    string
	CNNEndpoint   string
}

var config *Config

// GetConfig returns the application configuration
func GetConfig() *Config {
	if config == nil {
		config = &Config{
			StorageBucket: os.Getenv("STORAGE_BUCKET"),
			UploadPath:    os.Getenv("UPLOAD_PATH"),
			BarcodeAPI:    os.Getenv("BARCODE_API_ENDPOINT"),
			CNNEndpoint:   os.Getenv("CNN_API_ENDPOINT"),
		}
	}
	return config
}
