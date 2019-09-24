package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 向网络发送http get请求
func Get(url string, params map[string]string, headers map[string]string) (response string, err error) {
	client := &http.Client{Timeout: 5 * time.Second}
	for k, v := range params {
		if strings.Contains(url, "?") {
			url += "&"
		} else {
			url += "?"
		}
		url += k + "=" + v
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	// 构建header数据
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = string(body)
	return
}

// 向网络发送http post请求
func Post(url string, data interface{}, headers map[string]string) (response string, err error) {
	// 参数序列化
	jsonStr, _ := json.Marshal(data)

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	if err != nil {
		return
	}

	// 构建header数据
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	response = string(body)
	return
}
