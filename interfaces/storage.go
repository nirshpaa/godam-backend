package interfaces

import "github.com/nirshpaa/godam-backend/types"

// FileStorage defines the interface for saving files
type FileStorage interface {
	SaveBase64Image(base64Data, path string) (string, error)
	SaveFile(file interface{}, path string) (string, error)
}

// ImageRecognition defines the interface for image recognition
type ImageRecognition interface {
	ProcessImage(imagePath string) (*types.ImageRecognitionResult, error)
}

type RecognitionResult struct {
	RecognitionSuccess bool   `json:"recognition_success"`
	RecognitionData    string `json:"recognition_data"`
}
