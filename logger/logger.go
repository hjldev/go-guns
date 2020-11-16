package logger

import (
	"go-guns/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func Setup() {
	l := &lumberjack.Logger{
		Filename:   config.LogCfg.Path,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	log.SetOutput(l)
}

func Info(v ...interface{}) {
	go log.Println(v)
}
