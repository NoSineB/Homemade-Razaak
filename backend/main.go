package main

import (
	"github.com/gofiber/fiber/v2"
	"homemade.com/routes"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Listen(":3000")
}