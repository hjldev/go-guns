package router

import (
	"go-guns/app/controller/system"
	"go-guns/tools"

	"github.com/gin-gonic/gin"
)

func InitNoLoginRouter(g *gin.RouterGroup) {
	g.GET("/ping", func(c *gin.Context) {
		tools.R(c)
	})

	g.GET("/captcha", system.GenerateCaptcha)
	g.POST("/login", system.Login)
}
