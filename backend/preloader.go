package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Permission int

const (
	Public Permission = iota
	Login
	Admin
)

type PreloaderBaseConfig struct {
	DB     *gorm.DB
	JWTKey string
}

type PreloaderBaseConfiger interface {
	GetBaseConfig() PreloaderBaseConfig
}

type PreloaderConfig struct {
	*PreloaderBaseConfig
	Bind       BindConfig
	Permission Permission
	Tables     []string
}

type BindConfig struct {
	Param     bool
	Query     bool
	JSON      bool
	Multipart bool
}

type HandlerFunc[T any] func(c *gin.Context, u *User, r *T)

func Preload[T any](
	pc *PreloaderConfig,
	def *T,
	hf HandlerFunc[T],
) gin.HandlerFunc {
	return func(c *gin.Context) {

		var u *User

		if pc.Permission >= Login {
			rawToken := c.GetHeader("Authorization")
			if len(rawToken) < 8 {
				c.AbortWithStatusJSON(400, Resp("token无效", nil, nil))
				return
			}

			token := rawToken[7:]

			jwtToken, err := jwt.Parse(
				token,
				func(t *jwt.Token) (any, error) {
					return []byte(pc.JWTKey), nil
				},
			)
			if err != nil {
				c.AbortWithStatusJSON(400, Resp("token秘钥不正确", err, nil))
				return
			}

			claims, ok := jwtToken.Claims.(jwt.MapClaims)
			if !ok || !jwtToken.Valid {
				c.AbortWithStatusJSON(400, Resp("token格式不正确", nil, nil))
				return
			}

			userId := uint(claims["userId"].(float64))

			newToken, err := GetJwt(userId, pc.JWTKey)
			if err != nil {
				c.AbortWithStatusJSON(400, Resp("生成新token失败", err, nil))
				return
			}
			c.Header("Authorization", newToken)

			var user User
			user.ID = userId

			query := pc.DB
			if pc.Tables != nil {
				for _, value := range pc.Tables {
					query = query.Preload(value)
				}
			}

			if err := query.First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				c.AbortWithStatusJSON(400, Resp("用户不存在", err, nil))
				return
			} else if err != nil {
				c.AbortWithStatusJSON(500, Resp("查询用户失败", err, nil))
				return
			}

			if pc.Permission >= Admin && !u.Admin {
				c.AbortWithStatusJSON(403, Resp("你不是管理员", nil, nil))
				return
			}

			u = &user
		}

		if pc.Bind.Param {
			if err := c.ShouldBindUri(&def); err != nil {
				c.JSON(400, Resp("路径参数有误", err, nil))
				return
			}
		}

		if pc.Bind.Query {
			if err := c.ShouldBindQuery(&def); err != nil {
				c.JSON(400, Resp("查询字符串参数有误", err, nil))
				return
			}
		}

		if pc.Bind.JSON {
			if err := c.ShouldBindJSON(&def); err != nil {
				c.JSON(400, Resp("请求体有误", err, nil))
				return
			}
		}

		if pc.Bind.Multipart {
			if err := c.ShouldBind(&def); err != nil {
				c.JSON(400, Resp("请求表单错误", err, nil))
				return
			}
		}

		hf(c, u, def)
	}

}
