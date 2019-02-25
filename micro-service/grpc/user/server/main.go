//go:generate protoc -I ../user --go_out=plugins=grpc:../user ../user/user.proto

package main

import (
	"context"
	"log"
	"net"

	"github.com/flexzuu/benchmark/micro-service/grpc/stats"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/repo"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/repo/inmemmory"
	pb "github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

// server is used to implement user.UserServiceServer
type server struct {
	userRepo repo.User
	stats.ServiceHelper
}

// GetUser implements user.UserServiceServer
func (s *server) GetById(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	u, err := s.userRepo.Get(in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get failed")
	}
	s.Count++
	return ToProto(u), nil
}

// CreateUser implements user.UserServiceServer
func (s *server) Create(ctx context.Context, in *pb.CreateUserRequest) (*pb.User, error) {
	u, err := s.userRepo.Create(in.Email, in.Name)
	if err != nil {
		return nil, errors.Wrap(err, "create failed")
	}
	s.Count++
	return ToProto(u), nil
}

// DeleteUser implements user.UserServiceServer
func (s *server) Delete(ctx context.Context, in *pb.DeleteUserRequest) (*empty.Empty, error) {
	err := s.userRepo.Delete(in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "delete failed")
	}
	s.Count++
	return &empty.Empty{}, nil
}

func main() {
	userRepo := inmemmory.NewRepo()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	ServiceHelper := stats.ServiceHelper{Count: 0}

	pb.RegisterUserServiceServer(s, &server{
		userRepo,
		ServiceHelper,
	})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
