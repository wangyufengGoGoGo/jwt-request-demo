package main

import (
	"jwt-request-demo/db"
	"jwt-request-demo/route"
)

/**
 * @Author wyf
 * @Date 2021/7/8 11:05
 **/

func main() {
	mysqlConn, _ := db.InitMysql()

	route.InitRoute(mysqlConn)
}
