package controllers

import "github.com/kataras/iris"

import (
	. "api/curl/models"
	"strconv"
)

type UserController struct {
	userModel *User
}

func (user *UserController) Index(c iris.Context) {

	user.userModel.Name = c.Request().FormValue("name")
	user.userModel.Email = c.Request().FormValue("email")
	userList, err := user.userModel.GetUserList()

	if err != nil {

		c.JSON(iris.Map{
			"code":    -1,
			"message": err,
			"data":    userList,
		})
		return
	}

	c.JSON(iris.Map{
		"code":    1,
		"message": "",
		"data":    userList,
	})
}

func (user *UserController) Store(c iris.Context) {

	user.userModel.Email = c.Request().FormValue("email")
	user.userModel.Name = c.Request().FormValue("name")
	user.userModel.Password = c.Request().FormValue("password")
	userId, err := user.userModel.Insert()

	result := map[string]interface{}{"code": 1, "message": "", "data": ""}

	if err != nil {
		result["code"] = -1
		result["message"] = err
	}

	result["data"] = userId

	c.JSON(iris.Map(result))

}

func (user *UserController) Destory(c iris.Context) {

	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)

	//fmt.Printf("before userModel:%v\n", user.userModel)
	err = user.userModel.Delete(id)

	//fmt.Printf("after userModel:%v\n", user.userModel)

	result := map[string]interface{}{"code": 1, "message": "", "data": ""}

	if err != nil {
		result["code"] = -1
		result["message"] = err

	}
	result["data"] = user.userModel

	c.JSON(iris.Map(result))

}

func (user *UserController) Update(c iris.Context) {

	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)
	user.userModel.Name = c.Request().FormValue("username")
	user.userModel.Email = c.Request().FormValue("email")
	user.userModel.Password = c.Request().FormValue("password")
	updatedUserData, err := user.userModel.Update(id)

	result := map[string]interface{}{"code": 1, "message": "", "data": ""}

	if err != nil {
		result["code"] = -1
		result["message"] = err

	}
	result["data"] = updatedUserData

	c.JSON(iris.Map(result))

}

func (user *UserController) Show(c iris.Context) {

	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)

	userInfo, err := user.userModel.GetUserInfoById(id)

	result := map[string]interface{}{"code": 1, "message": "", "data": ""}

	if err != nil {
		result["code"] = -1
		result["message"] = err

	}
	result["data"] = userInfo

	c.JSON(iris.Map(result))

}
