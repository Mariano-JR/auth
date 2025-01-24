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

	u, _ := GetUser(user.Email)

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

	return c.Redirect("/home.html", fiber.StatusFound)
}

func CreateUser(c *fiber.Ctx) error {
	user := c.Locals("validatedModel").(User)

	if _, err := Save(user.Email, user.Name, user.Password); err != nil {
		return c.Redirect("/signup.html", fiber.StatusFound)
	}

	return c.Redirect("/", fiber.StatusFound)
}

func DeleteUser(c *fiber.Ctx) error {
	if _, err := Delete(c.Cookies("user_email")); err != nil {
		return err
	}

	c.ClearCookie("user_id", "user_email", "user_name")

	return nil
}
