
package handlers

import (
	"net"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"brokerx/backend/internal/models"
	"brokerx/backend/internal/services"
	"brokerx/backend/internal/store"
)

type SignupReq struct { Email string `json:"email"`; Password string `json:"password"`; Name string `json:"name"` }
func Signup(c *fiber.Ctx) error {
	var req SignupReq
	if err := c.BodyParser(&req); err != nil { return c.Status(400).JSON(fiber.Map{"error":"bad payload"}) }
	if req.Email == "" || req.Password == "" { return c.Status(400).JSON(fiber.Map{"error":"email/password required"}) }
	hash, _ := services.HashPassword(req.Password)
	acc := models.Account{Email: strings.ToLower(req.Email), Password: hash, Name: req.Name, Status: "active"}
	if err := store.DB.Create(&acc).Error; err != nil { return c.Status(400).JSON(fiber.Map{"error":"email exists?"}) }
	store.DB.Create(&models.AuditLog{ActorID: &acc.ID, Action:"signup", IP: clientIP(c), Success:true})
	return c.Status(201).JSON(acc)
}

type LoginReq struct { Email string `json:"email"`; Password string `json:"password"` }
func Login(c *fiber.Ctx) error {
	var req LoginReq
	if err := c.BodyParser(&req); err != nil { return c.Status(400).JSON(fiber.Map{"error":"bad payload"}) }
	var acc models.Account
	if err := store.DB.Where("email = ?", strings.ToLower(req.Email)).First(&acc).Error; err != nil {
		if err == gorm.ErrRecordNotFound { return c.Status(401).JSON(fiber.Map{"error":"invalid creds"}) }
		return c.Status(500).JSON(fiber.Map{"error":"db error"})
	}
	if !services.CheckPassword(acc.Password, req.Password) { return c.Status(401).JSON(fiber.Map{"error":"invalid creds"}) }
	tok, err := services.IssueJWT(acc.ID, "client")
	if err != nil { return c.Status(500).JSON(fiber.Map{"error":"jwt issue failed"}) }
	store.DB.Create(&models.AuditLog{ActorID: &acc.ID, Action:"login", IP: clientIP(c), Success:true})
	return c.JSON(fiber.Map{"token": tok, "account": fiber.Map{"id": acc.ID, "email": acc.Email, "name": acc.Name}})
}

func clientIP(c *fiber.Ctx) string {
	ip := c.IP()
	if ip == "" { ip, _, _ = net.SplitHostPort(c.Context().RemoteAddr().String()) }
	return ip
}
