package commons

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const password = "MEU_TOKEN"

func GenerateToken(login string) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   login,
		Issuer:    "authorization",
		ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(password))
}

func ValidateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(password), nil
	})

	return err
}
