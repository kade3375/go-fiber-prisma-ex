package api

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func New(app *fiber.App) {
	// GetAllBooks is a function to get all books data from database
	// @Summary Get all books
	// @Description Get all books
	// @Tags books
	// @Accept json
	// @Produce json
	// @Success 200 {object} ResponseHTTP{}
	// @Failure 503 {object} ResponseHTTP{}
	// @Router / [get]
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/docs/*", swagger.Handler) // swagger default
}