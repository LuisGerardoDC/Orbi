package main

import (
	"log"
	"net"

	pb "github.com/LuisGerardoDC/Orbi/NotificationService/src/api/proto"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/infra/rabbitmq"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/service"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		rabbitmq.StartConsumer()
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar : %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNotificationServiceServer(grpcServer, &service.NotificationService{})

	log.Println("Servidor gRPC escuchando en el puerto 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar servidor :%v", err)
	}
}
