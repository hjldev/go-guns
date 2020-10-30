package middleware

import (
	"github.com/gin-gonic/gin"
	"go-guns/tools"
	"net/http"
)

func NoFound(c *gin.Context) {
	tools.R(c, nil, http.StatusNotFound, "page not found")
}
