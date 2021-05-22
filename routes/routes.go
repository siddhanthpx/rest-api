package routes

import (
	"rest-api/handlers"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes to our API
func SetupRoutes(app *fiber.App) {

	app.Get("/games/", handlers.GetAllGames)      // Get all games in the database
	app.Get("/games/:id", handlers.GetGame)       // Get a single game with given ID
	app.Post("/games/new", handlers.AddGame)      // Add new game
	app.Delete("/games/:id", handlers.DeleteGame) // Delete game with given ID
	app.Put("/games/:id", handlers.UpdateGame)    // Update game with given ID

	// Serving API documentation on `localhost:8080/docs`
	app.Static("/swagger.yaml", "./swagger.yaml")
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	app.Get("/docs", adaptor.HTTPHandler(sh))

}
