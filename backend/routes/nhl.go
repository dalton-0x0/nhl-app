package routes

import (
	"github.com/dalton-0x0/nhl-app/backend/scraper"
	"github.com/gofiber/fiber/v2"
)

func GetLiveGames(c *fiber.Ctx) error {
	games, err := scraper.FetchLiveGames()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"games": games,
	})
}
