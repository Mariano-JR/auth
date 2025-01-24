package middlewares

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateMiddleware(model interface{}) fiber.Handler {

	return func(c *fiber.Ctx) error {
		instance := reflect.New(reflect.TypeOf(model)).Interface()

		if err := c.BodyParser(instance); err != nil {
			return c.RedirectBack("/", fiber.StatusFound)
		}

		if err := validate.Struct(instance); err != nil {
			return c.RedirectBack("/", fiber.StatusFound)
		}

		c.Locals("validatedModel", instance)

		return c.Next()
	}
}
