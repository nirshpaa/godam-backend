package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	fb "firebase.google.com/go/v4"
	"github.com/nirshpaa/godam-backend/libraries/firebase"
)

type Product struct {
	ID                   string    `json:"id" firestore:"id,omitempty"`
	Code                 string    `json:"code" firestore:"code"`
	Name                 string    `json:"name" firestore:"name"`
	PurchasePrice        int64     `json:"purchase_price" firestore:"purchase_price"`
	SalePrice            int64     `json:"sale_price" firestore:"sale_price"`
	MinimumStock         int64     `json:"minimum_stock" firestore:"minimum_stock"`
	ImageURL             string    `json:"image_url" firestore:"image_url,omitempty"`
	CompanyID            string    `json:"company_id" firestore:"company_id,omitempty"`
	BrandID              string    `json:"brand_id" firestore:"brand_id,omitempty"`
	CategoryID           string    `json:"product_category_id" firestore:"product_category_id,omitempty"`
	BarcodeValue         string    `json:"barcode_value" firestore:"barcode_value,omitempty"`
	ImageRecognitionData string    `json:"image_recognition_data" firestore:"image_recognition_data,omitempty"`
	CreatedAt            time.Time `json:"created_at" firestore:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" firestore:"updated_at"`
}

type IProductHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	ScanProduct(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	app *fb.App
}

func NewProductHandler(app *fb.App) IProductHandler {
	return &productHandler{app: app}
}

func (h *productHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uid := ctx.Value("uid").(string)

	// Create a logger that writes to stdout
	logger := log.New(os.Stdout, "[Product Handler] ", log.LstdFlags|log.Lshortfile)

	logger.Printf("\n=== Starting Product List Handler ===\n")
	logger.Printf("User ID from context: %s\n", uid)

	// Initialize Firestore client
	client := firebase.GetFirestore()
	if client == nil {
		logger.Println("ERROR: Firestore client is nil")
		http.Error(w, "Failed to initialize database", http.StatusInternalServerError)
		return
	}

	// First, get the user document to get the company_id
	logger.Println("\nFetching user document...")
	logger.Printf("Looking for user document with ID: %s\n", uid)

	userDoc, err := client.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		logger.Printf("ERROR fetching user document: %v\n", err)
		http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
		return
	}

	logger.Printf("User document exists: %v\n", userDoc.Exists())
	logger.Printf("User document data: %+v\n", userDoc.Data())

	userData := userDoc.Data()
	companyID, exists := userData["company_id"]
	if !exists {
		logger.Println("ERROR: company_id not found in user document")
		logger.Printf("Available fields in user document: %v\n", userData)
		http.Error(w, "User company not found", http.StatusInternalServerError)
		return
	}

	companyIDStr, ok := companyID.(string)
	if !ok {
		logger.Printf("ERROR: company_id is not a string: %T\n", companyID)
		http.Error(w, "Invalid company ID format", http.StatusInternalServerError)
		return
	}

	logger.Printf("Found company_id: %s\n", companyIDStr)

	// Now fetch products for this company
	logger.Printf("\nFetching products for company_id: %s\n", companyIDStr)
	query := client.Collection("products").Where("company_id", "==", companyIDStr)

	// Log the query for debugging
	logger.Printf("Query: %+v\n", query)

	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		logger.Printf("ERROR fetching company products: %v\n", err)
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	logger.Printf("Found %d products for company %s\n", len(docs), companyIDStr)

	var products []Product
	for _, doc := range docs {
		data := doc.Data()
		logger.Printf("\nCompany Product Document ID: %s\n", doc.Ref.ID)
		logger.Printf("Raw data: %+v\n", data)

		// Create a new Product struct
		product := Product{
			ID:                   doc.Ref.ID,
			Code:                 getString(data, "code"),
			Name:                 getString(data, "name"),
			PurchasePrice:        getInt64(data, "purchase_price"),
			SalePrice:            getInt64(data, "sale_price"),
			MinimumStock:         getInt64(data, "minimum_stock"),
			ImageURL:             getString(data, "image_url"),
			CompanyID:            getString(data, "company_id"),
			BrandID:              getString(data, "brand_id"),
			CategoryID:           getString(data, "product_category_id"),
			BarcodeValue:         getString(data, "barcode_value"),
			ImageRecognitionData: getString(data, "image_recognition_data"),
			CreatedAt:            getTime(data, "created_at"),
			UpdatedAt:            getTime(data, "updated_at"),
		}

		// Log each field's value and type
		logger.Printf("Mapped Product Details:")
		logger.Printf("ID: %s", product.ID)
		logger.Printf("Code: %s (type: %T)", product.Code, product.Code)
		logger.Printf("Name: %s (type: %T)", product.Name, product.Name)
		logger.Printf("PurchasePrice: %d (type: %T)", product.PurchasePrice, product.PurchasePrice)
		logger.Printf("SalePrice: %d (type: %T)", product.SalePrice, product.SalePrice)
		logger.Printf("MinimumStock: %d (type: %T)", product.MinimumStock, product.MinimumStock)
		logger.Printf("CompanyID: %s (type: %T)", product.CompanyID, product.CompanyID)
		logger.Printf("BrandID: %s (type: %T)", product.BrandID, product.BrandID)
		logger.Printf("CategoryID: %s (type: %T)", product.CategoryID, product.CategoryID)

		products = append(products, product)
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		logger.Printf("ERROR encoding response: %v\n", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	logger.Println("\n=== End of Product List Handler ===")
}

// Helper functions to safely get values from Firestore data
func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case string:
			return v
		case int, int64, float64:
			return fmt.Sprintf("%v", v)
		default:
			return ""
		}
	}
	return ""
}

func getInt64(data map[string]interface{}, key string) int64 {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case int64:
			return v
		case int:
			return int64(v)
		case float64:
			return int64(v)
		case string:
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				return i
			}
		}
	}
	return 0
}

func getTime(data map[string]interface{}, key string) time.Time {
	if val, ok := data[key]; ok {
		switch v := val.(type) {
		case time.Time:
			return v
		case string:
			if t, err := time.Parse(time.RFC3339, v); err == nil {
				return t
			}
		}
	}
	return time.Time{}
}

func (h *productHandler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get product
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement create product
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update product
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement delete product
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *productHandler) ScanProduct(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement scan product
	w.WriteHeader(http.StatusNotImplemented)
}
