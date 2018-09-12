package main

import (
	"github.com/kataras/iris"
	. "godemo/routes"
)

func main() {

	routes := NewUserRoutes()
	//userRoutes := routes.UserRoutes{}
	//app := userRoutes.InitRoute()

	app := routes.InitRoutes()
	app.Run(iris.Addr(":8080"))
}


