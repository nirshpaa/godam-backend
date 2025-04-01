package controllers

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
)

// ProductController handles product-related operations
type ProductController struct {
	dbRef *db.Ref
}

// NewProductController creates a new product controller
func NewProductController() *ProductController {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	dbClient, err := app.Database(context.Background())
	if err != nil {
		panic(err)
	}

	return &ProductController{
		dbRef: dbClient.NewRef("products"),
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
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var product map[string]interface{}
	if err := p.dbRef.Child(id).Get(context.Background(), &product); err != nil {
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
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	var product map[string]interface{}
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := p.dbRef.Child(id).Set(context.Background(), product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	product["id"] = id
	c.JSON(http.StatusOK, product)
}

// Delete removes a product
func (p *ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product ID is required"})
		return
	}

	if err := p.dbRef.Child(id).Delete(context.Background()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

// ScanImage handles product image scanning
func (p *ProductController) ScanImage(c *gin.Context) {
	// TODO: Implement image scanning with Firebase Storage
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Image scanning not implemented yet"})
}
