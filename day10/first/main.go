package main

import "github.com/gofiber/fiber/v2"

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Text: "Learging golang", Done: false},
	{ID: 2, Text: "Connect to fiber", Done: false},
}

func main() {
	app := fiber.New()

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		var t Todo

		if err := c.BodyParser(&t); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		t.ID = len(todos) + 1
		todos = append(todos, t)
		return c.Status(fiber.StatusCreated).JSON(t)
	})

	app.Listen(":3000")
}
