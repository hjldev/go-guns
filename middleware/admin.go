package middleware

import (
	"github.com/gin-gonic/gin"
)

func InitAdminMiddleware(r *gin.Engine) {
	// 公共的middleware
	InitMiddleware(r)
	// 记录到日志
	r.Use(LoggerToFile())
}
