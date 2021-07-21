package router

import (
	"go-guns/middleware"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine) {
	g := r.Group("/")

	// 这里加载不需要登录认证的路由
	InitNoLoginRouter(g)

	// 加载登录认证中间件
	g.Use(middleware.InitAuth())
	InitNoRoleRouter(g)

	// 加载权限中间件
	g.Use(middleware.InitRBAC())
	InitRoleRouter(g)

}
