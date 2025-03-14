package utils

import (
	"log"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/rabbitmq"
)

func GetRabbitMQ() *rabbitmq.RabbitMQ {
	rabbit, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}

	return rabbit
}
