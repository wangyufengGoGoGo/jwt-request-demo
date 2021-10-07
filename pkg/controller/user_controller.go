package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	model2 "jwt-request-demo/pkg/model"
	repository2 "jwt-request-demo/pkg/repository"
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
	repository2.UserRepository
}

func (u userController) AddUser(ctx *gin.Context) {
	defer func() {
		if p := recover(); p != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": p.(error).Error()})
		}
	}()
	sysUser := &model2.SysUser{}

	if err := ctx.ShouldBindJSON(sysUser); err != nil {
		panic(fmt.Errorf("cannot parse struct"))
	}
	user, err := u.UserRepository.AddUser(sysUser)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func NewUserController(userRepository repository2.UserRepository) UserController {
	return userController{userRepository}
}
