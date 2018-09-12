package services

import (
	"erp/repository"
	"erp/models"
)

type UserSerive struct {
	userRepository repository.Repository
}

func NewUserSerive() *UserSerive {

	userSerive := &UserSerive{}

	userSerive.userRepository = repository.NewUserRepository()

	return userSerive

}

func (userSerivce *UserSerive) GetUserList() ([]models.User, error) {

	return userSerivce.userRepository.GetAll()
}
func (userSerivce *UserSerive) GetUserInfoById(userId int64) (models.User, error) {

	return userSerivce.userRepository.GetUserInfoById(userId)
}

func (userSerive *UserSerive) RegisterUser(inputFormParms map[string][]string) (models.User, error) {

	//userInfo := userSerive.GetRegisterUserParmas(inputFormParms)

	return userSerive.userRepository.RegisterUser(inputFormParms)
}
func (userSerive *UserSerive) DeleteUserInfoById(userId int64) (models.User, error) {

	return userSerive.userRepository.DeleteUserInfoById(userId)
}
