package paths

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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

func TestUserRoute_GETUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockedUseCase := new(usecase.MockUserUseCase)
	userID := 1
	mockUser := &entity.User{ID: userID, Name: "John Doe"}
	mockedUseCase.On("GetUser", userID).Return(mockUser, nil)

	MockNewUserHandler := handlers.GetUserHandler{
		UseCase: mockedUseCase,
	}

	router.GET("/user/:id", MockNewUserHandler.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/user/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestUserRoute_POSTUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockedUseCase := new(usecase.MockUserUseCase)
	userID := 1

	mockUser := &entity.User{ID: userID, Name: "John Doe", Email: "correo@correo.co"}
	mockRequest := entity.UserRequest{Name: "John Doe", Email: "correo@correo.co"}

	mockedUseCase.On("CreateUser", mockRequest).Return(mockUser, nil)

	MockNewUserHandler := handlers.NewUserHandler{
		UseCase: mockedUseCase,
	}

	router.GET("/user", MockNewUserHandler.Handle)

	reqBody := `{"name":"John Doe","email":"correo@correo.co"}`
	req, _ := http.NewRequest(http.MethodGet, "/user", bytes.NewBufferString(reqBody))
	resp := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUserRoute_PUTUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockedUseCase := new(usecase.MockUserUseCase)
	userID := 1

	mockUser := &entity.User{ID: userID, Name: "John Doe", Email: "correo2@correo.co"}
	mockRequest := entity.UserRequest{ID: userID, Name: "John Doe", Email: "correo2@correo.co"}

	mockedUseCase.On("UpdateUser", mockRequest).Return(mockUser, nil)

	mockRabbitMQ := new(mocks.MockRabbitMQClient)
	mockRabbitMQ.On("PublishMessage", mock.Anything).Return(nil)

	handler := handlers.UpdateUserHandler{
		UseCase: mockedUseCase,
		Rabbit:  mockRabbitMQ,
	}

	router.PUT("/user/", handler.Handle)

	reqBody := `{"id":1,"name":"John Doe","email":"correo2@correo.co"}`
	req, _ := http.NewRequest(http.MethodPut, "/user/", bytes.NewBufferString(reqBody))
	resp := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestUserRoute_DELETEUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockedUseCase := new(usecase.MockUserUseCase)
	mockedDeleteAt := time.Date(2025, time.December, 15, 12, 0, 0, 0, time.UTC)
	userID := 1
	mockUser := &entity.User{
		ID:        1,
		Name:      "user 1",
		Email:     "correo@correo.vo",
		DeletedAt: &mockedDeleteAt,
	}

	mockedUseCase.On("DeleteUser", userID).Return(mockUser, nil)

	mockRabbitMQ := new(mocks.MockRabbitMQClient)
	mockRabbitMQ.On("PublishMessage", mock.Anything).Return(nil)

	handler := handlers.DeleteUserHandler{
		UseCase: mockedUseCase,
		Rabbit:  mockRabbitMQ,
	}

	router.DELETE("/user/:id", handler.Handle)

	req, _ := http.NewRequest(http.MethodDelete, "/user/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

}
