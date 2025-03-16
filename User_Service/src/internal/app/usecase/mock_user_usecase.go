package usecase

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) CreateUser(user entity.UserRequest) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}
func (m *MockUserUseCase) GetUser(id int) (*entity.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}
func (m *MockUserUseCase) UpdateUser(user entity.UserRequest) (*entity.User, error) {
	args := m.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}
func (m *MockUserUseCase) DeleteUser(id int) (*entity.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}
