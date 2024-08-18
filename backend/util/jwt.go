package util

import (
	"time"

	"github.com/bestcb2333/FloodGuard/database"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func GetJwt(userId database.CustomUint) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": uint(userId),
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}).SignedString([]byte(viper.GetString("JWT_KEY")))
}
