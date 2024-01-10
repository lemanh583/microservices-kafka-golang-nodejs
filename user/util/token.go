package util

import (
	"time"
	"user-services/config"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(payload map[string]interface{}, timeExp *time.Time) (*string, error) {
	exp := time.Now().Add(30 * 24 * time.Hour)
	if timeExp != nil {
		exp = *timeExp
	}
	claims := jwt.MapClaims{
		"payload": payload,
		"exp":     exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Cfg.TokenSecretKey))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return config.Cfg.TokenSecretKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*jwt.MapClaims)
	return claims, nil
}
