package handlers

import (
	"encoding/json"
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
	user, err := h.useCase.CreateUser(newUser)

	if err != nil {
		c.JSON(500, entity.Response{
			Succes:  false,
			Message: err.Error(),
		})
		return
	}

	rabbitMesage := entity.Message{
		User:   *user,
		Action: "create",
	}
	jsonBytes, err := json.Marshal(rabbitMesage)
	if err != nil {
		log.Printf("Error marshalling user: %s", err)
	}
	message := string(jsonBytes)

	err = h.rabbit.PublishMessage(message)
	if err != nil {
		log.Printf("Error publishing message: %s", err)
	}

	c.JSON(200, entity.Response{
		Succes:  true,
		Message: "User created",
	})
}
