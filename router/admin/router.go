package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/middleware/admin"
)

func InitAdminRouter(r *gin.Engine) {
	g := r.Group("/")

	// 这里加载不需要登录认证的路由
	InitNoLoginRouter(g)

	// 加载登录认证中间件
	g.Use(admin.InitAuth())
	InitNoRoleRouter(g)

	// 加载权限中间件
	g.Use(admin.InitRBAC())
	InitRoleRouter(g)

}
