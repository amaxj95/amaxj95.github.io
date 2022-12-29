package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/amaxj95/pgdoc_connect/pgdoc_connect/database"
    "github.com/amaxj95/pgdoc_connect/pgdoc_connect/models"
)

func Home(c *fiber.Ctx) error {
    return c.SendString("Docker Alpine & PostgresQL for Crunchy DB MGMT ")
}

func CreateFact(c *fiber.Ctx) error {
    prereq := new(models.Crunchy_prereq)
    if err := c.BodyParser(prereq); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": err.Error(),
        })
    }

    database.DB.Db.Create(&prereq)

    return c.Status(200).JSON(prereq)
}
