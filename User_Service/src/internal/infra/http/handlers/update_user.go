package handlers

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type UpdateUserHandler struct {
	useCase usecase.UserUseCase
}

func (h *UpdateUserHandler) Handle(c *gin.Context) {
	var newUser entity.UserRequest

	if err := c.ShouldBindBodyWithJSON(&newUser); err != nil {
		c.JSON(400, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	user, err := h.useCase.UpdateUser(newUser)
	if err != nil {
		c.JSON(500, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, entity.Response{
		Succes:  true,
		Message: "User Updated",
		User:    user,
	})
}
