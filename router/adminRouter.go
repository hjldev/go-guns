package router

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/controller/system"
	"go-guns/middleware"
	"go-guns/tools"
)

func InitAdminRouter(r *gin.Engine) {
	g := r.Group("/")

	// 这里加载不需要登录认证的路由
	g.GET("/ping", func(c *gin.Context) {
		tools.R(c)
	})
	g.GET("/captcha", system.GenerateCaptcha)
	g.POST("/login", system.Login)

	// 加载登录认证中间件
	g.Use(middleware.InitAuth())
	g.GET("/userInfo", system.UserInfo)

	// 加载权限中间件
	g.Use(middleware.InitRBAC())

}
