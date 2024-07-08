package service

import (
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/model"
)

type IUserService interface {
	GetUserById(int32) (model.User, error)
	GetUserByIds([]int32) ([]model.User, error)
	SearchUsers(SearchUsersRequest) ([]model.User, error)
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

func (us *UserService) GetUserById(id int32) (model.User, error) {
	user, err := us.db.GetUser(id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (us *UserService) GetUserByIds(ids []int32) ([]model.User, error) {
	users, err := us.db.GetUserList(ids)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (us *UserService) SearchUsers(searchReq SearchUsersRequest) ([]model.User, error) {
	// if not search criteria provided, return empty list
	if searchReq.Fname == "" && searchReq.City == "" && searchReq.Phone == 0 && searchReq.MinHeight == 0 && searchReq.MaxHeight == 0 && searchReq.Married == nil {
		return []model.User{}, nil
	}

	// get all users
	users, err := us.db.GetUsers()
	if err != nil {
		return nil, err
	}

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

	return filteredUsers, nil
}
