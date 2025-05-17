package main

import (
	"log"

	"github.com/dalton-0x0/nhl-app/backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/nhl/games/live", routes.GetLiveGames)

	log.Println("🚀🚀🚀 Server running on http://localhost:3000 ")
	log.Fatal(app.Listen(":3000"))
}
