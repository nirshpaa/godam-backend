package main

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/cmd/server/setup"
	apiTest "github.com/nirshpaa/godam-backend/controllers/tests"
	"github.com/nirshpaa/godam-backend/libraries/config"
	"github.com/nirshpaa/godam-backend/libraries/firebase"
	"github.com/nirshpaa/godam-backend/services"
)

var token string

func TestMain(t *testing.T) {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	// Initialize Firebase
	_, err := firebase.Initialize()
	if err != nil {
		t.Fatal(err)
	}

	// Create Firebase service
	firebaseService, err := services.NewFirebaseService(context.Background(), "firebase-credentials.json")
	if err != nil {
		t.Fatal(err)
	}

	// Initialize Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Setup routes
	setup.SetupRoutes(router, firebaseService)

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// api test for auths
	{
		auths := apiTest.Auths{App: router}
		t.Run("ApiLogin", auths.Login)
		token = auths.Token
	}

	// api test for users
	{
		users := apiTest.Users{App: router, Token: token}
		t.Run("APiUsersList", users.List)
		t.Run("APiUsersCrud", users.Crud)
	}

	// api test for access
	{
		access := apiTest.Access{App: router, Token: token}
		t.Run("APiAccessList", access.List)
	}

	// api test for roles
	{
		roles := apiTest.Roles{App: router, Token: token}
		t.Run("APiRolesList", roles.List)
		t.Run("APiRolesCrud", roles.Crud)
	}

	// api test for regions
	{
		regions := apiTest.Regions{App: router, Token: token}
		t.Run("APiRegionsTest", regions.Run)
	}

	// api test for brands
	{
		brands := apiTest.Brands{App: router, Token: token}
		t.Run("APiBrandsCrud", brands.Run)
	}

	// api test for product categories
	{
		productCategories := apiTest.ProductCategories{App: router, Token: token}
		t.Run("APiProductCategoriesCrud", productCategories.Run)
	}

	// api test for products
	{
		products := apiTest.Products{App: router, Token: token}
		t.Run("APiProductsCrud", products.Run)
	}
}
