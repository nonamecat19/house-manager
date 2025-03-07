package api

import (
	"github.com/gofiber/fiber/v2"
	"main/db"
	"main/json"
)

type GetAllItemsResponse struct {
	Items []json.ItemJson `json:"items"`
}

func GetAllItemsHandler(c *fiber.Ctx) error {
	queries := db.GetQueries()

	items, err := queries.ListItems(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch items",
		})
	}

	return c.JSON(GetAllItemsResponse{
		Items: json.GetItemList(items),
	})
}
