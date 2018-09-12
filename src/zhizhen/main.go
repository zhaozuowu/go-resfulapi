package main

import (
	"fmt"
	"github.com/kataras/iris/sessions"
)

type User struct {
	ID    uint64
	Name  string
	Email string
}

func (user *User) SetName(name string) {

	user.Name = name
}

type UserController struct {
	UserModel User
	Session sessions.Session
}

func (user *UserController) SetName(name string)  {

	user.UserModel.SetName(name)


	//vists := user.Session.Increment("visit",1)


	fmt.Printf("USER NAME IS :%v,vists:%v\n",user.UserModel.Name,vists)


}
func main() {
     user := UserController{}

     user.SetName("test")

}
