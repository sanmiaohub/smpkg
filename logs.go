package smpkg

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
)

type Logger struct {
	Origin *log.Logger
}

type LoggerFormat struct {
	Rid     string      `json:"rid"`
	Time    string      `json:"time"`
	Level   string      `json:"level"`
	Content interface{} `json:"content"`
}

func NewLog(channel string) *Logger {
	var err error
	var write io.Writer

	switch channel {
	case "file":
		write, err = os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(any("打开日志文件异常:" + err.Error()))
		}
		break
	case "console":
	default:
		write = os.Stdout
	}
	// 赋值应用日志工具
	return &Logger{
		Origin: log.New(write, "", 0),
	}
}

func (l *Logger) Debug(ctx CtxContext, data interface{}) {
	l.write(ctx, data, "debug")
}

func (l *Logger) Info(ctx CtxContext, data interface{}) {
	l.write(ctx, data, "info")
}

func (l *Logger) Error(ctx CtxContext, data interface{}) {
	l.write(ctx, data, "error")
}

func (l *Logger) write(ctx CtxContext, data interface{}, level string) {
	dataFormat := LoggerFormat{
		Rid:     ctx.RequestID(),
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Level:   level,
		Content: data,
	}

	writeData, err := json.Marshal(dataFormat)
	if err != nil {
		panic(any("写入日志发生异常:" + err.Error()))
	}
	l.Origin.Print(string(writeData))
}
