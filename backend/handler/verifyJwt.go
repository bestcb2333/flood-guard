package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func VerifyJwt(c *gin.Context) {
	JwtToken, err := jwt.Parse(
		c.GetHeader("Authorization")[7:],
		func(t *jwt.Token) (any, error) {
			return []byte(viper.GetString("JwtKey")), nil
		},
	)
	if err != nil {

	}
}
