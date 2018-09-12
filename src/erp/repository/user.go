package repository

import . "erp/models"

type UserRepository struct {
	userModel *User
}

func NewUserRepository() Repository {

	userRep := &UserRepository{}

	userRep.userModel = NewUser()

	return userRep
}

func (userRepository *UserRepository) GetAll() ([]User, error) {

	return userRepository.userModel.GetAll()
}

func (userRepository *UserRepository) GetUserInfoById(userId int64) (User, error) {

	return userRepository.userModel.Find(userId)
}

func (userRepository *UserRepository) DeleteUserInfoById(userId int64) (User, error) {

	return userRepository.userModel.Delete(userId)
}
func (userRepository *UserRepository) RegisterUser(inputParams map[string][]string) (User, error) {

	userRepository.FormatUserModelParmas(inputParams)

	return userRepository.userModel.RegisterUser()
}

func (userRepository *UserRepository) FormatUserModelParmas(inputFormParms map[string][]string) {

	if inputFormParms["name"] != nil {

		userRepository.userModel.Name = inputFormParms["name"][0]

	}

	if inputFormParms["password"] != nil {

		userRepository.userModel.Password = inputFormParms["password"][0]
	}

	if inputFormParms["email"] != nil {

		userRepository.userModel.Email = inputFormParms["email"][0]
	}

}
