package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jwt-request-demo/pkg/controller"
	"jwt-request-demo/pkg/repository"
	"log"
)

func BaseRegister(httpRouter *gin.Engine) {
	httpRouter.GET("/token", controller.Login)
}

func UserRegister(routerGroup *gin.RouterGroup, db *gorm.DB) {
	userRepository := repository.NewRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	userController := controller.NewUserController(userRepository)

	routerGroup.GET("/ping", controller.Ping)
	routerGroup.POST("/user", userController.AddUser)
}
