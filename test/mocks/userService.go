package mocks

import (
	"assignment-totality-corp/internal/model"
	"assignment-totality-corp/internal/service"

	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of IUserService
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUserById(id int32) model.User {
	args := m.Called(id)
	return args.Get(0).(model.User)
}

func (m *MockUserService) GetUserByIds(ids []int32) []model.User {
	args := m.Called(ids)
	return args.Get(0).([]model.User)
}

func (m *MockUserService) SearchUsers(req service.SearchUsersRequest) []model.User {
	args := m.Called(req)
	return args.Get(0).([]model.User)
}
