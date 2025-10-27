
package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/ansrivas/fiberprometheus/v2"

	"brokerx/backend/internal/routes"
	"brokerx/backend/internal/store"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" { port = "8000" }

	if err := store.InitStore(); err != nil { log.Fatalf("store init: %v", err) }

	app := fiber.New(fiber.Config{ ServerHeader: "BrokerX+ Go" })
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	prom := fiberprometheus.New("brokerx_api")
	prom.RegisterAt(app, "/metrics")
	app.Use(prom.Middleware)

	app.Get("/health", func(c *fiber.Ctx) error { return c.JSON(fiber.Map{"status":"ok"}) })

	routes.Register(app)

	log.Printf("listening on :%s", port)
	if err := app.Listen("0.0.0.0:" + port); err != nil { log.Fatal(err) }
}
