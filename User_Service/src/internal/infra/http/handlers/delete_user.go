package handlers

import (
	"strconv"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type DeleteUserHandler struct {
	useCase usecase.UserUseCase
}

func (h *DeleteUserHandler) Handle(c *gin.Context) {
	userID := c.Param("id")
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(400, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return

	}
	user, err := h.useCase.DeleteUser(userIDInt)

	if err != nil {
		c.JSON(500, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, entity.Response{
		Succes:  true,
		Message: "User deleted",
		User:    user,
	})
}
