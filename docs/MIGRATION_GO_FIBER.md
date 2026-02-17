# Migración a Go + Fiber — guía rápida

Este directorio contiene un scaffold inicial para migrar tu backend Hyperf (PHP) a Go con Fiber.

Arquitectura propuesta (ya creada parcialmente):
- `cmd/server` (main)
- `internal/api` — handlers y registro de rutas
- `internal/service` — lógica de negocio (ej: AuthService)
- `internal/repository` — acceso a datos (GORM)
- `internal/model` — DTOs y structs de dominio
- `internal/infra/auth` — utilidades JWT
- `internal/middleware` — middleware (auth, tenant)

Qué hay funcionando ahora:
- Routes registradas en `internal/api/routes.go`.
- Handler de login minimal en `internal/api/handlers/auth.go` (usa credenciales de ejemplo `test@example.com` / `password`).
- Middleware de autenticación que parsea JWT y pone `userData` y `tenant` en el contexto de la request.

Pasos recomendados para continuar la migración:
1. Instala dependencias (desde la carpeta `backgo_tpdp_np`):

```powershell
go get github.com/gofiber/fiber/v2
go get github.com/golang-jwt/jwt/v5
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

2. Implementa la inicialización de conexiones a Postgres en `internal/infra/db` (se sugiere usar un map de pools por tenant si usas conexiones separadas por tenant).
3. Implementa repositorios con GORM (`internal/repository/*`) y precisa los modelos GORM (tags `gorm:"..."`).
4. Sustituye los stubs de `AuthService` y `CategoryRepo` por consultas reales a la DB.
5. Añade pruebas unitarias y de integración pequeñas para `AuthService` y `CategoryService`.

Si quieres, puedo:
- Parchar la lógica de `UserService::loginService` y `Authservice` a un `AuthService` real en Go que valide credenciales en Postgres y genere tokens compatibles.
- Implementar la selección por tenant (map de conexiones) y proponer migraciones SQL.
