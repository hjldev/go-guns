package user

import (
	"github.com/gin-gonic/gin"
	"go-guns/global"
	"go-guns/tools"
)

func Info(c *gin.Context) {
	claims, exist := c.Get(global.AUTH_CLAIMS)
	if exist {
		user := (claims).(*tools.JwtAuth)
		tools.R(c, user)
	}
}
