package usecase

import (
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/infra/rabbitmq"
)

type NotificationUseCase struct {
	rabbitMQ rabbitmq.RabbitMQ
}

func NewNotificationUseCase(rabbitMQ rabbitmq.RabbitMQ) *NotificationUseCase {
	return &NotificationUseCase{rabbitMQ: rabbitMQ}
}

func (uc *NotificationUseCase) SendNotification(notification entity.Notification) error {
	// Lógica de negocio, puede involucrar validaciones o transformaciones
	return uc.rabbitMQ.Publish(notification)

}
