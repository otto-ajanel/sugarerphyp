package handlers

import (
	"sugarerpgo/internal/service"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(c fiber.Ctx) error {
	// Placeholder: en una futura iteración se puede implementar listado de usuarios.
	return c.JSON(fiber.Map{"message": "GetUsers - pendiente de migrar lógica"})
}

// GetPermissionsByUser obtiene las permissions del usuario autenticado.
func GetPermissionsByUser(c fiber.Ctx) error {
	userData := c.Locals("userData")
	if userData == nil {
		return c.Status(401).JSON(fiber.Map{"error": "userData not found in context"})
	}
	dataMap, ok := userData.(map[string]interface{})
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "invalid userData format"})
	}

	var userID int
	switch v := dataMap["id_user"].(type) {
	case float64:
		userID = int(v)
	case int:
		userID = v
	default:
		return c.Status(400).JSON(fiber.Map{"error": "invalid id_user type in token"})
	}

	us := service.NewUserService()
	perms, err := us.GetPermissionsByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(perms)
}
