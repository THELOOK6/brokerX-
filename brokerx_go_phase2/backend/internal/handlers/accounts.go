
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"brokerx/backend/internal/models"
	"brokerx/backend/internal/store"
)

type AccountCreate struct { Email string `json:"email"`; Password string `json:"password"`; Name string `json:"name"` }
func CreateAccount(c *fiber.Ctx) error {
	var req AccountCreate
	if err := c.BodyParser(&req); err != nil { return c.Status(400).JSON(fiber.Map{"error":"bad payload"}) }
	acc := models.Account{Email: req.Email, Name: req.Name, Status: "active"}
	if err := store.DB.Create(&acc).Error; err != nil { return c.Status(400).JSON(fiber.Map{"error":"create failed"}) }
	return c.Status(201).JSON(acc)
}
func GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")
	var acc models.Account
	if err := store.DB.First(&acc, "id = ?", id).Error; err != nil { return c.Status(404).JSON(fiber.Map{"error":"not found"}) }
	return c.JSON(acc)
}
