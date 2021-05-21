package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

func UpdateGame(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	db := database.ConnectDB()
	var fetchedGame data.Game
	if err := db.Where(&data.Game{ID: id}).First(&fetchedGame).Error; err != nil {
		return err
	}

	if err := ctx.BodyParser(&fetchedGame); err != nil {
		return err
	}
	db.Save(&fetchedGame)
	return nil
}
