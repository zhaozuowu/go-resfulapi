package repositories

import (
	. "goapp/models"
)

type Repositories interface {
	GetAll() (user []User, err error)
	//GetOne(id int) (user User, err error)
	//Insert() (id int, err error)
	//Update(id int) (updateUser User, err error)
	//Delete(id int) (deleteUser User, err error)
}

type UserRepository struct {
	userModel User
}

func (userRepository UserRepository) GetAll() (user []User, err error) {

	return userRepository.userModel.GetAll()
}

