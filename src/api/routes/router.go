package routes

import "github.com/kataras/iris"
import . "api/controllers"

func InitRoutes()   *iris.Application  {

	app := iris.New()

	app.Get("/users", Users)
	app.Post("/users", Store)
	app.Put("/users/{id:int}", Update)
	app.Delete("/user/{id:int}", Destory)

	return app
}