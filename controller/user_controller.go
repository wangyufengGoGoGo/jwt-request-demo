package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jwt-request-demo/model"
	"jwt-request-demo/repository"
	"net/http"
)

/**
 * @Author wyf
 * @Date 2021/7/8 19:16
 **/

type UserController interface {
	AddUser(*gin.Context)
}

type userController struct {
	repository.UserRepository
}

func (u userController) AddUser(ctx *gin.Context) {
	defer func() {
		if p := recover(); p != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": p.(error).Error()})
		}
	}()
	sysUser := &model.SysUser{}

	if err := ctx.ShouldBindJSON(sysUser); err != nil {
		panic(fmt.Errorf( "cannot parse struct"))
	}
	user, err := u.UserRepository.AddUser(sysUser)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data":user})
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return userController{userRepository}
}
