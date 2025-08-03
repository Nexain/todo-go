package main

import (
	"log"
	"todo-go/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api.SetupRoutes(app)
	log.Println("🚀 Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
