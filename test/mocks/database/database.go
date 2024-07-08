package dbMocks

import (
	"assignment-totality-corp/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock implementation of the IDatabase interface
type MockDatabase struct {
	mock.Mock
}

func (m *MockDatabase) GetUser(id int32) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockDatabase) GetUserList(ids []int32) ([]model.User, error) {
	args := m.Called(ids)
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockDatabase) GetUsers() ([]model.User, error) {
	args := m.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockDatabase) AddUser(user model.User) (model.User, error) {
	args := m.Called(user)
	return args.Get(0).(model.User), args.Error(1)
}

func (m *MockDatabase) RemoveUser(id int32) (model.User, error) {
	args := m.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
