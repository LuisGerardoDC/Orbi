package handlers

import "github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"

var (
	// UseCase
	userUseCases = usecase.UserUseCase{}

	// Handlers
	ImpNewUserHandler = NewUserHandler{
		useCase: userUseCases,
	}
)
