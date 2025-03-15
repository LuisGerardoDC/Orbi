package handlers

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/utils"
)

var (
	// services
	dbConnection = utils.ConnectDB()
	rabbit       = utils.GetRabbitMQ()

	// UseCase
	userUseCases = &usecase.UserUseCase{
		DB: dbConnection,
	}

	// Handlers
	ImpNewUserHandler = NewUserHandler{
		useCase: userUseCases,
		rabbit:  rabbit,
	}

	ImpGetUserHandler = GetUserHandler{
		UseCase: userUseCases,
	}

	ImpUpdateUserHandler = UpdateUserHandler{
		useCase: userUseCases,
		rabbit:  rabbit,
	}

	ImpDeleteUserHandler = DeleteUserHandler{
		useCase: userUseCases,
		rabbit:  rabbit,
	}
)
