package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

// Wrap route handlers in timeout
func timeoutWrapper(h fiber.Handler) fiber.Handler {
	return timeout.New(h, 5*time.Second)
}

func Setup(app *fiber.App) {
	app.Get("/", timeoutWrapper(func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}))

	app.Get("/healthcheck", timeoutWrapper(func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	}))
}
