package setup

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nirshpaa/godam-backend/handlers"
	"github.com/nirshpaa/godam-backend/middleware"
	"github.com/nirshpaa/godam-backend/models"
	"github.com/nirshpaa/godam-backend/services"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(router *gin.Engine, firebaseService *services.FirebaseService) {
	// Initialize services
	basePath := os.Getenv("UPLOAD_PATH")
	if basePath == "" {
		basePath = "./assets/products"
	}
	fileStorage := services.NewFileStorageService(basePath)

	barcodeEndpoint := os.Getenv("BARCODE_API_ENDPOINT")
	cnnEndpoint := os.Getenv("CNN_API_ENDPOINT")
	imageRecognition := services.NewImageRecognitionService(firebaseService.GetFirestore(), barcodeEndpoint, cnnEndpoint)

	imageTrainingService := services.NewImageTrainingService()

	// Health check route (no auth required)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Add company list route BEFORE auth middleware (publicly accessible)
	companyFirebase := models.NewCompanyFirebase(firebaseService.GetFirestore())
	if companyFirebase == nil {
		log.Fatalf("Failed to initialize company model")
	}
	companyHandler := handlers.NewCompanyHandler(companyFirebase)
	router.GET("/companies", companyHandler.List) // Public access to company list

	// Apply auth middleware to all routes except health check
	router.Use(middleware.AuthMiddleware(firebaseService))

	// Initialize Firebase models with error handling
	initModel := func(name string, initFunc func() error) {
		if err := initFunc(); err != nil {
			log.Fatalf("Failed to initialize %s model: %v", name, err)
		}
	}

	// Initialize models
	initModel("supplier", func() error {
		supplierFirebase := models.NewSupplierFirebase(firebaseService.GetFirestore())
		if supplierFirebase == nil {
			return fmt.Errorf("failed to create supplier model")
		}
		supplierHandler := handlers.NewSupplierHandler(supplierFirebase)
		suppliers := router.Group("/suppliers")
		{
			suppliers.GET("", supplierHandler.List)
			suppliers.GET("/:id", supplierHandler.Get)
			suppliers.POST("", supplierHandler.Create)
			suppliers.PUT("/:id", supplierHandler.Update)
			suppliers.DELETE("/:id", supplierHandler.Delete)
		}
		return nil
	})

	initModel("product", func() error {
		productFirebase, err := models.NewProductFirebase(firebaseService.GetFirestore())
		if err != nil {
			return fmt.Errorf("failed to create product model: %v", err)
		}
		if productFirebase == nil {
			return fmt.Errorf("failed to create product model")
		}
		productHandler := handlers.NewProductHandler(productFirebase, fileStorage, imageRecognition, imageTrainingService)
		products := router.Group("/products")
		{
			products.GET("", productHandler.List)
			products.GET("/:code", productHandler.Get)
			products.POST("", productHandler.Create)
			products.PUT("/:code", productHandler.Update)
			products.DELETE("/:code", productHandler.Delete)
			products.POST("/scan", productHandler.ScanProduct)
			products.GET("/barcode/:barcode", productHandler.FindByBarcode)
			products.GET("/company/:companyId", productHandler.FindByCompany)
			products.PUT("/:code/image", productHandler.UpdateImage)
			products.POST("/:code/image", productHandler.UploadImage)
			products.POST("/upload", productHandler.Upload)
		}
		return nil
	})

	initModel("brand", func() error {
		brandFirebase := models.NewBrandFirebase(firebaseService.GetFirestore())
		if brandFirebase == nil {
			return fmt.Errorf("failed to create brand model")
		}
		brandHandler := handlers.NewBrandHandler(brandFirebase)
		brands := router.Group("/brands")
		{
			brands.GET("", brandHandler.List)
			brands.GET("/:id", brandHandler.Get)
			brands.POST("", brandHandler.Create)
			brands.PUT("/:id", brandHandler.Update)
			brands.DELETE("/:id", brandHandler.Delete)
		}
		return nil
	})

	initModel("product category", func() error {
		productCategoryFirebase := models.NewProductCategoryFirebase(firebaseService.GetFirestore())
		if productCategoryFirebase == nil {
			return fmt.Errorf("failed to create product category model")
		}
		productCategoryHandler := handlers.NewProductCategoryHandler(productCategoryFirebase)
		categories := router.Group("/product-categories")
		{
			categories.GET("", productCategoryHandler.List)
			categories.GET("/:id", productCategoryHandler.Get)
			categories.POST("", productCategoryHandler.Create)
			categories.PUT("/:id", productCategoryHandler.Update)
			categories.DELETE("/:id", productCategoryHandler.Delete)
		}
		return nil
	})

	initModel("user", func() error {
		userFirebase := models.NewUserFirebase(firebaseService.GetFirestore(), firebaseService.GetAuthClient())
		if userFirebase == nil {
			return fmt.Errorf("failed to create user model")
		}
		userHandler := handlers.NewUserHandler(firebaseService, userFirebase)
		users := router.Group("/users")
		{
			users.GET("", userHandler.ListUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
			users.POST("/:id/join-company", userHandler.JoinCompany)
		}
		return nil
	})

	// Initialize remaining models
	initModel("company", func() error {
		companyFirebase := models.NewCompanyFirebase(firebaseService.GetFirestore())
		if companyFirebase == nil {
			return fmt.Errorf("failed to create company model")
		}
		companyHandler := handlers.NewCompanyHandler(companyFirebase)
		companies := router.Group("/companies")
		{
			companies.GET("/:id", companyHandler.Get)
			companies.POST("", companyHandler.Create)
			companies.PUT("/:id", companyHandler.Update)
			companies.DELETE("/:id", companyHandler.Delete)
		}
		return nil
	})

	initModel("customer", func() error {
		customerFirebase := models.NewCustomerFirebase(firebaseService.GetFirestore())
		if customerFirebase == nil {
			return fmt.Errorf("failed to create customer model")
		}
		customerHandler := handlers.NewCustomerHandler(customerFirebase)
		customers := router.Group("/customers")
		{
			customers.GET("", customerHandler.List)
			customers.GET("/:id", customerHandler.Get)
			customers.POST("", customerHandler.Create)
			customers.PUT("/:id", customerHandler.Update)
			customers.DELETE("/:id", customerHandler.Delete)
		}
		return nil
	})

	initModel("salesman", func() error {
		salesmanFirebase := models.NewSalesmanFirebase(firebaseService.GetFirestore())
		if salesmanFirebase == nil {
			return fmt.Errorf("failed to create salesman model")
		}
		salesmanHandler := handlers.NewSalesmanHandler(salesmanFirebase)
		salesmen := router.Group("/salesmen")
		{
			salesmen.GET("", salesmanHandler.List)
			salesmen.GET("/:id", salesmanHandler.Get)
			salesmen.POST("", salesmanHandler.Create)
			salesmen.PUT("/:id", salesmanHandler.Update)
			salesmen.DELETE("/:id", salesmanHandler.Delete)
		}
		return nil
	})

	initModel("shelve", func() error {
		shelveFirebase := models.NewShelveFirebase(firebaseService.GetFirestore())
		if shelveFirebase == nil {
			return fmt.Errorf("failed to create shelve model")
		}
		shelveHandler := handlers.NewShelveHandler(shelveFirebase)
		shelves := router.Group("/shelves")
		{
			shelves.GET("", shelveHandler.List)
			shelves.GET("/:id", shelveHandler.Get)
			shelves.POST("", shelveHandler.Create)
			shelves.PUT("/:id", shelveHandler.Update)
			shelves.DELETE("/:id", shelveHandler.Delete)
		}
		return nil
	})

	initModel("region", func() error {
		regionFirebase := models.NewRegionFirebase(firebaseService.GetFirestore())
		if regionFirebase == nil {
			return fmt.Errorf("failed to create region model")
		}
		regionHandler := handlers.NewRegionHandler(regionFirebase)
		regions := router.Group("/regions")
		{
			regions.GET("", regionHandler.List)
			regions.GET("/:id", regionHandler.Get)
			regions.POST("", regionHandler.Create)
			regions.PUT("/:id", regionHandler.Update)
			regions.DELETE("/:id", regionHandler.Delete)
		}
		return nil
	})

	initModel("role", func() error {
		roleFirebase := models.NewRoleFirebase(firebaseService.GetFirestore())
		if roleFirebase == nil {
			return fmt.Errorf("failed to create role model")
		}
		roleHandler := handlers.NewRoleHandler(roleFirebase)
		roles := router.Group("/roles")
		{
			roles.GET("", roleHandler.List)
			roles.GET("/:id", roleHandler.Get)
			roles.POST("", roleHandler.Create)
			roles.PUT("/:id", roleHandler.Update)
			roles.DELETE("/:id", roleHandler.Delete)
		}
		return nil
	})

	initModel("access", func() error {
		accessFirebase := models.NewAccessFirebase(firebaseService.GetFirestore())
		if accessFirebase == nil {
			return fmt.Errorf("failed to create access model")
		}
		accessHandler := handlers.NewAccessHandler(accessFirebase)
		access := router.Group("/access")
		{
			access.GET("", accessHandler.List)
			access.GET("/:id", accessHandler.Get)
			access.POST("", accessHandler.Create)
			access.PUT("/:id", accessHandler.Update)
			access.DELETE("/:id", accessHandler.Delete)
		}
		return nil
	})

	initModel("branch", func() error {
		branchFirebase := models.NewBranchFirebase(firebaseService.GetFirestore())
		if branchFirebase == nil {
			return fmt.Errorf("failed to create branch model")
		}
		branchHandler := handlers.NewBranchHandler(branchFirebase)
		branches := router.Group("/branches")
		{
			branches.GET("", branchHandler.List)
			branches.GET("/:id", branchHandler.Get)
			branches.POST("", branchHandler.Create)
			branches.PUT("/:id", branchHandler.Update)
			branches.DELETE("/:id", branchHandler.Delete)
		}
		return nil
	})

	initModel("purchase", func() error {
		purchaseFirebase := models.NewPurchaseFirebase(firebaseService.GetFirestore())
		if purchaseFirebase == nil {
			return fmt.Errorf("failed to create purchase model")
		}
		purchaseHandler := handlers.NewPurchaseHandler(purchaseFirebase)
		purchases := router.Group("/purchases")
		{
			purchases.GET("", purchaseHandler.List)
			purchases.GET("/:id", purchaseHandler.Get)
			purchases.POST("", purchaseHandler.Create)
			purchases.PUT("/:id", purchaseHandler.Update)
			purchases.DELETE("/:id", purchaseHandler.Delete)
		}
		return nil
	})

	initModel("sales order", func() error {
		salesOrderFirebase := models.NewSalesOrderFirebase(firebaseService.GetFirestore())
		if salesOrderFirebase == nil {
			return fmt.Errorf("failed to create sales order model")
		}
		productFirebase, err := models.NewProductFirebase(firebaseService.GetFirestore())
		if err != nil {
			return fmt.Errorf("failed to create product model: %v", err)
		}
		salesOrderHandler := handlers.NewSalesOrderHandler(salesOrderFirebase, productFirebase)
		salesOrders := router.Group("/sales-orders")
		{
			salesOrders.GET("", salesOrderHandler.List)
			salesOrders.GET("/:id", salesOrderHandler.Get)
			salesOrders.POST("", salesOrderHandler.Create)
			salesOrders.PUT("/:id", salesOrderHandler.Update)
			salesOrders.DELETE("/:id", salesOrderHandler.Delete)
			salesOrders.GET("/stats", salesOrderHandler.Stats)
			salesOrders.POST("/update-stock", salesOrderHandler.UpdateProductStock)
		}
		return nil
	})

	initModel("receive", func() error {
		receiveFirebase := models.NewReceiveFirebase(firebaseService.GetFirestore())
		if receiveFirebase == nil {
			return fmt.Errorf("failed to create receive model")
		}
		receiveHandler := handlers.NewReceiveHandler(receiveFirebase)
		receives := router.Group("/receives")
		{
			receives.GET("", receiveHandler.List)
			receives.GET("/:id", receiveHandler.Get)
			receives.POST("", receiveHandler.Create)
			receives.PUT("/:id", receiveHandler.Update)
			receives.DELETE("/:id", receiveHandler.Delete)
		}
		return nil
	})

	initModel("delivery", func() error {
		deliveryFirebase := models.NewDeliveryFirebase(firebaseService.GetFirestore())
		if deliveryFirebase == nil {
			return fmt.Errorf("failed to create delivery model")
		}
		deliveryHandler := handlers.NewDeliveryHandler(deliveryFirebase)
		deliveries := router.Group("/deliveries")
		{
			deliveries.GET("", deliveryHandler.List)
			deliveries.GET("/:id", deliveryHandler.Get)
			deliveries.POST("", deliveryHandler.Create)
			deliveries.PUT("/:id", deliveryHandler.Update)
			deliveries.DELETE("/:id", deliveryHandler.Delete)
		}
		return nil
	})

	initModel("purchase return", func() error {
		purchaseReturnFirebase := models.NewPurchaseReturnFirebase(firebaseService.GetFirestore())
		if purchaseReturnFirebase == nil {
			return fmt.Errorf("failed to create purchase return model")
		}
		purchaseReturnHandler := handlers.NewPurchaseReturnHandler(purchaseReturnFirebase)
		purchaseReturns := router.Group("/purchase-returns")
		{
			purchaseReturns.GET("", purchaseReturnHandler.List)
			purchaseReturns.GET("/:id", purchaseReturnHandler.Get)
			purchaseReturns.POST("", purchaseReturnHandler.Create)
			purchaseReturns.PUT("/:id", purchaseReturnHandler.Update)
			purchaseReturns.DELETE("/:id", purchaseReturnHandler.Delete)
		}
		return nil
	})

	initModel("sales order return", func() error {
		salesOrderReturnFirebase := models.NewSalesOrderReturnFirebase(firebaseService.GetFirestore())
		if salesOrderReturnFirebase == nil {
			return fmt.Errorf("failed to create sales order return model")
		}
		salesOrderReturnHandler := handlers.NewSalesOrderReturnHandler(salesOrderReturnFirebase)
		salesOrderReturns := router.Group("/sales-order-returns")
		{
			salesOrderReturns.GET("", salesOrderReturnHandler.List)
			salesOrderReturns.GET("/:id", salesOrderReturnHandler.Get)
			salesOrderReturns.POST("", salesOrderReturnHandler.Create)
			salesOrderReturns.PUT("/:id", salesOrderReturnHandler.Update)
			salesOrderReturns.DELETE("/:id", salesOrderReturnHandler.Delete)
		}
		return nil
	})

	initModel("receive return", func() error {
		receiveReturnFirebase := models.NewReceiveReturnFirebase(firebaseService.GetFirestore())
		if receiveReturnFirebase == nil {
			return fmt.Errorf("failed to create receive return model")
		}
		receiveReturnHandler := handlers.NewReceiveReturnHandler(receiveReturnFirebase)
		receiveReturns := router.Group("/receive-returns")
		{
			receiveReturns.GET("", receiveReturnHandler.List)
			receiveReturns.GET("/:id", receiveReturnHandler.Get)
			receiveReturns.POST("", receiveReturnHandler.Create)
			receiveReturns.PUT("/:id", receiveReturnHandler.Update)
			receiveReturns.DELETE("/:id", receiveReturnHandler.Delete)
		}
		return nil
	})

	initModel("delivery return", func() error {
		deliveryReturnFirebase := models.NewDeliveryReturnFirebase(firebaseService.GetFirestore())
		if deliveryReturnFirebase == nil {
			return fmt.Errorf("failed to create delivery return model")
		}
		deliveryReturnHandler := handlers.NewDeliveryReturnHandler(deliveryReturnFirebase)
		deliveryReturns := router.Group("/delivery-returns")
		{
			deliveryReturns.GET("", deliveryReturnHandler.List)
			deliveryReturns.GET("/:id", deliveryReturnHandler.Get)
			deliveryReturns.POST("", deliveryReturnHandler.Create)
			deliveryReturns.PUT("/:id", deliveryReturnHandler.Update)
			deliveryReturns.DELETE("/:id", deliveryReturnHandler.Delete)
		}
		return nil
	})

	initModel("predictive", func() error {
		productFirebase, err := models.NewProductFirebase(firebaseService.GetFirestore())
		if err != nil {
			return fmt.Errorf("failed to create product model: %v", err)
		}
		salesOrderFirebase := models.NewSalesOrderFirebase(firebaseService.GetFirestore())
		if salesOrderFirebase == nil {
			return fmt.Errorf("failed to create sales order model")
		}
		predictiveService := services.NewPredictiveService(productFirebase, salesOrderFirebase)
		if predictiveService == nil {
			return fmt.Errorf("failed to create predictive service")
		}
		predictiveHandler := handlers.NewPredictiveHandler(predictiveService)
		predictive := router.Group("/predictive")
		{
			predictive.GET("/stock-recommendations/:companyId", predictiveHandler.GetStockRecommendations)
			predictive.GET("/sales-predictions/:companyId", predictiveHandler.GetSalesPredictions)
			predictive.GET("/product-history/:productCode", predictiveHandler.GetProductHistory)
			predictive.GET("/sales-report/:companyId", predictiveHandler.GetSalesReport)
		}
		return nil
	})
}
