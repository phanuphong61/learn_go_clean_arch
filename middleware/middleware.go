package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func CorsMiddleware(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Credentials", "true")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	if c.Method() == "OPTIONS" {
		c.Status(fiber.StatusOK)
		return nil
	}
	return c.Next()
}
