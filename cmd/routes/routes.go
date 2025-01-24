package routes

import (
	"github.com/Mariano-JR/auth/internal/user"
	"github.com/Mariano-JR/auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static/")

	auth := app.Group("/auth")

	//GET

	//POST
	auth.Post("/login", user.LoginUser)
	auth.Post("/register", middlewares.ValidateMiddleware(user.User{}), user.CreateUser)
	//PUT

	//DELETE
	auth.Delete("/delete", user.DeleteUser)
}
