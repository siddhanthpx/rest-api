package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

// AddGame adds a new game to the database
// by parsing JSON from the request body
func AddGame(ctx *fiber.Ctx) error {

	// Parse the JSON object into empty game struct
	var game data.Game
	if err := ctx.BodyParser(&game); err != nil {
		return err
	}

	// Validate the posted JSON
	if err := game.Validate(); err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	// Connect to the database
	db := database.ConnectDB()

	// Insert the object into the database
	if err := db.Create(&game).Error; err != nil {
		return err
	}

	ctx.JSON(game)

	return nil
}
