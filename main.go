package main

import (
	"log"

	"github.com/auleki/go-fiber-todo/config"
	"github.com/auleki/go-fiber-todo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
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

	routes.TodoRoute(api.Group("/todos"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// fetch environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env files")
	}

	// connect to Database
	config.ConnectDB()

	// setup routes
	setupRoutes(app)

	err = app.Listen(":8000")

	if err != nil {
		panic(err)
	}

}
