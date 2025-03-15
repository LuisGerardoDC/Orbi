package rabbitmq

type InterfaecRabbitMQ interface {
	PublishMessage(message string) error
	Close()
}
