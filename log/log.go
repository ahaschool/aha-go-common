package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Dir string // log file dir
	LogType string //日志类型 file文件 std控制台
	InfoFileName string //info 文件名称
	ErrorFileName string // error 文件名称
}

var (
	ilogger *log.Logger
	elogger *log.Logger
	LogType string
	ifile string
	efile string
	Dir string
)

func Init(c *Config) {
	LogType = c.LogType
	Dir = c.Dir
	os.Mkdir(c.Dir, 0744)
	ifile = fmt.Sprintf("%s/%s", c.Dir, c.InfoFileName)
	efile = fmt.Sprintf("%s/%s", c.Dir, c.ErrorFileName)
	os.Mkdir(filepath.Dir(ifile)+"/history", 0744)
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
	elogger.SetFlags(log.Llongfile)
	ilogger.SetFlags(log.Llongfile)
}

func Debug(format string, args ...interface{})  {
	ilogger.SetPrefix(time.Now().Format("2006-01-02 15:04:05")+"[Debug]")
	if LogType == "file" {
		ilogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}

func Info(format string, args ...interface{}) {
	ilogger.SetPrefix(time.Now().Format("2006-01-02 15:04:05")+"[Info]")
	if LogType == "file" {
		ilogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}

func Warn(format string, args ...interface{})  {
	ilogger.SetPrefix(time.Now().Format("2006-01-02 15:04:05")+"[Warn]")
	if LogType == "file" {
		ilogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}

func Error(format string, args ...interface{})  {
	elogger.SetPrefix(time.Now().Format("2006-01-02 15:04:05")+"[Error]")
	if LogType == "file" {
		elogger.Output(2, fmt.Sprintf(format, args...))
	}
	if LogType == "std" {
		log.Println(fmt.Sprintf(format, args...))
	}
}
