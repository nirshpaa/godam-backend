package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	apiTest "github.com/nishanpandit/inventory/controllers/tests"
	"github.com/nishanpandit/inventory/libraries/config"
	"github.com/nishanpandit/inventory/routing"
	"github.com/nishanpandit/inventory/schema"
	testUtil "github.com/nishanpandit/inventory/tests"
)

var token string

func TestMain(t *testing.T) {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	db, teardown := testUtil.NewUnit(t)
	defer teardown()

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	// Initialize Gin router
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Setup routes
	router.Use(routing.SetupRoutes())

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
