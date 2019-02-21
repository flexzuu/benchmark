//go:generate protoc -I ../post --go_out=plugins=grpc:../post ../post/post.proto

package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/flexzuu/benchmark/micro-service/grpc/post/post"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo"
	"github.com/flexzuu/benchmark/micro-service/grpc/post/repo/inmemmory"
	"github.com/flexzuu/benchmark/micro-service/grpc/user/user"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement post.PostServiceServer
type server struct {
	postRepo   repo.Post
	userClient user.UserServiceClient
}

// GetPost implements post.PostServiceServer
func (s *server) GetById(ctx context.Context, in *pb.GetPostRequest) (*pb.Post, error) {
	p, err := s.postRepo.GetById(in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "get failed")
	}
	return ToProto(p), nil
}

// ListPostsOfAuthor implements post.PostServiceServer
func (s *server) ListOfAuthor(ctx context.Context, in *pb.ListPostsOfAuthorRequest) (*pb.ListPostsResponse, error) {
	ps, err := s.postRepo.ListOfAuthor(in.AuthorID)
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}
	posts := make([]*pb.Post, len(ps))
	for i, p := range ps {
		posts[i] = ToProto(p)
	}
	return &pb.ListPostsResponse{
		Posts: posts,
	}, nil
}

// CreatePost implements post.PostServiceServer
func (s *server) Create(ctx context.Context, in *pb.CreatePostRequest) (*pb.Post, error) {
	//Validate AuthorID with user service
	_, err := s.userClient.GetById(ctx, &user.GetUserRequest{
		ID: in.AuthorID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "invalid author id")
	}
	p, err := s.postRepo.Create(in.AuthorID, in.Headline, in.Content)
	if err != nil {
		return nil, errors.Wrap(err, "create failed")
	}
	return ToProto(p), nil
}

// DeletePost implements post.PostServiceServer
func (s *server) Delete(ctx context.Context, in *pb.DeletePostRequest) (*empty.Empty, error) {
	err := s.postRepo.Delete(in.ID)
	if err != nil {
		return nil, errors.Wrap(err, "delete failed")
	}
	return &empty.Empty{}, nil
}

func main() {
	postRepo := inmemmory.NewRepo()
	// Set up a connection to the server.
	userServiceAdress := os.Getenv("USER_SERVICE")
	if userServiceAdress == "" {
		log.Fatalln("please provide USER_SERVICE as env var")
	}
	conn, err := grpc.Dial(userServiceAdress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to user service: %v", err)
	}
	defer conn.Close()
	userClient := user.NewUserServiceClient(conn)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, &server{
		postRepo,
		userClient,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
