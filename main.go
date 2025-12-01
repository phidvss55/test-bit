package main

import (
	"go-circleci/api"
	"go-circleci/logger"
	"go-circleci/repository"
	"go-circleci/services"
	"log"
)

func main() {
	// Initialize SQLite database connection
	db, err := services.InitDatabase("file:./app.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Create products table on startup
	// if err := services.CreateProductsTable(db); err != nil {
	// 	log.Fatalf("Failed to create products table: %v", err)
	// }

	// Create product repository instance
	productRepo := repository.NewSQLiteProductRepository(db)

	// Create product service instance
	productService := services.NewProductService(productRepo)

	// Create cat fact service instance
	catFactService := services.NewCatFactService("https://catfact.ninja/fact")

	// Create composite service that supports both CatFact and Product operations
	compositeService := services.NewCompositeService(catFactService.(*services.CatFactService), productService)

	// Wrap with logging
	service := logger.NewLoggingService(compositeService)

	// Pass composite service to API server
	apiServer := api.NewApiServer(service)

	log.Fatal(apiServer.Start(":5000"))
}
