package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

// GetAtributes devuelve todos los atributos.
func GetAtributes(c *fiber.Ctx) error {
    as := service.NewAtributeService()
    arr, err := as.GetAllAtributes()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(arr)
}

// CreateAtribute crea un nuevo atributo.
func CreateAtribute(c *fiber.Ctx) error {
    var req service.CreateAtributeRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
    }
    as := service.NewAtributeService()
    a, err := as.CreateAtribute(req)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(a)
}
