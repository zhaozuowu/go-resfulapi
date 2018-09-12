package main

import (
	"github.com/kataras/iris"
	orm "api/curl/databases"
	. "api/curl/routes"
	_ "api/curl/databases"
)

func main() {

	defer orm.Eloument.Close()
	route := Route{}

	app := route.InitRoute()

	app.Run(iris.Addr("demo.go.cc:9890"))

}
