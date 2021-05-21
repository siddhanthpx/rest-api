package handlers

import (
	"rest-api/data"
	"rest-api/database"

	"github.com/gofiber/fiber/v2"
)

func DeleteGame(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	db := database.ConnectDB()
	if err := db.Delete(&data.Game{}, id).Error; err != nil {
		return err
	}

	return nil
}
