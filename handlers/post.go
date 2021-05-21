package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

func AddGame(ctx *fiber.Ctx) error {

	var game data.Game
	if err := ctx.BodyParser(&game); err != nil {
		return err
	}

	db := database.ConnectDB()
	if err := db.Create(&game).Error; err != nil {
		return err
	}

	ctx.JSON(game)

	return nil
}
