package boot

import "github.com/spf13/viper"

type Jwt struct {
	Secret  string
	Timeout int
}

var JwtCfg Jwt

func InitJwt(cfg *viper.Viper) {
	JwtCfg = Jwt{
		Secret:  cfg.GetString("secret"),
		Timeout: cfg.GetInt("timeout"),
	}
}
