package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

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
