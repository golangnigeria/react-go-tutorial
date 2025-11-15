package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var db *sql.DB

func main() {
	fmt.Println("Starting server...")

	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	POSTGRES_URI := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("pgx", POSTGRES_URI)
	if err != nil {
		log.Fatal("Error connecting to Neon DB:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot ping database:", err)
	}

	fmt.Println("Connected to Neon DB")

	createTable()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173/",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", toggleTodo)  // toggle completed
	app.Patch("/api/todos", bulkUpdateTodos) // bulk update
	app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		body TEXT NOT NULL,
		completed BOOLEAN DEFAULT FALSE
	);
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal("Failed to create todos table:", err)
	}
}

// Get all todos
func getTodos(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, body, completed FROM todos ORDER BY id DESC")
	if err != nil {
		return err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Body, &t.Completed); err != nil {
			return err
		}
		todos = append(todos, t)
	}

	return c.JSON(todos)
}

// Create a new todo
func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	var id int
	err := db.QueryRow(
		"INSERT INTO todos (body) VALUES ($1) RETURNING id",
		todo.Body,
	).Scan(&id)
	if err != nil {
		return err
	}

	todo.ID = id
	todo.Completed = false

	return c.Status(201).JSON(todo)
}

// Toggle a todo's completed state
func toggleTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	var current bool
	err = db.QueryRow("SELECT completed FROM todos WHERE id=$1", todoID).Scan(&current)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	newState := !current
	_, err = db.Exec("UPDATE todos SET completed=$1 WHERE id=$2", newState, todoID)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "completed": newState})
}

// Bulk update todos by IDs
func bulkUpdateTodos(c *fiber.Ctx) error {
	var payload struct {
		IDs       []int `json:"ids"`
		Completed bool  `json:"completed"`
	}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if len(payload.IDs) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No IDs provided"})
	}

	query := "UPDATE todos SET completed=$1 WHERE id = ANY($2)"
	_, err := db.Exec(query, payload.Completed, payload.IDs)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"success": true, "updated_ids": payload.IDs})
}

// Delete a todo
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	res, err := db.Exec("DELETE FROM todos WHERE id=$1", todoID)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	}

	return c.JSON(fiber.Map{"success": true})
}
