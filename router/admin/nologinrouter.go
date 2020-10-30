package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/apis/admin/system/user"
	"go-guns/tools"
)

func InitNoLoginRouter(g *gin.RouterGroup) {
	g.GET("/ping", func(c *gin.Context) {
		tools.R(c)
	})

	g.GET("/login", user.Login)
}
