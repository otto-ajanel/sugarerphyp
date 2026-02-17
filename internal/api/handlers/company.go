package handlers

import (
	"sugarerpgo/internal/service"

	"github.com/gofiber/fiber/v3"
)

func GetCompanies(c fiber.Ctx) error {
	cs := service.NewCompanyService()
	companies, err := cs.GetAllCompanies()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(companies)
}
