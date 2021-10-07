package main

import (
	"jwt-request-demo/api"
	"jwt-request-demo/init/route"
	"jwt-request-demo/pkg/db"
	"jwt-request-demo/pkg/middleware"
)

func main() {
	mysqlConn, _ := db.InitMysql()

	httpRouter := route.InitRoute()
	api.BaseRegister(httpRouter)

	routerGroup := httpRouter.Group("/api", middleware.AuthorizeJWT())

	api.UserRegister(routerGroup, mysqlConn)

	if err := httpRouter.Run(":1234"); err != nil {
		panic(err)
	}
}
