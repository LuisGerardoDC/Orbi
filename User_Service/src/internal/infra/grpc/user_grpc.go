package grpc

import (
	"context"

	"github.com/LuisGerardoDC/Orbi/UserService/src/api/proto"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/app/usecase"
	"github.com/LuisGerardoDC/Orbi/UserService/src/internal/domain/entity"
)

type UserServiceGRPC struct {
	proto.UnimplementedUserServiceServer
	useCase *usecase.UserUseCase
}

func NewUserServiceGRPC(useCase *usecase.UserUseCase) *UserServiceGRPC {
	return &UserServiceGRPC{useCase: useCase}
}

func (s *UserServiceGRPC) CreateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	user := entity.User{ID: req.Id, Name: req.Name, Email: req.Email}
	if err := s.useCase.CreateUser(user); err != nil {
		return nil, err
	}
	return &proto.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserServiceGRPC) GetUser(ctx context.Context, req *proto.UserIDRequest) (*proto.UserResponse, error) {
	user, err := s.useCase.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserServiceGRPC) UpdateUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	user := entity.User{ID: req.Id, Name: req.Name, Email: req.Email}
	if err := s.useCase.UpdateUser(user); err != nil {
		return nil, err
	}
	return &proto.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

func (s *UserServiceGRPC) DeleteUser(ctx context.Context, req *proto.UserIDRequest) (*proto.DeleteResponse, error) {
	if err := s.useCase.DeleteUser(req.Id); err != nil {
		return nil, err
	}
	return &proto.DeleteResponse{Status: "Usuario eliminado"}, nil
}
