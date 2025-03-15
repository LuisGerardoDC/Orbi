package handlers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/infra/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserHandler_Handle_Success(t *testing.T) {
	mockUseCase := usecase.MockUserUseCase{}
	handler := handlers.GetUserHandler{
		UseCase: &mockUseCase,
	}

	userID := 1
	user := entity.User{ID: userID, Name: "John Doe"}
	mockUseCase.On("GetUser", userID).Return(user, nil)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:id", handler.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/user/"+strconv.Itoa(userID), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertCalled(t, "GetUser", userID)
}

func TestGetUserHandler_Handle_InvalidID(t *testing.T) {
	mockUseCase := usecase.MockUserUseCase{}
	handler := handlers.GetUserHandler{
		UseCase: &mockUseCase,
	}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:id", handler.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/user/invalid", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockUseCase.AssertNotCalled(t, "GetUser")
}

func TestGetUserHandler_Handle_UserNotFound(t *testing.T) {
	mockUseCase := usecase.MockUserUseCase{}
	handler := handlers.GetUserHandler{
		UseCase: &mockUseCase,
	}

	userID := 1
	mockUseCase.On("GetUser", userID).Return(entity.User{}, errors.New("user not found"))

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/user/:id", handler.Handle)

	req, _ := http.NewRequest(http.MethodGet, "/user/"+strconv.Itoa(userID), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUseCase.AssertCalled(t, "GetUser", userID)
}
