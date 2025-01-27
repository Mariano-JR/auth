package auth

import (
	"github.com/gofiber/fiber/v2"
)

func Google(c *fiber.Ctx) error {
	c.Redirect(GoogleOAuthConfig.AuthCodeURL(OAuthStateString))
	return nil
}

func Callback(c *fiber.Ctx) error {
	if c.Query("state") != OAuthStateString {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid OAuth state")
	}

	code := c.Query("code")
	_, err := GoogleOAuthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to exchange token: " + err.Error())
	}

	c.Redirect("/home.html")
	return nil
}
