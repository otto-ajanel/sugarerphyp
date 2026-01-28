package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

func GetSuppliers(c fiber.Ctx) error {
	ss := service.NewSupplierService()
	res, err := ss.GetAllSuppliers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
