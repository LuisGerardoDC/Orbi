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
		UseCase: userUseCases,
	}

	ImpGetUserHandler = GetUserHandler{
		UseCase: userUseCases,
	}

	ImpUpdateUserHandler = UpdateUserHandler{
		UseCase: userUseCases,
		Rabbit:  rabbit,
	}

	ImpDeleteUserHandler = DeleteUserHandler{
		UseCase: userUseCases,
		Rabbit:  rabbit,
	}
)
