package usecase

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/LuisGerardoDC/Orbi/NotificationService/src/internal/domain/entity"
)

type NotificationUseCase struct {
}

func (uc *NotificationUseCase) SendNotification(notification string) error {
	message := entity.Message{}
	jsonBytes := []byte(notification)
	err := json.Unmarshal(jsonBytes, &message)

	if err != nil {
		return err
	}

	basePath := os.Getenv("EMAIL_STORAGE_PATH")
	if basePath == "" {
		basePath = "/data/emails"
	}
	dirPath := filepath.Join(basePath, fmt.Sprintf("NotificationsID_%d", message.User.ID))
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s_%s.html", message.Action, time.Now().Format("2006-01-02T15:04:05"))
	filePath := filepath.Join(dirPath, fileName)

	template := getTemplate(message)
	err = os.WriteFile(filePath, []byte(template), os.ModePerm)
	if err != nil {
		return err
	}

	logPath := os.Getenv("LOG_PATH")
	err = os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(fmt.Sprintf("%slogs.txt", logPath), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintln(notification))

	return err
}

func getTemplate(message entity.Message) string {
	return fmt.Sprintf("<h1>%s</h1>", message.Action)
}
