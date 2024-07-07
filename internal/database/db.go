package database

import "assignment-totality-corp/internal/model"

type IDatabase interface {
	GetUser(int32) (model.User, error)
	GetUserList([]int32) ([]model.User, error)
	GetUsers() ([]model.User, error)
	AddUser(model.User) (model.User, error)
	RemoveUser(int32) (model.User, error)
}

type Database struct {
	Users map[int32]model.User
}

func NewDatabase() Database {
	// popuplate the database with some dummy data
	users := make(map[int32]model.User)

	users[1] = model.User{
		ID:      1,
		FName:   "Alice",
		City:    "New York",
		Phone:   1234567890,
		Height:  5.5,
		Married: false,
	}

	users[2] = model.User{
		ID:      2,
		FName:   "Bob",
		City:    "Los Angeles",
		Phone:   1234567890,
		Height:  5.5,
		Married: true,
	}

	users[3] = model.User{
		ID:      3,
		FName:   "Charlie",
		City:    "Chicago",
		Phone:   1234567890,
		Height:  5.5,
		Married: false,
	}

	users[4] = model.User{
		ID:      4,
		FName:   "David",
		City:    "New York",
		Phone:   1234567890,
		Height:  5.5,
		Married: true,
	}

	return Database{Users: users}
}

func (db *Database) GetUser(id int32) (model.User, error) {
	user, ok := db.Users[id]
	if !ok {
		return model.User{}, nil
	}
	return user, nil
}

func (db *Database) GetUserList(ids []int32) ([]model.User, error) {
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

func (db *Database) RemoveUser(id int32) (model.User, error) {
	user, ok := db.Users[id]
	if !ok {
		return model.User{}, nil
	}
	delete(db.Users, id)
	return user, nil
}

func (db *Database) GetUsers() ([]model.User, error) {
	users := make([]model.User, 0)
	for _, user := range db.Users {
		users = append(users, user)
	}
	return users, nil
}
