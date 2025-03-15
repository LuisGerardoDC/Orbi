package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http/handlers"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/rabbitmq/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(usecase.MockUserUseCase)
	mockedDeleteAt := time.Date(2025, time.December, 15, 12, 0, 0, 0, time.UTC)
	userID := 1
	mockUser := &entity.User{
		ID:        1,
		Name:      "user 1",
		Email:     "correo@correo.vo",
		DeletedAt: &mockedDeleteAt,
	}
	mockUseCase.On("DeleteUser", userID).Return(mockUser, nil)

	mockRabbitMQ := new(mocks.MockRabbitMQClient)

	mockRabbitMQ.On("PublishMessage", mock.Anything).Return(nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(userID)}}

	handler := handlers.DeleteUserHandler{
		UseCase: mockUseCase,
		Rabbit:  mockRabbitMQ,
	}
	handler.Handle(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}
