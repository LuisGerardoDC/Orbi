package rabbitmq

import (
	"log"
)

func StartConsumer() {
	rabbit, err := NewRabbitMQ()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbit.Close()

	messages, err := rabbit.ConsumeMessages()
	if err != nil {
		log.Fatalf("Error al consumir mensajes: %v", err)
	}

	for msg := range messages {
		log.Printf("Mensaje recibido: %s", msg.Body)
		// Aquí puedes agregar lógica para manejar el mensaje, como enviar una notificación
	}
}
