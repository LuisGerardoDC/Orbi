package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"notifications",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{connection: conn, channel: ch, queue: q}, nil
}

func (r *RabbitMQ) ConsumeMessages() (<-chan amqp.Delivery, error) {
	messages, err := r.channel.Consume(
		r.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *RabbitMQ) Close() {
	r.channel.Close()
	r.connection.Close()
}
