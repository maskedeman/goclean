package grpc

import (
	"context"
	"errors"
	"goclean/internal/models"
	interfaces "goclean/pkg/v1"
	pb "goclean/proto"

	"google.golang.org/grpc"
)

type UserServStruct struct {
	useCase interfaces.UseCaseInterface
	pb.UnimplementedUserServiceServer
}

func NewServer(grpcServer *grpc.Server, usecase interfaces.UseCaseInterface) {
	userGrpc := &UserServStruct{useCase: usecase}
	pb.RegisterUserServiceServer(grpcServer, userGrpc)
}

func (srv *UserServStruct) Create(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserProfileResponse, error) {
	data := srv.transformUserRPC(req)
	if data.Email == "" || data.Name == "" {
		return &pb.UserProfileResponse{}, errors.New("Fill all the fields")
	}

	user, err := srv.useCase.Create(data)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

func (srv *UserServStruct) Read(ctx context.Context, req *pb.SingleUserRequest) (*pb.UserProfileResponse, error) {
	id := req.GetId()
	if id == "" {
		return &pb.UserProfileResponse{}, errors.New("Id cannot be blank")
	}
	print(req.Id)
	user, err := srv.useCase.Read(id)
	if err != nil {
		return &pb.UserProfileResponse{}, err
	}
	return srv.transformUserModel(user), nil
}

func (srv *UserServStruct) transformUserRPC(req *pb.CreateUserRequest) models.User {
	return models.User{Name: req.GetName(), Email: req.GetEmail()}
}

func (srv *UserServStruct) transformUserModel(user models.User) *pb.UserProfileResponse {
	return &pb.UserProfileResponse{Id: string(user.ID), Name: user.Name, Email: user.Email}
}
