
package routes

import (
	"github.com/gofiber/fiber/v2"
	"brokerx/backend/internal/handlers"
	"brokerx/backend/internal/middlewares"
)

func Register(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/v1/auth/signup", handlers.Signup)
	api.Post("/v1/auth/login", handlers.Login)

	api.Get("/v1/accounts/:id", middlewares.RequireJWT(), handlers.GetAccount)

	api.Post("/v1/orders", middlewares.RequireJWT(), handlers.CreateOrder)
	api.Get("/v1/orders/:id", middlewares.RequireJWT(), handlers.GetOrder)

	api.Get("/v1/portfolios/:id", middlewares.RequireJWT(), handlers.GetPortfolio)

	api.Get("/v1/quotes", middlewares.RequireJWT(), handlers.GetQuote)
	api.Get("/v1/stream/quotes", middlewares.RequireJWT(), handlers.StreamQuotes)
}
