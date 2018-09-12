package main

import "github.com/jinzhu/gorm"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

type User struct {

	ID uint `gorm:"type:int;not null;auto_increment"`
	Name string `gorm:"type:varchar(32);not null;default:''"`
	Email string `gorm:"type:varchar(32);not null;default:''; unique_index:users_email_unique"`
	Password string `gorm:"type:varchar(32);not null;default:'';"`

}

func (User) TableName()  string {

	return "users"
}

func (user *User) beforeCreate(scope *gorm.Scope)  {

	scope.SetColumn("created_at",time.Now())
}

func main()  {

	orm ,err := gorm.Open("mysql","homestead:secret@tcp(192.168.10.10:3306)/bbs?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err !=nil {

		fmt.Printf("mysql connect fail %v\n",err)
	}

	defer orm.Close()
	//requestParm := map[string]interface{}{"name":"张三","email":"test@126.com","password":"403132804"}
	//
	////user := User{}
	//
	//for key,value := range  requestParm {
	//
	//	fmt.Printf("key is :%v,value is :%v\n",key,value)
	//}
	//
	//fmt.Printf(requestParm)
	////
	//user.Name = requestParm["name"]
	//user.Password = requestParm["password"]
	//user.Email = requestParm["email"]
	//
	//orm.Create(&user)
	//
	//user := User{}
	//
	//orm.First(&user,"李四")
	//
	////orm.First(&user,"name = ?","张三")
	//
	//orm.First(&user,"name = ?","张三")
	//fmt.Printf("the first user is :%v\n",user)
	//
	//orm.Model(&user).Update("name","李四")
	//
	//fmt.Printf("the update user is :%v\n",user)

	//users := []User{}

	//orm.Where(map[string]interface{}{"name":"李四"}).Find(&user)

	//orm.Where([]int{5,6,4}).Find(&users)

	//orm.Not("NAME",[]string{"李四2"}).Find(&users)
	//orm.Not(map[string]interface{}{"name":"李四2"}).Find(&users)
	//orm.Where("id",[]uint{5,6}).Find(&users)

	user := User{}
	orm.First(&user,5)

	//orm.Model(&user).Where("name = ?","111").Update("name","222")

	orm.Delete(&user)
	fmt.Printf("the user is :%v",user)

















}
