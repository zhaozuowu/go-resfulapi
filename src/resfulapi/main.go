package main

import (
	"resfulapi/routes"
	"github.com/kataras/iris"
	orm "resfulapi/databases"
)


func main()  {

	defer orm.Eloument.Close()
	app := routes.InitRoutes()

	app.Run(iris.Addr(":8080"))
}
