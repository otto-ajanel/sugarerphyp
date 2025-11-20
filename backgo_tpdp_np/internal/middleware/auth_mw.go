package middleware

import (
    "strings"

    "github.com/gofiber/fiber/v2"
    authinfra "github.com/otto-ajanel/backgo_tpdp_np/internal/infra/auth"
)

// AuthRequired valida JWT, setea userData y tenant en locals.
func AuthRequired() fiber.Handler {
    return func(c *fiber.Ctx) error {
        header := c.Get("Authorization")
        if header == "" {
            return c.Status(401).JSON(fiber.Map{"error": "missing authorization header"})
        }
        token := strings.TrimPrefix(header, "Bearer ")
        claims, err := authinfra.ParseToken(token)
        if err != nil {
            return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
        }
        // En el token original se guardaba "data" con userData.
        if data, ok := claims["data"]; ok {
            c.Locals("userData", data)
            // intentar obtener name_tenant si existe
            if m, ok := data.(map[string]interface{}); ok {
                if tn, ok2 := m["name_tenant"].(string); ok2 {
                    c.Locals("tenant", tn)
                }
            }
        }
        return c.Next()
    }
}
