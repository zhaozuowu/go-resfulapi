package models

import (
	"github.com/jinzhu/gorm"
	. "godemo/databases"
)

type Model struct {
	orm *gorm.DB
}

func (model *Model) InitConnect() {

	db := NewMsql()

	model.orm = db.InitConnect()

}
