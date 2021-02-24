package admin

import (
	"github.com/gin-gonic/gin"
	"go-guns/app/admin/system/user"
)

func InitNoRoleRouter(g *gin.RouterGroup) {
	g.GET("/userInfo", user.Info)
}
