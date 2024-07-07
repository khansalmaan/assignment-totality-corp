package service

import (
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/model"
)

type IUserService interface {
	GetUserById(int32) model.User
	GetUserByIds([]int32) []model.User
}

type UserService struct {
	db database.IDatabase
}

// new user service
func NewUserService(db database.IDatabase) IUserService {
	return &UserService{db: db}
}

func (us *UserService) GetUserById(id int32) model.User {
	user, _ := us.db.GetUser(id)
	return user
}

func (us *UserService) GetUserByIds(ids []int32) []model.User {
	users, _ := us.db.GetUserList(ids)
	return users
}
