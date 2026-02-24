package api

import (
	"sugarerpgo/internal/api/handlers"
	"sugarerpgo/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

// RegisterRoutes registra las rutas públicas y privadas de la API.
func RegisterRoutes(app *fiber.App) {
	// Ruta pública
	app.Get("/", handlers.Welcome)
	app.Post("/api/login", handlers.Login)

	// Grupo /api/v1 protegido por JWT/tenant middleware
	v1 := app.Group("/api/v1", middleware.AuthRequired())
	v1.Get("/user", handlers.GetUsers)
	v1.Post("/user", handlers.CreateUser)
	v1.Get("/permissionsbyuser", handlers.GetPermissionsByUser)
	v1.Get("/categories", handlers.GetCategories)
	v1.Post("/createcategory", handlers.CreateCategory)
	v1.Post("/product", handlers.CreateProduct)
	v1.Get("/products", handlers.GetProducts)
	v1.Get("/all-products", handlers.GetAllProducts)
	v1.Get("/atributes", handlers.GetAtributes)
	v1.Post("/atributes", handlers.CreateAtribute)
	v1.Post("/createatribute", handlers.CreateAtribute)
	v1.Get("/atributedetails", handlers.GetAtributeDetail)
	v1.Post("/createatributedetail", handlers.CreateAtributeDetail)
	v1.Post("/uploadproductimage", handlers.UploadProductImage)
	v1.Get("/companies", handlers.GetCompanies)
	v1.Get("/stores", handlers.GetStores)
	v1.Get("/incomes", handlers.GetIncomes)
	v1.Post("/incomes", handlers.CreateIncome)
	v1.Get("/suppliers", handlers.GetSuppliers)
	v1.Get("/productsaviable", handlers.GetProductsAviable)
	v1.Get("/getimageproduct", handlers.GetImageProduct)
}
