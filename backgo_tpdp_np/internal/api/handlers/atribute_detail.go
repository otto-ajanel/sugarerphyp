package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

// GetAtributes devuelve todos los atributos.
func GetAtributeDetail(c *fiber.Ctx) error {
	as := service.NewAtributeDetailService()
	// Soporta filtro opcional ?atribute_id=123 para devolver solo los detalles de ese atributo
	if idStr := c.Query("atribute_id"); idStr != "" {
		// parsear int
		aID, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid atribute_id"})
		}
		arr, err := as.GetAtributeDetailsByAtributeID(aID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(arr)
	}
	arr, err := as.GetAllAtributeDetails()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(arr)
}

// CreateAtribute crea un nuevo atributo.
func CreateAtributeDetail(c *fiber.Ctx) error {
	var req service.CreateAtributeDetailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	as := service.NewAtributeDetailService()
	a, err := as.CreateAtributeDetail(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(a)
}
