package main

import (
	"fmt"
	"log"
	"net"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/api/proto"
	notigrpc "github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/infra/grpc"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/infra/rabbitmq"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		rabbitmq.StartConsumer()
	}()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterNotificationServiceServer(grpcServer, &notigrpc.NotificationServer{})

	fmt.Println("Notification Service running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
