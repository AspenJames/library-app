package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

// Port on which to run the API server.
var port int

// Add middlewares:
//   - Compress:  compress responses
//   - ETag:      caching header
//   - Limiter:   rate limiter
//   - Logger:    logging middleware
//   - Monitor:   expose common metrics
//   - Recover:   handle panics with app.ErrorHandler
//   - RequestID: add unique X-Request-ID header
func setupMiddlewares(app *fiber.App) *fiber.App {
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 10 * time.Second,
	}))
	app.Use(logger.New(logger.Config{
		TimeZone: "UTC",
	}))
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(requestid.New())

	// Setup metrics endpoint
	app.Get("/metrics", monitor.New(monitor.Config{
		Title: "Library-app server metrics",
	}))
	return app
}

// Wrap route handlers in timeout
func timeoutWrapper(h fiber.Handler) fiber.Handler {
	return timeout.New(h, 5*time.Second)
}

func main() {
	flag.IntVar(&port, "p", 8080, "Port on which to listen")
	flag.Parse()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
				"code":  code,
			})
		},
	})
	setupMiddlewares(app)
	app.Get("/", timeoutWrapper(func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}))

	addr := fmt.Sprintf(":%d", port)
	app.Listen(addr)
}
