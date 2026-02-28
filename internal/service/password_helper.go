package service

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword genera un hash seguro de la contraseña usando bcrypt
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost = 10 (excelente balance entre seguridad y velocidad)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword compara una contraseña plana con su hash bcrypt
func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
