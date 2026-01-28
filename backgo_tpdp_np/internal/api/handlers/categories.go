package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

// GetCategories lista todas las categorías.
func GetCategories(c fiber.Ctx) error {
	cs := service.NewCategoryService()
	cats, err := cs.GetAllCategories()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cats)
}

// CreateCategory crea una nueva categoría desde el body { "newCategory": "Nombre" }.
func CreateCategory(c fiber.Ctx) error {
	var req service.CreateCategoryRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	cs := service.NewCategoryService()
	cat, err := cs.CreateCategory(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(cat)
}
