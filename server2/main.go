package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/aspenjames/library-app/server/middlewares"
	"github.com/aspenjames/library-app/server/routes"
	"github.com/gofiber/fiber/v2"
)

// Port on which to run the API server.
var port int

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
	middlewares.Setup(app)
	routes.Setup(app)
	addr := fmt.Sprintf(":%d", port)
	app.Listen(addr)
}
