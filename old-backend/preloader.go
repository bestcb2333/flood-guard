package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Permission int

const (
	Public Permission = iota
	Login
	Admin
)

type PreloaderConfig struct {
	DB          *gorm.DB
	JWTKey      string
	Binding     binding.Binding
	BindingBody binding.BindingBody
	Permission  Permission
	Tables      []string
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

			user := User{ID: userId}

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

		if def != nil {

			if pc.Binding != nil {
				if err := c.ShouldBindWith(&def, pc.Binding); err != nil {
					c.JSON(400, Resp("请求格式有误", err, nil))
					return
				}
			}

			if pc.BindingBody != nil {
				if err := c.ShouldBindBodyWith(&def, pc.BindingBody); err != nil {
					c.JSON(400, Resp("请求格式错误", err, nil))
					return
				}
			}

		}

		hf(c, u, def)
	}

}
