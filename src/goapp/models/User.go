package models

import "time"

type User struct {
	ID        uint64 `gorm:"not null;primary_key;auto_increment"`
	Name      string `gorm:"type:varchar(16);not null;default:'姓名'"`
	Email     string `gorm:"type:varchar(16);not null;default:'';comment:'邮箱'";unique_index:user_email_unique`
	Password  string `gorm:"type:varchar(60);not null;default:''comment:'密码'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time
}

func (user *User) GetAll() (userList []User, err error) {

	if err = Orm.Find(&userList).Error; err != nil {

		return
	}

	return
}

func (user *User) GetUserInfoById(id int) (userInfo User, err error) {

	if err = Orm.First(&user, id).Error; err != nil {

		return
	}

	userInfo = *user

	return
}
