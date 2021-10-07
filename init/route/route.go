package route

import (
	"github.com/gin-gonic/gin"
)

/**
 * @Author wyf
 * @Date 2021/7/8 11:07
 **/

func InitRoute() *gin.Engine {
	httpRouter := gin.New()
	gin.SetMode(gin.DebugMode)
	httpRouter.Use(
		gin.Recovery(),
		gin.LoggerWithWriter(gin.DefaultErrorWriter),
	)

	return httpRouter
}
