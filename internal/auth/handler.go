package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Provider(c *fiber.Ctx) error {
	provider := c.Params("provider")
	return c.SendString("Redirecting to OAuth provider: " + provider)
}

func Callback(c *fiber.Ctx) error {
	return c.SendString("Callback received from OAuth provider")
}
