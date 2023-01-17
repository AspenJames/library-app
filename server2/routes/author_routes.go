package routes

import (
	"github.com/aspenjames/library-app/server/models"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthorRouter(router fiber.Router) {
	// Get all authors
	router.Get("/", timeoutWrapper(func(c *fiber.Ctx) error {
		authors, err := models.AllAuthors()
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"authors": authors,
		})
	}))

	// Get an author
	router.Get("/:id", timeoutWrapper(func(c *fiber.Ctx) error {
		author, err := models.FindAuthor(c.Params("id"))
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"author": author,
		})
	}))
}
