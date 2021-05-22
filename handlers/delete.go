package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

// DeleteGame deletes a game object from the database
// with the matching id
func DeleteGame(ctx *fiber.Ctx) error {

	// Parse id from query
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	// Connect to the database
	db := database.ConnectDB()

	// Delete the game object found in the table using the id
	if err := db.Delete(&data.Game{}, id).Error; err != nil {
		return err
	}

	return nil
}
