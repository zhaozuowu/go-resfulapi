package controllers

import (
	"github.com/kataras/iris"
	. "goapp/services"
	_ "goapp/services"
	"fmt"
)

type UserController struct {
	userService UserService
}

func (user *UserController) Index(c iris.Context) {

	response := map[string]interface{}{"code": 1, "message": "", "data": ""}
	userList, err := user.userService.GetUserList()

	fmt.Printf("userlist:%v,err:%v",userList,err)

	response["data"] = userList

	if err != nil {

		response["code"] = -1
		response["messsage"] = err
		response["data"] = ""
	}

	c.JSON(iris.Map(response))
}
