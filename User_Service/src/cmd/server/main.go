package main

import (
	"log"
	"net"

	pb "github.com/LuisGerardoDC/Orbi/UserService/src/api/proto"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	usergrpc "github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/grpc"
	Server "github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	go func() {
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
	}()

	gin.SetMode(gin.DebugMode)
	router := Server.GetRouter()
	router.Run(":8080")

}
