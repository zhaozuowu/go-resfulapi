package models

import (
	orm "resfulapi/databases"
)

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

func (user *User) Users() (users []User, err error) {

	if err = orm.Eloument.Find(&users).Error; err != nil {
		return
	}

	return

}

func (user *User) Insert() (id uint, err error) {

	result := orm.Eloument.Create(&user)

	//fmt.Printf("result:is%v",result.Error)

	id = user.ID

	if result.Error != nil {

		err = result.Error
		return
	}

	return
}

func (user *User) Destory(id int64) (result User, err error) {

	//orm.Eloument.Where(map[string]interface{}{"id":id}).Delete(&user)

	if err = orm.Eloument.First(&user, id).Error; err != nil {

		return
	}

	if err = orm.Eloument.Delete(&user).Error; err != nil {

		return
	}

	result = *user

	return

}

func (user *User) Update(id int64) (result User, err error) {

	if err = orm.Eloument.Select([]string{"id", "name", "password", "email"}).First(&result, id).Error; err != nil {
		return
	}

	if err = orm.Eloument.Model(&result).Update(&user).Error; err != nil {
		return
	}

	return
}
func (user *User) GetUserInfoById(id int64) (userInfo User, err error) {

	if err = orm.Eloument.First(&user, id).Error; err != nil {

		return
	}

	userInfo = *user

	return

}
