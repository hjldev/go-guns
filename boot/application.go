package boot

import "github.com/spf13/viper"

type Application struct {
	Name     string
	Port     string
	Mode     string
	Enabledp bool
}

var AppCfg Application

func InitApplication(cfg *viper.Viper) {
	AppCfg = Application{
		Name:     cfg.GetString("name"),
		Port:     cfg.GetString("port"),
		Mode:     cfg.GetString("mode"),
		Enabledp: cfg.GetBool("enabledp"),
	}
}
