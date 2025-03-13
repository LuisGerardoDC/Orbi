package service

import (
	"context"
	"log"

	pb "github.com/LuisGerardoDC/Orbi/NotificationService/src/api/proto"
)

type NotificationService struct {
	pb.UnimplementedNotificationServiceServer
}

func (s *NotificationService) SendNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	log.Printf("Notificación recibida: %s", req.Message)
	return &pb.NotificationResponse{Status: "Enviado"}, nil
}
