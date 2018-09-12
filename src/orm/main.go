package main

import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {

	Name string `gorm:"default:''"`
	Email string
	Password string
}

func (user User) TableName() string  {

	return "admin_users"
}
func main()  {

	 db,err := gorm.Open("mysql","homestead:secret@tcp(192.168.10.10:3306)/bbs?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	 defer db.Close()

	 if err !=nil {
	 	fmt.Printf("mysql connect fail %v\n",err)
	 }


	 //db.CreateTable(&User{})

	 //fmt.Printf("%v\n",user)
	 //user := User{Name:"zhangsan",Email:"gaofandi1236677@126.com",Password:"xxx"}
	 ////
	 //result := db.Create(&user)
	 ////
	 //if result.Error !=nil {
	 //	fmt.Printf("insert row errro %v",result.Error)
	 //}

		//getUser := User{}
	//
	////getUser.Name = "xxxx"
	////getUser.Password = "xxxx"
	////getUser.Email = "XXX"
	////db.Save(getUser)
	//
	//db.Select([]string{"name","password"}).First(&getUser)
	//fmt.Printf("getUser:%v\n",getUser)
	//
	//getUser.Email = "XXXXXXX"
	//getUser.Name = "xxxxxxxx"
	//getUser.Password = "XXXXXX"
	//db.Save(&getUser)

	user := []User{}

	db.Find(&user)

	fmt.Printf("user list:%v\n",user)

	getUser := User{}

	db.Select([]string{"name","password"}).First(&getUser)

	fmt.Printf("row list:%v\n",getUser)

	db.Delete(&getUser)



}