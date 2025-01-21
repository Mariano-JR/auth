package routes

import (
	"github.com/Mariano-JR/auth/internal/user"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static")

	//GET
	app.Get("auth/users", user.GetUsers)

	//POST
	app.Post("/auth/create", user.CreateUser)

	//PUT

	//DELETE
}
