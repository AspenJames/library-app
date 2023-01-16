package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// Add middlewares:
//   - Compress:  compress responses
//   - ETag:      caching header
//   - Limiter:   rate limiter
//   - Logger:    logging middleware
//   - Monitor:   expose common metrics
//   - Recover:   handle panics with app.ErrorHandler
//   - RequestID: add unique X-Request-ID header
func Setup(app *fiber.App) *fiber.App {
	app.Use(recover.New())
	app.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 10 * time.Second,
	}))
	app.Use(logger.New(logger.Config{
		TimeZone: "UTC",
		// Don't log healthcheck requests.
		Next: func(c *fiber.Ctx) bool {
			return c.OriginalURL() == "/healthcheck"
		},
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
