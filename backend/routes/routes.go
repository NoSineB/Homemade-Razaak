package routes

import (
	"github.com/gofiber/fiber/v2"
	"homemade.com/handlers"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)

	app.Post("/cart", handlers.AddToCart)
	app.Post("/cart/total", handlers.GetCartTotal)
	app.Post("/cart/checkout", handlers.Checkout)
	app.Post("/cart/len", handlers.GetCartLength)
}
