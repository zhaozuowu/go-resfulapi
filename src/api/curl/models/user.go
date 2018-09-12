package models

import (
	orm "api/curl/databases"
)

type User struct {
	ID       uint64 `gorm:"not null;auto_increment;primary_key"`
	Name     string `gorm:"type:varchar(32);not null;default:'test';"`
	Email    string `gorm:"type:varchar(32);not null;default:'test';unique_index:user_email_unique"`
	Password string `gorm:"type:varchar(60);not null;default:'test';"`
}

func (user *User) GetUserList() (userList []User, err error) {

	if err = orm.Eloument.Where(user).Find(&userList).Error; err != nil {

		return
	}

	return
}

func (user *User) Insert() (userId uint64, err error) {

	if err = orm.Eloument.Create(&user).Error; err != nil {
		return
	}

	userId = user.ID

	return
}

func (user *User) Delete(userId int64) (err error) {

	if err = orm.Eloument.First(&user, userId).Error; err != nil {

		return
	}

	if err = orm.Eloument.Delete(&user).Error; err != nil {

		return
	}
	return
}

func (user *User) Update(userId int64) (updatedUserData User, err error) {

	if err = orm.Eloument.First(&updatedUserData, userId).Error; err != nil {

		return
	}

	updateData := map[string]interface{}{"name": user.Name, "email": user.Email, "password": user.Password}

	if err = orm.Eloument.Model(&updatedUserData).Update(updateData).Error; err != nil {

		return
	}

	return
}

func (user *User) GetUserInfoById(userId int64) (userInfo User, err error) {

	if err = orm.Eloument.First(&user, userId).Error; err != nil {
		return
	}

	userInfo = *user

	return
}
