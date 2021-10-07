package main

import (
	"jwt-request-demo/api"
	route2 "jwt-request-demo/init/route"
	db2 "jwt-request-demo/pkg/db"
	"jwt-request-demo/pkg/middleware"
)

func main() {
	mysqlConn, _ := db2.InitMysql()

	httpRouter := route2.InitRoute()
	api.BaseRegister(httpRouter)

	routerGroup := httpRouter.Group("/api", middleware.AuthorizeJWT())

	api.UserRegister(routerGroup, mysqlConn)

	if err := httpRouter.Run(":1234"); err != nil {
		panic(err)
	}
}
