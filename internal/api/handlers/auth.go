package handlers

import (
	"sugarerpgo/internal/service"

	"github.com/gofiber/fiber/v3"
)

var authService = service.NewAuthService()

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login maneja POST /api/login
func Login(c fiber.Ctx) error {
	var req loginReq
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}
	user, token, err := authService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"user": user, "token": token})
}
