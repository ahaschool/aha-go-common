package http

import (
	"github.com/ahaschool/aha-go-common/log"
	"fmt"
	"github.com/gin-gonic/gin"
)


func NewService(engine *gin.Engine, method string, path string, handler HandlerFunc)  {
	ctl := func(c *gin.Context) {
		ctx := NewContext(c)
		ctx.CH = make(chan bool)

		// 捕获代码中的直接抛错
		rec := func() {
			if r := recover(); r != nil {
				ctx.Res["code"] = 1221
				ctx.Res["message"] = "服务出现异常"
				ctx.Res["error"] = r
				content := fmt.Sprintf("panic: %s;", r)
				log.Error(content)
			}
			ctx.CH <- true
		}
		go func() {
			defer rec()
			handler(ctx)
		}()
		value := <-ctx.CH
		close(ctx.CH)
		c.JSON(200, ctx.Res)
		ctx.Tail["value"] = value
	}

	switch method {
	case "get":
		engine.GET(path, ctl)
	case "post":
		engine.POST(path, ctl)
	case "head":
		engine.HEAD(path, ctl)
	}
}
