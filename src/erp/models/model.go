package models

import (
	"github.com/jinzhu/gorm"
	. "erp/databases"
)

type Model struct {
	orm *gorm.DB
}


func (model *Model) InitConnect() {

	dbFactory := DbFactory{}

	db := dbFactory.GetDbDriver()

	model.orm = db.InitConnect()

}
