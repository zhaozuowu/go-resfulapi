package databases

import "github.com/jinzhu/gorm"
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Mysql struct {
}

func NewMysql() Db {

	return &Mysql{}
}

func (mysql *Mysql) InitConnect() *gorm.DB {

	if err := godotenv.Load(".env"); err != nil {

		fmt.Printf("数据库配置加载文件有误\n")

		os.Exit(1)
	}

	db, err := gorm.Open("mysql", "homestead:secret@tcp(192.168.10.10)/bbs?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {

		fmt.Printf("链接数据库失败:%v\n", err)
		os.Exit(500)
	}

	return db
}
