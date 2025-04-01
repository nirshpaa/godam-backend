package controllers

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"github.com/gin-gonic/gin"
)

// CompanyController handles company-related operations
type CompanyController struct {
	dbRef *db.Ref
}

// NewCompanyController creates a new company controller
func NewCompanyController() *CompanyController {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	dbClient, err := app.Database(context.Background())
	if err != nil {
		panic(err)
	}

	return &CompanyController{
		dbRef: dbClient.NewRef("companies"),
	}
}

// List retrieves all companies
func (c *CompanyController) List(ctx *gin.Context) {
	var companies []map[string]interface{}
	if err := c.dbRef.Get(context.Background(), &companies); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch companies"})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}

// Get retrieves a single company by ID
func (c *CompanyController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	var company map[string]interface{}
	if err := c.dbRef.Child(id).Get(context.Background(), &company); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

// Create creates a new company
func (c *CompanyController) Create(ctx *gin.Context) {
	var company map[string]interface{}
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRef, err := c.dbRef.Push(context.Background(), company)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	company["id"] = newRef.Key
	ctx.JSON(http.StatusCreated, company)
}

// Update updates an existing company
func (c *CompanyController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	var company map[string]interface{}
	if err := ctx.ShouldBindJSON(&company); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.dbRef.Child(id).Set(context.Background(), company); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	company["id"] = id
	ctx.JSON(http.StatusOK, company)
}

// Delete removes a company
func (c *CompanyController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	if err := c.dbRef.Child(id).Delete(context.Background()); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
