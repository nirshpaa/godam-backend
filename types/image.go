package types

// ImageRecognitionResult represents the result of image recognition
type ImageRecognitionResult struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

// ImageTrainingRequest represents a request to train on a new image
type ImageTrainingRequest struct {
	ProductName string `json:"product_name"`
	ImageBase64 string `json:"image_base64"`
}
