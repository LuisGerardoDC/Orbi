package handlers

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

var (
	// UseCase
	userUseCases = usecase.UserUseCase{
		DB: make(map[string]entity.User),
	}

	// Handlers
	ImpNewUserHandler = NewUserHandler{
		useCase: userUseCases,
	}

	ImpGetUserHandler = GetUserHandler{
		useCase: userUseCases,
	}

	ImpUpdateUserHandler = UpdateUserHandler{
		useCase: userUseCases,
	}

	ImpDeleteUserHandler = DeleteUserHandler{
		useCase: userUseCases,
	}
)
