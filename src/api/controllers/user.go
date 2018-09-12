package controllers

import (
	"github.com/kataras/iris"
	"api/models"
	db "api/database"
)

func Users(c iris.Context) {

	//username := c.Params().Get("username")

	//password := c.Params().Get("password")

	db := db.Mysql{}

	db.Init()

	users := []models.User{}
	db.Db.Find(&users)

	c.JSON(iris.Map{
		"code":0,
		"data":users,
	})

}

func Store(c iris.Context) {

}

func Update(c iris.Context) {

}

func Destory(c iris.Context) {

}
