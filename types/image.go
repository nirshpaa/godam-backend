package types

// CreateProductData represents the data needed to create a new product
type CreateProductData struct {
	Action        string  `json:"action"`
	SuggestedName string  `json:"suggested_name"`
	Confidence    float64 `json:"confidence"`
}

// ImageRecognitionResult represents the result of image recognition
type ImageRecognitionResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"` // Can be either string (product ID) or CreateProductData
}

// ImageTrainingRequest represents a request to train on a new image
type ImageTrainingRequest struct {
	ProductName string `json:"product_name"`
	ImageBase64 string `json:"image_base64"`
}
