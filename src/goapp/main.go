package main

import (
	"github.com/kataras/iris"
	. "goapp/routes"
)

func main() {

	//var re repositories.Repositories
	//
	//user := &repositories.UserRepository{}//针对接口类型的引用，这里是要&
	//
	//re = user
	//
	//fmt.Printf("re:%v\n",re)

	route := UserRoutes{}
	app := route.InitRoute()


	app.Run(iris.Addr(":8080"))
}
