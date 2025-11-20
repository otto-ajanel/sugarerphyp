package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/otto-ajanel/backgo_tpdp_np/internal/service"
)

var authService = service.NewAuthService()

type loginReq struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

// Login maneja POST /api/login
func Login(c *fiber.Ctx) error {
    var req loginReq
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
    }
    user, token, err := authService.Login(req.Email, req.Password)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"user": user, "token": token})
}
