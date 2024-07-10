package util

import "math/rand"

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
