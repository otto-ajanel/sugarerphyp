package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Welcome devuelve un mensaje público de bienvenida.
func Welcome(c *fiber.Ctx) error {

	return c.SendString("Welcome Go backend Sugar ERP by Otto Ajanel")
}
