package aha_service

type ReqPassport struct {
	Host  string
	Token string
}


// GetUserInfoByAuth 根据auth_token 获取用户信息
func (q *ReqPassport) GetUserInfoByAuth(xenvStr, authToken string) (string, error) {
	data := map[string]string{
		"access_token": authToken,
	}
	conf := &Config{
		ReqHost: q.Host,
		XToken:  q.Token,
	}
	res, err := ahaGet(conf, xenvStr, "/authorization/verify", data)
	if err != nil {
		return "", err
	}
	return res, nil
}
