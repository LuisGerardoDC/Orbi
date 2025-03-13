package usecase

import (
	"errors"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

type UserUseCase struct {
	users map[string]entity.User
}

func NewUserUseCase() *UserUseCase {
	return &UserUseCase{users: make(map[string]entity.User)}
}

func (uc *UserUseCase) CreateUser(user entity.User) error {
	if _, exists := uc.users[user.ID]; exists {
		return errors.New("usuario ya existe")
	}
	uc.users[user.ID] = user
	return nil
}

func (uc *UserUseCase) GetUser(id string) (entity.User, error) {
	user, exists := uc.users[id]
	if !exists {
		return entity.User{}, errors.New("usuario no encontrado")
	}
	return user, nil
}

func (uc *UserUseCase) UpdateUser(user entity.User) error {
	if _, exists := uc.users[user.ID]; !exists {
		return errors.New("usuario no encontrado")
	}
	uc.users[user.ID] = user
	return nil
}

func (uc *UserUseCase) DeleteUser(id string) error {
	if _, exists := uc.users[id]; !exists {
		return errors.New("usuario no encontrado")
	}
	delete(uc.users, id)
	return nil
}
