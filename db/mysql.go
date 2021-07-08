package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @Author wyf
 * @Date 2021/7/8 18:38
 **/

func InitMysql() (*gorm.DB, error) {

	USER := "root"
	PASS := "123456"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "test"

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)

	return gorm.Open(mysql.Open(url))
}
