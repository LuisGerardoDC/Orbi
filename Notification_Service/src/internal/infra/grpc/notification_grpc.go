package grpc

import (
	"context"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/api/proto"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/usecase"
)

type NotificationServiceGRPC struct {
	useCase *usecase.NotificationUseCase
}

func NewNotificationServiceGRPC(useCase *usecase.NotificationUseCase) *NotificationServiceGRPC {
	return &NotificationServiceGRPC{useCase: useCase}
}

func (s *NotificationServiceGRPC) SendNotification(ctx context.Context, req *proto.NotificationRequest) (*proto.NotificationResponse, error) {
	notification := entity.Notification{Message: req.Message}
	if err := s.useCase.SendNotification(notification); err != nil {
		return nil, err
	}

	return &proto.NotificationResponse{Status: "Enviado"}, nil
}
