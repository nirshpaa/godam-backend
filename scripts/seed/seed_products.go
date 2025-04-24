package seed

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Product struct {
	Code         string    `firestore:"code"`
	Name         string    `firestore:"name"`
	Description  string    `firestore:"description"`
	BarcodeValue string    `firestore:"barcode_value"`
	Category     string    `firestore:"category"`
	Price        float64   `firestore:"price"`
	CreatedAt    time.Time `firestore:"created_at"`
	UpdatedAt    time.Time `firestore:"updated_at"`
}

func SeedProducts() {
	ctx := context.Background()

	// Initialize Firestore client
	client, err := firestore.NewClient(ctx, "your-project-id", option.WithCredentialsFile("firebase-credentials.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sample products data
	products := []Product{
		{
			Code:         "PRD001",
			Name:         "Premium Coffee Beans",
			Description:  "Arabica coffee beans from Colombia",
			BarcodeValue: "123456789012",
			Category:     "Beverages",
			Price:        24.99,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Code:         "PRD002",
			Name:         "Organic Green Tea",
			Description:  "High-quality organic green tea leaves",
			BarcodeValue: "234567890123",
			Category:     "Beverages",
			Price:        19.99,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Code:         "PRD003",
			Name:         "Wireless Headphones",
			Description:  "Noise-cancelling wireless headphones",
			BarcodeValue: "345678901234",
			Category:     "Electronics",
			Price:        199.99,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Code:         "PRD004",
			Name:         "Smart Watch",
			Description:  "Fitness tracking smart watch",
			BarcodeValue: "456789012345",
			Category:     "Electronics",
			Price:        149.99,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			Code:         "PRD005",
			Name:         "Protein Powder",
			Description:  "Whey protein powder 2kg",
			BarcodeValue: "567890123456",
			Category:     "Supplements",
			Price:        49.99,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	// Add products to Firestore
	for _, product := range products {
		_, err := client.Collection("products").Doc(product.Code).Set(ctx, product)
		if err != nil {
			log.Printf("Failed to add product %s: %v", product.Code, err)
		} else {
			log.Printf("Successfully added product %s", product.Code)
		}
	}

	log.Println("Finished seeding products")
}
