package database

import "assignment-totality-corp/internal/model"

type IDatabase interface {
	GetUser(int) (model.User, error)
	GetUserList([]int) ([]model.User, error)
	AddUser(model.User) (model.User, error)
	RemoveUser(int) (model.User, error)
}

type Database struct {
	Users map[int]model.User
}

func NewDatabase() IDatabase {
	return &Database{Users: make(map[int]model.User)}
}

func (db *Database) GetUser(id int) (model.User, error) {
	user, ok := db.Users[id]
	if !ok {
		return model.User{}, nil
	}
	return user, nil
}

func (db *Database) GetUserList(ids []int) ([]model.User, error) {
	users := make([]model.User, 0)
	for _, id := range ids {
		user, ok := db.Users[id]
		if !ok {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *Database) AddUser(user model.User) (model.User, error) {
	db.Users[user.ID] = user
	return user, nil
}

func (db *Database) RemoveUser(id int) (model.User, error) {
	user, ok := db.Users[id]
	if !ok {
		return model.User{}, nil
	}
	delete(db.Users, id)
	return user, nil
}
