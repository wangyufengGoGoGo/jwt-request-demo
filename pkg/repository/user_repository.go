package repository

import (
	"gorm.io/gorm"
	model2 "jwt-request-demo/pkg/model"
	"log"
)

/**
 * @Author wyf
 * @Date 2021/7/8 18:57
 **/

type userRepository struct {
	*gorm.DB
}

type UserRepository interface {
	AddUser(user *model2.SysUser) (*model2.SysUser, error)
	Migrate() error
}

func NewRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db,
	}
}

func (u userRepository) AddUser(user *model2.SysUser) (*model2.SysUser, error) {
	return user, u.Create(user).Error
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.AutoMigrate(&model2.SysUser{})
}
