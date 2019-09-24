package log

import (
	"testing"
)

func TestLog(t *testing.T)  {
	Init(&Config{
		Dir: "/tmp/",
	})
	testInfo()
	testWarn()
	testDebug()
	testError()
}

func testInfo()  {
	Info("这是一条INFO日志")
}

func testDebug() {
	Debug("这是一条DEBUG日志")
}

func testError()  {
	Error("这是一条ERROR日志")
}

func testWarn()  {
	Warn("这是一条WARN日志")
}
