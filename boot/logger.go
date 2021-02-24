package boot

import (
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

type Logger struct {
	Path string
}

var LogCfg Logger

func InitLog(cfg *viper.Viper) {
	logger := Logger{
		Path: cfg.GetString("path"),
	}

	logger.Setup()
}

func (logger *Logger) Setup() {
	l := &lumberjack.Logger{
		Filename:   logger.Path,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	log.SetOutput(l)
}
