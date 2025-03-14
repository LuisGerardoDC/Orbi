package notifgrpc

import (
	"context"
	"log"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/api/proto"
)

type NotificationServer struct {
	proto.UnimplementedNotificationServiceServer
}

func (s *NotificationServer) SendNotification(ctx context.Context, req *proto.NotificationRequest) (*proto.NotificationResponse, error) {
	log.Printf("Received notification for user %s: %s", req.UserId, req.Message)

	return &proto.NotificationResponse{Success: true}, nil
}
