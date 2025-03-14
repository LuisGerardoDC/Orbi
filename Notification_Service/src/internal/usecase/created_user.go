package usecase

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/domain/entity"
)

type NotificationNewUser struct {
}

func (uc *NotificationNewUser) SendNotification(user entity.User) error {

	basePath := os.Getenv("EMAIL_STORAGE_PATH")
	if basePath == "" {
		basePath = "/data/emails"
	}
	dirPath := filepath.Join(basePath, fmt.Sprintf("NotificationsID_%d", user.ID))
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s_%s.html", "new user", time.Now().Format("2006-01-02T15:04:05"))
	filePath := filepath.Join(dirPath, fileName)

	template := fmt.Sprintf("<h1>Bienvenido %d</h1>", user.ID)
	err = os.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
