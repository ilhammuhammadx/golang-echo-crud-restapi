package main

import (
	"echo-rest-api-mysql/config"
	"echo-rest-api-mysql/domain/item/controllers"
	"echo-rest-api-mysql/domain/item/services"
	"echo-rest-api-mysql/routes"
	"echo-rest-api-mysql/validators"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load database configuration
	cfg := config.LoadDBConfig()

	// Initialize database connection
	db := config.InitDB(cfg)

	// Initialize repositories and services
	itemService := services.NewItemService(db)

	// Initialize controllers
	itemController := controllers.NewItemController(itemService, validators.CustomValidator{Validator: validator.New()})

	// Set up Echo server and routes
	e := echo.New()
	routes.RegisterItemRoutes(e, itemController)

	// Start the server
	e.Logger.Fatal(e.Start(":8081"))
}
