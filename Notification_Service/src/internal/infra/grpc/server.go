package notifgrpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/api/proto"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/usecase"
)

type NotificationServer struct {
	proto.UnimplementedNotificationServiceServer
	usecase *usecase.NotificationNewUser
}

func (s *NotificationServer) SendNotification(ctx context.Context, req *proto.NotificationRequest) (*proto.NotificationResponse, error) {
	log.Printf("Received notification for user %s: %s", req.UserId, req.Message)
	baseurl := os.Getenv("USER_SERVICE_URL")
	baseport := os.Getenv("USER_SERVICE_PORT")

	url := fmt.Sprintf("http://%s:%s/user/%s", baseurl, baseport, req.UserId)
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Failed to get user: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to get user: %v", resp.Status)
	}
	var response entity.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Printf("Failed to decode user: %v", err)
	}

	s.usecase.SendNotification(*response.User)

	return &proto.NotificationResponse{Success: true}, nil
}
