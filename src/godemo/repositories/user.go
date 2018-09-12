package repositories

import (
	. "godemo/models"
)

type UserRepository struct {
	userModel *User
}

func NewUserRepository() Repository {

	userRepository := &UserRepository{}

	userRepository.userModel = NewModel()
	return userRepository
}

func (user *UserRepository) GetAll() ([]User, error) {

	return user.userModel.GetAll()
}

func (user *UserRepository) GetUserInfoById(id int64) (User, error) {

	return user.userModel.GetUserInfoById(id)
}
