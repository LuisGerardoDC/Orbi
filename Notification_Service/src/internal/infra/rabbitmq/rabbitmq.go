package rabbitmq

import (
	"log"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/domain/entity"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
		return nil, err
	}
	return &RabbitMQ{connection: conn}, nil
}

func (r *RabbitMQ) Publish(notification entity.Notification) error {
	// Enviar la notificación a la cola de RabbitMQ
	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.Publish(
		"",      // exchange
		"queue", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(notification.Message),
		},
	)
	return err
}
