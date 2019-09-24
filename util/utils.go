package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"unicode/utf8"
)

// sign 加签处理
func Sign(secret, timestamp string, params map[string]interface{}) string {
	res := make(map[string]interface{})
	str := ""
	if len(params) == 0 {
		str = MD5(secret + timestamp)
	} else {
		var keys []string
		for k := range params {
			keys = append(keys, k)
		}
		for v, k := range keys {
			res[k] = v
		}
		buf, err := json.Marshal(res) //格式化编码
		if err != nil {
			return ""
		}
		str = MD5(secret + FilterEmoji(string(buf)) + timestamp)
	}
	return str
}

// filterEmoji 去掉emoji表情
func FilterEmoji(content string) string {
	newContent := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			newContent += string(value)
		}
	}
	return newContent
}

// MD5 md5加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}