package repository

import (
	"gorm.io/gorm"
	"jwt-request-demo/model"
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
	AddUser(user *model.SysUser) (*model.SysUser, error)
	Migrate() error
}

func NewRepository(db *gorm.DB) UserRepository {
	return userRepository{
		db,
	}
}

func (u userRepository) AddUser(user *model.SysUser) (*model.SysUser, error) {
	return user, u.Create(user).Error
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]...Migrate")
	return u.AutoMigrate(&model.SysUser{})
}