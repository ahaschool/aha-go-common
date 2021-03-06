package aha_service

import (
	"github.com/ahaschool/aha-go-common/http"
	"github.com/ahaschool/aha-go-common/util"
)

type Config struct {
	XToken  string
	ReqHost string
}

func AhaPost(c *Config, xenvStr string, path string, data interface{}) (response string, err error) {
	// decode xenv value
	xenv, err := util.XenvDecode(xenvStr)
	if err != nil {
		xenv = make(map[string]string)
	}
	// build header params
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["X-Token"] = c.XToken

	// build xenv header
	for k, v := range xenv {
		headers["X-Env-"+k] = v
	}
	// return http response
	return http.Post(c.ReqHost+path, data, headers)
}

func AhaGet(c *Config, xenvStr string, path string, params map[string]string) (response string, err error) {

	// decode xenv value
	xenv, err := util.XenvDecode(xenvStr)
	if err != nil {
		xenv = make(map[string]string)
	}

	// build header params
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["X-Token"] = c.XToken

	// build xenv header
	for k, v := range xenv {
		headers["X-Env-"+k] = v
	}

	// return http response
	return http.Get(c.ReqHost+path, params, headers)
}
