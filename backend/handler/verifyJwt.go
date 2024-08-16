package handler

import (
	"errors"

	db "github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func VerifyJwt(c *gin.Context) (*db.User, error) {
	JwtToken, err := jwt.Parse(
		c.GetHeader("Authorization")[7:],
		func(t *jwt.Token) (any, error) {
			return []byte(viper.GetString("JWT_KEY")), nil
		},
	)
	if err != nil {
		return nil, errors.New("token签名格式不正确: " + err.Error())
	}
	claims, ok := JwtToken.Claims.(jwt.MapClaims)
	if !ok || !JwtToken.Valid {
		return nil, errors.New("token身份信息有误")
	}

	var user db.User
	if err := DB.First(
		&user, uint(claims["userId"].(float64)),
	).Error; err != nil {
		util.HandleQueryErr(c, "找不到用户", err)
		return nil, err
	}

	return &user, nil
}
