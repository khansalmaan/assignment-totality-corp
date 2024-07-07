package service

import (
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/model"
)

type IUserService interface {
	GetUserById(int32) model.User
	GetUserByIds([]int32) []model.User
	SearchUsers(string, string, int64, float64, float64, bool) []model.User
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

func (us *UserService) SearchUsers(fname, city string, phone int64, minHeight, maxHeight float64, married bool) []model.User {
	// get all users
	users, _ := us.db.GetUsers()

	filteredUsers := make([]model.User, 0)

	// filter users based on search criteria
	for i := 0; i < len(users); i++ {
		if fname != "" && users[i].FName != fname {
			continue
		}
		if city != "" && users[i].City != city {
			continue
		}
		if phone != 0 && users[i].Phone != phone {
			continue
		}
		if minHeight != 0 && users[i].Height < minHeight {
			continue
		}
		if maxHeight != 0 && users[i].Height > maxHeight {
			continue
		}
		if married && !users[i].Married {
			continue
		}

		filteredUsers = append(filteredUsers, users[i])
	}

	return filteredUsers
}
