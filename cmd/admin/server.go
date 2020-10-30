package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-guns/casbin"
	"go-guns/config"
	"go-guns/database"
	"go-guns/global"
	"go-guns/logger"
	admin2 "go-guns/middleware/admin"
	"go-guns/router/admin"
	"os"
	"os/signal"
)

var (
	StartCmd = &cobra.Command{
		Use:          "admin",
		Short:        "Start API server",
		Example:      "go-guns server",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

var echoTimes int

func init() {
	// 全局环境设置
	StartCmd.PersistentFlags().StringVarP(&global.Env, "env", "e", "dev", "server env ; eg:dev,test,prod")
}

func setup() {
	//1. 读取配置
	config.SetupEnv("admin", global.Env)
	//2. 设置日志
	logger.Setup()
	//3. 初始化数据库链接
	database.Setup()
	//4. 接口访问控制加载
	casbin.Setup()
}

func run() error {
	logger.Info("guns run server")
	if global.Env == global.PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 加载中间件
	admin2.InitAdminMiddleware(r)
	// 注册系统路由
	admin.InitAdminRouter(r)

	// 启动
	err := r.Run(":" + config.AppCfg.Port)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s \n\n", `欢迎使用admin后台管理系统`)
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	return nil
}
