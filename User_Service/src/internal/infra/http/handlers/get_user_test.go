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

func TestGetUserHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(usecase.MockUserUseCase)
	userID := 1
	mockUser := &entity.User{ID: userID, Name: "John Doe"}
	mockUseCase.On("GetUser", userID).Return(mockUser, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(userID)}}

	handler := handlers.GetUserHandler{UseCase: mockUseCase}
	handler.Handle(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetUserHandler_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(usecase.MockUserUseCase)
	userID := 2
	mockUseCase.On("GetUser", userID).Return(nil, errors.New("user not found"))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: strconv.Itoa(userID)}}

	handler := handlers.GetUserHandler{UseCase: mockUseCase}
	handler.Handle(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockUseCase.AssertExpectations(t)
}

func TestGetUserHandler_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := new(usecase.MockUserUseCase)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "invalid"}}

	handler := handlers.GetUserHandler{UseCase: mockUseCase}
	handler.Handle(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
