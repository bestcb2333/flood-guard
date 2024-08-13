package handler

import (
	"bytes"
	"html/template"
	"net/smtp"
	"time"

	"github.com/bestcb2333/FloodGuard/asset"
	db "github.com/bestcb2333/FloodGuard/database"
	"github.com/bestcb2333/FloodGuard/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type MailData struct {
	Receiver   string
	Authcode   string
	Expiration string
	Location   string
}

type MailSentValue struct {
	Receiver   string
	Expiration time.Time
}

var MailSent = make(map[string]MailSentValue)

func GetMail(c *gin.Context) {

	email := c.Query("email")

	if err := DB.First(
		new(db.User), "email = ?", email,
	).Error; err == nil {
		util.Error(c, 400, "这个邮箱已经被注册过了", nil)
		return
	}

	authcode := util.RandStr(6)
	expiration := time.Now().Add(10 * time.Minute)
	MailSent[authcode] = MailSentValue{
		Receiver:   email,
		Expiration: expiration,
	}

	address, err := util.LocateAddress(c.ClientIP())
	if err != nil {
		util.Error(c, 500, "获取客户端住址失败", err)
		return
	}

	mailBody, err := getMailBody(&MailData{
		Receiver:   email,
		Authcode:   authcode,
		Expiration: expiration.Format("2006-01-02 15:04"),
		Location:   address,
	})
	if err != nil {
		util.Error(c, 500, "邮件内容生成失败", err)
		return
	}

	if err := smtp.SendMail(
		viper.GetString("SMTP_SERVER")+":"+viper.GetString("SMTP_PORT"),
		smtp.PlainAuth("",
			viper.GetString("SMTP_MAIL"),
			viper.GetString("SMTP_PASSWORD"),
			viper.GetString("SMTP_SERVER"),
		),
		viper.GetString("SMTP_MAIL"),
		[]string{email},
		mailBody,
	); err != nil {
		util.Error(c, 500, "邮件发送失败", nil)
	}

	util.Info(c, 200, "邮件发送成功", nil)
}

// 生成邮件内容体的函数
func getMailBody(maildata *MailData) ([]byte, error) {

	tmpl, err := template.New("email").ParseFS(asset.FS, "email.html")
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	buffer.Write([]byte(
		"From: Axolotland Gaming Club <axolotland@163.com>\r\n" +
			"To: " + maildata.Receiver + "\r\n" +
			"Subject: 验证码邮件\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n",
	))

	if err := tmpl.Execute(&buffer, maildata); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// 验证邮箱的函数
func AuthMail(authcode, email string) bool {
	value := MailSent[authcode]
	if value.Receiver != email || time.Now().After(value.Expiration) {
		return false
	}
	delete(MailSent, authcode)
	return true
}
