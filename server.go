package main

import (
	"log"
	"rest-api/data"
	"rest-api/database"
	"rest-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db := database.ConnectDB()

	if err := db.AutoMigrate(data.Game{}); err != nil {
		log.Fatal("Could not migrate:", err)
	}

	// Initialize router instance
	app := fiber.New()

	// Setup routes (check ./routes for setup function)

	routes.SetupRoutes(app)

	// Serving on localhost:8080/

	log.Fatal(app.Listen(":8080"))
}
