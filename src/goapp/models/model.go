package models

import (
	"github.com/jinzhu/gorm"
	. "goapp/databases"
)

var Orm *gorm.DB

var dirver = "mysql"

func init() {

	dbFactory := DbFactory{}

	dbDriver := dbFactory.GetConnect(dirver)

	connect := dbDriver.InitConnect()


	Orm = connect


}
