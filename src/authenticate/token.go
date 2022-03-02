package authenticate

import (
	"api/src/config"
	"time"

	"github.com/golang-jwt/jwt"
)

// TokenGenerate gera um token assinado digitalmente como forma de autenticação
func TokenGenerate(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}
