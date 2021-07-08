package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jwt-request-demo/controller"
	"jwt-request-demo/middleware"
	"jwt-request-demo/repository"
	"jwt-request-demo/util"
	"log"
	"net/http"
)

/**
 * @Author wyf
 * @Date 2021/7/8 11:07
 **/

func InitRoute(db *gorm.DB) {
	httpRouter := gin.New()
	gin.SetMode(gin.DebugMode)
	httpRouter.Use(gin.Recovery())
	httpRouter.GET("/token", func(ctx *gin.Context) {
		token := util.GenerateToken("admin", "123456")
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":  http.StatusOK,
			"token": token,
		})
	})
	routerGroup := httpRouter.Group("/api", middleware.AuthorizeJWT())

	userRepository := repository.NewRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	userController := controller.NewUserController(userRepository)
	routerGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"msg":  "Hello JWT",
		})
	})

	routerGroup.POST("/user", userController.AddUser)

	if err := httpRouter.Run(":1234"); err != nil {
		panic(err)
	}
}
