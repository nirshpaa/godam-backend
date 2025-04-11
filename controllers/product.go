package controllers

import (
	"context"
	"net/http"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/services"
)

// ProductController handles product-related operations
type ProductController struct {
	dbRef                   *db.Ref
	imageRecognitionService *services.ImageRecognitionService
}

// NewProductController creates a new product controller
func NewProductController(imageRecognitionService *services.ImageRecognitionService) *ProductController {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	dbClient, err := app.Database(context.Background())
	if err != nil {
		panic(err)
	}

	return &ProductController{
		dbRef:                   dbClient.NewRef("products"),
		imageRecognitionService: imageRecognitionService,
	}
}

// List retrieves all products
func (p *ProductController) List(c *gin.Context) {
	var products []map[string]interface{}
	if err := p.dbRef.Get(context.Background(), &products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Get retrieves a single product by ID
func (p *ProductController) Get(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var product map[string]interface{}
	if err := p.dbRef.Child(code).Get(context.Background(), &product); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Create creates a new product
func (p *ProductController) Create(c *gin.Context) {
	var product map[string]interface{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRef, err := p.dbRef.Push(context.Background(), product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	product["id"] = newRef.Key
	c.JSON(http.StatusCreated, product)
}

// Update updates an existing product
func (p *ProductController) Update(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var product map[string]interface{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.dbRef.Child(code).Set(context.Background(), product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	product["code"] = code
	c.JSON(http.StatusOK, product)
}

// Delete removes a product
func (p *ProductController) Delete(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	if err := p.dbRef.Child(code).Delete(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// ScanProduct handles product scanning requests
func (c *ProductController) ScanProduct(ctx *gin.Context) {
	// Parse multipart form
	if err := ctx.Request.ParseMultipartForm(10 << 20); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form: " + err.Error()})
		return
	}

	// Get file from form
	file, _, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided: " + err.Error()})
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded image
	tempFile, err := os.CreateTemp("", "scan-*.png")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file: " + err.Error()})
		return
	}
	defer os.Remove(tempFile.Name())

	// Copy the uploaded file to the temporary file
	if _, err := tempFile.ReadFrom(file); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file: " + err.Error()})
		return
	}

	// Process the image
	result, err := c.imageRecognitionService.ProcessImage(tempFile.Name())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process image: " + err.Error()})
		return
	}

	// Return the recognition result
	ctx.JSON(http.StatusOK, gin.H{
		"success": result.Success,
		"data":    result.Data,
	})
}
