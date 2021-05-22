package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

// swagger:route PUT /games/:id games updateGame
// Updates fields in the game
// responses:
//	200: gamesResponse

// UpdateGame updates fields int the game with the given id
// and saves to database
func UpdateGame(ctx *fiber.Ctx) error {

	// Parse id from query
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	// Connect to the database
	db := database.ConnectDB()

	// Find the game with the given ID and encode into an empty game instance
	var fetchedGame data.Game
	if err := db.Where(&data.Game{ID: id}).First(&fetchedGame).Error; err != nil {
		return err
	}

	// Parse fields from post body and encode to fetched game instance
	if err := ctx.BodyParser(&fetchedGame); err != nil {
		return err
	}

	// Validate post body
	if err := fetchedGame.Validate(); err != nil {
		ctx.JSON(fiber.Map{
			"message": err,
		})
		return err
	}

	// Send back updated game instance to the database
	if err := db.Save(&fetchedGame).Error; err != nil {
		return err
	}

	return nil
}
