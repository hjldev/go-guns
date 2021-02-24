package boot

import (
	"fmt"
	"github.com/spf13/viper"
)

func AdminBoot(module, env string) {
	configFile := "config/" + module + "-" + env + ".yml"
	vi := viper.New()
	vi.SetConfigFile(configFile)
	if err := vi.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	// 初始化应用相关信息
	appCfg := vi.Sub("application")
	InitApplication(appCfg)

	// 初始化数据库信息
	dbCfg := vi.Sub("database")
	InitDatabase(dbCfg)

	// 初始化日志信息
	logCfg := vi.Sub("logger")
	InitLog(logCfg)

	InitJwt(vi.Sub("jwt"))

	CasBinSetup()
}
