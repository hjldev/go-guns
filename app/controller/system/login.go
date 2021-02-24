package system

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
	"go-guns/app/model"
	"go-guns/tools"
)

func GenerateCaptcha(c *gin.Context) {
	var store = base64Captcha.DefaultMemStore
	e := model.ConfigJsonBody{}
	e.Id = uuid.New().String()
	e.DriverDigit = base64Captcha.DefaultDriverDigit
	driver := e.DriverDigit
	cap := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _ := cap.Generate()
	captcha := &model.Captcha{
		Id:   id,
		B64s: b64s,
	}
	tools.R(c, captcha)
}

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
