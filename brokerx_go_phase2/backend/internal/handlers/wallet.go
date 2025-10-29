package handlers

import (
	"brokerx/backend/internal/services"
	"github.com/gofiber/fiber/v2"
)

func Deposit(c *fiber.Ctx) error {
	accountID := c.Locals("account_id").(string)

	var req struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := services.Deposit(accountID, req.Amount); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "deposit successful"})
}

func GetBalance(c *fiber.Ctx) error {
	accountID := c.Locals("account_id").(string)

	balance, err := services.GetBalance(accountID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch balance"})
	}

	return c.JSON(fiber.Map{"balance": balance})
}
