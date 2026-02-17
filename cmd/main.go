package main

import (
	"database/sql"
	"fmt"
	"os"
	"sugarerpgo/internal/application"
	"sugarerpgo/internal/infrastructure"
	"sugarerpgo/internal/interfaces"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "host=localhost port=5432 user=postgres password=49720012 dbname=erpsugar sslmode=disable"
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	repo := &infrastructure.PostgresUserRepository{DB: db}
	eventBus := &infrastructure.SimpleEventBus{}
	handler := &interfaces.UserHandler{QueryHandler: &application.GetUsersHandler{Repo: repo, EventBus: eventBus}}

	app.Get("/users", handler.GetUsers)
	// ---

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting server on port %s\n", port)
	app.Listen(":" + port)
}
