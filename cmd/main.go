package main

import (
	"github.com/dowingows/quiz-go/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectRedis()

	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
