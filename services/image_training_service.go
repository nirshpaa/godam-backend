package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/nirshpaa/godam-backend/types"
)

type ImageTrainingService struct {
	apiURL string
}

func NewImageTrainingService() *ImageTrainingService {
	return &ImageTrainingService{
		apiURL: os.Getenv("TRAINING_API_ENDPOINT"),
	}
}

// SaveProductImage sends the image to the Python API for training
func (s *ImageTrainingService) SaveProductImage(productName, imageBase64 string) error {
	// Prepare the request payload
	payload := types.ImageTrainingRequest{
		ProductName: productName,
		ImageBase64: imageBase64,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	// Send request to Python API
	resp, err := http.Post(
		fmt.Sprintf("%s/api/train/image", s.apiURL),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return fmt.Errorf("failed to send request to training API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("training API returned status %d", resp.StatusCode)
	}

	return nil
}

// ScheduleTraining schedules the training process to run daily
func (s *ImageTrainingService) ScheduleTraining() {
	go func() {
		for {
			// Run training at 2 AM every day
			now := time.Now()
			next := now.AddDate(0, 0, 1)
			next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location())
			time.Sleep(next.Sub(now))

			// Execute training via API
			if err := s.runTraining(); err != nil {
				fmt.Printf("Training failed: %v\n", err)
			}
		}
	}()
}

func (s *ImageTrainingService) runTraining() error {
	// Send request to Python API to start training
	resp, err := http.Post(
		fmt.Sprintf("%s/api/train/start", s.apiURL),
		"application/json",
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to start training: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("training API returned status %d", resp.StatusCode)
	}

	return nil
}
