package handler

import (
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {

	id := captcha.New()
	c.Header("Content-Type", "image/png")
	c.Header("X-Captcha-Id", id)

	if err := captcha.WriteImage(
		c.Writer, id, captcha.StdWidth, captcha.StdHeight,
	); err != nil {
		util.Error(c, 500, "验证码绘制失败", err)
		return
	}
}

func AuthCaptcha(c *gin.Context) {

	var request struct {
		CaptchaId    string
		CaptchaValue string
	}

	if util.ParseJSON(c, &request) != nil {
		util.Error(c, 400, "无法读取你的请求", nil)
		return
	}

	if !captcha.VerifyString(request.CaptchaId, request.CaptchaValue) {
		util.Error(c, 400, "验证码不正确", nil)
		return
	}
}
