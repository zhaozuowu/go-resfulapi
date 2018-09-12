package main

import (
	"resfulapi/routes"
	"github.com/kataras/iris"
)

func main() {

	app := routes.InitRoutes()

	app.Run(iris.Addr(":8080"))
}
