package middlewares

import (
	"github.com/Mariano-JR/auth/internal/user"

	"github.com/gofiber/fiber/v2"
)

func CookiesMiddleware(c *fiber.Ctx) error {
	var data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if u, err := user.GetUser(data.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not founded",
		})
	} else if u.Password == data.Password {
		c.Cookie(&fiber.Cookie{
			Name:  "user_id",
			Value: u.ID,
			Path:  "/",
		})
		c.Cookie(&fiber.Cookie{
			Name:  "user_email",
			Value: u.Email,
			Path:  "/",
		})
		c.Cookie(&fiber.Cookie{
			Name:  "user_name",
			Value: u.Name,
			Path:  "/",
		})
	}

	return c.Next()
}
