package interfaces

import (
	"context"
	"sugarerpgo/internal/application"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	QueryHandler *application.GetUsersHandler
}

func (h *UserHandler) GetUsers(c fiber.Ctx) error {
	tenantID := "1"

	if tenantID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "tenant id required"})
	}
	users, err := h.QueryHandler.Handle(context.Background(), application.GetUsersQuery{TenantID: tenantID})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
