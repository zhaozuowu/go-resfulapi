package controllers

import (
	"erp/services"
	"github.com/kataras/iris"
)

type UserController struct {
	userSerive *services.UserSerive
}

func (userController *UserController) Index(c iris.Context) {

	userList, err := userController.userSerive.GetUserList()

	c.JSON(iris.Map{
		"code":    1,
		"message": err,
		"data":    userList,
	})
}

func (userController *UserController) Show(c iris.Context) {

	userId, err := c.Params().GetInt64("id")

	userInfo, err := userController.userSerive.GetUserInfoById(userId)

	c.JSON(iris.Map{
		"code":    1,
		"message": err,
		"data":    userInfo,
	})

}

func (userController *UserController) Store(c iris.Context) {

	userCreateInput := c.FormValues()

	userInfo, err := userController.userSerive.RegisterUser(userCreateInput)

	c.JSON(iris.Map{
		"code":    1,
		"message": err,
		"data":    userInfo,
	})

}

func (userController *UserController) Destory(c iris.Context) {

	userId, err := c.Params().GetInt64("id")

	userInfo, err := userController.userSerive.DeleteUserInfoById(userId)

	c.JSON(iris.Map{
		"code":    1,
		"message": err,
		"data":    userInfo,
	})

}

func NewUserController() *UserController {

	userController := &UserController{}
	userController.userSerive = services.NewUserSerive()

	return userController
}
