package controllers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/auleki/go-fiber-todo/config"
	"github.com/auleki/go-fiber-todo/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// gets all todos
func GetTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	// query to filter
	query := bson.D{{}}

	cursor, err := todoCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went horribly wrong, toss your pc if not mac",
			"error":   err.Error(),
		})
	}

	var todos []models.Todo = make([]models.Todo, 0)

	// iterate the cursor and decode each item into a Todo
	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

func CreateTodo(c *fiber.Ctx) error {

	err := c.BodyParser(&body)

	// handle Error
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cant parse JSON",
		})
	}

	todo := &Todo{
		Id:        len(todos) + 1,
		Title:     body.Title,
		Completed: false,
	}

	todos = append(todos, todo)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func GetTodo(c *fiber.Ctx) error {
	// get parameter value
	paramId := c.Params("id")

	// convert params value from strings to int
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Could not parse Id",
		})
	}
	// find and return todo based on param
	for _, todo := range todos {
		if todo.Id == id {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"success": true,
				"data": fiber.Map{
					"todo": todo,
				},
			})
		}
	}
	// if no todo of passed ID
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	type Request struct {
		Title     *string `json:"title"`
		Completed *bool   `json:"completed"`
	}
	// extract parameter id
	paramId := c.Params("id")
	// convert id from string to int
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID",
		})
	}

	var body Request

	err = c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	var todo *Todo

	// loop over todos to find right one to update
	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Not Found",
		})
	}

	if body.Title != nil {
		todo.Title = *body.Title
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	// return a map with a success status and updated todo
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	paramId := c.Params("id")

	id, err := strconv.Atoi(paramId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
		})
	}

	// loop over todos to find a match and delete
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"status":  true,
				"message": "Deleted Successfully",
			})
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "Todo not found",
	})
}
