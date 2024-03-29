package grpchandler

import (
	"context"
	"fmt"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/ZaiPeeKann/puregrade/internal/service"
	pb "github.com/ZaiPeeKann/puregrade/internal/transport/grpc/grpchandler"
)

type GRPCServer struct {
	services *service.Service
	pb.UnimplementedAuthServer
}

func NewGRPCServer(services *service.Service) *GRPCServer {
	return &GRPCServer{services: services}
}

func (s *GRPCServer) mustEmbedUnimplementedAuthServer() { return }

func (s *GRPCServer) SingIn(ctx context.Context, req *pb.SingInRequest) (*pb.SingInResponse, error) {

	access, refresh, err := s.services.Authorization.GenerateTokens(req.GetUsername(), req.GetPassword())

	return &pb.SingInResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, err
}

func (s *GRPCServer) SingUp(ctx context.Context, req *pb.SingUpRequest) (*pb.SingUpResponse, error) {
	var user puregrade.User
	user.Username = req.GetUsername()
	user.Email = req.GetEmail()
	user.Password = req.GetPassword()
	user.Avatar = req.GetAvatar()
	for _, v := range req.GetRoles() {
		fmt.Print(int(v))
		user.Roles = append(user.Roles, int64(v))
	}
	id, err := s.services.Authorization.CreateUser(user)

	return &pb.SingUpResponse{
		Id: int32(id),
	}, err
}
