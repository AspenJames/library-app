package routes

import (
	"github.com/aspenjames/library-app/server/models"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRouter(router fiber.Router) {
	// Get all books
	router.Get("/", timeoutWrapper(func(c *fiber.Ctx) error {
		books, err := models.AllBooks()
		if err != nil {
			return err
		}
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
		if err := c.BodyParser(&book); err != nil {
			return err
		}
		if err := models.CreateBook(&book); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"book": book,
		})
	}))

	// Update a book
	router.Post("/:id", timeoutWrapper(func(c *fiber.Ctx) error {
		book := models.Book{}
		if err := c.BodyParser(&book); err != nil {
			return err
		}
		updated, err := models.UpdateBook(c.Params("id"), &book)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"book": updated,
		})
	}))

	// Delete a book
	router.Delete("/:id", timeoutWrapper(func(c *fiber.Ctx) error {
		if err := models.DeleteBook(c.Params("id")); err != nil {
			return err
		}
		return c.JSON(fiber.Map{})
	}))
}
