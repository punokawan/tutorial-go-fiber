package main

import (
	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo

func main() {
	app := fiber.New()

	// Create a new Todo
	app.Post("/todos", func(c *fiber.Ctx) error {
		var todo Todo
		if err := c.BodyParser(&todo); err != nil {
			return err
		}
		todo.ID = len(todos) + 1
		todos = append(todos, todo)
		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	// Get all Todos
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// Get a Todo by ID
	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		paramID := c.Params("id")
		// Convert paramID to integer
		// Handle error if the conversion fails
		// Retrieve the todo with the corresponding ID
		// Return the todo as JSON
		// Handle error if the todo is not found
		return c.JSON(paramID)
	})

	// Update a Todo by ID
	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		// Similar steps as the Get handler to find the todo by ID
		// Update the todo with the provided data
		// Return the updated todo as JSON
		// Handle error if the todo is not found
		return nil
	})

	// Delete a Todo by ID
	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		// Similar steps as the Get handler to find the todo by ID
		// Delete the todo from the todos slice
		// Return a success message
		// Handle error if the todo is not found
		return nil
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
