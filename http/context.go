package http

import (
	"encoding/json"
	"github.com/ahaschool/aha-go-common/errcode"
	"github.com/ahaschool/aha-go-common/log"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type Context struct {
	CH   chan bool
	Time time.Time
	Gin  *gin.Context
	Res  map[string]interface{}
	Tail map[string]interface{}
}

type HandlerFunc func(*Context)

func NewContext(gin *gin.Context) (context *Context) {
	res := make(map[string]interface{})
	res["code"] = 0
	res["message"] = "success"
	res["data"] = nil
	context = &Context{Gin: gin, Res: res, Time: time.Now()}
	context.Tail = make(map[string]interface{})
	context.Tail["api"] = gin.Request.URL.Path
	context.Tail["query"] = gin.Request.URL.RawQuery
	context.Tail["method"] = gin.Request.Method
	context.Tail["time"] = context.Time
	context.Tail["nano"] = context.Time.UnixNano()
	context.Tail["guid"] = gin.GetHeader("X-Env-Guniqid")
	context.Res["server_time"] = time.Now().Format("2006-01-02 15:04:05")
	return
}

// GetAhaUserID 获取用户编号
func (ctx *Context) GetAhaUserID() int {
	var str string
	str = ctx.Gin.DefaultQuery("user_id", "")
	UserID, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return UserID
}

func (ctx *Context) GetHeader(key string) string {
	return ctx.Gin.GetHeader(key)
}

func (ctx *Context) BindQuery(obj interface{}) (err error) {
	err = ctx.Gin.ShouldBindQuery(obj)
	if err != nil {
		ctx.Res["error"] = err
	}
	return
}

func (ctx *Context) Bind(obj interface{}) (err error) {
	err = ctx.Gin.ShouldBindJSON(obj)
	if err != nil {
		ctx.Res["error"] = err
	}
	return
}

func (ctx *Context) Error(err interface{}) {
	switch err.(type) {
	case *errcode.Status:
		ctx.Res["code"] = err.(*errcode.Status).C
		ctx.Res["message"] = err.(*errcode.Status).M
	default:
		ctx.Res["code"] = 500
		ctx.Res["message"] = err.(error).Error()
	}
	sjson, _ := json.Marshal(ctx.Res)
	body, _ := ctx.Gin.Get("body")
	hjson, _ := json.Marshal(ctx.Gin.Request.Header)
	log.Info("| %s  %s | %s| %s| %s",
		ctx.Gin.Request.Method,
		ctx.Gin.Request.RequestURI,
		hjson,
		body,
		sjson,
	)
}

func (ctx *Context) Response(data interface{}) {
	ctx.Res["data"] = data
	status := errcode.Success
	ctx.Res["code"] = status.Code()
	ctx.Res["message"] = status.Message()
}
