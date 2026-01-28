package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

// espera en el body datos y un archivo imagen
func UploadProductImage(c fiber.Ctx) error {
	var req service.UploadProductImageRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	file, _ := c.FormFile("imageFile")
	req.ImageFile = file
	ps := service.NewProductPathService()
	res, err := ps.UploadProductImage(req)
	if err != nil {

		return c.Status(500).JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(res)
}
