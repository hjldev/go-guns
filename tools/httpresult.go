package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// data第一个参数为data，第二个参数是status， 第三个参数是message
func R(ctx *gin.Context, data ...interface{}) {
	if len(data) == 0 {
		res := Response{
			Status:  200,
			Message: "ok",
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
	status := 200
	if len(data) > 1 {
		status = data[1].(int)
	}
	message := ""
	if len(data) > 2 {
		message = data[2].(string)
	}
	res := Response{
		Status:  status,
		Message: message,
		Data:    data[0],
	}
	ctx.JSON(http.StatusOK, res)
}
