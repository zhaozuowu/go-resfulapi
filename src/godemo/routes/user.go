package routes

import (
	. "godemo/controllers"
	"github.com/kataras/iris"
)

type UserRoute struct {
	userController *UserController
}

func (userRoutes *UserRoute) InitRoutes() *iris.Application {

	app := iris.New()
	app.Get("/users", userRoutes.userController.Index)
	app.Get("/users/{id:int}", userRoutes.userController.Show)

	return app

}
func NewUserRoutes() *UserRoute {

	routes := &UserRoute{}

	routes.userController = NewUserController()
	return routes

}
