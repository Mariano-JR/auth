package routes

import (
	"github.com/Mariano-JR/auth/internal/auth"
	"github.com/Mariano-JR/auth/internal/user"
	"github.com/Mariano-JR/auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Static("/", "./static/")

	authGroup := app.Group("/auth")

	//GET
	authGroup.Get("/google", auth.Google)
	authGroup.Get("/google/callback", auth.Callback)

	//POST
	authGroup.Post("/login", user.LoginUser)
	authGroup.Post("/register", middlewares.ValidateMiddleware(user.User{}), user.CreateUser)
	//PUT

	//DELETE
	authGroup.Delete("/delete", user.DeleteUser)
}
