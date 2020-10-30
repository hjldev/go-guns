package router

import (
	"github.com/gin-gonic/gin"
	"go-guns/router/admin"
)

func InitRouter(r *gin.Engine) {

	// 注册系统路由
	admin.InitAdminRouter(r)
}
