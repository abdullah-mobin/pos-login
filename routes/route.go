package routes

import (
	"pos-login/handlers"
	"pos-login/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/auth")
	auth.Post("/login", handlers.Login).Name("Login-route")
	auth.Get("/refresh", middleware.JWTMiddleware, handlers.Refresh).Name("Refresh-route")
}
