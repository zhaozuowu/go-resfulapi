package databases

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

var Eloument *gorm.DB
func init() {

	db, err := gorm.Open("mysql", "homestead:secret@tcp(192.168.10.10)/bbs?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("connect mysql fail:%v\n", err)
	}

	Eloument = db
}
