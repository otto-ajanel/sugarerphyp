package handlers

import (
	"github.com/gofiber/fiber/v3"
)

// Welcome devuelve un mensaje público de bienvenida.
func Welcome(c fiber.Ctx) error {
	//Message de bienvenida público en consola y en la ruta /
	smg := "Welcome Go backend Sugar ERP by @Otto Ajanel"
	println(smg)
	return c.SendString("Welcome Go backend Sugar ERP by Otto Ajanel")
}
