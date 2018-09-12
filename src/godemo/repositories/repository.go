package repositories

import . "godemo/models"

type Repository interface {
	GetAll() ([]User, error)
	GetUserInfoById(userId int64) (User, error)
}
