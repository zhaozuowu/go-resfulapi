package databases

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysql struct {
}

func (mysql *Mysql) InitConnect() *gorm.DB {

	db, err := gorm.Open("mysql", "homestead:secret@tcp(192.168.10.10)/bbs?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {

		fmt.Printf("链接数据库失败")

		os.Exit(500)
	}

	return db
}

func NewMsql() Db {

	return &Mysql{}
}
