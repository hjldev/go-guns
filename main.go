package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-guns/boot"
	"go-guns/global"
	"go-guns/middleware"
	"go-guns/router/admin"
	"log"
	"os"
	"os/signal"
)

var serverName string
var env string

func init() {
	flag.StringVar(&serverName, "name", "admin", "admin service")
	flag.StringVar(&env, "env", "dev", "dev")
}

func main() {
	flag.Parse()

	switch serverName {
	case "admin":
		adminServer()
	default:
		fmt.Println("no service")
	}
}

func adminServer() {
	// 读取配置并初始化
	boot.AdminBoot("admin", "dev")

	log.Println("guns run server")
	if global.Env == global.PROD {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	// 加载中间件
	middleware.InitAdminMiddleware(r)
	// 注册系统路由
	admin.InitAdminRouter(r)

	// 启动
	err := r.Run(":" + boot.AppCfg.Port)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s \n\n", `欢迎使用admin后台管理系统`)
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
