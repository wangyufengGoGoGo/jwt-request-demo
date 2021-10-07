package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-request-demo/pkg/util"
	"net/http"
)

func Login(ctx *gin.Context) {
	token := util.GenerateToken("admin", "123456")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":  http.StatusOK,
		"token": token,
	})
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": http.StatusOK,
		"msg":  "Hello JWT",
	})
}
