package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

var (
	Game  data.Game
	Games []data.Game
)

// swagger:route GET /games games listGames
// Returns a list of games
// responses:
//	200: gamesResponse

// GetAllGames returns all the games in the database and
// encodes them into JSON before serving
func GetAllGames(ctx *fiber.Ctx) error {

	// Connect to the database
	db := database.ConnectDB()

	// Fetch the whole table from the database and apply to an array of games
	db.Table("games").Scan(&Games)

	// Send the JSON to the handler
	if err := ctx.JSON(Games); err != nil {
		return err
	}

	return nil
}

// GetGame returns a single game with the
// matching id as a JSON object
func GetGame(ctx *fiber.Ctx) error {

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

	// Send the JSON to the handler
	if err = ctx.JSON(fetchedGame); err != nil {
		return err
	}

	return nil
}
