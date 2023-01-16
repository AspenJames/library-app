package routes

import (
	"github.com/aspenjames/library-app/server/models"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRouter(router fiber.Router) {
	// Get all books
	router.Get("/", timeoutWrapper(func(c *fiber.Ctx) error {
		books := models.AllBooks()
		return c.JSON(fiber.Map{
			"books": books,
		})
	}))

	// Get a book
	router.Get("/:id", timeoutWrapper(func(c *fiber.Ctx) error {
		book, err := models.FindBook(c.Params("id"))
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"book": book,
		})
	}))

	// Create a book
	router.Post("/", timeoutWrapper(func(c *fiber.Ctx) error {
		book := models.Book{}
		err := c.BodyParser(&book)
		if err != nil {
			return err
		}
		err = models.CreateBook(&book)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"book": book,
		})
	}))
}
