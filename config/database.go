package config

import "github.com/spf13/viper"

type Database struct {
	Driver string
	Source string
}

var DbCfg Database

func InitDatabase(cfg *viper.Viper) {
	DbCfg = Database{
		Driver: cfg.GetString("driver"),
		Source: cfg.GetString("source"),
	}
}
