package controllers

import (
	"github.com/kataras/iris"
	. "godemo/services"
	"strconv"
)

type UserController struct {
	userSerive *UserService
}

func (userController *UserController) Index(c iris.Context) {

	userList, err := userController.userSerive.GetUserList()


	response := map[string]interface{}{"code": 1, "message": "", "data": ""}

	response["data"] = userList

	if err != nil {

		response["code"] = -1
		response["message"] = err
		response["data"] = ""
	}

	c.JSON(iris.Map(response))
}

func (userController *UserController) Show(c iris.Context) {

	userId, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)
	userInfo, err := userController.userSerive.GetUserInfoById(userId)
	response := map[string]interface{}{"code": 1, "message": "", "data": ""}

	response["data"] = userInfo

	if err != nil {

		response["code"] = -1
		response["message"] = err
		response["data"] = ""
	}

	c.JSON(iris.Map(response))
}

func NewUserController() *UserController {

	user := &UserController{}

	user.userSerive = NewUserService()

	return user
}
