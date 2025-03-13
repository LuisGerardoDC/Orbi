package handlers

import (
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type GetUserHandler struct {
	useCase usecase.UserUseCase
}

func (h *GetUserHandler) Handle(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.useCase.GetUser(userID)
	if err != nil {
		c.JSON(500, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, entity.Response{
		Succes: true,
		User:   user,
	})
}
