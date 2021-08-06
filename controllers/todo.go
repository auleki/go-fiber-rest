package controllers

import "github.com/gofiber/fiber/v2"

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// static typed todos
var todos = []*Todo{
	{
		Id:        1,
		Title:     "Record a verse",
		Completed: false,
	},
	{
		Id:        3,
		Title:     "Fucking breath",
		Completed: true,
	},
	{
		Id:        2,
		Title:     "Complete data entry profile",
		Completed: false,
	},
}

// gets all todos
func GetTodos(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}
