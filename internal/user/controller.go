package user

import (
	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if _, err := Login(user.Email, user.Password); err != nil {
		return c.Redirect("/", fiber.StatusFound)

	}
	return c.Redirect("/home.html", fiber.StatusFound)
}

func CreateUser(c *fiber.Ctx) error {
	var user struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if _, err := Save(user.Email, user.Name, user.Password); err != nil {
		return c.Redirect("/signup.html", fiber.StatusFound)
	}

	return c.Redirect("/", fiber.StatusFound)
}
