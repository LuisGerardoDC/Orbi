package usecase

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

type InterfaceUserUseCase interface {
	CreateUser(user entity.UserRequest) (*entity.User, error)
	GetUser(id int) (*entity.User, error)
	UpdateUser(user entity.UserRequest) (*entity.User, error)
	DeleteUser(id int) (*entity.User, error)
}
