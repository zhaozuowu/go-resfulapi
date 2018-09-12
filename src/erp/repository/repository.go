package repository

import ."erp/models"

type Repository interface {

	GetAll()([]User,error)
	GetUserInfoById(userId int64) (User, error)
	RegisterUser(inputParams map[string][]string) (User, error)
	DeleteUserInfoById(userId int64) (User, error)
	

}
