package mocks

import (
	"github.com/stretchr/testify/mock"
)

// MockRabbitMQClient es el mock de RabbitMQClient
type MockRabbitMQClient struct {
	mock.Mock
}

func (m *MockRabbitMQClient) PublishMessage(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockRabbitMQClient) Close() {

}
