package utils

import (
	"strings"

	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/models"
	"github.com/Edilberto-Vazquez/golang-rest-and-websockets/server"
	"github.com/golang-jwt/jwt/v4"
)

func RouteNeedToken(route string, NO_AUTH_NEEDED []string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func ProcessToken(authorization string, s server.Server) (*jwt.Token, error) {
	tokenString := strings.TrimSpace(authorization)
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	return token, err
}
