package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"main/db"
)

func DeleteOneItemHandler(c *fiber.Ctx) error {
	queries := db.GetQueries()

	id := c.Params("id")
	itemId := uuid.Must(uuid.Parse(id))

	err := queries.DeleteItem(c.Context(), itemId)
	if err != nil {
		log.Error(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete item",
		})
	}

	return c.SendString(id)
}
