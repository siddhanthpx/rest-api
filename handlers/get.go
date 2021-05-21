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

func GetAllGames(ctx *fiber.Ctx) error {

	db := database.ConnectDB()
	db.Table("games").Scan(&Games)

	if err := ctx.JSON(Games); err != nil {
		return err
	}

	return nil
}

func GetGame(ctx *fiber.Ctx) error {

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}

	db := database.ConnectDB()

	var fetchedGame data.Game
	if err := db.Where(&data.Game{ID: id}).First(&fetchedGame).Error; err != nil {
		return err
	}

	if err = ctx.JSON(fetchedGame); err != nil {
		return err
	}

	return nil
}
