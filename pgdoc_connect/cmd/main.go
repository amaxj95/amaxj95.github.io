package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/amaxj95/pgdoc_connect/pgdoc_connect/database"
)

func main() {
    database.ConnectDb()
    app := fiber.New()

    setupRoutes(app)

    app.Listen(":3000")
}
