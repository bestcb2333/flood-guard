package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 通过IP地址定位住址的函数
func LocateAddress(ip string) (string, error) {

	resp, err := http.Get(fmt.Sprintf(
		"https://whois.pconline.com.cn/ipJson.jsp?ip=%s&json=true", ip,
	))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("IP接口状态异常")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data["addr"], nil
}

// 通过用户ID获取JWT字符串
func GetJwt(userId uint) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		},
	).SignedString([]byte(Config.JWTKey))
}

// 制作响应体
func Resp(message string, err error, data any) gin.H {
	var errString *string
	if err != nil {
		data := err.Error()
		errString = &data
	}
	return gin.H{
		"message": message,
		"error":   errString,
		"data":    data,
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成伪随机字符串
func RandStr(n int) string {
	result := make([]byte, n)
	lettersLen := len(letters)
	for i := range result {
		result[i] = letters[rand.Intn(lettersLen)]
	}
	return string(result)
}

// 通过JWT字符串和预加载的表名获取用户对象
func GetUserByJwt(token string, tables ...string) (*User, error) {
	JwtToken, err := jwt.Parse(
		token[7:],
		func(t *jwt.Token) (any, error) {
			return []byte(Config.JWTKey), nil
		},
	)
	if err != nil {
		return nil, errors.New("token签名格式不正确: " + err.Error())
	}
	claims, ok := JwtToken.Claims.(jwt.MapClaims)
	if !ok || !JwtToken.Valid {
		return nil, errors.New("token身份信息有误")
	}

	var user User
	query := DB
	if len(tables) != 0 {
		for _, table := range tables {
			query.Preload(table)
		}
	}

	if err := query.First(
		&user, uint(claims["userId"].(float64)),
	).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetBodyByCtx(c *gin.Context, dest any) error {

	dataAny, exists := c.Get("body")
	if !exists {
		return errors.New("没有body")
	}

	data, ok := dataAny.([]byte)
	if !ok {
		return errors.New("body不为byteslice")
	}
	return json.Unmarshal(data, dest)
}

func ReadFileHeader[T any](fh *multipart.FileHeader) (*T, error) {

	file, err := fh.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	data = removeBOM(data)

	var dest T
	if err := json.Unmarshal(data, &dest); err != nil {
		return nil, err
	}

	return &dest, nil
}

func removeBOM(data []byte) []byte {
	if len(data) >= 3 && data[0] == 0xEF && data[1] == 0xBB && data[2] == 0xBF {
		return data[3:]
	}
	return data
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
