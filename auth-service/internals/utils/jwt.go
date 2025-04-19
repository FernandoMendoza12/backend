package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internals/config"
	"auth-service/internals/domain"

)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *domain.User) (string, error) {
	cfg := config.LoadConfig()
	secret := []byte(cfg.JWTSecretKey)

	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Subject:   fmt.Sprint(user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func ValidateToken(token string) (*Claims, error) {
	cfg := config.LoadConfig()
	secret := []byte(cfg.JWTSecretKey)

	parsed, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("metodo de firma invalido")
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(*Claims)
	if !ok || !parsed.Valid {
		return nil, errors.New("token invalido")
	}

	return claims,nil
}
