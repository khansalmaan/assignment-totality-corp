package service

import (
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/model"
)

type IUserService interface {
	GetUserById(int) model.User
	GetUserByIds([]int) []model.User
}

type UserService struct {
	db database.IDatabase
}

// new user service
func NewUserService(db database.IDatabase) IUserService {
	return &UserService{db: db}
}

func (us *UserService) GetUserById(id int) model.User {
	user, _ := us.db.GetUser(id)
	return user
}

func (us *UserService) GetUserByIds(ids []int) []model.User {
	users, _ := us.db.GetUserList(ids)
	return users
}
