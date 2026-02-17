package auth

import (
    "errors"

    "github.com/golang-jwt/jwt/v5"
)

// Nota: usamos la misma clave base64 que en el backend Hyperf para compatibilidad.
var jwtKey = []byte("hiG8DlOKvtih6AxlZn5XKImZ06yu8I3mkOzaJrEuW8yAv8Jnkw330uMt8AEqQ5LB")

// ParseToken parsea y valida un token JWT (HS256). Devuelve los claims en un map.
func ParseToken(tokenStr string) (map[string]interface{}, error) {
    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Convertir a map[string]interface{}
        m := make(map[string]interface{})
        for k, v := range claims {
            m[k] = v
        }
        return m, nil
    }
    return nil, errors.New("invalid token")
}

// GenerateToken crea un JWT simple con los claims pasados.
func GenerateToken(claims map[string]interface{}) (string, error) {
    c := jwt.MapClaims{}
    for k, v := range claims {
        c[k] = v
    }
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    return t.SignedString(jwtKey)
}
