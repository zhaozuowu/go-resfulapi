package services

import (
	. "goapp/repositories"
	. "goapp/models"
)

type UserService struct {
	userRepositry UserRepository
}

func (userService *UserService) GetUserList() ([]User, error) {

	return userService.userRepositry.GetAll()
}


