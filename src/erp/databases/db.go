package databases

import "github.com/jinzhu/gorm"

type Db interface {

	InitConnect() *gorm.DB
}
