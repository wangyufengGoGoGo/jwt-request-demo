package model

import (
	"gorm.io/gorm"
)

/**
 * @Author wyf
 * @Date 2021/7/8 18:36
 **/

type SysUser struct {
	gorm.Model
	Username string	`json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (SysUser) TableName() string {
	return "sys_users"
}

