package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

// GetProducts devuelve todos los productos.
func GetProducts(c *fiber.Ctx) error {
	ps := service.NewProductService()
	products, err := ps.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

// CreateProduct crea un producto nuevo.
func CreateProduct(c *fiber.Ctx) error {
	var req service.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	ps := service.NewProductService()
	res, err := ps.CreateProduct(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
