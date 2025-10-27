
package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RequireJWT() fiber.Handler {
	return func(c *fiber.Ctx) error {
		h := c.Get("Authorization")
		if len(h) < 8 || h[:7] != "Bearer " {
			return c.Status(401).JSON(fiber.Map{"error":"missing bearer token"})
		}
		secret := os.Getenv("JWT_SECRET")
		parsed, err := jwt.Parse(h[7:], func(t *jwt.Token) (interface{}, error) { return []byte(secret), nil })
		if err != nil || !parsed.Valid {
			return c.Status(401).JSON(fiber.Map{"error":"invalid token"})
		}
		c.Locals("claims", parsed.Claims)
		return c.Next()
	}
}
