package api

import (
	"strconv"
	"todo-go/core"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/todos", func(c *fiber.Ctx) error {
		todos, err := core.ListTask()
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(todos)
	})

	app.Post("/todos", func(c *fiber.Ctx) error {
		type Request struct {
			Task string `json:"task"`
		}
		var body Request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}
		if body.Task == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Task tidak boleh kosong"})
		}
		if len(body.Task) > 255 {
			return c.Status(400).JSON(fiber.Map{"error": "Task terlalu panjang, maksimal 255 karakter"})
		}
		if err := core.AddTask(body.Task); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(201).JSON(fiber.Map{"message": "Todo added"})
	})

	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		type Request struct {
			Task string `json:"task"`
		}
		var body Request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}
		updated, err := core.UpdateTask(id, body.Task)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if !updated {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}
		return c.JSON(fiber.Map{"message": "Todo updated"})
	})

	app.Patch("/todos/:id/complete", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		status, err := core.CompleteTask(id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if status == "not_found" {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}
		if status == "already_done" {
			return c.JSON(fiber.Map{"message": "Todo already completed"})
		}
		return c.JSON(fiber.Map{"message": "Todo completed"})
	})

	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}
		deleted, err := core.DeleteTask(id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if !deleted {
			return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
		}
		return c.JSON(fiber.Map{"message": "Todo deleted"})
	})
}
