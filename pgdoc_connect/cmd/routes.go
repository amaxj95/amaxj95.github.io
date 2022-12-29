package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/amaxj95/pgdoc_connect/pgdoc_connect/handlers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/", handlers.Home)

    app.Post("/prereq", handlers.CreateFact)
}
