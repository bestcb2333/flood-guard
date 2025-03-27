package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type HandlerConfig int

const (
	LOGIN HandlerConfig = 1 << iota
	ADMIN
	PARAM
	JSON
	BIND
)

type LogicFunc[T any] func(c *gin.Context, u *User, params T)

func Preload[T any](
	logicFunc LogicFunc[T], config HandlerConfig, preloads ...string,
) gin.HandlerFunc {
	return func(c *gin.Context) {

		var u *User
		var r T

		if config&LOGIN != 0 {
			rawToken := c.GetHeader("Authorization")
			if len(rawToken) < 8 {
				c.JSON(401, Resp("token无效", nil, nil))
				return
			}

			token := rawToken[7:]

			jwtToken, err := jwt.Parse(
				token,
				func(t *jwt.Token) (any, error) {
					return []byte(Config.JWTKey), nil
				},
			)
			if err != nil {
				c.JSON(401, Resp("token签名密钥不正确", err, nil))
				return
			}

			claims, ok := jwtToken.Claims.(jwt.MapClaims)
			if !ok || !jwtToken.Valid {
				c.JSON(401, Resp("token格式不正确", nil, nil))
				return
			}

			userId := uint(claims["userId"].(float64))
			newToken, err := GetJwt(userId)
			if err != nil {
				c.JSON(500, Resp("生成新token失败", err, nil))
				return
			}
			c.Header("Authorization", newToken)

			user := User{ID: userId}

			query := DB.Model(new(User))
			for _, preload := range preloads {
				query = query.Preload(preload)
			}

			if err := query.First(&user).Error; err == gorm.ErrRecordNotFound {
				c.JSON(401, Resp("用户不存在", nil, nil))
				return
			} else if err != nil {
				c.JSON(401, Resp("查询用户失败", err, nil))
				return
			}

			u = &user
		}

		if config&ADMIN != 0 && !u.Admin {
			c.JSON(403, Resp("你不是管理员", nil, nil))
			return
		}

		if config&JSON != 0 {
			if err := c.BindJSON(&r); err != nil {
				c.JSON(400, Resp("请求格式有误", err, nil))
				return
			}
		}

		if config&PARAM != 0 {
			if err := c.BindQuery(&r); err != nil {
				c.JSON(400, Resp("URL参数格式有误", err, nil))
				return
			}
		}

		if config&BIND != 0 {
			if err := c.Bind(&r); err != nil {
				c.JSON(400, Resp("表单格式有误", err, nil))
				return
			}
		}

		logicFunc(c, u, r)
	}
}
