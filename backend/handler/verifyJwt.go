package handler

import (
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
		util.Error(c, 400, "token签名格式不正确!", nil)
		return nil, err
	}
	claims, ok := JwtToken.Claims.(jwt.MapClaims)
	if !ok || !JwtToken.Valid {
		util.Error(c, 401, "token身份信息有误", nil)
		return nil, err
	}

	var user db.User
	if err := DB.First(
		&user, claims["userId"].(uint),
	).Error; err != nil {
		util.HandleQueryErr(c, "找不到用户", err)
		return nil, err
	}

	return &user, nil
}
