package models

import "time"

type User struct {
	Model
	ID        uint64 `gorm:"not null;primary_key;auto_increment;comment:'主键';json:'id'"`
	Name      string `gorm:"type:varchar(32);not null;default:'';comment:'姓名';json:''name"`
	Email     string `gorm:"type:varchar(32);not null;default:'';comment:'邮箱';json:'email'"`
	Password  string `gorm:"type:varchar(60);not null; default:'';comment:'密码';json:'password'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (user *User) TableName() string {

	return "erp_users"
}

func (user *User) GetAll() (userList []User, err error) {

	if err = user.orm.Find(&userList).Error; err != nil {

		return
	}

	return
}

func (user *User) CreateTable() {

}
func (user *User) GetUserInfoById(id int64) (userInfo User, err error) {

	if err = user.orm.First(&user, id).Error; err != nil {

		return
	}

	userInfo = *user
	return
}

func NewModel() *User {

	model := &User{}
	model.InitConnect()
	return model
}
