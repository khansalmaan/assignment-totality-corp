package service

import (
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/model"
)

type IUserService interface {
	GetUserById(int32) model.User
	GetUserByIds([]int32) []model.User
	SearchUsers(SearchUsersRequest) []model.User
}

type SearchUsersRequest struct {
	Fname     string
	City      string
	Phone     int64
	MinHeight float64
	MaxHeight float64
	Married   *bool
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

func (us *UserService) SearchUsers(searchReq SearchUsersRequest) []model.User {
	// get all users
	users, _ := us.db.GetUsers()

	filteredUsers := make([]model.User, 0)

	// filter users based on search criteria
	for i := 0; i < len(users); i++ {
		if searchReq.Fname != "" && users[i].FName != searchReq.Fname {
			continue
		}
		if searchReq.City != "" && users[i].City != searchReq.City {
			continue
		}
		if searchReq.Phone != 0 && users[i].Phone != searchReq.Phone {
			continue
		}
		if searchReq.MinHeight != 0 && users[i].Height < searchReq.MinHeight {
			continue
		}
		if searchReq.MaxHeight != 0 && users[i].Height > searchReq.MaxHeight {
			continue
		}
		if searchReq.Married != nil && users[i].Married != *searchReq.Married {
			continue
		}

		filteredUsers = append(filteredUsers, users[i])
	}

	return filteredUsers
}
