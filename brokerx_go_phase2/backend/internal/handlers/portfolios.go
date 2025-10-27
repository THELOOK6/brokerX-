
package handlers

import "github.com/gofiber/fiber/v2"

func GetPortfolio(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"cash": 100000.0, "positions": map[string]int{"AAPL": 10}})
}
