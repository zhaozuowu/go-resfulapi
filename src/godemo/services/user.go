package services

import (
	. "godemo/models"
	. "godemo/repositories"
)

type UserService struct {
	userRepository Repository
}

func (userService *UserService) GetUserList() ([]User, error) {


	return userService.userRepository.GetAll()

}
func (userService *UserService) GetUserInfoById(id int64) (User, error) {

	return userService.userRepository.GetUserInfoById(id)
}

func NewUserService() *UserService {

	userService := &UserService{}

	userService.userRepository = NewUserRepository()

	return userService
}
