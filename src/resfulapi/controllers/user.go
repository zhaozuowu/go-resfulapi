package controllers

import (
	"github.com/kataras/iris"
	"resfulapi/models"
	"fmt"
	"strconv"
)

type User struct {
	orm models.User
}

func (user *User) Index(c iris.Context) {

	result, err := user.orm.Users()

	if err != nil {
		c.JSON(iris.Map{
			"code":    -1,
			"message": "查询数据有误",
		})
	}

	c.JSON(iris.Map{
		"code": 200,
		"data": result,
	})
}

func (user *User) Store(c iris.Context) {

	//fmt.Printf("this key :%v",c.Request().FormValue("name"));
	user.orm.Name = c.Request().FormValue("name")
	user.orm.Password = c.Request().FormValue("password")
	user.orm.Email = c.Request().FormValue("email")

	id, err := user.orm.Insert()

	fmt.Printf("FAIL id:%v,err:%v", id, err)

	if err != nil {

		c.JSON(iris.Map{
			"code":    -1,
			"message": "添加失败",
		})
		return
	}

	c.JSON(iris.Map{
		"code": 1,
		"data": id,
	})
}

func (user *User) Destory(c iris.Context) {

	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)
	result, err := user.orm.Destory(id)

	if err != nil {

		c.JSON(iris.Map{
			"code":    -1,
			"message": "删除失败",
		})
		return
	}

	c.JSON(iris.Map{
		"code": 1,
		"data": result,
	})
}

func (user *User) Update(c iris.Context) {

	id, err := strconv.ParseInt(c.Params().Get("id"), 10, 64)
	user.orm.Name = c.Request().FormValue("name")
	user.orm.Password = c.Request().FormValue("password")
	user.orm.Email = c.Request().FormValue("email")
	result, err := user.orm.Update(id)
	if err != nil {

		c.JSON(iris.Map{
			"code":    -1,
			"message": "更新失败",
			"result":  result,
			"user":    user.orm,
		})
		return
	}

	c.JSON(iris.Map{
		"code":    1,
		"message": "更新成功",
		"result":  result,
		"user":    user.orm,
	})
}

func (user *User) Show(c iris.Context) {

	id, err := strconv.ParseInt(c.Params().Get("id"),10,64)

	userInfo, err := user.orm.GetUserInfoById(id)

	if err != nil {

		c.JSON(iris.Map{
			"code":    -1,
			"message": "查询失败",
		})
		return
	}

	c.JSON(iris.Map{
		"code": 1,
		"data": userInfo,
	})

}
