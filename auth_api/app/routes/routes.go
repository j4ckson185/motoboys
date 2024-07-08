package routes

import (
	"github.com/carloshomar/vercardapio/app/handlers"
	"github.com/carloshomar/vercardapio/app/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users/register", handlers.CreateUser)
	app.Post("/users/", handlers.)
	app.Get("/users/:id", ProtectedRoute, handlers.GetUser)

	app.Get("/establishments", handlers.ListEstablishments)
	app.Put("/establishments/status/handler/:id", handlers.HandlerEstablishmentStatus)
	app.Get("/establishments/:id", handlers.GetEstablishments)
	app.Put("/establishments/:id", handlers.UpdateEstablishment)
	app.Get("/establishments/:id/users", ProtectedRoute, handlers.GetUserByEstablishment)

	app.Post("/delivery-man/login", handlers.LoginDeliveryMan)
	app.Post("/delivery-man/register", handlers.CreateDeliveryMan)

}

func ProtectedRoute(c *fiber.Ctx) error {

	_, err := middlewares.ValidateJWT(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	return c.Next()
}
