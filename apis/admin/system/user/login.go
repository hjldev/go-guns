package user

import (
	"github.com/gin-gonic/gin"
	"go-guns/models"
	"go-guns/tools"
)

func Login(ctx *gin.Context) {
	token, err := tools.GenerateToken(tools.JwtAuth{
		Username: "test",
		Role:     "admin",
	})
	tools.HasError(err)
	tools.R(ctx, models.Token{
		Token: token,
	})
}
