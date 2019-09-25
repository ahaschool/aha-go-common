package log

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Dir string // log file dir
	LogType string //日志类型 file文件 std控制台
}

var (
	ilogger *log.Logger
	elogger *log.Logger
	LogType string
)

func Init(c *Config) {
	LogType = c.LogType
	ifile := fmt.Sprintf("%s/server.log", c.Dir)
	efile := fmt.Sprintf("%s/server-error.log", c.Dir)
	infofile, err := os.OpenFile(ifile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		log.Fatalln("打开或者创建info日志文件失败")
	}
	errorfile, err := os.OpenFile(efile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		log.Fatalln("打开或者创建error日志文件失败")
	}
	//defer file.Close() // 暂时不频繁的关闭文件句柄

	// 得到一个新的文件 默认采用debug模式
	ilogger = log.New(infofile, "[Info]", log.LstdFlags)
	elogger = log.New(errorfile, "[Error]", log.LstdFlags)

	// 设置记录格式
	ilogger.SetFlags(log.LstdFlags | log.Lshortfile)
	elogger.SetFlags(log.LstdFlags | log.Lshortfile)
}

func Debug(format string, args ...interface{})  {
	ilogger.SetPrefix("[Debug]")
	if LogType == "file" {
		ilogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}

func Info(format string, args ...interface{}) {
	ilogger.SetPrefix("[Info]")
	if LogType == "file" {
		ilogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}

func Warn(format string, args ...interface{})  {
	ilogger.SetPrefix("[Warn]")
	if LogType == "file" {
		ilogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}

func Error(format string, args ...interface{})  {
	elogger.SetPrefix("[Error]")
	if LogType == "file" {
		elogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}
