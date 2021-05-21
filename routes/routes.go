package routes

import (
	"rest-api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/games/", handlers.GetAllGames)
	app.Get("/games/:id", handlers.GetGame)

}
