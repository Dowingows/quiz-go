package main

import (
	"github.com/dowingows/quiz-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.Create)
	app.Get("/fact/3", handlers.GetFact)
}
