package user

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	var user struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if user.Email == "" || user.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email and Name are required",
		})
	}

	if _, err := Save(user.Email, user.Name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func GetUsers(c *fiber.Ctx) error {
	users := Users()

	return c.JSON(users)
}
