package config

import "github.com/spf13/viper"

type Logger struct {
	Path string
}

var LogCfg Logger

func InitLog(cfg *viper.Viper) {
	LogCfg = Logger{
		Path: cfg.GetString("path"),
	}
}
