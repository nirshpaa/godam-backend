package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
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

type SalesOrder struct {
	ID                string             `firestore:"id"`
	Code              string             `firestore:"code"`
	Date              string             `firestore:"date"`
	CustomerID        string             `firestore:"customer_id"`
	SalesmanID        string             `firestore:"salesman_id"`
	CompanyID         string             `firestore:"company_id"`
	TotalAmount       float64            `firestore:"total_amount"`
	Discount          float64            `firestore:"discount"`
	AdditionalDisc    float64            `firestore:"additional_disc"`
	Status            string             `firestore:"status"`
	SalesOrderDetails []SalesOrderDetail `firestore:"sales_order_details"`
}

type SalesOrderDetail struct {
	ProductID   string  `firestore:"product_id"`
	ProductCode string  `firestore:"product_code"`
	Quantity    float64 `firestore:"quantity"`
	UnitPrice   float64 `firestore:"unit_price"`
	TotalPrice  float64 `firestore:"total_price"`
	Discount    float64 `firestore:"discount"`
}

func main() {
	// Get the absolute path to the firebase-credentials.json file
	credPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "nirshpaa", "godam-backend", "firebase-credentials.json")

	// Set the GOOGLE_APPLICATION_CREDENTIALS environment variable
	if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credPath); err != nil {
		log.Fatalf("Failed to set GOOGLE_APPLICATION_CREDENTIALS: %v", err)
	}

	ctx := context.Background()

	// Initialize Firestore client
	client, err := firestore.NewClient(ctx, "godam-45852", option.WithCredentialsFile("firebase-credentials.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Seed products
	log.Println("Seeding products...")
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

	// Seed sales orders
	log.Println("Seeding sales orders...")
	salesOrders := []SalesOrder{
		{
			ID:             uuid.New().String(),
			Code:           "SO-001",
			Date:           time.Now().AddDate(0, 0, -2).Format(time.RFC3339),
			CustomerID:     "CUST001",
			SalesmanID:     "SALES001",
			CompanyID:      "COMP001",
			TotalAmount:    0, // Will be calculated
			Discount:       5.0,
			AdditionalDisc: 2.0,
			Status:         "completed",
			SalesOrderDetails: []SalesOrderDetail{
				{
					ProductID:   "PRD001",
					ProductCode: "PRD001",
					Quantity:    2,
					UnitPrice:   24.99,
					TotalPrice:  49.98,
					Discount:    0,
				},
				{
					ProductID:   "PRD002",
					ProductCode: "PRD002",
					Quantity:    3,
					UnitPrice:   19.99,
					TotalPrice:  59.97,
					Discount:    0,
				},
			},
		},
		{
			ID:             uuid.New().String(),
			Code:           "SO-002",
			Date:           time.Now().AddDate(0, 0, -1).Format(time.RFC3339),
			CustomerID:     "CUST002",
			SalesmanID:     "SALES002",
			CompanyID:      "COMP001",
			TotalAmount:    0, // Will be calculated
			Discount:       10.0,
			AdditionalDisc: 0,
			Status:         "pending",
			SalesOrderDetails: []SalesOrderDetail{
				{
					ProductID:   "PRD003",
					ProductCode: "PRD003",
					Quantity:    1,
					UnitPrice:   199.99,
					TotalPrice:  199.99,
					Discount:    0,
				},
				{
					ProductID:   "PRD004",
					ProductCode: "PRD004",
					Quantity:    2,
					UnitPrice:   149.99,
					TotalPrice:  299.98,
					Discount:    0,
				},
			},
		},
		{
			ID:             uuid.New().String(),
			Code:           "SO-003",
			Date:           time.Now().Format(time.RFC3339),
			CustomerID:     "CUST003",
			SalesmanID:     "SALES001",
			CompanyID:      "COMP001",
			TotalAmount:    0, // Will be calculated
			Discount:       0,
			AdditionalDisc: 0,
			Status:         "completed",
			SalesOrderDetails: []SalesOrderDetail{
				{
					ProductID:   "PRD005",
					ProductCode: "PRD005",
					Quantity:    5,
					UnitPrice:   49.99,
					TotalPrice:  249.95,
					Discount:    0,
				},
			},
		},
	}

	// Add sales orders to Firestore
	for _, order := range salesOrders {
		// Calculate total amount
		var totalAmount float64
		for _, detail := range order.SalesOrderDetails {
			totalAmount += detail.TotalPrice
		}
		order.TotalAmount = totalAmount - order.Discount - order.AdditionalDisc

		_, err := client.Collection("sales_orders").Doc(order.ID).Set(ctx, order)
		if err != nil {
			log.Printf("Failed to add sales order %s: %v", order.Code, err)
		} else {
			log.Printf("Successfully added sales order %s", order.Code)
		}
	}

	log.Println("Finished seeding all data")
}
