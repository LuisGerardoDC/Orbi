package usergrcp

import (
	"context"
	"fmt"
	"log"

	"github.com/LuisGerardoDC/Orbi/UserService/src/api/proto/notificationpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NotifyUserCreated(id int) {
	conn, err := grpc.NewClient("notification-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("No se pudo conectar con NotificationService: %v", err)
	}
	defer conn.Close()

	client := notificationpb.NewNotificationServiceClient(conn)

	notificationRequest := &notificationpb.NotificationRequest{
		UserId:  fmt.Sprint(id),
		Message: "New User",
	}

	resp, err := client.SendNotification(context.Background(), notificationRequest)
	if err != nil {
		log.Printf("Error al enviar notificación: %v", err)
		return
	}

	if resp.GetSuccess() {
		log.Println("Notificación enviada con éxito")
	} else {
		log.Println("Fallo al enviar la notificación")
	}
}
