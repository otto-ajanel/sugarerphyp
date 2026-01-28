package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

// GetProducts devuelve todos los productos.
func GetProducts(c fiber.Ctx) error {
	ps := service.NewProductService()
	// parse pagination query params
	pageStr := c.Query("page", "1")
	perPageStr := c.Query("per_page", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	resp, err := ps.GetAllProductsPaginated(page, perPage)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

// CreateProduct crea un producto nuevo.
func CreateProduct(c fiber.Ctx) error {
	var req service.CreateProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	ps := service.NewProductService()
	res, err := ps.CreateProduct(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
