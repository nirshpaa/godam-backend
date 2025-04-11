package services

import (
	"cloud.google.com/go/firestore"
)

type ProductService struct {
	ImageRecognition *ImageRecognitionService
	FileStorage      *FileStorageService
}

func NewProductService(client *firestore.Client, barcodeEndpoint, cnnEndpoint string) *ProductService {
	return &ProductService{
		ImageRecognition: NewImageRecognitionService(client, barcodeEndpoint, cnnEndpoint),
		FileStorage:      NewFileStorageService("assets"),
	}
}
