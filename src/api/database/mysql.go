package database

import "github.com/jinzhu/gorm"
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Mysql struct {

	Db * gorm.DB
}

func (mysql *Mysql) Init() {

	db,err := gorm.Open("mysql","homestead:secret@tcp(192.168.10.10:3306)/bbs?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect fail %v\n",err)
	}

	mysql.Db = db

}