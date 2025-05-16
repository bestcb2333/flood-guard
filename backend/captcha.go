package main

import (
	p "github.com/bestcb2333/gin-gorm-preloader/preloader"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func RegCaptchahandler(r *gin.Engine, bc *p.BaseConfig, cfg *Config) {

	r.GET("/captcha", func(c *gin.Context) {

		id := captcha.New()
		c.Header("Content-Type", "image/png")
		c.Header("X-Captcha-Id", id)

		if err := captcha.WriteImage(
			c.Writer, id, captcha.StdWidth, captcha.StdHeight,
		); err != nil {
			c.JSON(500, Resp("验证码绘制失败", err, nil))
			return
		}
	})

}

func AuthCaptcha(c *gin.Context) {

	captchaId := c.GetHeader("X-Captcha-Id")
	captchaValue := c.GetHeader("X-Captcha-Value")

	if !captcha.VerifyString(captchaId, captchaValue) {
		c.JSON(400, Resp("验证码不正确", nil, nil))
		return
	}
}
