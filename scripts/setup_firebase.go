package main

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Initialize Firebase
	opt := option.WithCredentialsFile("firebase-credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	firestore, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Failed to get Firestore client: %v", err)
	}
	defer firestore.Close()

	// Get Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to get Auth client: %v", err)
	}

	// Initialize with sample data
	if err := initializeCollections(firestore, authClient); err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	}

	log.Println("Successfully initialized Firebase with sample data")
}

func initializeCollections(firestore *firestore.Client, authClient *auth.Client) error {
	ctx := context.Background()

	// Initialize companies
	company := map[string]interface{}{
		"code":       "DM",
		"name":       "Dummy",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	companyRef, _, err := firestore.Collection("companies").Add(ctx, company)
	if err != nil {
		return err
	}

	// Initialize branches
	branch := map[string]interface{}{
		"company_id": companyRef.ID,
		"code":       "123",
		"name":       "Toko Bagus",
		"type":       "s",
		"address":    "jalan jalan",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	branchRef, _, err := firestore.Collection("branches").Add(ctx, branch)
	if err != nil {
		return err
	}

	// Initialize shelves
	shelf := map[string]interface{}{
		"branch_id":  branchRef.ID,
		"code":       "SHV-01",
		"capacity":   1000,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	_, _, err = firestore.Collection("shelves").Add(ctx, shelf)
	if err != nil {
		return err
	}

	// Initialize categories
	category := map[string]interface{}{
		"name": "Accesories",
	}
	categoryRef, _, err := firestore.Collection("categories").Add(ctx, category)
	if err != nil {
		return err
	}

	// Initialize product categories
	productCategory := map[string]interface{}{
		"company_id":  companyRef.ID,
		"category_id": categoryRef.ID,
		"name":        "Furniture",
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}
	productCategoryRef, _, err := firestore.Collection("product_categories").Add(ctx, productCategory)
	if err != nil {
		return err
	}

	// Initialize brands
	brand := map[string]interface{}{
		"company_id": companyRef.ID,
		"code":       "BRAND-01",
		"name":       "Brand Test",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	brandRef, _, err := firestore.Collection("brands").Add(ctx, brand)
	if err != nil {
		return err
	}

	// Initialize suppliers
	supplier := map[string]interface{}{
		"company_id": companyRef.ID,
		"code":       "SUP_01",
		"name":       "Supplier Test",
		"address":    "jalan supplier",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	_, _, err = firestore.Collection("suppliers").Add(ctx, supplier)
	if err != nil {
		return err
	}

	// Initialize products
	products := []map[string]interface{}{
		{
			"company_id":             companyRef.ID,
			"brand_id":               brandRef.ID,
			"product_category_id":    productCategoryRef.ID,
			"code":                   "PROD-01",
			"name":                   "Product Satu",
			"purchase_price":         0,
			"sale_price":             1000,
			"minimum_stock":          25,
			"image_url":              "",
			"barcode_value":          "",
			"image_recognition_data": "",
			"created_at":             time.Now(),
			"updated_at":             time.Now(),
		},
		{
			"company_id":             companyRef.ID,
			"brand_id":               brandRef.ID,
			"product_category_id":    productCategoryRef.ID,
			"code":                   "PROD-02",
			"name":                   "Product Dua",
			"purchase_price":         0,
			"sale_price":             500,
			"minimum_stock":          1000,
			"image_url":              "",
			"barcode_value":          "",
			"image_recognition_data": "",
			"created_at":             time.Now(),
			"updated_at":             time.Now(),
		},
	}
	for _, product := range products {
		_, _, err := firestore.Collection("products").Add(ctx, product)
		if err != nil {
			return err
		}
	}

	// Initialize access
	access := map[string]interface{}{
		"name":       "root",
		"alias":      "root",
		"created_at": time.Now(),
	}
	accessRef, _, err := firestore.Collection("access").Add(ctx, access)
	if err != nil {
		return err
	}

	// Initialize roles
	role := map[string]interface{}{
		"name":       "superadmin",
		"company_id": companyRef.ID,
		"created_at": time.Now(),
	}
	roleRef, _, err := firestore.Collection("roles").Add(ctx, role)
	if err != nil {
		return err
	}

	// Initialize access_roles
	accessRole := map[string]interface{}{
		"access_id":  accessRef.ID,
		"role_id":    roleRef.ID,
		"created_at": time.Now(),
	}
	_, _, err = firestore.Collection("access_roles").Add(ctx, accessRole)
	if err != nil {
		return err
	}

	// Initialize users
	users := []struct {
		email    string
		password string
		userData map[string]interface{}
	}{
		{
			email:    "admin@admin.com",
			password: "admin123", // You should change this password
			userData: map[string]interface{}{
				"username":   "admin",
				"email":      "admin@admin.com",
				"is_active":  true,
				"company_id": companyRef.ID,
				"created_at": time.Now(),
				"updated_at": time.Now(),
			},
		},
		{
			email:    "nishan@gmail.com",
			password: "nishan123", // You should change this password
			userData: map[string]interface{}{
				"username":   "nishanpandit",
				"email":      "nishan@gmail.com",
				"is_active":  true,
				"company_id": companyRef.ID,
				"branch_id":  branchRef.ID,
				"created_at": time.Now(),
				"updated_at": time.Now(),
			},
		},
	}

	for _, user := range users {
		// Create user in Firebase Authentication
		params := (&auth.UserToCreate{}).
			Email(user.email).
			Password(user.password).
			EmailVerified(true).
			Disabled(false)

		firebaseUser, err := authClient.CreateUser(ctx, params)
		if err != nil {
			return err
		}

		// Store user data in Firestore using the Firebase Auth UID
		_, err = firestore.Collection("users").Doc(firebaseUser.UID).Set(ctx, user.userData)
		if err != nil {
			return err
		}

		// Initialize roles_users
		roleUser := map[string]interface{}{
			"role_id":    roleRef.ID,
			"user_id":    firebaseUser.UID,
			"created_at": time.Now(),
		}
		_, _, err = firestore.Collection("roles_users").Add(ctx, roleUser)
		if err != nil {
			return err
		}

		log.Printf("Created user: %s with email: %s\n", firebaseUser.UID, user.email)
	}

	return nil
}
