package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/controller/system"
	"go-guns/tools"
)

func InitNoLoginRouter(g *gin.RouterGroup) {
	g.GET("/ping", func(c *gin.Context) {
		tools.R(c)
	})

	g.GET("/captcha", system.GenerateCaptcha)
	g.POST("/login", system.Login)
}
