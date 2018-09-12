package routes

import (
	"github.com/kataras/iris"
	. "resfulapi/controllers"
)

func InitRoutes() *iris.Application {

	app := iris.New()

	user := User{}
	app.Get("/users", user.Index)
	app.Post("/users", user.Store)
	app.Put("/users/{id:int}", user.Update)
	app.Delete("/users/{id:int}", user.Destory)
	app.Get("/users/{id:int}",user.Show)
	return app

}
