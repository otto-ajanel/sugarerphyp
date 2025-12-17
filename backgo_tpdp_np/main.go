package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/api"
)

func main() {
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Registrar CORS globalmente (ajusta AllowOrigins si quieres restringir)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Registrar rutas desde internal/api
	api.RegisterRoutes(app)

	if err := app.Listen("0.0.0.0:3000"); err != nil {
		log.Fatalf("error al iniciar servidor: %v", err)
	}
}
