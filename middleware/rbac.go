package middleware

import (
	"github.com/gin-gonic/gin"
	"go-guns/boot"
	"go-guns/global"
	"go-guns/tools"
	"net/http"
)

func InitRBAC() gin.HandlerFunc {
	return func(c *gin.Context) {

		if claims, ok := c.Get(global.AUTH_CLAIMS); ok {
			e, err := boot.CasBin()
			tools.HasError(err, 500)
			res, err := e.Enforce((claims).(*tools.JwtAuth).Role, c.Request.URL.Path, c.Request.Method)

			if res || (claims).(*tools.JwtAuth).Role == "admin" {
				c.Next()
				return
			}
		}
		tools.R(c, nil, http.StatusForbidden, "对不起，您没有该接口访问权限，请联系管理员")
		c.Abort()
	}
}
