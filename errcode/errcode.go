package errcode

import (
	//"aha-api-server/conf"
	"github.com/micro/go-micro/errors"
	"net/http"
)

// 全部的错误返回代号
var (

	// 基本代号
	Success      = Add(0, "success")
	Unauthorized = Add(401, "认证失败")
	NothingFound = Add(404, "数据不存在")
	Canceled     = Add(498, "客户端取消请求")
	ServerErr    = Add(500, "服务器错误")
	Deadline     = Add(504, "服务调用超时")

	// 业务错误代号
	ParamsError = Add(10001, "参数错误")
)

type Status struct {
	C int32
	M string
}

// 获取错误代码
func (s *Status) Code() int32 {
	return int32(s.C)
}

// 获取错误内容
func (s *Status) Message() string {
	return s.M
}

// 返回错误状态
func Add(code int32, message string) *Status {
	return &Status{
		C: code,
		M: message,
	}
}

// 返回错误状态
func ServerAdd(code int32, message string) *errors.Error {
	return &errors.Error{
		Code:   code,
		Detail: message,
		Status: http.StatusText(200),
	}
}

// 返回错误状态
func Error(err *errors.Error, message string) *errors.Error {
	if message != "" {
		err.Detail = message
	}
	return &errors.Error{
		Code:   err.Code,
		Detail: err.Detail,
		Status: http.StatusText(200),
	}
}
