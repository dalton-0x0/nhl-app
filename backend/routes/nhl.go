package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetLiveGames(c *fiber.Ctx) error {
	// Placeholder response
	data := fiber.Map{
		"games": []fiber.Map{
			{
				"homeTeam":  "Maple Leafs",
				"awayTeam":  "Bruins",
				"score":     "2-3",
				"period":    "3rd",
				"remaining": "05:34",
				"status":    "LIVE",
			},
		},
	}
	return c.JSON(data)
}
