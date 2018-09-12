package routes

import (
	"erp/controllers"
	"github.com/kataras/iris"
)

type UserRoutes struct {
	userController *controllers.UserController
}

func (userRoutes *UserRoutes) InitUserRoutes() *iris.Application {

	app := iris.New()

	app.Get("/users", userRoutes.userController.Index)
	app.Get("/users/{id:int}", userRoutes.userController.Show)
	app.Post("/users", userRoutes.userController.Store)
	app.Delete("/users/{id:int}", userRoutes.userController.Destory)

	return app

}

func NewUserRoutes() *UserRoutes {

	routes := &UserRoutes{}

	routes.userController = controllers.NewUserController()

	return routes
}
