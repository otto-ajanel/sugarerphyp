package handlers

import (
	"strconv"
	"sugarerpgo/internal/service"

	"github.com/gofiber/fiber/v3"
)

func GetUsers(c fiber.Ctx) error {
	ps := service.NewUserServCrud()
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
	resp, err := ps.GetUsersPaginate(page, perPage)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(resp)
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

func CreateUser(c fiber.Ctx) error {
	var reqFormUser map[string]interface{}
	err := c.Bind().Body(&reqFormUser)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})

	}

	// echo back the request for now
	return c.JSON(reqFormUser)
}
