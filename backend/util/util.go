package util

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

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

func ParseJSON(c *gin.Context, object any) error {
	if bodyBytes, exists := c.Get("requestBody"); exists {
		if err := json.Unmarshal(bodyBytes.([]byte), object); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("没有请求体")
}
