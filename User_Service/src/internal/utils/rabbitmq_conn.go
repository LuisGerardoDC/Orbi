package utils

import (
	"log"
	"os"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/rabbitmq"
)

func GetRabbitMQ() *rabbitmq.RabbitMQ {
	Enable := os.Getenv("ENABLE_MQ")
	if Enable == "" {
		return nil
	}

	rabbit, err := rabbitmq.NewRabbitMQ()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	return rabbit
}
