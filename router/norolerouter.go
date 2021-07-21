package router

import (
	"go-guns/app/controller/system"

	"github.com/gin-gonic/gin"
)

func InitNoRoleRouter(g *gin.RouterGroup) {
	g.GET("/userInfo", system.UserInfo)
}
