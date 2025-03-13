package usecase

import (
	"errors"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

type UserUseCase struct {
	DB map[string]entity.User
}

func (uc *UserUseCase) CreateUser(user entity.UserRequest) error {
	if _, exists := uc.DB[user.Email]; exists {
		return errors.New("usuario ya existe")
	}
	uc.DB[user.Email] = entity.User{
		ID:    len(uc.DB),
		Name:  user.Name,
		Email: user.Email,
	}

	return nil
}

func (uc *UserUseCase) GetUser(email string) (entity.User, error) {
	user, exists := uc.DB[email]
	if !exists {
		return entity.User{}, errors.New("usuario no encontrado")
	}
	return user, nil
}

func (uc *UserUseCase) UpdateUser(user entity.UserRequest) error {
	if _, exists := uc.DB[user.Email]; !exists {
		return errors.New("usuario no encontrado")
	}
	uc.DB[user.Email] = entity.User{
		Name:  user.Name,
		Email: user.Email,
	}
	return nil
}

func (uc *UserUseCase) DeleteUser(id string) error {
	if _, exists := uc.DB[id]; !exists {
		return errors.New("usuario no encontrado")
	}
	delete(uc.DB, id)
	return nil
}
