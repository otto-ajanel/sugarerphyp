package handlers

import (
	"sugarerpgo/internal/service"

	"github.com/gofiber/fiber/v3"
)

// GetStores devuelve todas las tiendas.
func GetStores(c fiber.Ctx) error {
	ss := service.NewStoreService()
	stores, err := ss.GetAllStores()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(stores)
}
