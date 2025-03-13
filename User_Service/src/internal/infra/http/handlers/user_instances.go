package handlers

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/utils"
)

var (
	// database
	dbConnection = utils.ConnectDB()

	// UseCase
	userUseCases = usecase.UserUseCase{
		DB: dbConnection,
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
