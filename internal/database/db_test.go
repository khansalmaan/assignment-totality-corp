package database

import (
	"testing"

	"assignment-totality-corp/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase(t *testing.T) {
	db := NewDatabase()
	assert.Equal(t, 4, len(db.Users), "Expected 4 users in the database")
}

func TestGetUser(t *testing.T) {
	db := NewDatabase()

	// Test retrieving an existing user
	user, err := db.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), user.ID, "Expected user ID to be 1")

	// Test retrieving a non-existing user
	user, err = db.GetUser(99)
	assert.Error(t, err, "Expected an error for non-existing user")
	assert.Equal(t, model.User{}, user, "Expected empty user for non-existing user")
}

func TestGetUserList(t *testing.T) {
	db := NewDatabase()

	// Test retrieving a list of users
	users, err := db.GetUserList([]int32{1, 2, 99})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(users), "Expected 2 users in the result")

	// Verify that the correct users are returned
	expectedIDs := map[int32]bool{1: true, 2: true}
	for _, user := range users {
		assert.True(t, expectedIDs[user.ID], "Expected user ID to be 1 or 2")
	}
}

func TestAddUser(t *testing.T) {
	db := NewDatabase()

	// Create a new user
	newUser := model.User{
		ID:      5,
		FName:   "Eve",
		City:    "San Francisco",
		Phone:   9876543210,
		Height:  5.8,
		Married: false,
	}

	// Add the new user
	user, err := db.AddUser(newUser)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user, "Expected the added user to be returned")

	// Retrieve the newly added user
	user, err = db.GetUser(5)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user, "Expected the retrieved user to match the added user")
}

func TestRemoveUser(t *testing.T) {
	db := NewDatabase()

	// Remove an existing user
	user, err := db.RemoveUser(1)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), user.ID, "Expected removed user ID to be 1")

	// Verify that the user is removed
	user, err = db.GetUser(1)
	assert.Error(t, err, "Expected an error for non-existing user")
	assert.Equal(t, model.User{}, user, "Expected empty user for non-existing user")

	// Remove a non-existing user
	user, err = db.RemoveUser(99)
	assert.NoError(t, err)
	assert.Equal(t, model.User{}, user, "Expected empty user for non-existing user")
}

func TestGetUsers(t *testing.T) {
	db := NewDatabase()

	// Retrieve the list of users
	users, err := db.GetUsers()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(users), "Expected 4 users in the result")

	// Verify that the retrieved users match the database users
	ids := make(map[int32]bool)
	for _, user := range users {
		ids[user.ID] = true
	}
	for id := range db.Users {
		assert.True(t, ids[id], "Expected user ID to be present in the result")
	}
}
