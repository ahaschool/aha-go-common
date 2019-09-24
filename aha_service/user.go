package aha_service

type ReqUser struct {
	Host  string
	Token string
}

// 根据mobile向用户中心获取一个用户
func (q *ReqUser) PostUserConfirm(xenvStr string, mobile string, attributes map[string]string) (response string, err error) {
	data := attributes
	data["mobile"] = mobile
	conf := &Config{
		ReqHost: q.Host,
		XToken:  q.Token,
	}
	return ahaPost(conf, xenvStr, "/user/create", data)
}

