package handlers

import (
	"sugarerpgo/internal/service"

	"github.com/gofiber/fiber/v3"
)

func GetSuppliers(c fiber.Ctx) error {
	ss := service.NewSupplierService()
	res, err := ss.GetAllSuppliers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
