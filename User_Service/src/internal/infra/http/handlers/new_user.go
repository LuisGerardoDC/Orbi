package handlers

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	usergrcp "github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/grcp"
	"github.com/gin-gonic/gin"
)

type NewUserHandler struct {
	UseCase usecase.InterfaceUserUseCase
}

func (h *NewUserHandler) Handle(c *gin.Context) {
	var newUser entity.UserRequest

	if err := c.ShouldBindBodyWithJSON(&newUser); err != nil {
		c.JSON(400, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}
	user, err := h.UseCase.CreateUser(newUser)

	if err != nil {
		c.JSON(500, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	go usergrcp.NotifyUserCreated(user.ID)

	c.JSON(200, entity.Response{
		Succes:  true,
		Message: "User created",
	})
}
