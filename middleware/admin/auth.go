package admin

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-guns/global"
	"go-guns/tools"
	"net/http"
)

func InitAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.GetHeader("Authorization")
		if len(bearer) < 7 {
			tools.R(c, nil, http.StatusUnauthorized, "no Authorization")
			c.Abort()
			return
		}
		token := bearer[7:]
		claims, err := tools.ParseToken(token)
		if err != nil {
			e := (err).(*jwt.ValidationError)
			tools.R(c, nil, http.StatusUnauthorized, e.Inner.Error())
			c.Abort()
			return
		}
		if claims != nil {
			c.Set(global.AUTH_CLAIMS, claims)
			c.Next()
		}
	}
}
