package routes

import (
	"github.com/kataras/iris"
	. "goapp/controllers"
)

type UserRoutes struct {
	userController UserController

}

func (user UserRoutes) initRoute() {

	app := iris.New()

	app.Get("/users",)
}
