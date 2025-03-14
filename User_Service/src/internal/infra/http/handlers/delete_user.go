package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/rabbitmq"
	"github.com/gin-gonic/gin"
)

type DeleteUserHandler struct {
	useCase usecase.UserUseCase
	rabbit  *rabbitmq.RabbitMQ
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

	rabbitMesage := entity.Message{
		User:   *user,
		Action: "delete",
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
		Message: "User deleted",
		User:    user,
	})
}
