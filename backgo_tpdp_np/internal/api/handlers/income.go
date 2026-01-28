package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

func GetIncomes(c fiber.Ctx) error {
	is := service.NewIncomeService()
	res, err := is.GetAllIncomes()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}
func CreateIncome(c fiber.Ctx) error {
	var req = make(map[string]interface{})
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	is := service.NewIncomeService()
	res, err := is.CreateIncome(req, c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}
