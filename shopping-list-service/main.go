package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	api "main/api/items"
)

func main() {
	app := fiber.New()

	items := app.Group("/items")

	items.Get("/", api.GetAllItemsHandler)
	items.Post("/", api.AddOneItemHandler)
	items.Delete("/:id<guid>", api.DeleteOneItemHandler)

	_ = app.Listen(":10001")
}
