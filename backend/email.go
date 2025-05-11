package main

import (
	"bytes"
	"html/template"
	"net/smtp"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type MailData struct {
	Receiver   string
	Authcode   string
	Expiration string
	Location   string
}

var EmailCodeData struct {
	Data map[string]EmailCodeDataValue
	Mu   sync.Mutex
}

type EmailCodeDataValue struct {
	Code       string
	Expiration time.Time
}

type SendEmailDTO struct {
	Email string `form:"email"`
}

// 生成邮件内容体的函数
func getMailBody(maildata *MailData) ([]byte, error) {

	tmplPath := "/data/email.html"
	var tmpl *template.Template
	var err error

	if _, err = os.Stat(tmplPath); err == nil {
		tmpl, err = template.ParseFiles(tmplPath)
	} else {
		tmpl, err = template.New("default").Parse(`<html><body>
				{{.Receiver}}，你的验证码是 {{.Authcode}}。
				于{{.Expiration}}前有效，IP地址位于{{.Location}}。
		</body></html>`)
	}
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	buffer.Write([]byte(
		"From: Flood Guard <floodguard@126.com>\r\n" +
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

func AddSendEmailRoutes(r *gin.Engine, pbc *PreloaderBaseConfig, config *SMTPConfig) {

	r.GET("/email/:id", Preload(
		&PreloaderConfig{Bind: BindConfig{Param: true}},
		&SendEmailDTO{},
		func(c *gin.Context, _ *User, r *SendEmailDTO) {

			if ok := isValidEmail(r.Email); !ok {
				c.JSON(400, Resp("邮箱格式有误", nil, nil))
				return
			}

			if err := pbc.DB.First(
				new(User), "email = ?", r.Email,
			).Error; err == nil {
				c.JSON(400, Resp("这个邮箱已经被注册过了", err, nil))
				return
			}

			authcode := RandStr(6)
			expiration := time.Now().Add(10 * time.Minute)
			EmailCodeData.Mu.Lock()
			EmailCodeData.Data[r.Email] = EmailCodeDataValue{
				Code:       authcode,
				Expiration: expiration,
			}
			EmailCodeData.Mu.Unlock()

			address, err := LocateAddress(c.ClientIP())
			if err != nil {
				c.JSON(500, Resp("获取客户端地址失败", err, nil))
				return
			}

			mailBody, err := getMailBody(&MailData{
				Receiver:   r.Email,
				Authcode:   authcode,
				Expiration: expiration.Format("2006-01-02 15:04"),
				Location:   address,
			})
			if err != nil {
				c.JSON(500, Resp("邮件内容生成失败", err, nil))
				return
			}

			if err := smtp.SendMail(
				config.Server+":"+config.Port,
				smtp.PlainAuth("",
					config.Mail,
					config.Password,
					config.Server,
				),
				config.Mail,
				[]string{r.Email},
				mailBody,
			); err != nil {
				c.JSON(500, Resp("邮件发送失败", err, nil))
				return
			}

			c.JSON(200, Resp("邮件发送成功", nil, nil))
		},
	))
}
