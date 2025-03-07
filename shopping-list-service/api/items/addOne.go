package api

import (
	"github.com/gofiber/fiber/v2"
	"main/db"
	"main/json"
)

type AddOneItemBody struct {
	Name string `json:"name"`
}

type AddOneItemResponse struct {
	NewItem json.ItemJson `json:"newItem"`
}

func AddOneItemHandler(c *fiber.Ctx) error {
	queries := db.GetQueries()

	body := new(AddOneItemBody)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if body.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is required",
		})
	}

	item, err := queries.AddItem(c.Context(), body.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add item",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(AddOneItemResponse{
		NewItem: json.GetItem(item),
	})
}
