package router

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/admin/system/user"
	"go-guns/app/controller/admin"
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
	g.POST("/login", user.Login)

	// 加载登录认证中间件
	g.Use(middleware.InitAuth())
	g.GET("/userInfo", user.Info)

	// 加载权限中间件
	g.Use(middleware.InitRBAC())

}
