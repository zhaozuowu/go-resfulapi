package routes

import "github.com/kataras/iris"

import (
	. "api/curl/controllers"
	"github.com/kataras/iris/mvc"
	."api/controllers"
)

type Route struct {
	userController UserController
}

func (route *Route) InitRoute() *iris.Application {

	app := iris.New()
	app.Get("/users", route.userController.Index)
	app.Post("/users/", route.userController.Store)
	app.Delete("/users/{id:int}", route.userController.Destory)
	app.Put("/users/{id:int}", route.userController.Update)
	app.Get("/users/{id:int}", route.userController.Show)
	mvc.Configure(app.Party("/articles"), func(app *mvc.Application) {
		app.Handle(new(Article))

	})
	return app
}
