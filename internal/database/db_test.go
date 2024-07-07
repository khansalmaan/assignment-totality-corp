package database

import (
	"testing"

	"assignment-totality-corp/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	db := NewDatabase()

	assert.Equal(t, 4, len(db.Users))
}

func TestGetUser(t *testing.T) {
	db := NewDatabase()

	user, err := db.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), user.ID)

	user, err = db.GetUser(99)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user)
}

func TestGetUserList(t *testing.T) {
	db := NewDatabase()

	users, err := db.GetUserList([]int32{1, 2, 99})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users))

	for _, user := range users {
		assert.Contains(t, []int32{1, 2}, user.ID)
	}
}

func TestAddUser(t *testing.T) {
	db := NewDatabase()

	newUser := model.User{
		ID:      5,
		FName:   "Eve",
		City:    "San Francisco",
		Phone:   9876543210,
		Height:  5.8,
		Married: false,
	}

	user, err := db.AddUser(newUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user)

	user, err = db.GetUser(5)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestRemoveUser(t *testing.T) {
	db := NewDatabase()

	user, err := db.RemoveUser(1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), user.ID)

	user, err = db.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user)

	user, err = db.RemoveUser(99)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user)
}

func TestGetUsers(t *testing.T) {
	db := NewDatabase()

	users, err := db.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(users))

	ids := make(map[int32]bool)
	for _, user := range users {
		ids[user.ID] = true
	}

	for id := range db.Users {
		assert.True(t, ids[id])
	}
}
