package handlers

import (
	"log"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/rabbitmq"
	"github.com/gin-gonic/gin"
)

type NewUserHandler struct {
	useCase usecase.UserUseCase
	rabbit  *rabbitmq.RabbitMQ
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

	if err := h.useCase.CreateUser(newUser); err != nil {
		c.JSON(500, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	message := "User created" + newUser.Name
	err := h.rabbit.PublishMessage(message)
	if err != nil {
		log.Printf("Error publishing message: %s", err)
	}

	c.JSON(200, entity.Response{
		Succes:  true,
		Message: "User created",
	})
}
