package service

import (
	"errors"
	"fmt"

	authinfra "sugarerpgo/internal/infra/auth"

	"sugarerpgo/internal/infra/db"
	"sugarerpgo/internal/model"
)

type AuthService struct{}

func NewAuthService() *AuthService { return &AuthService{} }

// Login consulta la base de datos y devuelve el usuario + token si las credenciales son válidas.
func (s *AuthService) Login(email, password string) (map[string]interface{}, string, error) {
	// Conectar a la DB (usa variables de entorno)
	gdb, err := db.Get()
	if err != nil {
		return nil, "", fmt.Errorf("db connect error: %w", err)
	}

	// Estructura para recibir los datos (incluye name_tenant desde join)
	var result struct {
		model.User
		NameTenant string `gorm:"column:name_tenant"`
	}

	// Consulta: join tenants para obtener name_tenant
	q := gdb.Table("users").Select("users.*, tenants.name_tenant").Joins("join tenants on tenants.id_tenant = users.id_tenant").Where("users.email = ? AND users.password = ? AND users.active = ?", email, password, true)
	if err := q.First(&result).Error; err != nil {
		return nil, "", errors.New("invalid credentials or user not found")
	}

	// Preparar respuesta pública (no exponer password)
	userMap := map[string]interface{}{
		"id_user":     result.IDUser,
		"username":    result.Username,
		"email":       result.Email,
		"active":      result.Active,
		"name_tenant": result.NameTenant,
	}

	// Generar token con el claim "data" similar al backend PHP
	token, err := authinfra.GenerateToken(map[string]interface{}{"data": userMap, "name_tenant": result.NameTenant})
	if err != nil {
		return nil, "", fmt.Errorf("token gen error: %w", err)
	}

	return userMap, token, nil
}
