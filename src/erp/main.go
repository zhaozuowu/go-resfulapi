package main

import (
	. "erp/routes"

	"github.com/kataras/iris"
)

func main() {

	routes := NewUserRoutes()

	app := routes.InitUserRoutes()

	app.Run(iris.Addr(":8080"))
}
