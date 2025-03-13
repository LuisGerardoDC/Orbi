package main

import (
	"log"
	"net"

	pb "github.com/LuisGerardoDC/Orbi/UserService/src/api/proto"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	usergrpc "github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/grpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	useCase := usecase.NewUserUseCase()
	pb.RegisterUserServiceServer(grpcServer, usergrpc.NewUserServiceGRPC(useCase))

	log.Println("UserService escuchando en el puerto 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al iniciar servidor: %v", err)
	}
}
