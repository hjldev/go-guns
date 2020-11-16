package user

import (
	"github.com/gin-gonic/gin"
	"go-guns/model"
	"go-guns/tools"
)

func Login(ctx *gin.Context) {
	token, err := tools.GenerateToken(tools.JwtAuth{
		UserId: 1,
		Role:   "admin",
	})
	tools.HasError(err)
	tools.R(ctx, model.Token{
		Token: token,
	})
}
