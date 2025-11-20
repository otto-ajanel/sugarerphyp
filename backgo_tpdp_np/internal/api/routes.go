package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/api/handlers"
	"github.com/otto-ajanel/backgo_tpdp_np/internal/middleware"
)

// RegisterRoutes registra las rutas públicas y privadas de la API.
func RegisterRoutes(app *fiber.App) {
	// Ruta pública
	app.Get("/", handlers.Welcome)
	app.Post("/api/login", handlers.Login)

	// Grupo /api/v1 protegido por JWT/tenant middleware
	v1 := app.Group("/api/v1", middleware.AuthRequired())
	v1.Get("/user", handlers.GetUsers)
	v1.Get("/permissionsbyuser", handlers.GetPermissionsByUser)
	v1.Get("/categories", handlers.GetCategories)
	v1.Post("/createcategory", handlers.CreateCategory)
	v1.Post("/product", handlers.CreateProduct)
	v1.Get("/products", handlers.GetProducts)
	v1.Get("/atributes", handlers.GetAtributes)
	v1.Post("/createatribute", handlers.CreateAtribute)
}
