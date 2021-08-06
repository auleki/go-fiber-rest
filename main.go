package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App) {
	//route to get all todos
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Home endpoint",
		})
	})

	api := app.Group("/api")

	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "AT THE NEW ENDPOINT",
		})
	})
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	// setup Routes
	setupRoutes(app)

	// listen on port 8000 for any network requests
	err := app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}
