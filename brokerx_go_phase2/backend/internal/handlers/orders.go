
package handlers

import (
	"encoding/json"
	"time"
	"github.com/gofiber/fiber/v2"
	"brokerx/backend/internal/models"
	"brokerx/backend/internal/store"
	"brokerx/backend/internal/services"
)

type OrderCreate struct { AccountID string `json:"account_id"`; Symbol string `json:"symbol"`; Side string `json:"side"`; Qty int `json:"qty"`; Price float64 `json:"price"` }

func CreateOrder(c *fiber.Ctx) error {
	idem := c.Get("Idempotency-Key")
	if idem == "" { return c.Status(400).JSON(fiber.Map{"error":"Missing Idempotency-Key"}) }
	if val, err := store.RDB.Get(store.Ctx, "idem:"+idem).Result(); err == nil {
		c.Response().Header.SetContentType("application/json")
		return c.Status(201).SendString(val)
	}
	var req OrderCreate
	if err := c.BodyParser(&req); err != nil { return c.Status(400).JSON(fiber.Map{"error":"bad payload"}) }
	order := models.Order{ AccountID: req.AccountID, Symbol: req.Symbol, Side: req.Side, Qty: req.Qty, Price: req.Price, Status: "ACCEPTED" }
	if err := store.DB.Create(&order).Error; err != nil { return c.Status(400).JSON(fiber.Map{"error":"create failed"}) }
	b, _ := json.Marshal(order)
	_ = store.RDB.Set(store.Ctx, "idem:"+idem, string(b), time.Hour).Err()
	return c.Status(201).JSON(order)
}

func GetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var o models.Order
	if err := store.DB.First(&o, "id = ?", id).Error; err != nil { return c.Status(404).JSON(fiber.Map{"error":"not found"}) }
	return c.JSON(o)
}

func PlaceOrder(c *fiber.Ctx) error {
	accountID := c.Locals("account_id").(string)

	var req struct {
		Symbol string  `json:"symbol"`
		Side   string  `json:"side"`
		Qty    int     `json:"qty"`
		Price  float64 `json:"price"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	order, err := services.PlaceOrder(accountID, req.Symbol, req.Side, req.Qty, req.Price)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(order)
}

func GetOrders(c *fiber.Ctx) error {
	accountID := c.Locals("account_id").(string)

	orders, err := services.GetOrders(accountID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch orders"})
	}

	return c.JSON(orders)
}
