package util

import (
	"github.com/golang-jwt/jwt/v4"
	"go-admin/app/wechat/config"
	"time"
)

func GenerateToken(openid string) (string, error) {
	claims := jwt.MapClaims{
		"openid": openid,
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtSecret))
}
