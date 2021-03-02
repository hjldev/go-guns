package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/controller/system"
)

func InitNoRoleRouter(g *gin.RouterGroup) {
	g.GET("/userInfo", system.UserInfo)
}
