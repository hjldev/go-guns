package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/middleware"
)

func InitAdminMiddleware(r *gin.Engine) {
	// 公共的middleware
	middleware.InitMiddleware(r)
	// 记录到日志
	r.Use(middleware.LoggerToFile())
}
